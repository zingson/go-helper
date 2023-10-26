package hcron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

/*
使用说明：
var h = hcron.New(storage)
h.AddFunc(key,spec,func)

spec格式： Second | Minute | Hour | Dom | Month | Dow
spec示例： 0 0 0 * * *

*/

// Storage 任务存储，如：存储到数据库表
type Storage interface {
	Find() []*Entry                                  // 1. 查询有变更的任务 2.当前时间大于下一次执行时间的任务
	Save(entry *Entry)                               // 保存任务，key已存在则不更新
	Lock(key string, prev, next time.Time) (ok bool) // 根据Key+prev加锁
	Log(entry *Entry)                                // 记录日志
}

// New 计划任务
func New(storage Storage) *Cron {
	t := &Cron{c: cron.New(cron.WithSeconds()), storage: storage, entry: sync.Map{}, handler: sync.Map{}}
	t.AddFunc("RefreshJob", "33 0/3 * * * *", t.refreshJob) // 定时从存储中读取任务加入
	t.c.Start()
	return t
}

// Cron 计划任务
type Cron struct {
	c       *cron.Cron //
	storage Storage    // 任务存储
	entry   sync.Map   // 任务对象  key -> Entry
	handler sync.Map   // 任务处理  name -> Handler
}

func (c *Cron) AddFunc(key, spec string, f func()) {
	c.Handler(newCustomHandler(key, f)).addJob(key, key, spec, nil)
	println(fmt.Sprintf("hcron AddFunc key=%s spec=%s ", key, spec))
}

func (c *Cron) AddJob(key, spec string, job cron.Job) {
	c.AddFunc(key, spec, job.Run)
}

// Handler 任务处理
type Handler interface {
	Name() string
	Run(v []byte)
}

func (c *Cron) Handler(handlers ...Handler) *Cron {
	for _, handler := range handlers {
		c.handler.Store(handler.Name(), handler)
	}
	return c
}

// Entry 任务对象
type Entry struct {
	Key     string       // 唯一标识
	Label   string       // 显示名字
	Spec    string       // 时间表达式
	Handler string       // 处理函数类型
	Args    []byte       // 参数
	Err     error        // 错误
	Enable  bool         // true=正常 false=禁用
	Next    time.Time    // 下一次时间
	Prev    time.Time    // 上一次时间
	EntryID cron.EntryID // ID，仅相当于当前实例唯一

	schedule cron.Schedule
}

func (c *Cron) GetEntry(key string) (entry *Entry) {
	v, ok := c.entry.Load(key)
	if !ok {
		v = &Entry{Key: key}
		c.entry.Store(key, v)
	}
	entry = v.(*Entry)
	return
}

// 新增计划任务
func (c *Cron) addJob(key, handler, spec string, args []byte) {
	if spec == "" || key == "" || handler == "" {
		return
	}
	entry := c.GetEntry(key)
	if entry.Spec == spec && len(entry.Args) == len(args) {
		return
	}

	c.Remove(key)
	entryID, err := c.c.AddJob(spec, cron.NewChain(cron.SkipIfStillRunning(cron.DefaultLogger), lock(c, key)).Then(run(c, handler, args)))
	if err != nil {
		println("【 key =", key, "】 AddJobError:", err.Error())
	}

	entry.EntryID = entryID
	ey := c.c.Entry(entryID)
	if ey.Schedule != nil {
		entry.schedule = ey.Schedule
		entry.Next = entry.schedule.Next(time.Now())
	}
	entry.Spec = spec
	entry.Err = err
	entry.Handler = handler
	entry.Args = args
	entry.Enable = true
	c.storage.Save(entry)
}

// 运行
func run(c *Cron, handler string, args []byte) cron.FuncJob {
	return func() {
		if v, ok := c.handler.Load(handler); ok {
			v.(Handler).Run(args)
		}
	}
}

// 全局锁+运行日志
func lock(c *Cron, key string) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		return cron.FuncJob(func() {
			entry := c.GetEntry(key)
			if entry.schedule == nil {
				return
			}
			if entry.Next.IsZero() {
				entry.Next = entry.schedule.Next(time.Now())
			}
			entry.Prev = entry.Next
			entry.Next = entry.schedule.Next(entry.Prev)

			// 加全局锁
			if ok := c.storage.Lock(entry.Key, entry.Prev, entry.Next); !ok {
				return
			}

			//记录日志
			defer func() {
				if r := recover(); r != nil {
					entry.Err = r.(error)
				}
				c.storage.Log(entry)
			}()
			j.Run()
		})
	}
}

// Run 立即执行任务
func (c *Cron) Run(key string) {
	e := c.GetEntry(key)
	run(c, e.Handler, e.Args)()
}

// Remove 删除任务
func (c *Cron) Remove(key string) {
	c.c.Remove(c.GetEntry(key).EntryID)
}

// 读取任务加入到计划
func (c *Cron) refreshJob() {
	list := c.storage.Find()
	if list == nil || len(list) == 0 {
		return
	}
	for _, item := range list {
		if !item.Enable {
			c.Remove(item.Key)
			continue
		}
		c.addJob(item.Key, item.Handler, item.Spec, item.Args)
	}
	return
}

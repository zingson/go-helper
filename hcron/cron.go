package hcron

import (
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

// New 计划任务
func New(storage Storage) *Cron {
	t := &Cron{Cron: cron.New(cron.WithSeconds()), storage: storage, entry: sync.Map{}, handler: sync.Map{}}
	t.AddFunc("RefreshJob", "33 0/3 * * * *", t.refreshJob) // 定时从存储中读取任务加入
	t.Start()
	return t
}

// Cron 计划任务
type Cron struct {
	*cron.Cron          //
	storage    Storage  // 任务存储
	entry      sync.Map // 任务对象  key -> Entry
	handler    sync.Map // 任务处理  name -> Handler
}

// Handler 任务处理
type Handler interface {
	Name() string
	Run(v []byte)
}

// Entry 任务对象
type Entry struct {
	Key        string // 唯一标识
	Spec       string //
	Handler    string //
	Args       []byte //
	Err        error  //
	Enable     bool   // true=正常 false=禁用
	cron.Entry        //
}

// Storage 任务存储，如：存储到数据库表
type Storage interface {
	Find() []*Entry                                  // 1. 查询有变更的任务 2.解锁当前时间大于下一次执行时间的任务
	Save(entry *Entry)                               // 保存任务，key已存在则不更新
	Lock(key string, prev, next time.Time) (ok bool) // 根据Key与前后时间加锁
	Log(entry *Entry, ms int64)                      // 记录日志
}

func (c *Cron) Handler(handlers ...Handler) *Cron {
	for _, handler := range handlers {
		c.handler.Store(handler.Name(), handler)
	}
	return c
}

type customHandler struct {
	name string
	f    func()
}

func newCustomHandler(name string, f func()) *customHandler { return &customHandler{name, f} }

func (o *customHandler) Name() string { return o.name }

func (o *customHandler) Run([]byte) { o.f() }

func (c *Cron) PutEntry(entry ...*Entry) {
	for _, item := range entry {
		c.entry.Store(item.Key, item)
	}
}

func (c *Cron) GetEntry(key string) (entry *Entry) {
	v, ok := c.entry.Load(key)
	if ok {
		entry = v.(*Entry)
	}
	entry = &Entry{}
	return
}

func (c *Cron) AddFunc(key, spec string, f func()) {
	c.Handler(newCustomHandler(key, f)).addJob(key, key, spec, nil)
}

func (c *Cron) AddJob(key, spec string, job cron.Job) {
	c.Handler(newCustomHandler(key, job.Run)).addJob(key, key, spec, nil)
}

func (c *Cron) addJob(key, handler, spec string, args []byte) {
	entry := c.GetEntry(key)
	if entry.Spec == spec && len(entry.Args) == len(args) {
		return
	}
	c.Remove(key)
	entryID, err := c.Cron.AddJob(spec, cron.NewChain(skipIfStillRunning(c, key), log(c, key)).Then(run(c, handler, args)))
	if err != nil {
		println("【 key =", key, "】 AddJobError:", err.Error())
	}
	entry = &Entry{Key: key, Spec: spec, Err: err, Handler: handler, Args: args, Enable: true, Entry: c.Entry(entryID)}
	c.PutEntry(entry)
	c.storage.Save(entry)
}

func run(c *Cron, handler string, args []byte) cron.FuncJob {
	return func() {
		if v, ok := c.handler.Load(handler); ok {
			v.(Handler).Run(args)
		}
	}
}

// skipIfStillRunning 如果任务在执行则跳过
func skipIfStillRunning(c *Cron, key string) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		return cron.FuncJob(func() {
			entry := c.GetEntry(key)
			if ok := c.storage.Lock(entry.Key, entry.Prev, entry.Next); !ok {
				return
			}
			j.Run()
		})
	}
}

// log 记任务执行日志
func log(c *Cron, key string) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		return cron.FuncJob(func() {
			var entry = c.GetEntry(key)
			var milli = time.Now().UnixMilli()
			var err error
			defer func() {
				if r := recover(); r != nil {
					err = r.(error)
				}
				entry.Err = err
				entry.Entry = c.Entry(entry.ID)
				c.storage.Log(entry, time.Now().UnixMilli()-milli)
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
	c.Cron.Remove(c.GetEntry(key).ID)
}

// 读取任务加入到计划
func (c *Cron) refreshJob() {
	list := c.storage.Find()
	if list == nil || len(list) == 0 {
		return
	}
	for _, item := range list {
		if item.Enable {
			c.addJob(item.Key, item.Handler, item.Spec, item.Args)
			continue
		}
		c.Remove(item.Key)
	}
	return
}

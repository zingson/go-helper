package hcron

// 自定义Handler
type customHandler struct {
	name string
	f    func()
}

func newCustomHandler(name string, f func()) *customHandler { return &customHandler{name, f} }

func (o *customHandler) Name() string { return o.name }

func (o *customHandler) Run([]byte) { o.f() }

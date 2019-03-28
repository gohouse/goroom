package goroom

import (
	"html/template"
	"sync"
)
type View struct {
	*template.Template
}

var view *View
var view_once sync.Once

func GetViewInstance() *View {
	view_once.Do(func() {
		view = new(View)
	})
	return view
}

func (v *View) Boot(args ...interface{}) func(*GoRoom) {
	return func(srv *GoRoom) {
		// 初始化数据库链接
		// 这一步是为了确保单例初始化
		v = GetViewInstance()
		v.Template = new(template.Template)
		srv.View = v
	}
}

func (v *View) New(arg string) *View {
	v.Template = template.New(arg)
	return v
}

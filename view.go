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

func NewView() *View {
	view_once.Do(func() {
		view = new(View)
	})
	return view
}

func DefaultView() func(*GoRoom) {
	return func(srv *GoRoom) {
		// 初始化数据库链接
		// 这一步是为了确保单例初始化
		view = NewView()
		view.Template = template.New("")
	}
}

package goroom

type IGoRoom interface {
	Boot(args ...interface{}) func(room *Engin)
}

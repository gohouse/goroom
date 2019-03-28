package goroom

import (
	"testing"
)

func TestNew(t *testing.T) {
	engin := New()
	t.Log(engin)
}

func TestVersion(t *testing.T) {
	t.Log(Version)
}
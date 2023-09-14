package smog5

import (
	"testing"
)

func TestUniverse(t *testing.T) {
	u := NewUniverse()
	if u == nil {
		t.Error("NewUniverse() returned nil")
	}
}

func TestHash(t *testing.T) {
	if hash("foo") != hash("foo") {
		t.Error("hash(\"foo\") != hash(\"foo\")")
	}
	if hash("foo") == hash("bar") {
		t.Error("hash(\"foo\") == hash(\"bar\")")
	}

	t1 := hash("foo")
	t2 := hash("bar")
	t3 := hash("kris")
	t4 := hash("kristofer")
	t5 := hash("when in the course of human events")
	t6 := hash("when in the course of human events, it becomes necessary for one people to dissolve the political bands which have connected them with another, and to assume among the powers of the earth, the separate and equal station to which the laws of nature and of nature's god entitle them, a decent respect to the opinions of mankind requires that they should declare the causes which impel them to the separation.")
	t7 := hash("hen in the course of human events, it becomes necessary for one people to dissolve the political bands which have connected them with another, and to assume among the powers of the earth, the separate and equal station to which the laws of nature and of nature's god entitle them, a decent respect to the opinions of mankind requires that they should declare the causes which impel them to the separation.")
	t.Error("\nt1:", t1, "t2:", t2, "\nt3:", t3, "t4:", t4, "\nt5:", t5, "\nt6:", t6, "\nt7:", t7)
}

func TestUniverseNewObject(t *testing.T) {
	u := NewUniverse().Initialize()
	c := u.NewClass("Foo")
	if c == nil {
		t.Error("NewClass() returned nil")
	}
	if c.Name != "Foo" {
		t.Error("NewClass() returned object with wrong class")
	}
	o := u.NewObject("Foo")
	if o == nil {
		t.Error("NewObject() returned nil")
	}
	if o.Clazz.Name != "Foo" {
		t.Error("NewObject() returned object with wrong class")
	}
}

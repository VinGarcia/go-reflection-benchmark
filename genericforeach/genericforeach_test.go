package genericforeach

import (
	"fmt"
	"testing"
)

var result string

var listLen = 100
var numLists = 50

func BenchmarkReflectionWithPtr(b *testing.B) {
	lists := [][]*Foo{}
	for i := 0; i < numLists; i++ {
		list := []*Foo{}
		for j := 0; j < listLen; j++ {
			list = append(list, &Foo{
				A: i + j,
				B: fmt.Sprint("foobar", i, "_", j),
			})
		}
		lists = append(lists, list)
	}

	b.Run("testing GenericForEach", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			genericForEach(lists[i%numLists], func(f interface{}) {
				result = f.(*Foo).B
			})
		}
	})

	b.Run("testing StaticForEachWithPtr", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			staticForEachWithPtr(lists[i%numLists], func(f interface{}) {
				result = f.(*Foo).B
			})
		}
	})

	iLists := [][]interface{}{}
	for i := 0; i < numLists; i++ {
		list := []interface{}{}
		for j := 0; j < listLen; j++ {
			list = append(list, &Foo{
				A: i + j,
				B: fmt.Sprint("foobar", i, "_", j),
			})
		}
		iLists = append(iLists, list)
	}

	b.Run("testing ForEachWithInterface", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			forEachWithInterface(iLists[i%numLists], func(f interface{}) {
				result = f.(*Foo).B
			})
		}
	})
}

func BenchmarkReflection(b *testing.B) {
	lists := [][]Foo{}
	for i := 0; i < numLists; i++ {
		list := []Foo{}
		for j := 0; j < listLen; j++ {
			list = append(list, Foo{
				A: i + j,
				B: fmt.Sprint("foobar", i, "_", j),
			})
		}
		lists = append(lists, list)
	}

	b.Run("testing GenericForEach", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			genericForEach(lists[i%numLists], func(f interface{}) {
				result = f.(Foo).B
			})
		}
	})

	b.Run("testing StaticForEach", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			staticForEach(lists[i%numLists], func(f interface{}) {
				result = f.(Foo).B
			})
		}
	})
}

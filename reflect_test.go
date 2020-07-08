package main

import (
	"fmt"
	"reflect"
	"testing"
)

var result interface{}

func BenchmarkReflection(b *testing.B) {
	fs := []Foo{}
	for i := 0; i < 100; i++ {
		fs = append(fs, Foo{
			A: i,
			B: fmt.Sprint("foobar", i),
			C: strp(fmt.Sprint("barfoo", i)),
			D: []string{"foo", "bar"},
			E: map[string]interface{}{"foo": "bar"},
			F: MyStruct{A: 10 + i, B: 22.2 + float64(i)},
			G: &MyStruct{A: 20 + i, B: 44.4 + float64(i)},
		})
	}

	b.Run("testing toMapWithNoReflection", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = toMapWithNoReflection(fs[i%100])
		}
	})

	b.Run("testing toMap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = toMap(fs[i%100])
		}
	})

	t := reflect.TypeOf(Foo{})
	b.Run("testing toMapWithCachedType", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = toMapWithCachedType(t, fs[i%100])
		}
	})

	b.Run("testing toMapUsingTag", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = toMapUsingTag(fs[i%100])
		}
	})

	b.Run("testing toMapUsingTagWithCachedType", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = toMapUsingTagWithCachedType(t, fs[i%100])
		}
	})

	b.Run("testing toMapUsingCachedTagNames", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = toMapUsingCachedTagNames(fs[i%100])
		}
	})

	is := []interface{}{}
	for i := 0; i < 100; i++ {
		is = append(is, Foo{
			A: i,
			B: fmt.Sprint("foobar", i),
			C: strp(fmt.Sprint("barfoo", i)),
			D: []string{"foo", "bar"},
			E: map[string]interface{}{"foo": "bar"},
			F: MyStruct{A: 10 + i, B: 22.2 + float64(i)},
			G: &MyStruct{A: 20 + i, B: 44.4 + float64(i)},
		})
	}
	type ToMapper interface {
		ToMap() map[string]interface{}
	}
	b.Run("testing toMapUsingMethod&DuckTyping", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result = is[i%100].(ToMapper).ToMap()
		}
	})
}

func strp(s string) *string {
	return &s
}

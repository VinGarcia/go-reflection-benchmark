package genericforeach

import "reflect"

type Foo struct {
	A int
	B string
}

func genericForEach(list interface{}, callback func(interface{})) {
	v := reflect.ValueOf(list)
	t := v.Type()

	if t.Kind() != reflect.Slice {
		panic("type must be a slice!")
	}

	for i := 0; i < v.Len(); i++ {
		callback(v.Index(i).Interface())
	}
}

func staticForEachWithPtr(list []*Foo, callback func(interface{})) {
	for i := 0; i < len(list); i++ {
		callback(list[i])
	}
}

func staticForEach(list []Foo, callback func(interface{})) {
	for i := 0; i < len(list); i++ {
		callback(list[i])
	}
}

func forEachWithInterface(list []interface{}, callback func(interface{})) {
	for i := 0; i < len(list); i++ {
		callback(list[i])
	}
}

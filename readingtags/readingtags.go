package readingtags

import (
	"reflect"
)

type Bar interface{}
type MyStruct struct {
	A int
	B float64
}

type Foo struct {
	A int                    `foo:"v1"`
	B string                 `foo:"v2"`
	C *string                `foo:"v3"`
	D []string               `foo:"v4"`
	E map[string]interface{} `foo:"v5"`
	F MyStruct               `foo:"v6"`
	G Bar                    `foo:"v7"`
}

func (f Foo) ToMap() map[string]interface{} {
	return toMapWithNoReflection(f)
}

func toMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(Foo{})
	v := reflect.ValueOf(obj)
	m := map[string]interface{}{}
	for i := 0; i < v.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}

	return m
}

func toMapWithCachedType(t reflect.Type, obj interface{}) map[string]interface{} {
	v := reflect.ValueOf(obj)
	m := map[string]interface{}{}
	for i := 0; i < v.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}

	return m
}

func toMapUsingTag(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(Foo{})
	v := reflect.ValueOf(obj)
	m := map[string]interface{}{}
	for i := 0; i < v.NumField(); i++ {
		m[t.Field(i).Tag.Get("foo")] = v.Field(i).Interface()
	}

	return m
}

func toMapUsingTagWithCachedType(t reflect.Type, obj interface{}) map[string]interface{} {
	v := reflect.ValueOf(obj)
	m := map[string]interface{}{}
	for i := 0; i < v.NumField(); i++ {
		m[t.Field(i).Tag.Get("foo")] = v.Field(i).Interface()
	}

	return m
}

// This function collects only the names
// that will be used from the type
// this should save several calls to `Field(i).Tag.Get("foo")`
// which might improve the performance by a lot.
func getTagNames(f interface{}) []string {
	t := reflect.TypeOf(f)
	resp := []string{}
	for i := 0; i < t.NumField(); i++ {
		resp = append(resp, t.Field(i).Tag.Get("foo"))
	}
	return resp
}

var tagNamesCache = map[reflect.Type][]string{}

func toMapUsingCachedTagNames(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	m := map[string]interface{}{}

	names, found := tagNamesCache[t]
	if !found {
		names = getTagNames(obj)
		tagNamesCache[t] = names
	}

	for i := 0; i < v.NumField(); i++ {
		m[names[i]] = v.Field(i).Interface()
	}

	return m
}

func toMapWithNoReflection(obj Foo) map[string]interface{} {
	return map[string]interface{}{
		"A": obj.A,
		"B": obj.B,
		"C": obj.C,
		"D": obj.D,
		"E": obj.E,
		"F": obj.F,
		"G": obj.G,
	}
}

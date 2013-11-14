package introspection

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type e interface{}

type m map[string]interface{}

func mult2(f e) e {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string) + f.(string) + f.(string)
	}
	return f
}

func Map(n []e, f func(e) e) []e {
	m := make([]e, len(n))
	for k, v := range n {
		m[k] = f(v)
	}
	return m
}

func ConvertToJson(m interface{}) []byte {
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func Structify(m m, s interface{}) {
	v := reflect.Indirect(reflect.ValueOf(s))

	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Name
		v.Field(i).Set(reflect.ValueOf(m[key]))
	}
}

func GetStructValues(m interface{}) (v []interface{}) {
	r := reflect.Indirect(reflect.ValueOf(m))
	v = make([]interface{}, 0)
	for i := 0; i < r.NumField(); i++ {
		v = append(v, r.Field(i).Interface())
	}
	return v[1:]
}

func ConvertToMap(s interface{}) m {
	typ := reflect.TypeOf(s)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	m := make(m)

	// Only structs are supported so return an empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
		return m
	}

	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Name
		val := v.Field(i).Interface()
		m[key] = val
	}
	return m
}

func MapValues(m[string]interface{} mm) {
	v := make([]string, 0, len(mm))

	for  _, value := range mm {
	   v = append(v, value)
	}
}

func Attributes(m interface{}) map[string]reflect.Type {
	typ := reflect.TypeOf(m)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// create an attribute data structure as a map of types keyed by a string.
	attrs := make(map[string]reflect.Type)
	// Only structs are supported so return an empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
		return attrs
	}

	// loop through the struct's fields and set the map
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			attrs[p.Name] = p.Type
		}
	}

	return attrs
}

func Types(m interface{}) []reflect.Type {
	typ := reflect.TypeOf(m)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	types := make([]reflect.Type, typ.NumField())
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
		return types
	}
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			types = append(types, p.Type)
		}
	}
	return types
}

func InterfaceName(i interface{}) string {
	v := reflect.TypeOf(i)
	return v.Name()
}

func FindMethod(recvType reflect.Type, funcVal *reflect.Value) *reflect.Method {
	// It is not possible to get the name of the method from the Func.
	// Instead, compare it to each method of the Controller.
	for i := 0; i < recvType.NumMethod(); i++ {
		method := recvType.Method(i)
		if method.Func == *funcVal {
			return &method
		}
	}
	return nil
}

package introspection

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type e interface{}

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

func Structify(m map[string]interface{}, s interface{}) {
	v := reflect.Indirect(reflect.ValueOf(s))

	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Name
		v.Field(i).Set(reflect.ValueOf(m[key]))
	}
}

func GetStructValues(m interface{}) (v []interface{}) {
	r := reflect.Indirect(reflect.ValueOf(m))
	v = make([]interface{}, 0)
	for i := 1; i < r.NumField(); i++ {
		v = append(v, r.Type().Field(i))
	}
	return v
}

func ConvertToMap(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Name
		val := v.Field(i).Interface()
		m[key] = val
	}
	return m
}

func InterfaceName(i interface{}) string {
	v := reflect.TypeOf(i)
	//for v.Kind() == reflect.Ptr {
	//	v = v.Elem()
	//}
	//return v.Kind().String()
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

/*func main() {
	f := map[string]interface{}{"Id": 1, "Title": "Wu Tang Clan", "Content": "biography about them",
		"UserId": 1, "Published": true, "Created": time.Now(), "Modified": time.Now()}
	var p Post
	Structify(f, &p)
	fmt.Println(p.Title)
	fmt.Println("------")
	vals := getStructValues(&p)
	for x := range vals {
		fmt.Println(vals[x])
	}
}test*/

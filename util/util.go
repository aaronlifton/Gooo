package util

import (
	"encoding/json"
	"fmt"
	"goweb/conversion"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// actually works somehow
func ToMap(m interface{}) map[string]interface{} {
	var f interface{}
	var b []byte
	var err error
	//v := reflect.TypeOf(m)

	b, err = json.Marshal(m)
	HandleErr(err)
	err = json.Unmarshal(b, &f)
	HandleErr(err)
	fmt.Println(conversion.StructName(m))
	return f.(map[string]interface{})
}

package util

import (
	"fmt"
	"os"
)

type m map[string]interface{}

func HandleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

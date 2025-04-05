package main

import (
	"fmt"
	"reflect"
)

type client struct {
	uuid       string
	operations []int
}

func main() {
	data1 := make([]client, 10)
	data2 := make([]client, 10)

	data1[1].operations = append(data1[1].operations, 10)
	isEqual := reflect.DeepEqual(data1, data2)
	fmt.Println("equal:", isEqual)
}

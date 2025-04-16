package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

func main() {
	channel := make(chan int, 1)
	channel2 := make(chan int, 2)

	vch := reflect.ValueOf(channel)
	vch2 := reflect.ValueOf(channel2)

	suceed := vch.TrySend(reflect.ValueOf(100))
	fmt.Println(suceed, vch.Len(), vch.Cap())

	suceed = vch2.TrySend(reflect.ValueOf(200))
	fmt.Println(suceed, vch2.Len(), vch2.Cap())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	vch3 := reflect.ValueOf(ctx.Done())
	defer cancel()
	for {
		branches := []reflect.SelectCase{
			{Dir: reflect.SelectRecv},
			{Dir: reflect.SelectRecv, Chan: vch},
			{Dir: reflect.SelectRecv, Chan: vch2},
			{Dir: reflect.SelectRecv, Chan: vch3},
		}
		index, value, ok := reflect.Select(branches)
		fmt.Println(index, value, ok)
		if !ok {
			break
		}
	}
}

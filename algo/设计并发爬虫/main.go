package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var battleNum int32

func main() {
	PaChong(10, 1)
}

func PaChong(f int, g int) {
	controlGNum := make(chan int, g)
	wg := &sync.WaitGroup{}
	tim := time.NewTicker(time.Second * 3)

	for i := 0; i < f; i++ {
		controlGNum <- i
		if atomic.LoadInt32(&battleNum) == 1 {
			<-tim.C
			atomic.StoreInt32(&battleNum, 0)
		}
		fmt.Println(i, "=====")

		wg.Add(1)
		ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
		go CheckIsDo(controlGNum, wg, ctx)
	}
	wg.Wait()
}

func CheckIsDo(controlGNum <-chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Println("超时退出")
		return
	default:
		atomic.AddInt32(&battleNum, 1)
		fmt.Println(atomic.LoadInt32(&battleNum))
		<-controlGNum
		Do()
		return
	}

}

func Do() {

}

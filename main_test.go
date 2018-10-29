package main

import (
	"testing"
	"fmt"
	"time"
)

func Test_Main(t *testing.T) {
	t.Parallel() //有并发
	start := time.Now()
	t.Log(start)
	for i := 0; i < 10; i++ {
		if i == 6 {
			t.Fatal() //直接退出
		}
		fmt.Println(i)
	}
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
	end := time.Now()
	t.Log(end.Sub(start))
	t.Error(1)

}

func Test_main1(t *testing.T) {
	defer fmt.Println("this is defer")
	a := 1
	if a == 1 {
		fmt.Println("this equal")
		return
	}

}

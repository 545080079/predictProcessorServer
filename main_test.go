/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: main_test
 * @Date: 2021/3/24 13:54
 */

package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestX(t *testing.T) {
}

func TestFF(t *testing.T) {
	var str1 string = "%saa"
	var str2 = "123"
	str3 := "321"
	fmt.Printf(str1, str2, str3)
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	n := 0

	startTime := time.Now()
	go func() {
		fmt.Printf("test: %v\n", n)
		n++
		time.Sleep(500 * time.Millisecond)
		wg.Done()
	}()
	go func() {
		fmt.Printf("test: %v\n", n)
		n++
		time.Sleep(500 * time.Millisecond)
		wg.Done()
	}()

	go func() {
		fmt.Printf("test: %v\n", n)
		n++
		time.Sleep(500 * time.Millisecond)
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("Done! cost time = %v, n = %v\n", time.Since(startTime).String(), n)
}
/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: cache
 * @Date: 2021/4/2 19:16
 */

package model

import "fmt"

type Cache struct {
	Name string	//节点名称
	RealResult string	//非预测模式下执行的结果
	RunResult string	//预测模式下执行的结果
}

var memoryCache []Cache

func init() {
	memoryCache = make([]Cache, 0)
}

func CacheAdd(c Cache) {
	memoryCache = append(memoryCache, c)
}

func CachePrint() {
	fmt.Println("-----------memoryCache START-----------")
	for i, v := range memoryCache {
		fmt.Printf("[%d] %v\n", i, v)
	}
	fmt.Println("-----------memoryCache END-----------")
}

func FindCacheLast() Cache {
	return memoryCache[len(memoryCache) - 1]
}
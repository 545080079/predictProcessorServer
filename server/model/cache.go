/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: cache
 * @Date: 2021/4/2 19:16
 */

package model


type ResultCache struct {
	Cache []Cache
}

type Cache struct {
	Name string	//节点名称
	RealResult string	//非预测模式下执行的结果
	RunResult string	//预测模式下执行的结果
}

func (resultCache ResultCache) Add(c Cache) {
	resultCache.Cache = append(resultCache.Cache, c)
}
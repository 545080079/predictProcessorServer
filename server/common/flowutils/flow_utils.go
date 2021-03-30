/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: flow_utils
 * @Date: 2021/3/12 10:20
 */

package flowutils

import (
	"github.com/gin-gonic/gin"
	"log"
	"predictProcessorServer/server/model"
	"time"
)

/*
检查入参是否符合语法
TODO
 */
func CheckInputParams(c *gin.Context) bool {
	if c.Query("input") == "" {
		return false
	}
	if c.Query("qrn") == "" {
		return false
	}
	return true
}


/*
	DAG顺序执行器
 */
func ProcessNormal(dag *model.DAG, userInput model.InputMap) error {
	//第1个节点， 入参由[用户]和[此节点的Parameters]提供
	//2~End节点，入参由[上一个节点计算结果]和[此节点的Parameters]提供
	lastInput := userInput

	//顺序执行节点
	p := dag.Next[0]
	for p != nil {

		if p.Resource == "" {
			continue
		}

		//[此节点的Parameters]参数追加到当前输入
		for k, v := range p.Parameters {
			//如果字段重名，目前默认新结果覆盖旧结果
			lastInput[k] = v
		}

		resp := Call(p.Resource, lastInput)
		log.Printf("[Process] exec node name[%v]: return [%v], cost time [%v]", p.Name, resp.result, resp.costTime)

		if p.Next == nil || p.Next[0] == nil {
			break
		}

		lastInput[p.Name] = resp.result
		p = p.Next[0]
	}
	return nil
}


/*
	计算函数入口
 */
type CallResp struct {
	result int64
	costTime time.Duration
}

func Call(resourceQRN string, input model.InputMap) *CallResp {
	resp := &CallResp{
		result:   -1,
	}

	var functionName string

	/*
	此处qrn:x:{num}为函数的resourceQRN写法
	States节点传入的Resource便对应该处
	 */
	switch resourceQRN {
	case "qrn:x:1":
		functionName = "sum"
	case "qrn:x:2":
		functionName = "find"
	default:
		functionName = "sum"
	}

	switch functionName {
	case "sum":
		return sum(input)
	case "find":
		return find(input)
	default:
		return resp
	}
}



/*
	计算函数实现（模拟耗时操作）
 */
func sum(input model.InputMap) *CallResp {

	arr := []float64{
		input["k1"].(float64),
		input["k2"].(float64),
		input["k3"].(float64),
	}
	var res float64 = 0
	startTime := time.Now()
	for _, v := range arr {
		time.Sleep(time.Millisecond * 100)
		res += v
	}
	return &CallResp{
		result:   int64(res),
		costTime: time.Since(startTime),
	}
}

func find(input model.InputMap) *CallResp {
	startTime := time.Now()
	arr := []float64{
		input["k1"].(float64),
		input["k2"].(float64),
		input["k3"].(float64),
	}
	target := input["target"].(float64)
	for i, v := range arr {
		time.Sleep(time.Millisecond * 100)
		if int64(v) == int64(target) {
			return &CallResp{
				result:   int64(i),
				costTime: time.Since(startTime),
			}
		}
	}

	return &CallResp{
		result:   -1,
		costTime: time.Since(startTime),
	}
}
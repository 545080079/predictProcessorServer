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
	"strconv"
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
	执行DAG
 */
func Process(dag *model.DAG, input *model.InputJSON) error {
	p := dag.Next[0]
	for p != nil {

		if p.Resource == "" {
			continue
		}

		//TODO 临时写法，后期补参数解析
		param1, _ := strconv.ParseInt(input.K1, 10, 64)
		param2, _ := strconv.ParseInt(input.K2, 10, 64)
		param3, _ := strconv.ParseInt(input.K3, 10, 64)
		arr := []int64{
			param1, param2, param3,
		}
		resp := Call(p.Resource, arr, param3)
		log.Printf("[Process] exec node name[%v]: return [%v], cost time [%v]", p.Name, resp.result, resp.costTime)

		if p.Next == nil || p.Next[0] == nil {
			break
		}
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

func Call(resourceQRN string, arr []int64, target int64) *CallResp {
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
		return sum(arr)
	case "find":
		return find(arr, target)
	default:
		return resp
	}
}



/*
	计算函数实现（模拟耗时操作）
 */
func sum(arr []int64) *CallResp {
	var res int64 = 0
	startTime := time.Now()
	for _, v := range arr {
		time.Sleep(time.Millisecond * 100)
		res += v
	}
	return &CallResp{
		result:   res,
		costTime: time.Since(startTime),
	}
}

func find(arr []int64, target int64) *CallResp {
	startTime := time.Now()
	for i, v := range arr {
		time.Sleep(time.Millisecond * 100)
		if v == target {
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
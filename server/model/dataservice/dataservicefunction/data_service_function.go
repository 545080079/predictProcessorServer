/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: function_data_service
 * @Date: 2021/4/2 14:39
 */

package dataservicefunction

import (
	"encoding/json"
	"log"
	"predictProcessorServer/server/common/parseutils"
	"predictProcessorServer/server/model"
	"strconv"
	"time"
)

/*
	计算函数入口
	Result: Marshal字符串
*/
type CallResp struct {
	Result string
	CostTime float64
}

func BuildResult(resultMap map[string]string) string {
	result, err := json.Marshal(resultMap)
	if err != nil {
		log.Fatal("[BuildResult] marshal failed, err:", err)
	}
	return string(result)
}


func Call(resourceQRN string, input model.InputMap) *CallResp {
	resp := &CallResp{
		Result:   "",
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
	case "qrn:x:3":
		functionName = "pass"
	default:
		functionName = "sum"
	}

	switch functionName {
	case "sum":
		return sum(input)
	case "find":
		return find(input)
	case "pass":
		return pass(input)
	default:
		return resp
	}
}



/*
	计算函数实现（模拟耗时操作）
*/

//求和
func sum(input model.InputMap) *CallResp {

	arr := []float64{
		parseutils.StrToFloat64(input["k1"]),
		parseutils.StrToFloat64(input["k2"]),
		parseutils.StrToFloat64(input["k3"]),
	}
	var res float64 = 0
	startTime := time.Now()
	for _, v := range arr {
		time.Sleep(time.Millisecond * 1000)
		res += v
	}
	return &CallResp{
		Result:   BuildResult(map[string]string {
			"Result": parseutils.Float64ToStr(res),
		}),
		CostTime: time.Since(startTime).Seconds(),
	}
}

//遍历查找
func find(input model.InputMap) *CallResp {
	startTime := time.Now()
	arr := []float64{
		parseutils.StrToFloat64(input["k1"]),
		parseutils.StrToFloat64(input["k2"]),
		parseutils.StrToFloat64(input["k3"]),
	}
	target := parseutils.StrToFloat64(input["target"])
	for i, v := range arr {
		time.Sleep(time.Millisecond * 100)
		if int64(v) == int64(target) {
			return &CallResp{
				Result:   BuildResult(map[string]string {
					"Result": strconv.FormatInt(int64(i), 10),
				}),
				CostTime: time.Since(startTime).Seconds(),
			}
		}
	}

	return &CallResp{
		Result:   BuildResult(map[string]string {
			"Result": strconv.FormatInt(-1, 10),
		}),
		CostTime: time.Since(startTime).Seconds(),
	}
}

//传递
func pass(input model.InputMap) *CallResp {
	startTime := time.Now()

	//将上一个节点输出的Result字段作为参数
	param := input["lastNode-Result"]
	time.Sleep(time.Millisecond * 1000)

	return &CallResp{
		Result:   BuildResult(map[string]string {
			"Result": param,
		}),
		CostTime: time.Since(startTime).Seconds(),
	}
}
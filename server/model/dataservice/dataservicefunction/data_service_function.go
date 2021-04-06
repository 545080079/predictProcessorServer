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

	log.Printf("[Call] resourceQRN: [%v]", resourceQRN)
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
		log.Printf("[Call] into sum.")
		return sum(input)
	case "find":
		log.Printf("[Call] into find.")
		return find(input)
	case "pass":
		log.Printf("[Call] into pass.")
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
	if input["lastNode-Result"] != "" {
		arr = append(arr, parseutils.StrToFloat64(input["lastNode-Result"]))
	}
	var res float64 = 0
	startTime := time.Now()
	for _, v := range arr {
		time.Sleep(time.Millisecond * 1000)
		res += v
	}

	log.Printf("[sum] call sum result:[%v]", res)
	return &CallResp{
		Result:   BuildResult(map[string]string {
			"lastNode-Result": parseutils.Float64ToStr(res),
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
			log.Printf("[find] call find result:index[%v]", i)
			return &CallResp{
				Result:   BuildResult(map[string]string {
					"lastNode-Result": strconv.FormatInt(int64(i), 10),
				}),
				CostTime: time.Since(startTime).Seconds(),
			}
		}
	}

	return &CallResp{
		Result:   BuildResult(map[string]string {
			"lastNode-Result": strconv.FormatInt(-1, 10),
		}),
		CostTime: time.Since(startTime).Seconds(),
	}
}

//传递
func pass(input model.InputMap) *CallResp {
	startTime := time.Now()

	//将上一个节点输出的Result字段作为参数
	lastNodeResult := input["lastNode-Result"]
	time.Sleep(time.Millisecond * 1000)

	log.Printf("[pass] call pass result:[%v]", lastNodeResult)

	return &CallResp{
		Result:   BuildResult(map[string]string {
			"lastNode-Result": lastNodeResult,
		}),
		CostTime: time.Since(startTime).Seconds(),
	}
}
/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: normal_function
 * @Date: 2021/4/7 17:33
 */

package dataservicefunction

import (
	"log"
	"predictProcessorServer/server/common/parseutils"
	"predictProcessorServer/server/model"
	"strconv"
	"time"
)

/*
	普通计算函数实现（模拟耗时操作）
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
		Result:   BuildResult(map[string]interface{} {
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
				Result:   BuildResult(map[string]interface{} {
					"lastNode-Result": strconv.FormatInt(int64(i), 10),
				}),
				CostTime: time.Since(startTime).Seconds(),
			}
		}
	}

	return &CallResp{
		Result:   BuildResult(map[string]interface{} {
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
		Result:   BuildResult(map[string]interface{} {
			"lastNode-Result": lastNodeResult,
		}),
		CostTime: time.Since(startTime).Seconds(),
	}
}

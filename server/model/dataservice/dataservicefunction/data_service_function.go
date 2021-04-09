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
	"predictProcessorServer/server/model"
)

type CallResp struct {
	Result string
	CostTime float64
}

func BuildResult(resultMap map[string]interface{}) string {
	result, err := json.Marshal(resultMap)
	if err != nil {
		log.Fatal("[BuildResult] marshal failed, err:", err)
	}
	return string(result)
}


/*
	计算函数入口
	Result: Marshal字符串
*/
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
	case "qrn:template:1":
		functionName = "checkToken"
	case "qrn:template:2":
		functionName = "checkServiceStatus"
	case "qrn:template:3":
		functionName = "checkFile"
	case "qrn:template:4":
		functionName = "transferMP3ToTXT"
	case "qrn:template:5":
		functionName = "writeResult"
	default:
		functionName = "sum"
	}

	log.Printf("[Call] into %v.", functionName)
	switch functionName {
	case "sum":
		return sum(input)
	case "find":
		return find(input)
	case "pass":
		return pass(input)
	case "checkToken":
		return checkToken(input)
	case "checkServiceStatus":
		return checkServiceStatus(input)
	case "checkFile":
		return checkFile(input)
	case "transferMP3ToTXT":
		return transferMP3ToTXT(input)
	case "writeResult":
		return writeResult(input)
	default:
		return resp
	}
}


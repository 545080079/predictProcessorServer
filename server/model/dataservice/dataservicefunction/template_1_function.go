/**
 * @Author: yutaoluo@tencent.com
 * @Description:音频转文字应用模板的函数实习
 * @File: template_1_function
 * @Date: 2021/4/7 17:34
 */

package dataservicefunction

import (
	"log"
	"os"
	"predictProcessorServer/server/model"
	"time"
)

/*
包含5个函数

命名qrn:template:1

1）检查权限

2）检查源文件合法（是否为MP3格式文件）

3）检查是否开通服务

4）执行具体服务（如：音频转文字 MP3 -> TXT)

5）写入容器
 */


// 检查权限
func checkToken(input model.InputMap) *CallResp {
	startTime := time.Now()

	return &CallResp{
		Result:   BuildResult(map[string]interface{}{
			"lastNode-Result": true,
		}),
		CostTime: time.Since(startTime).Seconds(),
	}
}
// 检查源文件合法（是否为MP3格式文件）
func checkFile(input model.InputMap) *CallResp {
	startTime := time.Now()

	return &CallResp{
		Result:   BuildResult(map[string]interface{}{
			"lastNode-Result": true,
		}),
		CostTime: time.Since(startTime).Seconds(),
	}
}
// 检查是否开通服务
func checkServiceStatus(input model.InputMap) *CallResp {
	startTime := time.Now()

	return &CallResp{
		Result:   BuildResult(map[string]interface{}{
			"lastNode-Result": true,
		}),
		CostTime: time.Since(startTime).Seconds(),
	}
}

// 执行具体服务（如：音频转文字 MP3 -> TXT)
func transferMP3ToTXT(input model.InputMap) *CallResp {
	startTime := time.Now()

	fileName := "./output-" + "123456"
	f, err := os.Create(fileName)
	if err != nil {
		log.Printf("[TransferMP3ToTXT] create file %v failed, err:%v", fileName, err)
	}
	_, err = f.WriteString("这是一段话")
	if err != nil {
		log.Printf("[TransferMP3ToTXT] transfer to txt failed, err:%v", err)
	}


	return &CallResp{
		Result:   BuildResult(map[string]interface{}{
			"lastNode-Result": "output-<randomPostfix>.txt",
		}),
		CostTime: time.Since(startTime).Seconds(),
	}
}

// 写入结果容器
func writeResult(input model.InputMap) *CallResp {
	startTime := time.Now()

	lastNodeResult := input["lastNode-Result"]
	//TODO 实现读取文件

	return &CallResp{
		Result:   BuildResult(map[string]interface{}{
			"lastNode-Result": lastNodeResult,
		}),
		CostTime: time.Since(startTime).Seconds(),
	}
}

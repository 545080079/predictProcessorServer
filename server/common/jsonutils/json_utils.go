/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: json_utils.go
 * @Date: 2021/3/11 11:32
 */

package jsonutils

import (
	"encoding/json"
	"log"
	"predictProcessorServer/server/model"
)

/*
	Definition String转GO JSON
 */
func ParseStringToJson(definition string) *model.DAGJson {
	var res *model.DAGJson
	log.Println("[ParseStringToJson] definition:", definition)
	if definition == "" {
		return res
	}
	err := json.Unmarshal([]byte(definition), &res)
	if err != nil {
		log.Fatal("[Unmarshal] ret error: ", err)
	}

	return res
}

func ParseStringToInputMap(input string) map[string]interface{} {
	var res *model.InputMap
	resMap := make(map[string]interface{}, 0)
	log.Println("[ParseStringToInputJSON] input:", input)
	if input == "" {
		return resMap
	}
	err := json.Unmarshal([]byte(input), &res)
	if err != nil {
		log.Fatal("[Unmarshal] ret error: ", err)
	}

	for k, v := range *res {
		resMap[k] = v
	}
	return resMap
}
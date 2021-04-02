/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: parse_utils.go
 * @Date: 2021/3/11 11:32
 */

package parseutils

import (
	"encoding/json"
	"log"
	"predictProcessorServer/server/model"
	"strconv"
)

func StrToFloat64(s string) float64 {
	float, _ := strconv.ParseFloat(s, 64)
	return float
}

func Float64ToStr(f float64) string {
	str := strconv.FormatFloat(f, 'f', -1, 64)
	return str
}

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

func ParseStringToInputMap(input string) map[string]string {
	var res *model.InputMap
	resMap := make(map[string]string, 0)
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
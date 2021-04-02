/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: page_service_dag
 * @Date: 2021/3/29 10:41
 */

package pageservicedag

import (
	"log"
	"predictProcessorServer/conf"
	"predictProcessorServer/server/common/parseutils"
	"predictProcessorServer/server/common/stateutils"
	"predictProcessorServer/server/model"
)

func ParseDefinition(definition string) (*model.DAG, error) {
	log.Println("[test]", definition, "[end]")
	dagJSON := parseutils.ParseStringToJson(definition)

	//读取状态机列表
	nameArray := make([]string, 0)
	for k, v := range dagJSON.States {
		nameArray = append(nameArray, k)
		log.Printf("[i]k: %s, v: %v", k, v)
	}
	log.Printf("[HandlerParseDefinition]状态机节点：%v", nameArray)

	//检查状态机列表
	ret := stateutils.CheckStateInfo(nameArray)
	if !ret {
		log.Fatalln("[HandlerParseDefinition]状态机列表名称重复")
	}

	//图dummy节点
	dummyDAGNode := &model.DAG {
		Next:        nil,
		Name:		 conf.DummyNodeName,
		Resource: 	 "",
		Type:        "dummy",
		Comment:     "",
		IsEnd:       false,
	}
	tail := dummyDAGNode

	for _, v := range nameArray {
		DAGNode := &model.DAG {
			Name:		 v,
			Resource: 	 dagJSON.States[v].Resource,
			Type:        dagJSON.States[v].Type,
			Comment:     dagJSON.States[v].Comment,
			IsEnd:       dagJSON.States[v].End,
			Parameters:  dagJSON.States[v].Parameters,
		}
		next := make([]*model.DAG, 1)
		next[0] = DAGNode
		tail.Next = next
		tail = DAGNode

		log.Printf("%v", DAGNode)
		//DAGNode.RLock()
		//defer DAGNode.RUnlock()
	}

	dummyDAGNode.Print()
	return dummyDAGNode, nil
}

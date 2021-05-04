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
	"predictProcessorServer/server/model"
)

/*
	根据Definition构造DAGraph
	return: *model.DAG dummy node address
 */
func ParseDefinition(definition string) (*model.DAG, error) {
	log.Println("[test]", definition, "[end]")
	dagJSON := parseutils.ParseStringToJson(definition)

	//存储链表映射关系, Name->next address
	nodeAddressLinks := make(map[string]*model.DAG, 0)
	//存储链表映射关系, Name->next.Name
	nodeLinks := make(map[string]string, 0)
	nodeLinks[conf.DummyNodeName] = dagJSON.StartAt

	//存储所有节点Name
	nameArray := make([]string, 0)
	//读取状态机列表
	for k, v := range dagJSON.States {
		nodeLinks[k] = v.Next
		nameArray = append(nameArray, k)
		log.Printf("[i]k: %s, v: %v", k, v)
	}
	log.Printf("[HandlerParseDefinition]状态机节点：%v", nameArray)

	//检查状态机列表
	ret := parseutils.CheckStateInfo(nameArray)
	if !ret {
		log.Fatalln("[HandlerParseDefinition]状态机列表名称重复")
	}

	//图dummy节点
	dummyDAGNode := &model.DAG {
		Next:        make([]*model.DAG, 1),
		Name:		 conf.DummyNodeName,
		Resource: 	 "",
		Type:        "dummy",
		Comment:     "Head node",
		IsEnd:       false,
	}
	dummyDAGNode.Next[0] = nil

	for _, v := range nameArray {
		node := &model.DAG {
			Name:		 v,
			Resource: 	 dagJSON.States[v].Resource,
			Type:        dagJSON.States[v].Type,
			Comment:     dagJSON.States[v].Comment,
			IsEnd:       dagJSON.States[v].End,
			Parameters:  dagJSON.States[v].Parameters,
			Next: 		 make([]*model.DAG, 1),
		}

		node.Next[0] = dummyDAGNode.Next[0]
		dummyDAGNode.Next[0] = node

		nodeAddressLinks[node.Name] = node
		log.Printf("[ParseDefinition] build node:[%v]", node)
		//DAGNode.RLock()
		//defer DAGNode.RUnlock()
	}

	//根据Next节点调整链表顺序
	p := dummyDAGNode
	var next *model.DAG
	for p != nil {
		if p.Next != nil && len(p.Next) != 0 {
			next = p.Next[0]
			p.Next[0] = nodeAddressLinks[nodeLinks[p.Name]]
		} else {
			p.Next[0] = nodeAddressLinks[nodeLinks[p.Name]]
			break
		}
		p = next
	}

	dummyDAGNode.Print()
	return dummyDAGNode, nil
}

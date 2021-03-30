/**
 * @Author: yutaoluo@tencent.com
 * @Description: DAG工作流相关定义
 * @File: DAG
 * @Date: 2021/3/11 16:32
 */

package model

import (
	"log"
	"sync"
)

/*
	解析后的工作流Graph Node链表
 */
type DAG struct {
	Next        []*DAG
	Name		string
	Type        string
	Resource	string
	Parameters  map[string]string
	Comment     string
	IsEnd       bool
	sync.RWMutex
}

/*
	TCSL语言描述格式
*/
type DAGJson struct {
	StartAt string `json:"startAt"`
	Resource string `json:"resource"`
	States map[string]State `json:"states"`
}

/*
	状态机的状态节点
*/
type State struct {
	Type string `json:"type"`
	Comment string `json:"comment"`
	Resource string `json:"resource"`
	Parameters  map[string]string `json:"parameters"`
	Next string `json:"next"`
	End bool `json:"end"`
}

//TODO 暂时先限定三个入参
type InputJSON struct {
	K1 string `json:"k1"`
	K2 string `json:"k2"`
	K3 string `json:"k3"`
}

/*
	打印图
 */
func (d *DAG) Print() {

	log.Println("-------------打印DAG-------------")

	ptr := d
	i := 0
	for ptr != nil {

		log.Printf("[%d]type:%v, comment:%v, resourceQRN:%v, isEnd:%v", i, ptr.Type, ptr.Comment, ptr.Resource, ptr.IsEnd)

		//访问节点的next
		if len(ptr.Next) == 0 {
			break
		}
		ptr = ptr.Next[0]
		i++
	}
	log.Println("-------------结束打印-------------")

}
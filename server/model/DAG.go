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
	工作流Graph
 */
type DAG struct {
	Next        []*DAG
	Name		string
	ResourceQRN string
	Type        string
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
	Next string `json:"next"`
	End bool `json:"end"`
}

/*
	打印图
 */
func (d *DAG) Print() {

	log.Println("-------------打印DAG-------------")

	ptr := d
	i := 0
	for ptr != nil {

		log.Printf("[%d]type:%v, comment:%v, resourceQRN:%v, isEnd:%v", i, ptr.Type, ptr.Comment, ptr.ResourceQRN, ptr.IsEnd)

		//访问节点的next
		if len(ptr.Next) == 0 {
			break
		}
		ptr = ptr.Next[0]
		i++
	}
	log.Println("-------------结束打印-------------")

}
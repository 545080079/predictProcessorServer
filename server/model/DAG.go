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

/*
状态机传递参数类型
输入输出都是InputMap
 */
type InputMap map[string]string

/*
	获取Graph节点数量
 */
func (d *DAG) LenExceptDummy() int {

	ptr := d
	cnt := 0
	for ptr != nil {

		//访问节点的next
		if len(ptr.Next) == 0 {
			break
		}
		ptr = ptr.Next[0]
		cnt++
	}

	//dummy节点不计入
	return cnt - 1
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
/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: t_flow_service
 * @Date: 2021/3/12 10:51
 */

package daoimpl

import (
	"predictProcessorServer/server/model"
	"strconv"
)

var DAGSet map[string]*model.DAG
var DAGIndex int

func init() {
	DAGSet = make(map[string]*model.DAG, 0)
	DAGIndex = 0
}

func Push(dag *model.DAG) string {
	DAGIndex++
	key := "qrn:" + strconv.Itoa(DAGIndex)
	DAGSet[key] = dag
	return key
}

func FindDAG(resourceQRN string) *model.DAG {

	return DAGSet[resourceQRN]
}
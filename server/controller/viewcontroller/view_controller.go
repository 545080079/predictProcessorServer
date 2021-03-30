/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: view_service
 * @Date: 2021/3/29 10:55
 */

package viewcontroller

import (
	"github.com/gin-gonic/gin"
	"log"
	"predictProcessorServer/server/common/viewutils"
	"predictProcessorServer/server/dao/daoimpl"
	"predictProcessorServer/server/model/pageservice/pageservicedag"
)

func HandlerGenerateGraph(c *gin.Context) {

	definition := c.Query("definition")
	if definition == "" {
		c.JSON(200, "param definition is nil.")
		return
	}

	dummyDAGNode, _ := pageservicedag.ParseDefinition(definition)
	//存至数据层
	key := daoimpl.Push(dummyDAGNode)
	log.Println("push dag key=", key)
	//渲染图
	viewutils.GenerateGraphByDAG(dummyDAGNode, "blue")
	log.Println("Generate Graph success.")
}

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
	"predictProcessorServer/server/model/pageservice/pageservicedag"
)

func HandlerGenerateGraph(c *gin.Context) {

	definition := c.Query("definition")
	if definition == "" {
		c.JSON(200, "param definition is nil.")
		return
	}

	dummyDAGNode, _ := pageservicedag.ParseDefinition(definition)
	//渲染图到ResponseWriter
	viewutils.GenerateGraphByLinkedList(dummyDAGNode, c.Writer)
	log.Println("Generate Graph success.")
}

/**
 * @Author: yutaoluo@tencent.com
 * @Description: Definition JSON解析器相关服务
 * @File: parser_service
 * @Date: 2021/3/11 11:29
 */

package parsercontroller

import (
	"github.com/gin-gonic/gin"
	"log"
	"predictProcessorServer/server/dao/daoimpl"
	"predictProcessorServer/server/model/pageservice/pageservicedag"
)

/*
	解析Definition JSON
	Return DAG graph
 */
func HandlerParseDefinition(c *gin.Context) {

	//解析JSON
	definition := c.Query("definition")
	if definition == "" {
		c.JSON(200, "param definition is nil.")
		return
	}

	dummyDAGNode, _ := pageservicedag.ParseDefinition(definition)

	//存至数据层
	key := daoimpl.Push(dummyDAGNode)
	log.Println("push dag key=", key)

	c.JSON(200, dummyDAGNode)
}
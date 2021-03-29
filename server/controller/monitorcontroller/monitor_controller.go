/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: monitor_server
 * @Date: 2021/3/10 17:36
 */

package monitorcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerTest(c *gin.Context) {
	c.String(http.StatusOK, "Hello.")
}

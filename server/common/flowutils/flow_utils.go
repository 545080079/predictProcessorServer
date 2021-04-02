/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: flow_utils
 * @Date: 2021/3/12 10:20
 */

package flowutils

import (
	"github.com/gin-gonic/gin"
)

/*
检查入参是否符合语法
TODO
 */
func CheckInputParams(c *gin.Context) bool {
	if c.Query("input") == "" {
		return false
	}
	if c.Query("qrn") == "" {
		return false
	}
	return true
}


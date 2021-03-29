/**
 * @Author: yutaoluo@tencent.com
 * @Description:	抽象接口预定义
 * @File: interface
 * @Date: 2021/3/12 10:06
 */

package model

/*
	格式化打印到log/console
 */
type CanPrint interface {
	Print()
}

/*
	toString格式化输出
 */
type String interface {
	String()
}
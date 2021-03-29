/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: base
 * @Date: 2021/3/12 10:52
 */
package daoimpl

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"predictProcessorServer/conf"
)

/*
	xorm base类

	cmd命令：
		xorm reverse mysql root:123456@tcp({ip:port})/ai_flow?charset=utf8 {输出的go模板文件目录}
 */


func Init() {
	var engine *xorm.Engine
	engine, err := xorm.NewEngine("mysql", conf.DataSourceName)
	if err != nil {
		log.Fatal("[xorm] engine init err:", err)
	}
	if err := engine.Ping(); err != nil {
		log.Fatal("[xorm] ping mysql failed, err: ", err)
	}
	defer engine.Close()
	log.Print("[xorm] link to mysql success.")
	
}
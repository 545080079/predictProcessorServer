/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: conf.go
 * @Date: 2021/3/10 17:25
 */

package conf

const PORT = ":8083"
const PORT_GRPC = ":8084"
const DataSourceName = "root:123456@tcp({ip:port})/ai_flow?charset=utf8"

//DAG定义
const DummyNodeName = "dummy"
const CreateModeInit = 1
const CreateModeModify = 2
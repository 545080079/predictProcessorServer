/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: flow_service
 * @Date: 2021/3/12 10:17
 */

package flowcontroller

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	pb "predictProcessorServer/pb3/asw/AswSvr"
	"predictProcessorServer/server/common/parseutils"
	"predictProcessorServer/server/dao/daoimpl"
	"predictProcessorServer/server/model/pageservice/pageserviceflow"
)

func HandlerFlowProcess(c *gin.Context) {
	if !parseutils.CheckInputParams(c) {
		log.Fatal("[HandlerFlowProcess] Param is nil.")
	}

	//此处resourceQRN是flow的QRN
	//为了方便测试，默认前端传进qrn:1,每一次测试后重启后台服务即可
	input, resourceQRN := c.Query("input"), c.Query("qrn")
	log.Printf("[HandlerFlowProcess] input: %v, resourcesQRN: %v", input, resourceQRN)

	//获得resourceQRN唯一对应的DAG flow
	dag := daoimpl.FindDAG(resourceQRN)
	log.Println("dag: ", dag)

	userInputMap := parseutils.ParseStringToInputMap(input)

	//传入工作流入参，调用普通顺序执行器执行DAGraph flow
	//错误暂时忽略
	//TODO
	output, _ := pageserviceflow.ProcessNormal(dag, userInputMap)
	ret := map[string]interface{} {
		"Result":   output["lastNode-Result"],
		"CostTime": output["costTime"],
	}
	c.JSON(200, ret)
}

func HandlerFlowProcessParallel(c *gin.Context) {
	if !parseutils.CheckInputParams(c) {
		log.Fatal("[HandlerFlowProcess] Param is nil.")
	}

	//此处resourceQRN是flow的QRN
	//为了方便测试，默认前端传进qrn:1,每一次测试后重启后台服务即可
	input, resourceQRN := c.Query("input"), c.Query("qrn")
	log.Printf("[HandlerFlowProcess] input: %v, resourcesQRN: %v", input, resourceQRN)

	//获得resourceQRN唯一对应的DAG flow
	dag := daoimpl.FindDAG(resourceQRN)
	log.Println("dag: ", dag)

	userInputMap := parseutils.ParseStringToInputMap(input)

	//传入工作流入参，调用普通顺序执行器执行DAGraph flow
	//错误暂时忽略TODO
	output, _ := pageserviceflow.ProcessParallel(dag, userInputMap)
	ret := map[string]interface{} {
		"Result":   output["lastNode-Result"],
		"CostTime": output["costTime"],
	}
	c.JSON(200, ret)
}



//暂时不用
type FlowManager struct {
	pb.UnimplementedFlowManagerServer
}
func (fm *FlowManager) CreateFlow(ctx context.Context, req *pb.CreateFlowReq) (*pb.CreateFlowRsp, error) {
	log.Println("-----------create flow-----------")
	ret := &pb.CreateFlowRsp {
		Flag: true,
	}
	return ret, nil
}

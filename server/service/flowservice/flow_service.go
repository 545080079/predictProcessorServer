/**
 * @Author: yutaoluo@tencent.com
 * @Description:GRPC service类
 * @File: flow_service
 * @Date: 2021/3/12 10:17
 */

package flowservice

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	pb "predictProcessorServer/pb3/asw/AswSvr"
)

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

func HandlerCreateFlowService(c *gin.Context) {
}


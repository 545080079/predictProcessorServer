package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"predictProcessorServer/conf"
	pb "predictProcessorServer/pb3/asw/AswSvr"
	"predictProcessorServer/server/controller/flowcontroller"
	"predictProcessorServer/server/controller/monitorcontroller"
	"predictProcessorServer/server/controller/parsercontroller"
	"predictProcessorServer/server/controller/viewcontroller"
)

func init() {
	log.Println("[PredictExecuteServer] Init server.")
}



func main() {

	//GRPC server init
	listenGRPC, err := net.Listen("tcp", conf.PORT_GRPC)
	if err != nil {
		log.Fatal("[GRPC] listen err: ", err)
	}
	serverGRPC := grpc.NewServer()
	pb.RegisterFlowManagerServer(serverGRPC, &flowcontroller.FlowManager{})

	go serverGRPC.Serve(listenGRPC)

	//GRPC Client访问测试
	conn, err := grpc.Dial("localhost" + conf.PORT_GRPC, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("[client test] can't connect. err: ", err)
	}
	defer conn.Close()
	c := pb.NewFlowManagerClient(conn)
	r, _ := c.CreateFlow(context.Background(), &pb.CreateFlowReq{MachineQRN: "qrn:1"})
	log.Print("[client test] return: Flag=", r.Flag)

	//Gin server init

	log.Printf("[PredictExecuteServerMain] Try to run server in :%v", conf.PORT)

	route := gin.Default()
	v1 := route.Group("/v1")
	{
		monitor := v1.Group("monitor")
		{
			monitor.GET("/test", monitorcontroller.HandlerTest)
		}
		parser := v1.Group("parser")
		{
			parser.GET("/parse", parsercontroller.HandlerParseDefinition)
		}
		view := v1.Group("view")
		{
			view.GET("/gen", viewcontroller.HandlerGenerateGraph)
		}
	}
	err = route.Run(conf.PORT)
	if err != nil {
		log.Fatal("[PredictExecuteServerMain] Run server error: ", err)
	}
}

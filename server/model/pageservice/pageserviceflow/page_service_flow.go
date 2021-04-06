/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: page_service_flow
 * @Date: 2021/4/2 16:39
 */

package pageserviceflow

import (
	"encoding/json"
	"log"
	"predictProcessorServer/conf"
	"predictProcessorServer/server/common/parseutils"
	"predictProcessorServer/server/model"
	"predictProcessorServer/server/model/dataservice/dataservicefunction"
	"predictProcessorServer/server/model/pageservice/pageserviceview"
)

/*
	调取外部预测模型接口，获取预测值
 */
func DescribeDAGPredict(dag *model.DAG, userInput model.InputMap) map[string]string {
	res := make(map[string]string)
	p := dag.Next[0]
	for p != nil {
		//预测值从模型取得，这里暂时写死
		res[p.Name] = "99"
		if p.Next == nil || p.Next[0] == nil {
			break
		}
		p = p.Next[0]
	}
	return res
}

/*
	DAG并行执行器
 */
func ProcessParallel(dag *model.DAG, userInput model.InputMap) (model.InputMap, error) {
	log.Printf("[ProcessParallel] userInput:[%v]", userInput)

	//初始化所有节点可视化
	pageserviceview.GenerateGraphByDAG(dag, conf.CreateModeInit, float32(0.1), "red")
	pageserviceview.ModifyNodeColor(conf.DummyNodeName, "blue")

	predictMap := DescribeDAGPredict(dag, userInput)
	log.Printf("[ProcessParallel] build predictMap is [%v]", predictMap)

	ch := make(chan string, dag.LenExceptDummy())
	costTimeCh := make(chan float64, dag.LenExceptDummy())
	defer close(ch)
	defer close(costTimeCh)

	log.Printf("[ProcessParallel] channel len is [%v]", dag.LenExceptDummy())
	p := dag.Next[0]
	lastName := p.Name
	for p != nil {
		if p.Resource == "" {
			continue
		}
		//为了方便演示，输出仅限制为一个数字
		userInput["lastNode-Result"] = predictMap[lastName]
		//此时该节点具备执行条件

		log.Printf("[ProcessParallel] p.Name:[%v], p.Resource:[%v]", p.Name, p.Resource)
		resource := p.Resource
		go func() {
			resp := dataservicefunction.Call(resource, userInput)
			log.Printf("[ProcessParallel Go routine] resp: [%v]", resp)
			respMap := make(map[string]string)
			err := json.Unmarshal([]byte(resp.Result), &respMap)
			if err != nil {
				log.Print("[Call] unmarshal err:", err)
			}

			log.Printf("[ProcessParallel Go routine] respMap len is: [%v]", len(respMap))
			for i, res := range respMap {
				log.Printf("[ProcessParallel Go routine] resMap[%v] is: [%v]", i, res)
				ch <- res
			}
			costTimeCh <- resp.CostTime
			log.Printf("[ProcessParallel Go routine] respMap.Result transform to ch.[SUCCESS]")
		}()

		if p.Next == nil || p.Next[0] == nil {
			break
		}
		p = p.Next[0]
	}

	//TODO
	//获取预测模式下执行的结果,写入缓存
	theMostLongestCostTime := 0.0
	for i := 0; i < dag.LenExceptDummy(); i++ {

		cache := model.Cache {
			Name:       "default",
			RealResult: "unknown",
			RunResult:  <- ch,
		}
		costTime := <-costTimeCh
		if costTime > theMostLongestCostTime {
			theMostLongestCostTime = costTime
		}
		model.CacheAdd(cache)
		model.CachePrint()
	}

	//TODO 获取结果,与缓存器对比,标记开始产生错误的节点,从该节点开始退化为顺序执行,若无,执行结束


	//标记全部执行完毕，执行时间为最久的一个协程允许函数所消耗的时间
	pageserviceview.GenerateGraphByDAG(dag, conf.CreateModeInit, float32(theMostLongestCostTime), "green")
	userInput["costTime"] =  parseutils.Float64ToStr(theMostLongestCostTime)
	userInput["lastNode-Result"] = model.FindCacheLast().RunResult
	userInput["lastNode-Name"] = model.FindCacheLast().Name
	log.Println("[Debug][ProcessNormal] finished. return:", userInput)
	return userInput, nil
}

/*
	DAG顺序执行器
*/
func ProcessNormal(dag *model.DAG, userInput model.InputMap) (model.InputMap, error) {
	//第1个节点， 入参由[用户]和[此节点的Parameters]提供
	//2~End节点，入参由[上一个节点计算结果]和[此节点的Parameters]提供
	lastInput := userInput

	//初始化所有节点可视化
	pageserviceview.GenerateGraphByDAG(dag, conf.CreateModeInit, float32(0.1), "red")
	pageserviceview.ModifyNodeColor(conf.DummyNodeName, "blue")

	//顺序执行节点
	p := dag.Next[0]
	for p != nil {
		if p.Resource == "" {
			continue
		}

		//[此节点的Parameters]参数追加到当前输入
		for k, v := range p.Parameters {
			//如果字段重名，目前默认新结果覆盖旧结果
			lastInput[k] = v
		}

		resp := dataservicefunction.Call(p.Resource, lastInput)
		log.Printf("[Process] exec node name[%v] finished: return [%v], cost time [%v]", p.Name, resp.Result, resp.CostTime)

		//执行完一个节点，就标记为绿色，且展示执行的时间
		pageserviceview.ModifyNodeColor(p.Name, "green")
		pageserviceview.ModifyNodeValue(p.Name, float32(resp.CostTime))
		pageserviceview.GenerateGraphByDAG(dag,conf.CreateModeModify, 0, "")

		//此节点的输出提供给下一个节点
		respMap := make(map[string]string)
		err := json.Unmarshal([]byte(resp.Result), &respMap)
		if err != nil {
			log.Print("unmarshal err:", err)
		}
		for k, v := range respMap {
			lastInput["lastNode-" + k] = v
		}
		//Duration
		lastInput["costTime"] =  parseutils.Float64ToStr(parseutils.StrToFloat64(lastInput["costTime"]) + resp.CostTime)
		log.Println("[Debug][ProcessNormal] lastInput=", lastInput)

		if p.Next == nil || p.Next[0] == nil {
			break
		}
		p = p.Next[0]
	}
	return lastInput, nil
}


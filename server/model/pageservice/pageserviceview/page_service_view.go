/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: view_utils
 * @Date: 2021/3/29 10:45
 */

package pageserviceview

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"log"
	"os"
	"predictProcessorServer/conf"
	"predictProcessorServer/server/model"
)

var page *components.Page
var graph *charts.Graph

//节点颜色、值
var colors map[string]string
var values map[string]float32

func init() {
	colors = make(map[string]string, 0)
	values = make(map[string]float32, 0)
}

//更改指定节点的颜色
func ModifyNodeColor(nodeName, color string) {
	colors[nodeName] = color
}
//指定节点Value（运行耗时）
func ModifyNodeValue(nodeName string, value float32) {
	values[nodeName] = value
}

//指定整个DAG的节点颜色
func modifyAllNodesColor(dag *model.DAG, color string) {
	if dag == nil {
		return
	}
	for p := dag; p != nil; p = p.Next[0] {
		ModifyNodeColor(p.Name, color)
		if  p.Next == nil || p.Next[0] == nil {
			break
		}
	}
	fmt.Println("[DEBUG]colors:", colors)
}

//指定整个DAG的节点Value
func modifyAllNodesValue(dag *model.DAG, value float32) {
	if dag == nil {
		return
	}
	for p := dag; p != nil; p = p.Next[0] {
		ModifyNodeValue(p.Name, value)
		if  p.Next == nil || p.Next[0] == nil {
			break
		}
	}
	fmt.Println("[DEBUG]values:", values)
}

/*
	生成DAG渲染图的入口方法
	[入参]
		color：节点颜色
		createMode: 创建模式， CreateModeInit: 初始化创建，需要指定初始值和节点颜色, CreateModeModify：修改个别节点的属性，无需传入initValue, initColor参数
 */
func GenerateGraphByDAG(dummy *model.DAG, createMode int, initValue float32, initColor string) {
	page = components.NewPage()
	graph = charts.NewGraph()
	if createMode == conf.CreateModeInit {
		modifyAllNodesValue(dummy, initValue)
		modifyAllNodesColor(dummy, initColor)
	}
	graph.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:         "DAGraph",
	}))

	graph.AddSeries("graph", genNodes(dummy), genLinks(dummy),
		charts.WithGraphChartOpts(
			opts.GraphChart{
				Layout:             "none",
				Force:              nil,
				Roam:               false,
			}),
		charts.WithLabelOpts(opts.Label{
			Show:      true,
			Color:     "",
			Position:  "top",
			Formatter: "{b} | 耗时:{c}s",
		}))

	//进行渲染
	render()
}

//生成echarts图
func render() {
	page.AddCharts(graph)
	file, err := os.Create("./vue/graph.html")//TODO: 由于这个相对路径，只能在根目录下执行./bin/文件
	err = page.Render(file)
	if err != nil {
		log.Fatal("[GenerateGraphByDAG] page render file failed: ", err)
	}
}

/*
根据DAG 生成echarts节点
 */
func genNodes(dummy *model.DAG) []opts.GraphNode {
	nodes := make([]opts.GraphNode, 0)
	var offset float32 = 100

	dagNodes := dummy.TraverseStartAtDummy()
	for _, n := range dagNodes {
		node := opts.GraphNode {
			Name:       n.Name,
			X:          50,
			Y:          offset,
			Value:      values[n.Name],
			Fixed:      false,
			Symbol:     "roundRect",
			SymbolSize: 20,
			ItemStyle:  &opts.ItemStyle{
				Color: colors[n.Name],
			},
		}
		nodes = append(nodes, node)
		offset += 50
	}

	return nodes
}

/*
根据DAG 生成echarts连线
*/
func genLinks(dummy *model.DAG) []opts.GraphLink {
	links := make([]opts.GraphLink, 0)

	dagNodes := dummy.TraverseStartAtDummy()
	for _, n := range dagNodes {
		if  n.Next == nil || len(n.Next) == 0 || n.Next[0] == nil {
			continue
		}
		link := opts.GraphLink{
			Source: n.Name,
			Target: n.Next[0].Name,
		}
		links = append(links, link)
	}

	return links
}
/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: view_utils
 * @Date: 2021/3/29 10:45
 */

package viewutils

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
	file, err := os.Create("./vue/graph.html")
	err = page.Render(file)
	if err != nil {
		log.Fatal("[GenerateGraphByLinkedList] page render file failed: ", err)
	}
}

/*
根据DAG 生成echarts节点
 */
func genNodes(dummy *model.DAG) []opts.GraphNode {

	nodes := make([]opts.GraphNode, 0)

	p := dummy
	var offset float32 = 100
	for p != nil {
		node := opts.GraphNode {
			Name:       p.Name,
			X:          50,
			Y:          offset,
			Value:      values[p.Name],
			Fixed:      false,
			Symbol:     "roundRect",
			SymbolSize: 20,
			ItemStyle:  &opts.ItemStyle{
				Color:        colors[p.Name],
			},
		}
		nodes = append(nodes, node)
		offset += 50

		if  p.Next == nil || len(p.Next) == 0 {
			break
		}
		p = p.Next[0]
	}

	return nodes
}

/*
根据DAG 生成echarts连线
*/
func genLinks(dummy *model.DAG) []opts.GraphLink {

	links := make([]opts.GraphLink, 0)

	p := dummy
	for p != nil {
		if  p.Next == nil || len(p.Next) == 0 {
			break
		}
		link := opts.GraphLink{
			Source: p.Name,
			Target: p.Next[0].Name,
		}
		links = append(links, link)

		p = p.Next[0]
	}
	return links
}
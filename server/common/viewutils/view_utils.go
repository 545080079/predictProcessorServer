/**
 * @Author: yutaoluo@tencent.com
 * @Description:
 * @File: view_utils
 * @Date: 2021/3/29 10:45
 */

package viewutils

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"log"
	"os"
	"predictProcessorServer/server/model"
)

var graph *charts.Graph

func init() {
	graph = charts.NewGraph()
}

/*
	生成DAG渲染图的入口方法
	[入参]
		color：节点颜色
 */
func GenerateGraphByDAG(dummy *model.DAG, color string) {

	graph.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:         "DAG Render Demo",
	}))

	graph.AddSeries("graph", genNodes(dummy, color), genLinks(dummy),
		charts.WithGraphChartOpts(
			opts.GraphChart{
				Layout:             "none",
				Force:              nil,
				Roam:               true,
			}),
		charts.WithLabelOpts(opts.Label{
			Show:      true,
			Color:     "",
			Position:  "top",
		}))

	render()
}

func render() {
	page := components.NewPage()
	page.AddCharts(graph)
	file, err := os.Create("./vue/graph.html")
	err = page.Render(file)
	if err != nil {
		log.Fatal("[GenerateGraphByLinkedList] page render file failed: ", err)
	}
}
func genNodes(dummy *model.DAG, color string) []opts.GraphNode {

	nodes := make([]opts.GraphNode, 0)

	p := dummy
	var offset float32 = 100
	for p != nil {
		node := opts.GraphNode {
			Name:       p.Name,
			X:          50,
			Y:          offset,
			Value:      0,
			Fixed:      false,
			Symbol:     "roundRect",
			SymbolSize: 20,
			ItemStyle:  &opts.ItemStyle{
				Color:        color,
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
func genFakeLinks(dummy *model.DAG) []opts.GraphLink {

	var links = []opts.GraphLink {
		{
			Source: 0,
			Target: 1,
		},
	}
	return links
}
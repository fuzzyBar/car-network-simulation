package graph

import (
	"math"
)

type Graph struct {
	Nodes       [][]int     // 节点坐标
	Matrix      [][]float64 // 邻接矩阵
	EdgeNodes   []int       //边缘节点
	CenterNodes []int       //中心节点, 即非边缘节点
}

func NewGraph() *Graph {
	// 初始化邻接矩阵
	matrix := [][]float64{
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, // 0
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, // 1
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, // 2
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}, // 3
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}, // 4
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, // 5
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, // 6
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, // 7
		{1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1}, // 8
		{0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0}, // 9
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1}, // 10
		{0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 1, 0}, // 11
	}

	return &Graph{
		Nodes: [][]int{
			{400, 1200},
			{800, 1200},
			{1200, 800},
			{1200, 400},
			{800, 0},
			{400, 0},
			{0, 400},
			{0, 800},
			{400, 800},
			{800, 800},
			{800, 400},
			{400, 400},
		},
		Matrix: matrix,
	}
}

func (g *Graph) InitializeGraph() *Graph {
	for i, node := range g.Matrix {
		flag := 0
		for _, value := range node {
			if value > 0 {
				flag++
			}
		}
		if flag == 1 {
			g.EdgeNodes = append(g.EdgeNodes, i)
		} else if flag > 1 {
			g.EdgeNodes = append(g.EdgeNodes, i)
		} else {
			panic("isolated node, check map setting")
		}
	}
	return g
}

// GetNeighbors 获取节点的邻居节点
func (g *Graph) GetNeighbors(node int) []int {
	neighbors := []int{}
	for i, weight := range g.Matrix[node] {
		if weight > 0 {
			neighbors = append(neighbors, i)
		}
	}
	return neighbors
}

// GetEdgeWeight 获取两个节点之间的边的权值
func (g *Graph) GetEdgeWeight(from, to int) float64 {
	if g.Matrix[from][to] > 1 {
		return g.Matrix[from][to]
	}
	g.Matrix[from][to] = GetCarDistance(
		float64(g.Nodes[from][0]),
		float64(g.Nodes[to][0]),
		float64(g.Nodes[from][1]),
		float64(g.Nodes[to][1]),
	)
	g.Matrix[to][from] = g.Matrix[from][to]
	return g.Matrix[from][to]
}

func GetCarDistance(x1, x2, y1, y2 float64) float64 {
	return math.Sqrt(
		(x1-x2)*(x1-x2) + (y1-y2)*(y1-y2),
	)
}

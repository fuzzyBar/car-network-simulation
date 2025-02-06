package graph

type Graph struct {
	Nodes  []int       // 节点集合
	Matrix [][]float64 // 邻接矩阵，存储边的权值
}

func NewGraph() *Graph {
	// 初始化邻接矩阵
	matrix := [][]float64{
		{0, 0, 0, 0, 0, 0, 0, 0, 400, 0, 0, 0},       // 0
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 400, 0, 0},       // 1
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 400, 0, 0},       // 2
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 400, 0},       // 3
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 400, 0},       // 4
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 400},       // 5
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 400},       // 6
		{0, 0, 0, 0, 0, 0, 0, 0, 400, 0, 0, 0},       // 7
		{400, 0, 0, 0, 0, 0, 0, 400, 0, 400, 0, 400}, // 8
		{0, 400, 400, 0, 0, 0, 0, 0, 400, 0, 400, 0}, // 9
		{0, 0, 0, 400, 400, 0, 0, 0, 0, 400, 0, 400}, // 10
		{0, 0, 0, 0, 0, 400, 400, 0, 400, 0, 400, 0}, // 11
	}

	return &Graph{
		Nodes:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		Matrix: matrix,
	}
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
	return g.Matrix[from][to]
}

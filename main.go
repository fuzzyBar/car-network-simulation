package main

import (
	"car-network-simulation/simulation"
)

func main() {
	// 初始化随机种子
	//rand.Seed(time.Now().UnixNano())

	// 创建模拟
	sim := simulation.NewSimulation()

	// 运行模拟
	sim.Run(20*30, 100) // 模拟 50 个时间步长，车辆生成频率为20%
}

// import (
// 	"car-network-simulation/graph"
// 	"fmt"
// )

// func main() {
// 	g := graph.NewGraph()
// 	fmt.Println(g.GetEdgeWeight(3, 5))
// }

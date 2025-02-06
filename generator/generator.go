package generator

import (
	"car-network-simulation/graph"
	"car-network-simulation/vehicle"
	"math/rand"
)

type VehicleGenerator struct {
	Graph *graph.Graph
	// ComputePower int
	NextID int
}

func NewVehicleGenerator(g *graph.Graph, computePower int) *VehicleGenerator {
	return &VehicleGenerator{
		Graph: g,
		// ComputePower: computePower,
		NextID: 1,
	}
}

func (g *VehicleGenerator) GenerateVehicle() *vehicle.Vehicle {
	// 随机选择一个边缘节点作为起点
	startNodes := []int{0, 1, 2, 3, 4, 5, 6, 7}
	currentNode := startNodes[rand.Intn(len(startNodes))]

	// 随机选择一个邻居节点作为下一个节点
	neighbors := g.Graph.GetNeighbors(currentNode)
	nextNode := neighbors[rand.Intn(len(neighbors))]

	// 随机选择下下个节点
	nextNextNeighbors := g.Graph.GetNeighbors(nextNode)
	nextNextNode := nextNextNeighbors[rand.Intn(len(nextNextNeighbors))]
	if len(nextNextNeighbors) > 1 {
		for nextNextNode == currentNode {
			nextNextNode = nextNextNeighbors[rand.Intn(len(nextNextNeighbors))]
		}
	}

	// 随机生成速度（10-30 米/秒）
	speed := rand.Intn(21) + 10

	// 初始化 DistanceLeft
	distanceLeft := g.Graph.GetEdgeWeight(currentNode, nextNode)

	// 初始化 computerpower
	computerPower := rand.Intn(150) + 150
	computerPowerleft := rand.Intn(computerPower)

	// 创建车辆
	v := vehicle.NewVehicle(g.NextID, currentNode, nextNode, nextNextNode, speed, computerPower, computerPowerleft, distanceLeft)
	g.NextID++

	// 随机添加任务
	task := &vehicle.Task{
		//ID:          rand.Intn(1000),
		Size:        rand.Intn(10) + 1,
		ResourceReq: rand.Intn(50) + 1,
		Remaining:   rand.Intn(5) + 1,
	}
	v.AddTask(task)

	return v
}

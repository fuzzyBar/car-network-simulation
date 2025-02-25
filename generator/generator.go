package generator

import (
	"car-network-simulation/graph"
	"car-network-simulation/vehicle"
	"math/rand"
)

type VehicleGenerator struct {
	Graph *graph.Graph
	// ComputePower int
	ID int
}

func NewVehicleGenerator(g *graph.Graph) *VehicleGenerator {
	return &VehicleGenerator{
		Graph: g,
		// ComputePower: computePower,
		ID: 1,
	}
}

func (g *VehicleGenerator) GenerateVehicle() *vehicle.Vehicle {

	v := vehicle.NewVehicle()
	v.ID = g.ID
	v.Active = true

	// 随机选择一个边缘节点作为起点
	currentNode := g.Graph.EdgeNodes[rand.Intn(len(g.Graph.EdgeNodes))]
	v.CurrentNode = currentNode

	// 随机选择一个邻居节点作为下一个节点
	nextNode := g.Graph.GetNeighbors(currentNode)[0]
	v.NextNode = nextNode

	// 随机选择下下个节点
	nextNextNeighbors := g.Graph.GetNeighbors(nextNode)
	nextNextNode := nextNextNeighbors[rand.Intn(len(nextNextNeighbors))]
	if len(nextNextNeighbors) > 1 {
		for nextNextNode == currentNode {
			nextNextNode = nextNextNeighbors[rand.Intn(len(nextNextNeighbors))]
		}
	}
	v.NextNextNode = nextNextNode

	// 随机生成速度（8 - 14 米/秒）30km/h - 50km/h 市区密集生活区等车辆一般速度范围
	speed := rand.Intn(6) + 8
	v.Speed = speed

	// 初始化 DistanceLeft
	distanceLeft := g.Graph.GetEdgeWeight(currentNode, nextNode)
	v.DistanceLeft = distanceLeft

	// 初始化 computerpower
	computerPower := rand.Intn(150) + 150
	computerPowerleft := computerPower
	v.ComputePower = computerPower
	v.ComputePowerleft = computerPowerleft

	g.ID++

	// 随机添加任务
	// on todo list...
	v.NewTask()

	return v
}

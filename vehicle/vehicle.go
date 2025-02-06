package vehicle

import (
	"car-network-simulation/graph"
	"fmt"
	"math/rand"
)

type Vehicle struct {
	ID           int
	CurrentNode  int     // 当前所在节点
	NextNode     int     // 下一个节点
	NextNextNode int     // 下下个节点
	Speed        int     // 车辆速度（米/秒）
	DistanceLeft float64 // 剩余距离（米）
	Tasks        []Task
	ComputePower int  // 车辆的计算资源
	Active       bool // 车辆是否仍在模拟中
}

func NewVehicle(id, currentNode, nextNode, nextNextNode, speed, computePower int, distanceLeft float64) *Vehicle {
	return &Vehicle{
		ID:           id,
		CurrentNode:  currentNode,
		NextNode:     nextNode,
		NextNextNode: nextNextNode,
		Speed:        speed,
		DistanceLeft: distanceLeft,
		ComputePower: computePower,
		Active:       true,
	}
}

func (v *Vehicle) AddTask(task Task) {
	v.Tasks = append(v.Tasks, task)
	fmt.Printf("Vehicle %d added task %d\n", v.ID, task.ID)
}

func (v *Vehicle) ProcessTasks() {
	for i := range v.Tasks {
		if v.Tasks[i].Remaining > 0 {
			if v.ComputePower >= v.Tasks[i].ResourceReq {
				v.Tasks[i].Remaining--
				fmt.Printf("Vehicle %d processing task %d, remaining time: %d\n", v.ID, v.Tasks[i].ID, v.Tasks[i].Remaining)
			} else {
				fmt.Printf("Vehicle %d failed to process task %d due to insufficient compute resources\n", v.ID, v.Tasks[i].ID)
			}
		}
	}
}

func (v *Vehicle) Move(graph *graph.Graph) {
	if !v.Active {
		return
	}

	// 打印车辆状态
	// fmt.Printf("%d %.0f %d %d %d\n", v.ID, v.DistanceLeft, v.CurrentNode, v.NextNode, v.NextNextNode)

	// 如果车辆在节点上，更新下一个节点和下下个节点
	if v.DistanceLeft <= 0 {
		// 如果下一个节点是终止节点（边缘节点），且车辆已经到达该节点，则销毁车辆
		if v.NextNode >= 0 && v.NextNode <= 7 { //&& v.CurrentNode == v.NextNode {
			v.Active = false
			fmt.Printf("Vehicle %d reached the endpoint at node %d and is no longer active\n", v.ID, v.NextNode)
			return
		}

		fmt.Printf("Vehicle %d reached node %d\n", v.ID, v.NextNode)

		// 更新当前节点、下一个节点和下下个节点
		v.CurrentNode = v.NextNode
		v.NextNode = v.NextNextNode

		// 随机选择下下个节点
		neighbors := graph.GetNeighbors(v.NextNode)
		if len(neighbors) == 0 {
			v.Active = false
			fmt.Printf("Vehicle %d has no further path and is no longer active\n", v.ID)
			return
		}

		// 排除当前节点，避免调头（除非只有一个邻居）
		nextNextNode := neighbors[rand.Intn(len(neighbors))]
		if len(neighbors) > 1 {
			for nextNextNode == v.CurrentNode {
				nextNextNode = neighbors[rand.Intn(len(neighbors))]
			}
		}

		v.NextNextNode = nextNextNode
		v.DistanceLeft = graph.GetEdgeWeight(v.NextNode, v.NextNextNode)
	}

	// 移动车辆
	v.DistanceLeft -= float64(v.Speed)
	// if v.DistanceLeft <= 0 {
	// 	v.DistanceLeft = 0
	// 	fmt.Printf("Vehicle %d reached node %d\n", v.ID, v.NextNode)
	// } else {
	// 	fmt.Printf("Vehicle %d is moving towards node %d, distance left: %.0f meters\n", v.ID, v.NextNode, v.DistanceLeft)
	// }
}

package vehicle

import (
	"car-network-simulation/graph"
	"fmt"
	"math/rand"
)

type Vehicle struct {
	ID               int
	CurrentNode      int     // 当前所在节点
	NextNode         int     // 下一个节点
	NextNextNode     int     // 下下个节点
	Speed            int     // 车辆速度（米/秒）
	DistanceLeft     float64 // 剩余距离（米）
	Tasks            []*Task
	ComputePower     int     // 车辆的计算资源
	ComputePowerleft int     // 车辆剩余计算资源
	Active           bool    // 车辆是否仍在模拟中
	TasksToSend      []*Task //发送通道
	TaskToReceive    []*Task //接收通道
}

type Task struct {
	//ID          int
	Size        int  // 任务大小
	ResourceReq int  // 任务需要的计算资源
	Remaining   int  // 剩余处理时间
	Started     bool // 是否开始执行
	Finished    bool
	//transporting bool
}

func NewVehicle(id, currentNode, nextNode, nextNextNode, speed, computePower, computePowerleft int, distanceLeft float64) *Vehicle {
	return &Vehicle{
		ID:               id,
		CurrentNode:      currentNode,
		NextNode:         nextNode,
		NextNextNode:     nextNextNode,
		Speed:            speed,
		DistanceLeft:     distanceLeft,
		ComputePower:     computePower,
		ComputePowerleft: computePowerleft,
		Active:           true,
	}
}

func (v *Vehicle) AddTask(task *Task) {
	v.Tasks = append(v.Tasks, task)
	fmt.Printf("Vehicle added task\n") //v.ID, task.ID)
}

// 车辆任务处理模拟, unfinished
func (v *Vehicle) HandleTasks() {
	for i := range v.Tasks {

		v.Process(v.Tasks[i])

		// if v.Tasks[i].Remaining > 0 {
		// 	if v.ComputePower >= v.Tasks[i].ResourceReq {
		// 		v.Tasks[i].Remaining--
		// 		fmt.Printf("Vehicle %d processing task %d, remaining time: %d\n", v.ID, v.Tasks[i].ID, v.Tasks[i].Remaining)
		// 	} else {
		// 		fmt.Printf("Vehicle %d failed to process task %d due to insufficient compute resources\n", v.ID, v.Tasks[i].ID)
		// 	}
		// }
	}
}

// 车辆移动模拟, finished
func (v *Vehicle) Move(graph *graph.Graph) {
	if !v.Active {
		return
	}

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

		nextNextNode := neighbors[rand.Intn(len(neighbors))]
		// // 排除当前节点，避免调头（除非只有一个邻居）
		// if len(neighbors) > 1 {
		// 	for nextNextNode == v.CurrentNode {
		// 		nextNextNode = neighbors[rand.Intn(len(neighbors))]
		// 	}
		// }

		v.NextNextNode = nextNextNode
		v.DistanceLeft = graph.GetEdgeWeight(v.NextNode, v.NextNextNode)
	}

	// 移动车辆
	v.DistanceLeft -= float64(v.Speed) / 20

}

// 单个任务的执行, unfinished
func (v *Vehicle) Process(mytask *Task) {
	//剩余资源不足, 无法执行
	if v.ComputePowerleft < mytask.ResourceReq {
		v.TasksToSend = append(v.TasksToSend, mytask)
		return
	}
	//开始任务, 占用资源
	if !mytask.Started {
		mytask.Started = true
		v.ComputePowerleft -= mytask.ResourceReq
	}

	mytask.Remaining -= 10

	//判断是否完成任务
	if mytask.Remaining <= 0 {
		mytask.Finished = true
	}

	//完成任务, 释放资源
	if mytask.Finished {
		v.ComputePowerleft += mytask.ResourceReq
		return
	}
}

func (v *Vehicle) Close() {
	for i := range v.Tasks {
		v.Tasks[i] = nil
	}
	for i := range v.TasksToSend {
		v.TasksToSend[i] = nil
	}
	for i := range v.TaskToReceive {
		v.TaskToReceive[i] = nil
	}
	v.Tasks = nil
	v.TasksToSend = nil
	v.TaskToReceive = nil
}

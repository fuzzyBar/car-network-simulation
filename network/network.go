// 该网络负责模拟车辆的移动, 通信交流, 任务处理等
// 不包含随机事件的生成功能
// 即不包含网络中进入新的车辆, 车辆任务的生成
// 随机事件由simulation进行控制

package network

import (
	"car-network-simulation/graph"
	"car-network-simulation/vehicle"
	"fmt"
)

type Network struct {
	Vehicles []*vehicle.Vehicle
	Graph    *graph.Graph
}

func NewNetwork(g *graph.Graph) *Network {
	return &Network{
		Graph: g,
	}
}

// 主模拟函数
func (n *Network) SimulateTimeStep() {
	for _, v := range n.Vehicles {
		if v.Active {
			v.Move(n.Graph)
			v.HandleTasks()
		}
	}
	n.RemoveInactiveVehicles() // 移除不再活跃的车辆
}

// 向模拟网络中添加车辆
func (n *Network) AddVehicle(v *vehicle.Vehicle) {
	n.Vehicles = append(n.Vehicles, v)
	fmt.Printf("Vehicle %d added to network\n", v.ID)
}

// 移除网络中已达重点的车辆
func (n *Network) RemoveInactiveVehicles() {
	activeVehicles := make([]*vehicle.Vehicle, 0)
	for i, v := range n.Vehicles {
		if v.Active {
			activeVehicles = append(activeVehicles, v)
		} else {
			n.Vehicles[i].Close()
			n.Vehicles[i] = nil
			fmt.Printf("Vehicle %d is no longer active and has been removed\n", v.ID)
		}
	}
	// n.Vehicles = nil
	n.Vehicles = activeVehicles
}

// 获得车辆距离
func (n *Network) GetCarDistance(car1, car2 *vehicle.Vehicle) float64 {

	//获得car1坐标
	scale1 := car1.DistanceLeft / n.Graph.Matrix[car1.CurrentNode][car1.NextNode]
	car1_x := float64(n.Graph.Nodes[car1.NextNode][0]) - scale1*float64(
		n.Graph.Nodes[car1.NextNode][0]-n.Graph.Nodes[car1.CurrentNode][0],
	)
	car1_y := float64(n.Graph.Nodes[car1.NextNode][1]) - scale1*float64(
		n.Graph.Nodes[car1.NextNode][1]-n.Graph.Nodes[car1.CurrentNode][1],
	)

	//获得car2坐标
	scale2 := car2.DistanceLeft / n.Graph.Matrix[car2.CurrentNode][car2.NextNode]
	car2_x := float64(n.Graph.Nodes[car2.NextNode][0]) - scale2*float64(
		n.Graph.Nodes[car2.NextNode][0]-n.Graph.Nodes[car2.CurrentNode][0],
	)
	car2_y := float64(n.Graph.Nodes[car2.NextNode][1]) - scale2*float64(
		n.Graph.Nodes[car2.NextNode][1]-n.Graph.Nodes[car2.CurrentNode][1],
	)

	return graph.GetCarDistance(
		car1_x,
		car2_x,
		car1_y,
		car2_y,
	)
}

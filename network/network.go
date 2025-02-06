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

func (n *Network) AddVehicle(v *vehicle.Vehicle) {
	n.Vehicles = append(n.Vehicles, v)
	fmt.Printf("Vehicle %d added to network\n", v.ID)
}

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

func (n *Network) SimulateTimeStep() {
	for _, v := range n.Vehicles {
		if v.Active {
			v.Move(n.Graph)
			v.HandleTasks()
		}
	}
	n.RemoveInactiveVehicles() // 移除不再活跃的车辆
}

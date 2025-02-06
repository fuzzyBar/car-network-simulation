package simulation

import (
	"car-network-simulation/generator"
	"car-network-simulation/graph"
	"car-network-simulation/network"
	"fmt"
	"math/rand"
)

type Simulation struct {
	Network          *network.Network
	VehicleGenerator *generator.VehicleGenerator
}

func NewSimulation(computePower int) *Simulation {
	g := graph.NewGraph()
	return &Simulation{
		Network:          network.NewNetwork(g),
		VehicleGenerator: generator.NewVehicleGenerator(g, computePower),
	}
}

func (s *Simulation) Run(steps int, vehicleSpawnRate int) {
	for i := 0; i < steps; i++ {
		fmt.Printf("Time step %d:\n", i+1)

		// 随机生成车辆
		if rand.Intn(100) < vehicleSpawnRate {
			v := s.VehicleGenerator.GenerateVehicle()
			s.Network.AddVehicle(v)
		}

		// 模拟时间步长
		s.Network.SimulateTimeStep()
		//time.Sleep(1 * time.Second)
	}
}

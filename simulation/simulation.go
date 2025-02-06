package simulation

import (
	"car-network-simulation/generator"
	"car-network-simulation/graph"
	"car-network-simulation/network"
	"car-network-simulation/vehicle"
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
		fmt.Printf("Time step %d:\n\n", i+1)

		// 每20step, 即1秒, 尝试随机生成新任务, 尝试生成新车辆
		if i%20 == 0 {
			//随机生成任务
			s.TasksGeneration(3, 50)

			// 随机生成车辆
			if rand.Intn(100) < vehicleSpawnRate {
				v := s.VehicleGenerator.GenerateVehicle()
				s.Network.AddVehicle(v)
			}
		}

		// 模拟时间步长
		s.Network.SimulateTimeStep()
	}
}

func (s *Simulation) TasksGeneration(tryNum, successRate int) {
	for _, car := range s.Network.Vehicles {
		// 尝试生成任务的次数
		for i := 0; i < tryNum; i++ {
			// 每个任务生成的成功率50%
			if rand.Intn(100) < successRate {
				car.AddTask(
					&vehicle.Task{
						//ID:          rand.Intn(1000),
						Size:        rand.Intn(10) + 1,
						ResourceReq: rand.Intn(50) + 1,
						Remaining:   rand.Intn(5) + 1,
					},
				)
			}
		}
	}
}

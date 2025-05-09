package swarm

import (
	"fmt"
	"reflect"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestUpdateGlobalBest(t *testing.T) {
	var dimensions int = 3
	var expectedResult float64 = 50

	var velocityWeight VelocityWeight = VelocityWeight{
		weight:                     newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		personalBestVelocityWeight: newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		globalBestVelocityWeight:   newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
	}

	swarm, err := CreateSwarm(velocityWeight, dimensions)
	if err != nil {
		t.Error(err)
	}

	swarm.Particles = append(swarm.Particles,
		Particle{
			dimension:         dimensions,
			position:          newDimensionDense(dimensions, []float64{3.0, 2.0, 1}),
			personalBest:      newDimensionDense(dimensions, []float64{1.0, 1.0, 1}),
			personalBestValue: 0,
			velocity:          newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
			velocityWeight:    velocityWeight,
		},
	)

	for i := 0; i < 4; i++ {
		swarm.Particles = append(swarm.Particles, swarm.Particles[0])
	}

	swarm.Particles[0].personalBestValue = 10
	swarm.Particles[1].personalBestValue = 50
	swarm.Particles[2].personalBestValue = 5
	swarm.Particles[3].personalBestValue = 50
	swarm.Particles[4].personalBestValue = 10

	err = swarm.UpdateGlobalBest()
	if err != nil {
		t.Error(err)
	}

	result := swarm.GlobalBestValue

	if !reflect.DeepEqual(expectedResult, result) {
		fmt.Printf("expectedResult : %v \n", expectedResult)
		fmt.Printf("result : %v \n", result)
		t.Error("global best value doesn't match ")
	}
}

func TestUpdateParticleBest(t *testing.T) {
	var dimensions int = 3
	var expectedResult float64 = 14

	var velocityWeight VelocityWeight = VelocityWeight{
		weight:                     newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		personalBestVelocityWeight: newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		globalBestVelocityWeight:   newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
	}

	swarm, err := CreateSwarm(velocityWeight, dimensions)
	if err != nil {
		t.Error(err)
	}

	swarm.Particles = append(swarm.Particles,
		Particle{
			dimension:         dimensions,
			position:          newDimensionDense(dimensions, []float64{3.0, 2.0, 1}),
			personalBest:      newDimensionDense(dimensions, []float64{1.0, 1.0, 1}),
			personalBestValue: 0,
			velocity:          newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
			velocityWeight:    velocityWeight,
		},
	)

	for i := 0; i < 4; i++ {
		swarm.Particles = append(swarm.Particles, swarm.Particles[0])
	}

	swarm.Particles[0].personalBestValue = 1
	swarm.Particles[1].personalBestValue = 1
	swarm.Particles[2].personalBestValue = 1
	swarm.Particles[3].personalBestValue = 5
	swarm.Particles[4].personalBestValue = 1

	swarm.Particles[0].position = newDimensionDense(dimensions, []float64{10.0, 1.0, 1})
	swarm.Particles[1].position = newDimensionDense(dimensions, []float64{10.0, 2.0, 1})
	swarm.Particles[2].position = newDimensionDense(dimensions, []float64{10.0, 3.0, 1})
	swarm.Particles[3].position = newDimensionDense(dimensions, []float64{10.0, 4.0, 1})
	swarm.Particles[4].position = newDimensionDense(dimensions, []float64{10.0, 5.0, 1})

	particleFitnessFunction := func(position *mat.Dense) float64 {
		return position.RawMatrix().Data[0] + position.RawMatrix().Data[1]
	}

	err = swarm.UpdateParticleBest(particleFitnessFunction)
	if err != nil {
		t.Error(err)
	}

	result := swarm.Particles[3].personalBestValue

	if !reflect.DeepEqual(expectedResult, result) {
		fmt.Printf("particle position : %v \n", swarm.Particles[3].position)
		fmt.Printf("expectedResult : %v \n", expectedResult)
		fmt.Printf("result : %v \n", result)
		t.Error("global best value doesn't match ")
	}
}

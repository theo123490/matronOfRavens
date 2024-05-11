package swarm

import (
	"fmt"
	"reflect"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestSubtractFunction(t *testing.T) {
	var dimensions int = 3
	var expectedResult *mat.Dense = newDimensionDense(dimensions, []float64{2, 1, 0})

	var velocityWeight VelocityWeight = VelocityWeight{
		weight:                     newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		personalBestVelocityWeight: newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		globalBestVelocityWeight:   newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
	}

	var particle Particle = Particle{
		dimension:      dimensions,
		position:       newDimensionDense(dimensions, []float64{3.0, 2.0, 1}),
		personalBest:   newDimensionDense(dimensions, []float64{1.0, 1.0, 1}),
		velocity:       newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocityWeight: velocityWeight,
	}

	result := newDimensionDense(dimensions, nil)
	result.Sub(particle.position, particle.personalBest)
	if !reflect.DeepEqual(expectedResult, result) {
		fmt.Printf("result %v \n", result)
		fmt.Printf("expected Result %v \n", expectedResult)
		t.Error("result and expectedresult are not the sames")
	}
}

func TestCalculateVelocity(t *testing.T) {
	var dimensions int = 3
	var expectedResult *mat.Dense = newDimensionDense(dimensions, []float64{2, 2, 0})

	var velocityWeight VelocityWeight = VelocityWeight{
		weight:                     newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		personalBestVelocityWeight: newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		globalBestVelocityWeight:   newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
	}

	var particle Particle = Particle{
		dimension:      dimensions,
		position:       newDimensionDense(dimensions, []float64{3.0, 2.0, 1}),
		personalBest:   newDimensionDense(dimensions, []float64{1.0, 1.0, 1}),
		velocity:       newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocityWeight: velocityWeight,
	}

	result := particle.calculateBestVelocity(particle.personalBest, particle.velocityWeight.personalBestVelocityWeight)
	if !reflect.DeepEqual(expectedResult, result) {
		fmt.Printf("result %v \n", result)
		fmt.Printf("expected Result %v \n", expectedResult)
		t.Error("result and expectedresult are not the sames")
	}
}

func TestCalculatePBestVelocity(t *testing.T) {
	var dimensions int = 3
	var expectedResult *mat.Dense = newDimensionDense(dimensions, []float64{2, 2, 0})

	var velocityWeight VelocityWeight = VelocityWeight{
		weight:                     newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		personalBestVelocityWeight: newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		globalBestVelocityWeight:   newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
	}

	var particle Particle = Particle{
		dimension:      dimensions,
		position:       newDimensionDense(dimensions, []float64{3.0, 2.0, 1}),
		personalBest:   newDimensionDense(dimensions, []float64{1.0, 1.0, 1}),
		velocity:       newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocityWeight: velocityWeight,
	}

	result := particle.calculatePersonalBestVelocity()
	if !reflect.DeepEqual(expectedResult, result) {
		fmt.Printf("result %v \n", result)
		fmt.Printf("expected Result %v \n", expectedResult)
		t.Error("result and expectedresult are not the sames")
	}
}

func TestCalculateGBestVelocity(t *testing.T) {
	var dimensions int = 3
	var expectedResult *mat.Dense = newDimensionDense(dimensions, []float64{4, 0, 2})

	var velocityWeight VelocityWeight = VelocityWeight{
		weight:                     newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		personalBestVelocityWeight: newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		globalBestVelocityWeight:   newDimensionDense(dimensions, []float64{2.0, 3.0, 2}),
	}

	var particle Particle = Particle{
		dimension:      dimensions,
		position:       newDimensionDense(dimensions, []float64{3.0, 2.0, 1}),
		personalBest:   newDimensionDense(dimensions, []float64{1.0, 1.0, 1}),
		velocity:       newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocityWeight: velocityWeight,
	}

	globalBest := newDimensionDense(dimensions, []float64{1.0, 2.0, 0})

	result := particle.calculateGlobalBestVelocity(globalBest)
	if !reflect.DeepEqual(expectedResult, result) {
		fmt.Printf("result %v \n", result)
		fmt.Printf("expected Result %v \n", expectedResult)
		t.Error("result and expectedresult are not the sames")
	}
}

func TestCalculateTotalVelocity(t *testing.T) {
	var dimensions int = 3
	var expectedResult *mat.Dense = newDimensionDense(dimensions, []float64{7, 4, 3})

	var velocityWeight VelocityWeight = VelocityWeight{
		weight:                     newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		personalBestVelocityWeight: newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		globalBestVelocityWeight:   newDimensionDense(dimensions, []float64{2.0, 3.0, 2}),
	}

	var particle Particle = Particle{
		dimension:      dimensions,
		position:       newDimensionDense(dimensions, []float64{3.0, 2.0, 1}),
		personalBest:   newDimensionDense(dimensions, []float64{1.0, 1.0, 1}),
		velocity:       newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocityWeight: velocityWeight,
	}

	globalBest := newDimensionDense(dimensions, []float64{1.0, 2.0, 0})

	result := particle.calculateTotalVelocity(globalBest)
	if !reflect.DeepEqual(expectedResult, result) {
		fmt.Printf("result %v \n", result)
		fmt.Printf("expected Result %v \n", expectedResult)
		t.Error("result and expectedresult are not the sames")
	}
}

func TestCalculatePositionValue(t *testing.T) {
	var dimensions int = 3
	var expectedResult float64 = 5

	var velocityWeight VelocityWeight = VelocityWeight{
		weight:                     newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		personalBestVelocityWeight: newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		globalBestVelocityWeight:   newDimensionDense(dimensions, []float64{2.0, 3.0, 2}),
	}

	var particle Particle = Particle{
		dimension:      dimensions,
		position:       newDimensionDense(dimensions, []float64{3.0, 2.0, 1}),
		personalBest:   newDimensionDense(dimensions, []float64{1.0, 1.0, 1}),
		velocity:       newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocityWeight: velocityWeight,
	}

	particleFitnessFunction := func(position *mat.Dense) float64 {
		return position.RawMatrix().Data[0] + position.RawMatrix().Data[1]
	}

	result := particle.calculatePositionValue(particleFitnessFunction)
	if !reflect.DeepEqual(expectedResult, result) {
		fmt.Printf("result %v \n", result)
		fmt.Printf("expected Result %v \n", expectedResult)
		t.Error("result and expectedresult are not the sames")
	}
}

func TestUpdateBestValue(t *testing.T) {
	var dimensions int = 3
	var expectedResult float64 = 5

	var velocityWeight VelocityWeight = VelocityWeight{
		weight:                     newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		personalBestVelocityWeight: newDimensionDense(dimensions, []float64{1.0, 2.0, 2}),
		globalBestVelocityWeight:   newDimensionDense(dimensions, []float64{2.0, 3.0, 2}),
	}

	var particle Particle = Particle{
		dimension:      dimensions,
		position:       newDimensionDense(dimensions, []float64{3.0, 2.0, 1}),
		personalBest:   newDimensionDense(dimensions, []float64{1.0, 1.0, 1}),
		velocity:       newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocityWeight: velocityWeight,
	}

	particleFitnessFunction := func(position *mat.Dense) float64 {
		return position.RawMatrix().Data[0] + position.RawMatrix().Data[1]
	}

	particle.updateBestValue(particleFitnessFunction)
	result := particle.personalBestValue
	if !reflect.DeepEqual(expectedResult, result) {
		fmt.Printf("result %v \n", result)
		fmt.Printf("expected Result %v \n", expectedResult)
		t.Error("result and expectedresult are not the sames")
	}
}

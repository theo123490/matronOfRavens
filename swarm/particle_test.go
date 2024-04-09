package swarm

import (
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
		globalBest:     newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocity:       newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocityWeight: velocityWeight,
	}

	result := newDimensionDense(dimensions, nil)
	result.Sub(particle.position, particle.personalBest)
	if !reflect.DeepEqual(expectedResult, result) {
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
		globalBest:     newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocity:       newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocityWeight: velocityWeight,
	}

	result := particle.calculateBestVelocity(particle.personalBest, particle.velocityWeight.personalBestVelocityWeight)
	if !reflect.DeepEqual(expectedResult, result) {
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
		globalBest:     newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocity:       newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocityWeight: velocityWeight,
	}

	result := particle.calculatePBestVelocity()
	if !reflect.DeepEqual(expectedResult, result) {
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
		globalBest:     newDimensionDense(dimensions, []float64{1.0, 2.0, 0}),
		velocity:       newDimensionDense(dimensions, []float64{1.0, 2.0, 1}),
		velocityWeight: velocityWeight,
	}

	result := particle.calculateGBestVelocity()
	if !reflect.DeepEqual(expectedResult, result) {
		t.Error("result and expectedresult are not the sames")
	}
}

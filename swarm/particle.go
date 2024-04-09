package swarm

type dimensionValues []float64

type Particle struct {
	position       dimensionValues
	personalBest   dimensionValues
	globalBest     dimensionValues
	velocity       dimensionValues
	velocityWeight VelocityWeight
}

type VelocityWeight struct {
	weight                     dimensionValues
	personalBestVelocityWeight dimensionValues
	globalBestVelocityWeight   dimensionValues
}

func (d dimensionValues) subtractWith(refD dimensionValues) dimensionValues {
	var subtracted dimensionValues = make(dimensionValues, len(d))
	for i := range d {
		subtracted[i] = d[i] - refD[i]
	}
	return subtracted
}

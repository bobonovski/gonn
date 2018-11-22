package gonn

import "fmt"

// compute the plane perpendicular to the line
// connecting the given two end points
func GetBoundary(p []float32, q []float32) ([]float32, error) {
	if len(p) != len(q) {
		return nil, fmt.Errorf("input vectors dimension incompatible")
	}
	// compute the middle point of the line
	var mid []float32
	for k := 0; k < len(p); k += 1 {
		mid = append(mid, (p[k]+q[k])/2.0)
	}
	// compute the plane coefficients
	var coef []float32
	intercept := float32(0.0)
	for k := 0; k < len(p); k += 1 {
		coef = append(coef, q[k]-p[k])
		intercept -= (q[k] - p[k]) * mid[k]
	}
	coef = append(coef, intercept)
	return coef, nil
}

// evalute formula value on input data
func EvalFormula(c []float32, d []float32) (float32, error) {
	if len(c) != len(d)+1 {
		return 0.0, fmt.Errorf("data dimension is incompatible with coefficients length")
	}
	result := float32(0.0)
	for k := 0; k < len(d); k += 1 {
		result += c[k] * d[k]
	}
	result += c[len(d)]
	return result, nil
}

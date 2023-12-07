package main

import (
	"fmt"
	"math"
)

func main() {
	times := []float64{42, 89, 91, 89}
	distance := []float64{308, 1170, 1291, 1467}

	// test
	// times := []float64{7, 15, 30}
	// distance := []float64{9, 40, 200}

	beating := 1.
	// d = t * (T-t)
	// 0 = t^2-tT+d
	// t = (T+-sqrt(T^2-4d))/2
	for i := 0; i < len(times); i++ {
		minBound := math.Floor((times[i]-math.Sqrt(times[i]*times[i]-4*distance[i]))/2) + 1
		maxBound := math.Ceil((times[i]+math.Sqrt(times[i]*times[i]-4*distance[i]))/2) - 1

		beating *= maxBound - minBound + 1
	}

	fmt.Println(int(beating))

	T := 42899189.
	d := 308117012911467.
	minBound := math.Floor((T-math.Sqrt(T*T-4*d))/2) + 1
	maxBound := math.Ceil((T+math.Sqrt(T*T-4*d))/2) - 1
	fmt.Println(int(maxBound - minBound + 1))
}

package main

import (
	"fmt"
	"time"
)

func main() {

	t0 := time.Now()
	t1 := time.Now()
	since := time.Since(t0)

	fmt.Println("t0:", t0)
	fmt.Println("t1:", t1)
	fmt.Println("time.Since(t0):", since)

	t2 := time.Now()

	timeSlices := make([]time.Duration, 0)
	for i := 0; i < 1000; i++ {
		timeSlices = append(timeSlices, time.Since(t2))
	}
	fmt.Println("Result of 1000 executed time.Since(t2) in for statement.")
	for i, value := range timeSlices {
		fmt.Println(fmt.Sprintf("%03d", i), value)
	}
}

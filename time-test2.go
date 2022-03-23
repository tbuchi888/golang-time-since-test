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

	timeSlices := make([]time.Time, 0)
	for i := 0; i < 1000; i++ {
		timeSlices = append(timeSlices, time.Now())
	}
	fmt.Println("Result of 1000 executed time.Now() in for statement.")
	for i, value := range timeSlices {
		fmt.Println(i, value)
	}
}

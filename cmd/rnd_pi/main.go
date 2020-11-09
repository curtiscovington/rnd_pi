package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

type point struct {
	X float64
	Y float64
}

type PiApproximation struct {
	pointsInCircle int64
	totalPoints    int64
	result         chan float64
}

func NewRandPoint() *point {
	x := rand.Float64()
	y := rand.Float64()
	return &point{X: x, Y: y}
}

func (p point) Distance() float64 {
	x2 := math.Pow(p.X, 2)
	y2 := math.Pow(p.Y, 2)

	return math.Sqrt(x2 + y2)
}

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	pa := PiApproximation{}
	go pa.Run()
	go func() {
		for {
			select {
			case r := <-pa.result:
				fmt.Println(r)
			default:
			}
		}
	}()
	<-sigs
	// Run()
}

func (pa *PiApproximation) Run() {
	pa.result = make(chan float64, 1)
	for {
		pa.result <- pa.Step()
	}
}

func (pa *PiApproximation) Step() float64 {
	for i := 0; i < 1000000; i++ {
		p := NewRandPoint()

		if p.Distance() < 1 {
			pa.pointsInCircle++
		}

		pa.totalPoints++
	}

	return float64(4) * float64(pa.pointsInCircle) / float64(pa.totalPoints)
}

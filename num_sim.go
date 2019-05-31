package main

import (
	"fmt"
	"math"
)

type M struct {
	NumSlotsSubFrm float64
	Numerology     float64
	NumSlotsFrm    float64
	SlotLen        float64
	FMax           float64
	FFTMax         float64
}

func SimulationConfig(scs float64, fc int) *M {
	var A = new(M)
	var Z = float64(scs)
	var Tc float64
	var TimeSubFrm float64
	FMax := 480 * (math.Pow(10, 3))
	A.FFTMax = 4096

	A.Numerology = (math.Log2(Z / 15000))

	A.NumSlotsSubFrm = math.Pow(2, A.Numerology)

	A.NumSlotsFrm = 10 * A.NumSlotsSubFrm

	Tc = 1 / (FMax * A.FFTMax)
	TimeSubFrm = (FMax * A.FFTMax * Tc) / 1000
	A.SlotLen = TimeSubFrm / A.NumSlotsSubFrm

	return A

}

func main() {
	scs := 30 * math.Pow(10, 3)
	BW := 10
	// fc := 10
	res := SimulationConfig(scs, BW)

	fmt.Println("Numerology: ", res.Numerology, "\nNumber of slots per sub frame: ", res.NumSlotsSubFrm, "\nNumber of slots per frame: ", res.NumSlotsFrm, "\nSlot Length(msec): ", res.SlotLen)

}

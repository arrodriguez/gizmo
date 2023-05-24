package main

import (
	"math"
	"time"
)

type BusyLoad struct {
	// Between 0 and 1
	Load          float64       `json:"load"`
	Duration      time.Duration `json:"duration"`
	TickFrecuency time.Duration `json:"tick_frequency"`
}

/*
Ticks:

*/

// TODO: We need to decide if we are going to query number of core & cpus and simulate full CPU exhaustion
func (bl *BusyLoad) Run() {
	hz := 0
	st := time.Now()

	t := time.NewTimer(bl.Duration)
	tck := time.NewTicker(bl.TickFrecuency)

Loop:
	for {
		select {
		case now := <-tck.C:
			diff := now.Sub(st)
			dst := bl.Duration - diff
			slt := time.Duration(math.Ceil((1.0 - bl.Load) * float64(bl.TickFrecuency)))
			if dst > 0 && dst < slt {
				slt = dst
			}
			time.Sleep(slt)
		case <-t.C:
			tck.Stop()
			break Loop
		default:
			hz++
		}
	}

}

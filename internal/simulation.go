package internal

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	ModeOne Mode = 1
	ModeTwo Mode = 2
)

type Mode uint8

type SimulationParams struct {
	Mode                 Mode
	NumberOfBalls        uint8
	NumberOfMinutesToRun uint
}

type simulation struct {
	clock  *clock
	params SimulationParams
}

func NewSimulation(params SimulationParams) *simulation {
	// setup tracks (defined in reverse order for linking)
	// hour - holds balls indicating how many hours have passed (max 12)
	hourTrack := &track{maxBalls: 12, name: "Hour"}

	// five minute - holds balls indicating how many minutes have passed in increments of five (max 12)
	minFiveTrack := &track{maxBalls: 12, name: "FiveMin", nextTrack: hourTrack}

	// minute - holds balls indicating how many minutes have passed (max 5)
	minTrack := &track{maxBalls: 5, name: "Min", nextTrack: minFiveTrack}

	// reservoir - holds the unused balls not currently in any other track
	reservoir := &track{name: "Main", nextTrack: minTrack}

	// load balls into reservoir
	for i := uint8(1); i <= params.NumberOfBalls; i++ {
		reservoir.add(i)
	}

	// setup simulation
	return &simulation{
		clock: &clock{
			reservoir: reservoir,
		},
		params: params,
	}
}

func (s *simulation) Run() {
	fmt.Printf("simulation started: %+v\n", s.params)

	switch s.params.Mode {
	case ModeOne:
		start := time.Now()
		nb := int(s.params.NumberOfBalls)
		for {
			s.clock.tick()

			// check if reservoir is full
			if s.clock.reservoir.currentLength == nb {
				// check if reservoir is in initial order
				isInOrder := true
				for i, v := range s.clock.reservoir.balls {
					if v != uint8(i+1) {
						isInOrder = false
						break
					}
				}
				if isInOrder {
					fmt.Printf("simulation took: %s\n", time.Since(start))
					fmt.Printf("simulation output: %d balls cycle after %d days\n", s.params.NumberOfBalls, int(s.clock.getElapsedDays()))
					break
				}
			}
		}
	case ModeTwo:
		start := time.Now()
		for i := uint(0); i < s.params.NumberOfMinutesToRun; i++ {
			s.clock.tick()
		}
		fmt.Printf("simulation took: %s\n", time.Since(start))
		stateJSON, _ := json.Marshal(s.clock.getCurrentTrackStates())
		fmt.Printf("simulation output: %s\n", string(stateJSON))
	default:
		fmt.Println("mode not supported yet")
	}
}

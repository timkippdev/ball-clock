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
	NumberOfBalls        uint
	NumberOfMinutesToRun uint
}

type simulation struct {
	clock  *clock
	params SimulationParams
}

func NewSimulation(params SimulationParams) *simulation {
	// setup tracks (defined in reverse order for linking)
	// hour - holds balls indicating how many hours have passed (max 12)
	hourTrack := &track{balls: &ballQueue{max: 12}, name: "Hour", flushDayRatio: 0.5} // flushDayRatio 0.5 because each flush is 1/2 day (12 hours)

	// five minute - holds balls indicating how many minutes have passed in increments of five (max 12)
	minFiveTrack := &track{balls: &ballQueue{max: 12}, name: "FiveMin", nextTrack: hourTrack}

	// minute - holds balls indicating how many minutes have passed (max 5)
	minTrack := &track{balls: &ballQueue{max: 5}, name: "Min", nextTrack: minFiveTrack}

	// reservoir - holds the unused balls not currently in any other track
	reservoir := &track{balls: &ballQueue{}, name: "Main", nextTrack: minTrack}

	// load balls into reservoir
	for i := uint(1); i <= params.NumberOfBalls; i++ {
		reservoir.balls.add(&ball{
			id: i,
		})
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

	start := time.Now()
	defer func() {
		fmt.Printf("simulation took: %s\n", time.Since(start))
	}()

	switch s.params.Mode {
	case ModeOne:
		for {
			s.clock.tick()

			// check if reservoir is full
			elements := s.clock.reservoir.balls.elements
			if len(elements) == int(s.params.NumberOfBalls) {
				// check if reservoir is in initial order
				isInOrder := true
				for i, v := range elements {
					if v.id != uint(i+1) {
						isInOrder = false
						break
					}
				}
				if isInOrder {
					fmt.Printf("simulation output: %d balls cycle after %d days\n", s.params.NumberOfBalls, int(s.clock.getElapsedDays()))
					break
				}
			}
		}
	case ModeTwo:
		for i := uint(0); i < s.params.NumberOfMinutesToRun; i++ {
			s.clock.tick()
		}
		stateJSON, _ := json.Marshal(s.clock.getCurrentTrackStates())
		fmt.Printf("simulation output: %s\n", string(stateJSON))
	default:
		fmt.Println("mode not supported yet")
	}
}

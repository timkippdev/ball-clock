package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/timkippdev/ball-clock/internal"
)

var (
	mode                 internal.Mode
	numberOfBalls        uint
	numberOfMinutesToRun uint
)

func init() {
	// accept input parameters to determine "mode"
	flag.UintVar(&numberOfBalls, "balls", 0, "number of balls")
	flag.UintVar(&numberOfMinutesToRun, "minutes", 0, "number of minutes to run")
	flag.Parse()

	if numberOfBalls == 0 {
		fmt.Println("please supply --balls argument")
		os.Exit(1)
	}

	// validate number of balls input is between 27 and 127 (inclusive)
	if numberOfBalls < 27 || numberOfBalls > 127 {
		fmt.Println("please make sure the --balls argument is between 27 and 127")
		os.Exit(1)
	}

	mode = internal.ModeOne
	if numberOfMinutesToRun > 0 {
		mode = internal.ModeTwo
	}
}

func main() {
	// run simulation
	internal.NewSimulation(internal.SimulationParams{
		Mode:                 mode,
		NumberOfBalls:        uint8(numberOfBalls),
		NumberOfMinutesToRun: numberOfMinutesToRun,
	}).Run()
}

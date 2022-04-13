package internal

import "testing"

func BenchmarkSimulationModeOne30Balls(b *testing.B) {
	fireSimulation(b, ModeOne, 30, 0)
}

func BenchmarkSimulationModeOne45Balls(b *testing.B) {
	fireSimulation(b, ModeOne, 30, 0)
}

func BenchmarkSimulationModeTwo30Balls(b *testing.B) {
	fireSimulation(b, ModeTwo, 30, 325)
}

func fireSimulation(b *testing.B, mode Mode, balls uint, minutes uint) {
	for n := 0; n < b.N; n++ {
		NewSimulation(SimulationParams{
			Mode:                 mode,
			NumberOfBalls:        balls,
			NumberOfMinutesToRun: minutes,
		}).Run()
	}
}

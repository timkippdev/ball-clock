package internal

import "testing"

func BenchmarkSimulationModeOne30Balls(b *testing.B) {
	fireSimulation(b, ModeOne, 30, 0)
}

func BenchmarkSimulationModeOne45Balls(b *testing.B) {
	fireSimulation(b, ModeOne, 45, 0)
}

func BenchmarkSimulationModeOne123Balls(b *testing.B) {
	fireSimulation(b, ModeOne, 123, 0)
}

func BenchmarkSimulationModeTwo30Balls325Minutes(b *testing.B) {
	fireSimulation(b, ModeTwo, 30, 325)
}

func fireSimulation(b *testing.B, mode Mode, balls uint8, minutes uint) {
	for n := 0; n < b.N; n++ {
		NewSimulation(SimulationParams{
			Mode:                 mode,
			NumberOfBalls:        balls,
			NumberOfMinutesToRun: minutes,
		}).Run()
	}
}

package internal

import "testing"

func BenchmarkAdd0ExistingBalls(b *testing.B) {
	benchmarkAdd(b, 0)
}

func BenchmarkAdd100ExistingBalls(b *testing.B) {
	benchmarkAdd(b, 100)
}

func benchmarkAdd(b *testing.B, ballCount uint8) {
	t := &track{balls: make([]uint8, ballCount)}
	for i := uint8(0); i < ballCount; i++ {
		t.balls[i] = i
	}

	for n := 0; n < b.N; n++ {
		t.add(0)
	}
}

func BenchmarkCanAdd100BallsNoMax(b *testing.B) {
	benchmarkCanAdd(b, 100, 0)
}

func benchmarkCanAdd(b *testing.B, ballCount uint8, max uint8) {
	t := &track{balls: make([]uint8, ballCount), maxBalls: max}
	for i := uint8(0); i < ballCount; i++ {
		t.balls[i] = i
	}

	for n := 0; n < b.N; n++ {
		t.canAdd()
	}
}

func BenchmarkFlush(b *testing.B) {
	t := &track{}

	for n := 0; n < b.N; n++ {
		t.flush()
	}
}

func BenchmarkGetNext0Balls(b *testing.B) {
	benchmarkGetNext(b, 0)
}

func BenchmarkGetNext100Balls(b *testing.B) {
	benchmarkGetNext(b, 100)
}

func benchmarkGetNext(b *testing.B, ballCount uint8) {
	t := &track{balls: make([]uint8, ballCount)}
	for i := uint8(0); i < ballCount; i++ {
		t.balls[i] = i
	}

	for n := 0; n < b.N; n++ {
		t.getNext()
	}
}

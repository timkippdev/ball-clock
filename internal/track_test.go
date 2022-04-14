package internal

import "testing"

func BenchmarkAdd0ExistingBalls(b *testing.B) {
	benchmarkAdd(b, 0)
}

func BenchmarkAdd100ExistingBalls(b *testing.B) {
	benchmarkAdd(b, 100)
}

func BenchmarkAdd10000ExistingBalls(b *testing.B) {
	benchmarkAdd(b, 10000)
}

func BenchmarkAdd1000000ExistingBalls(b *testing.B) {
	benchmarkAdd(b, 1000000)
}

func benchmarkAdd(b *testing.B, ballCount int) {
	t := &track{balls: make([]uint, ballCount)}
	for i := 0; i < ballCount; i++ {
		t.balls[i] = uint(i)
	}

	for n := 0; n < b.N; n++ {
		t.add(uint(n))
	}
}

func BenchmarkCanAdd100BallsNoMax(b *testing.B) {
	benchmarkCanAdd(b, 100, 0)
}

func BenchmarkCanAdd1000Balls100Max(b *testing.B) {
	benchmarkCanAdd(b, 1000, 100)
}

func benchmarkCanAdd(b *testing.B, ballCount int, max uint8) {
	t := &track{balls: make([]uint, ballCount), maxBalls: max}
	for i := 0; i < ballCount; i++ {
		t.balls[i] = uint(i)
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

func BenchmarkGetNext10000Balls(b *testing.B) {
	benchmarkGetNext(b, 10000)
}

func BenchmarkGetNext1000000Balls(b *testing.B) {
	benchmarkGetNext(b, 1000000)
}

func benchmarkGetNext(b *testing.B, ballCount int) {
	t := &track{balls: make([]uint, ballCount)}
	for i := 0; i < ballCount; i++ {
		t.balls[i] = uint(i)
	}

	for n := 0; n < b.N; n++ {
		t.getNext()
	}
}

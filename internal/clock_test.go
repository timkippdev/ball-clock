package internal

import (
	"testing"
)

func BenchmarkFlushTrack100(b *testing.B) {
	benchmarkFlushTrack(b, 100)
}

func BenchmarkFlushTrack1000(b *testing.B) {
	benchmarkFlushTrack(b, 1000)
}

func BenchmarkFlushTrack10000(b *testing.B) {
	benchmarkFlushTrack(b, 10000)
}

func BenchmarkFlushTrack100000(b *testing.B) {
	benchmarkFlushTrack(b, 100000)
}

func BenchmarkFlushTrack1000000(b *testing.B) {
	benchmarkFlushTrack(b, 1000000)
}

func benchmarkFlushTrack(b *testing.B, ballCount uint) {
	t := &track{balls: make([]uint, ballCount)}
	c := &clock{
		reservoir: &track{balls: make([]uint, 0), nextTrack: t},
	}
	for i := uint(0); i < ballCount; i++ {
		t.balls[i] = i
	}

	for n := 0; n < b.N; n++ {
		c.flushTrack(t)
	}
}

func BenchmarkTick100(b *testing.B) {
	benchmarkTick(b, 100)
}

func BenchmarkTick1000(b *testing.B) {
	benchmarkTick(b, 1000)
}

func BenchmarkTick10000(b *testing.B) {
	benchmarkTick(b, 10000)
}

func BenchmarkTick100000(b *testing.B) {
	benchmarkTick(b, 100000)
}

func BenchmarkTick1000000(b *testing.B) {
	benchmarkTick(b, 1000000)
}

func benchmarkTick(b *testing.B, ballCount uint) {
	t := &track{balls: make([]uint, ballCount)}
	c := &clock{
		reservoir: &track{balls: make([]uint, 0), nextTrack: t},
	}
	for i := uint(0); i < ballCount; i++ {
		t.balls[i] = i
	}

	for n := 0; n < b.N; n++ {
		c.tick()
	}
}

func BenchmarkMoveBallToTrack1TrackMax1(b *testing.B) {
	benchmarkMoveBallToTrack(b, 1, 1)
}

func BenchmarkMoveBallToTrack2TracksMax1(b *testing.B) {
	benchmarkMoveBallToTrack(b, 2, 1)
}

func BenchmarkMoveBallToTrack3TracksMax1(b *testing.B) {
	benchmarkMoveBallToTrack(b, 3, 1)
}

func BenchmarkMoveBallToTrack100TracksMax1(b *testing.B) {
	benchmarkMoveBallToTrack(b, 100, 1)
}

func BenchmarkMoveBallToTrack1TrackMax5(b *testing.B) {
	benchmarkMoveBallToTrack(b, 1, 5)
}

func BenchmarkMoveBallToTrack2TracksMax5(b *testing.B) {
	benchmarkMoveBallToTrack(b, 2, 5)
}

func BenchmarkMoveBallToTrack3TracksMax5(b *testing.B) {
	benchmarkMoveBallToTrack(b, 3, 5)
}

func BenchmarkMoveBallToTrack100TracksMax5(b *testing.B) {
	benchmarkMoveBallToTrack(b, 100, 5)
}

func BenchmarkMoveBallToTrack1TrackMax10(b *testing.B) {
	benchmarkMoveBallToTrack(b, 1, 10)
}

func BenchmarkMoveBallToTrack2TracksMax10(b *testing.B) {
	benchmarkMoveBallToTrack(b, 2, 10)
}

func BenchmarkMoveBallToTrack3TracksMax10(b *testing.B) {
	benchmarkMoveBallToTrack(b, 3, 10)
}

func BenchmarkMoveBallToTrack100TracksMax10(b *testing.B) {
	benchmarkMoveBallToTrack(b, 100, 10)
}

func benchmarkMoveBallToTrack(b *testing.B, trackCount int, max uint8) {
	tracks := make([]*track, trackCount)
	for i := trackCount - 1; i >= 0; i-- {
		tracks[i] = &track{balls: make([]uint, 0), maxBalls: max}
		if i < trackCount-1 {
			tracks[i].nextTrack = tracks[i+1]
		}
	}

	c := &clock{
		reservoir: &track{balls: make([]uint, 0), nextTrack: tracks[0]},
	}
	for n := 0; n < b.N; n++ {
		c.moveBallToTrack(uint(n), tracks[0])
	}
}

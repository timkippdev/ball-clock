package internal

import "testing"

func BenchmarkBallQueueAdd0ExistingElements(b *testing.B) {
	benchmarkBallQueueAdd(b, 0)
}

func BenchmarkBallQueueAdd100ExistingElements(b *testing.B) {
	benchmarkBallQueueAdd(b, 100)
}

func BenchmarkBallQueueAdd10000ExistingElements(b *testing.B) {
	benchmarkBallQueueAdd(b, 10000)
}

func BenchmarkBallQueueAdd1000000ExistingElements(b *testing.B) {
	benchmarkBallQueueAdd(b, 1000000)
}

func benchmarkBallQueueAdd(b *testing.B, elements int) {
	bq := &ballQueue{elements: make([]*ball, elements)}
	for i := 0; i < elements; i++ {
		bq.elements[i] = &ball{id: uint(i)}
	}

	for n := 0; n < b.N; n++ {
		bq.add(&ball{id: uint(n)})
	}
}

func BenchmarkBallQueueCanAdd100ElementsNoMax(b *testing.B) {
	benchmarkBallQueueCanAdd(b, 100, 0)
}

func BenchmarkBallQueueCanAdd1000Elements100Max(b *testing.B) {
	benchmarkBallQueueCanAdd(b, 1000, 100)
}

func benchmarkBallQueueCanAdd(b *testing.B, elements int, max uint8) {
	bq := &ballQueue{elements: make([]*ball, elements)}
	for i := 0; i < elements; i++ {
		bq.elements[i] = &ball{id: uint(i)}
	}

	for n := 0; n < b.N; n++ {
		bq.canAdd()
	}
}

func BenchmarkBallQueueFlush(b *testing.B) {
	bq := &ballQueue{}

	for n := 0; n < b.N; n++ {
		bq.flush()
	}
}

func BenchmarkBallQueueGetNext0Elements(b *testing.B) {
	benchmarkBallQueueGetNext(b, 0)
}

func BenchmarkBallQueueGetNext100Elements(b *testing.B) {
	benchmarkBallQueueGetNext(b, 100)
}

func BenchmarkBallQueueGetNext10000Elements(b *testing.B) {
	benchmarkBallQueueGetNext(b, 10000)
}

func BenchmarkBallQueueGetNext1000000Elements(b *testing.B) {
	benchmarkBallQueueGetNext(b, 1000000)
}

func benchmarkBallQueueGetNext(b *testing.B, elements int) {
	bq := &ballQueue{elements: make([]*ball, elements)}
	for i := 0; i < elements; i++ {
		bq.elements[i] = &ball{id: uint(i)}
	}

	for n := 0; n < b.N; n++ {
		bq.getNext()
	}
}

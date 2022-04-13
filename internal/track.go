package internal

type track struct {
	balls         *ballQueue
	flushDayRatio float64
	name          string
	nextTrack     *track
}

type ball struct {
	id uint
}

type ballQueue struct {
	elements   []*ball
	flushCount uint
	max        uint8
}

func (bq *ballQueue) add(b *ball) {
	bq.elements = append(bq.elements, b)
}

func (bq *ballQueue) canAdd() bool {
	// if no max is set, treat as unlimited storage
	if bq.max == 0 {
		return true
	}
	return uint8(len(bq.elements)) < bq.max-1
}

func (bq *ballQueue) flush() {
	bq.flushCount++
	bq.elements = make([]*ball, 0)
}

func (bq *ballQueue) getNext() *ball {
	if len(bq.elements) == 0 {
		return nil
	}
	next := bq.elements[0]
	bq.elements = bq.elements[1:]
	return next
}

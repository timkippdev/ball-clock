package internal

type track struct {
	balls           []uint
	flushCount      uint
	flushToDayRatio float64
	maxBalls        uint8
	name            string
	nextTrack       *track
}

func (t *track) add(id uint) {
	t.balls = append(t.balls, id)
}

func (t *track) canAdd() bool {
	// if no max is set, treat as unlimited storage
	if t.maxBalls == 0 {
		return true
	}
	return uint8(len(t.balls)) < t.maxBalls-1
}

func (t *track) flush() {
	t.flushCount++
	t.balls = make([]uint, 0)
}

func (t *track) getNext() uint {
	if len(t.balls) == 0 {
		return 0
	}
	next := t.balls[0]
	t.balls = t.balls[1:]
	return next
}

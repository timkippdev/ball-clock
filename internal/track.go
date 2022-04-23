package internal

type trackSlots byte

const (
	slot1 = 1 << iota
	slot2
	slot3
	slot4
)

type track struct {
	balls         []uint8
	trackSlots    trackSlots
	currentLength int
	maxBalls      uint8
	name          string
	nextTrack     *track
}

func (t *track) add(id uint8) {
	t.balls = append(t.balls, id)
	t.currentLength++
}

func (t *track) canAdd() bool {
	// if no max is set, treat as unlimited storage
	if t.maxBalls == 0 {
		return true
	}
	return t.currentLength < int(t.maxBalls-1)
}

func (t *track) flush() {
	t.balls = make([]uint8, 0)
	t.currentLength = 0
}

func (t *track) getNext() uint8 {
	if t.currentLength == 0 {
		return 0
	}
	next := t.balls[0]
	t.balls = t.balls[1:]
	t.currentLength--
	return next
}

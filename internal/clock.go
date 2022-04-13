package internal

type clock struct {
	reservoir *track
}

func (c *clock) flushTrack(t *track) {
	balls := t.balls

	// move existing balls back to reservoir in reverse order
	for i := len(balls.elements) - 1; i >= 0; i-- {
		c.moveBallToTrack(balls.elements[i], c.reservoir)
	}

	// clear the existing balls from the track
	balls.flush()
}

func (c *clock) getCurrentTrackStates() map[string][]uint {
	state := map[string][]uint{}
	addTrackToOutputMap(state, c.reservoir)
	return state
}

func (c *clock) getElapsedDays() float64 {
	elapsedDays := float64(0)
	addTrackToElapsedDays(&elapsedDays, c.reservoir)
	return elapsedDays
}

func (c *clock) tick() {
	// pull next ball from reservoir
	nextBallFromMainTrack := c.reservoir.balls.getNext()
	if nextBallFromMainTrack == nil {
		return
	}

	// move ball to next track
	c.moveBallToTrack(nextBallFromMainTrack, c.reservoir.nextTrack)
}

func (c *clock) moveBallToTrack(b *ball, t *track) {
	balls := t.balls
	if balls.canAdd() {
		balls.add(b)
		return
	}

	// flush track
	c.flushTrack(t)

	// move active ball to next track or back to reservoir
	if t.nextTrack != nil {
		c.moveBallToTrack(b, t.nextTrack)
		return
	}
	c.moveBallToTrack(b, c.reservoir)
}

func addTrackToElapsedDays(elapsedDays *float64, t *track) {
	if t.flushDayRatio != 0 {
		*elapsedDays += (float64(t.balls.flushCount) * t.flushDayRatio)
	}

	if t.nextTrack != nil {
		addTrackToElapsedDays(elapsedDays, t.nextTrack)
	}
}

func addTrackToOutputMap(outputMap map[string][]uint, t *track) {
	ballIds := make([]uint, 0)
	for _, v := range t.balls.elements {
		ballIds = append(ballIds, v.id)
	}
	outputMap[t.name] = ballIds

	if t.nextTrack != nil {
		addTrackToOutputMap(outputMap, t.nextTrack)
	}
}

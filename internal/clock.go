package internal

import (
	"fmt"
	"strings"
)

type clock struct {
	reservoir *track
}

func (c *clock) flushTrack(t *track) {
	balls := t.balls

	// move existing balls back to reservoir in reverse order
	for i := len(balls) - 1; i >= 0; i-- {
		c.moveBallToTrack(balls[i], c.reservoir)
	}

	// clear the existing balls from the track
	t.flush()
}

func (c *clock) getCurrentTrackStates() map[string]string {
	state := map[string]string{}
	track := c.reservoir
	for track != nil {
		state[track.name] = strings.Join(strings.Fields(fmt.Sprintf("%d", track.balls)), ",")
		track = track.nextTrack
	}
	return state
}

func (c *clock) getElapsedDays() float64 {
	elapsedDays := float64(0)
	track := c.reservoir
	for track != nil {
		if track.flushToDayRatio != 0 {
			elapsedDays += (float64(track.flushCount) * track.flushToDayRatio)
		}
		track = track.nextTrack
	}
	return elapsedDays
}

func (c *clock) tick() {
	// pull next ball from reservoir
	nextBallFromMainTrack := c.reservoir.getNext()
	if nextBallFromMainTrack == 0 {
		return
	}

	// move ball to next track
	c.moveBallToTrack(nextBallFromMainTrack, c.reservoir.nextTrack)
}

func (c *clock) moveBallToTrack(id uint8, t *track) {
	// check if ball can be added to track
	if t.canAdd() {
		t.add(id)
		return
	}

	// flush track
	c.flushTrack(t)

	// move active ball to next track or back to reservoir
	if t.nextTrack != nil {
		c.moveBallToTrack(id, t.nextTrack)
		return
	}
	c.moveBallToTrack(id, c.reservoir)
}

package clock

import (
	"sort"
)

type tray struct {
	max 		int
	balls 		[]int
	nextTray 	*tray
}

// TODO: Move Tray to it's own package
type BallClockTray interface {
	Add(queue []int) []int
	IsFull() bool
}

type RollingBallClock struct {
	ballsOnInit int
	queue 		[]int 
	minutes 	*tray 
	fiveMinutes *tray 
	hours 		*tray
}

func (t *tray) IsFull() bool {
	return len(t.balls) == t.max
}

func (t *tray) Add(ball int) []int {
	if t.IsFull() {
		reverse(t.balls)
		toAppend := t.balls[:]
		t.balls = nil
		if t.nextTray == nil {
			return append(toAppend, ball)
		}
		return append(toAppend, t.nextTray.Add(ball)...)

	} else {
		t.balls = append(t.balls, ball)
		return nil
	}
}

func New(balls int) *RollingBallClock { 
	var q []int
	for i := 1; i <= balls; i++ {
		q = append(q, i)
	}
	hourTray := tray{
		max: 11,
	}
	fiveMinuteTray := tray{
		max: 11,
		nextTray: &hourTray,
	}
	minuteTray := tray{
		max: 4,
		nextTray: &fiveMinuteTray,
	}
	return &RollingBallClock{
		ballsOnInit: balls, 
		queue: q,
		minutes: &minuteTray,
		fiveMinutes: &fiveMinuteTray,
		hours: &hourTray,
	}
}

// TODO: Figure out how to cleanly return a simple int using a channel
func (c *RollingBallClock) FindCycleDays() int {
	return c.findCycleDays(0)
}

func (c *RollingBallClock) findCycleDays(minutes int) int {

	toAppend := c.minutes.Add(c.queue[0])
	minutes += 1 // increment minute tracker
	
	// Delete ball from queue after adding it to another tray and add any balls that need to be appended to the queue
	c.queue = append(c.queue[1:], toAppend...)

	if c.matchesInitQueue() {
		return minutes / 60 / 24
	}
	return c.findCycleDays(minutes)
}

func (c *RollingBallClock) matchesInitQueue() bool {
	return len(c.queue) == c.ballsOnInit && sort.IsSorted(sort.IntSlice(c.queue))
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

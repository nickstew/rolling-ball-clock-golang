package clock

import "sort"

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

type ClockCounter interface {
	FindCycleDays() int
}

type QueueStateTracker interface {
	IsQueueUnchanged() bool
}

func (t *tray) IsFull() bool {
	return len(t.balls) == t.max
}

// TODO: take in the ball to add and return anything to be appended to the queue
func (t *tray) Add(queue []int) []int {
	if t.IsFull() {
		reverse(t.balls)
		queue = append(queue, t.balls...)
		t.balls = nil
		if t.nextTray == nil {
			return append(queue[1:], queue[0])
		}
		return t.nextTray.Add(queue)

	} else {
		t.balls = append(t.balls, queue[0])
		if(len(queue) > 1) {
			return queue[1:]
		}
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

// TODO: Figure out how to cleanly return a simple int using a channel instaed of bloating memory with all these GOROUTINES
func (c *RollingBallClock) FindCycleDays() int {
	return c.findCycleDays(0)
}

func (c *RollingBallClock) findCycleDays(minutes int) int {
	c.queue = c.minutes.Add(c.queue)
	minutes += 1
	if c.IsQueueUnchanged() {
		return minutes / 60 / 24
	}
	return c.findCycleDays(minutes)
}

func (c *RollingBallClock) IsQueueUnchanged() bool {
	return len(c.queue) == c.ballsOnInit && sort.IsSorted(sort.IntSlice(c.queue))
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

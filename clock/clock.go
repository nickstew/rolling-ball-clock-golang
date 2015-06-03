package clock

import (
	"fmt"
	"sort"
)

type tray struct {
	index    int
	max      int
	balls    []int
	nextTray *tray
}

type RollingBallClock struct {
	ballsOnInit int
	qIndex      int
	queue       []int
	minutes     *tray
	fiveMinutes *tray
	hours       *tray
}

func (t *tray) IsFull() bool {
	return t.index+1 == t.max
}

func (t *tray) Add(ball int) []int {
	if t.IsFull() {
		reverse(t.balls)
		toAppend := t.balls[:]
		for i := 0; i < t.max; i += 1 {
			t.balls[i] = 0
		}
		t.index = 0
		if t.nextTray == nil {
			return append(toAppend, ball)
		}
		return append(toAppend, t.nextTray.Add(ball)...)

	} else {
		t.balls[t.index] = ball
		t.index += 1
		return nil
	}
}

func New(balls int) *RollingBallClock {
	q := make([]int, balls)
	for i := 1; i <= balls; i++ {
		q[i-1] = i
	}
	index := balls - 1
	minBalls := make([]int, 4)
	fiveBalls := make([]int, 11)
	hBalls := make([]int, 11)

	hourTray := tray{
		max:   11,
		balls: hBalls,
	}
	fiveMinuteTray := tray{
		max:      11,
		balls:    fiveBalls,
		nextTray: &hourTray,
	}
	minuteTray := tray{
		max:      4,
		balls:    minBalls,
		nextTray: &fiveMinuteTray,
	}
	return &RollingBallClock{
		ballsOnInit: balls,
		qIndex:      index,
		queue:       q,
		minutes:     &minuteTray,
		fiveMinutes: &fiveMinuteTray,
		hours:       &hourTray,
	}
}

func (c *RollingBallClock) FindCycleDays() int {
	result := make(chan int)
	go c.findCycleDays(0, result)
	answer := <-result
	return answer
}

func (c *RollingBallClock) findCycleDays(minutes int, result chan int) {

	toAppend := c.minutes.Add(c.queue[0])

	minutes += 1 // increment minute tracker
	if (minutes)%(60*24) == 0 && 378*24*60 == minutes {
		fmt.Printf("day: %v\n", minutes/60/24)
		return
	}
	// Delete ball from queue after adding it to another tray and add any balls that need to be appended to the queue
	// 1. delete
	for i := 0; i < c.qIndex; i = i + 1 {
		c.queue[i] = c.queue[i+1]
	}
	c.queue[c.qIndex] = 0
	c.qIndex -= 1

	//c.queue = c.queue[1:]
	// 2. add each ball to queue
	for i := 0; i < len(toAppend); i, c.qIndex = i+1, c.qIndex+1 {
		//fmt.Println(c.qIndex)
		c.queue[c.qIndex] = toAppend[i]
	}

	//c.queue = append(c.queue[1:], toAppend...)

	if c.qIndex+1 == c.ballsOnInit && sort.IsSorted(sort.IntSlice(c.queue)) {
		result <- (minutes / 60 / 24)
		return
	}
	c.findCycleDays(minutes, result)
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

package meetup

import (
	"time"

	"github.com/thethanos/go-containers/containers"
)

type MeetupQueuer struct {
	returnChan chan<- Meetup
	queue      *containers.Heap[Meetup]
}

type Meetup struct {
	Time    time.Time
	Users   []string
	Message string
}

func NewMeetupQueuer(returnChan chan<- Meetup, queue *containers.Heap[Meetup]) MeetupQueuer {
	queuer := MeetupQueuer{
		returnChan: returnChan,
		queue:      queue,
	}

	return queuer
}

func (queuer *MeetupQueuer) RunMeetupQueue() {
	for {
		if !queuer.queue.Empty() && time.Now().After(queuer.queue.Top().Time) {
			queuer.returnChan <- queuer.queue.Top()
			queuer.queue.Pop()
		}
	}
}

func (queuer *MeetupQueuer) PushMeetup(meetup Meetup) {
	queuer.queue.Push(meetup)
}

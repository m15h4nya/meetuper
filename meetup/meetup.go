package meetup

import (
	"time"

	"github.com/thethanos/go-containers/containers"
)

type MeetupQueuer struct {
	returnChan chan<- Meetup
	queue      *containers.Heap[Meetup]
	drafts     map[string]Meetup
}

type Meetup struct {
	Name    string
	Message string
	Users   []string
	Time    time.Time
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

func (queuer *MeetupQueuer) GetDraft(id string) Meetup {
	return queuer.drafts[id]
}

func (queuer *MeetupQueuer) UpdateDraft(id string, updated Meetup) {
	queuer.drafts[id] = updated
}

func (queuer *MeetupQueuer) PushMeetup(id string) {
	queuer.queue.Push(queuer.drafts[id])
	delete(queuer.drafts, id)
}

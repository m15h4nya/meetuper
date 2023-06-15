package meetup

import (
	"time"

	"github.com/bwmarrin/discordgo"
	containers "github.com/m15h4nya/meetupper/heap"
)

type MeetupQueuer struct {
	returnChan chan<- Meetup
	queue      *containers.Heap[Meetup]
	drafts     map[string]Meetup
}

type Meetup struct {
	Name     string
	Message  string
	Users    discordgo.MessageComponentInteractionDataResolved
	Start    time.Time
	Interval time.Duration
}

func NewMeetupQueuer(returnChan chan<- Meetup, queue *containers.Heap[Meetup]) MeetupQueuer {
	queuer := MeetupQueuer{
		returnChan: returnChan,
		queue:      queue,
		drafts:     map[string]Meetup{},
	}

	return queuer
}

func (q *MeetupQueuer) RunMeetupQueue() {
	for {
		if !q.queue.Empty() && time.Now().After(q.queue.Top().Start) {
			q.returnChan <- q.queue.Top()
			q.queue.Pop()
		}
	}
}

func (queuer *MeetupQueuer) NewDraft(id string) Meetup {
	queuer.drafts[id] = Meetup{}
	return queuer.drafts[id]
}

func (queuer *MeetupQueuer) GetDraft(id string) Meetup {
	return queuer.drafts[id]
}

func (queuer *MeetupQueuer) UpdateDraft(id string, updated Meetup) {
	queuer.drafts[id] = updated
}

func (queuer *MeetupQueuer) PushDraft(id string) {
	queuer.queue.Push(queuer.drafts[id])
	delete(queuer.drafts, id)
}

func (queuer *MeetupQueuer) DeleteDraft(id string) {
	delete(queuer.drafts, id)
}

func (queuer *MeetupQueuer) PushMeetup(meetup Meetup) {
	queuer.queue.Push(meetup)
}

func (queuer *MeetupQueuer) DeleteFromQueue(name string) {
	newQueue := containers.NewHeap(
		func(a, b Meetup) bool {
			return b.Start.Before(a.Start)
		})
	for !queuer.queue.Empty() {
		if meetup := queuer.queue.Top(); meetup.Name != name {
			newQueue.Push(meetup)
		}
		queuer.queue.Pop()
	}
	queuer.queue = &newQueue
}

func (queuer *MeetupQueuer) GetAllMeetups() []Meetup {
	return queuer.queue.GetData()
}

func (m Meetup) AddDuration(duration time.Duration) Meetup {
	m.Start = m.Start.Add(duration)
	return m
}

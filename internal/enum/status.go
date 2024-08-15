package enum

type Status int

const (
	StatusPending = iota
	StatusDone
)

var statusName = map[Status]string{
	StatusPending: "pending",
	StatusDone:    "done",
}

func (s Status) String() string {
	return statusName[s]
}

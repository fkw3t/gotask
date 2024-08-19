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

func NewStatusFromInt(i int) Status {
	return Status(i)
}

func (s Status) String() string {
	return statusName[s]
}

func (s Status) Int() int {
	return int(s)
}

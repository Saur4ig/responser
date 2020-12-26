package responser

type Status string

const (
	StatusError   Status = "error"
	StatusSuccess Status = "ok"
)

func (s Status) String() string {
	return string(s)
}

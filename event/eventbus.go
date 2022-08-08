package event

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

type Bus interface {
	Subscribe(event string, priority Priority, fn func(*Event))
	Unsubscribe(event string, fn func(*Event))
	Publish(event *Event)
}

type Event struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

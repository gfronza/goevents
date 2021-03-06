package messaging

type EventHandler func(body []byte) bool

type Consumer interface {
	Subscribe(action string, handler EventHandler) error
	Unsubscribe(action string) error
	Consume()
	Close()
}

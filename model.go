package starsystem

type Actor interface {
	Receive(msg interface{}, context *ActorProp)
}

type Message interface {
	Name() string
	Params() map[string][]byte
}

package starsystem

type State interface {
	Execute(msg interface{}, context *ActorProp)
}

type NormalState struct{}

func (n NormalState) Execute(msg interface{}, context *ActorProp) {
	context.state.Execute(msg, context)
}

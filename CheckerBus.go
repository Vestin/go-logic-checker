package go_checker

type CheckerBus struct {
	dispatcher Dispatcher
}

func (cb *CheckerBus) AddChecker(c Checker) *CheckerBus {
	cb.dispatcher.AddChecker(c)
	return cb
}

func (cb *CheckerBus) SetDispatcher(d Dispatcher) {
	cb.dispatcher = d
}

func (cb *CheckerBus) Check() bool {
	return cb.dispatcher.Check()
}

func (cb *CheckerBus) GetError() []string {
	return cb.dispatcher.GetError()
}
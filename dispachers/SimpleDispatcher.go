package dispachers

import "go-checker"

type SimpleDispatcher struct {
	checkers []go_checker.Checker
	error    string
}

func (sp *SimpleDispatcher) AddChecker(c go_checker.Checker) {
	sp.checkers = append(sp.checkers, c)
}

func (sp *SimpleDispatcher) Check() bool {
	for _, c := range sp.checkers {
		if c.Check() {
			continue
		}
		sp.error = c.GetError()
		return false
	}
	return true
}

func (sp *SimpleDispatcher) GetError() []string {
	return []string{sp.error}
}

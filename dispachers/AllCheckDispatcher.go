package dispachers

import "go-checker"

type AllCheckDispatcher struct {
	checkers []go_checker.Checker
	errors   []string
}

func (sp *AllCheckDispatcher) AddChecker(c go_checker.Checker) {
	sp.checkers = append(sp.checkers, c)
}

func (sp *AllCheckDispatcher) Check() bool {
	pass := true
	for _, c := range sp.checkers {
		if c.Check() {
			continue
		}else{
			sp.errors = append(sp.errors, c.GetError())
			pass = false
		}
	}
	return pass
}

func (sp *AllCheckDispatcher) GetError() []string {
	return sp.errors
}

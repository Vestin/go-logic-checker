package go_checker

type Dispatcher interface {
	AddChecker(checker Checker)
	Check() bool
	GetError() []string
}
package go_checker

type Checker interface {
    Check() bool
	GetError() string
}

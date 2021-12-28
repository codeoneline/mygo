package mygo

type Subscription interface {
	Unsubscribe()
	Err() <-chan error
}

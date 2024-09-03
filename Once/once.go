package once

type Once struct {
	lock chan bool
	done bool
}

func New() *Once {
	return &Once{
		lock: make(chan bool, 1),
		done: false,
	}
}

func (o *Once) Do(f func()) {
	o.lock <- true
	defer func() {
		o.done = true
		<-o.lock
		if r := recover(); r != nil {
			panic(r)
		}
	}()
	if !o.done {
		f()
	}
}

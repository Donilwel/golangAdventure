package rwmutex

type RWMutex struct {
	lock chan int
}

func New() *RWMutex {
	var rw RWMutex
	rw.lock = make(chan int, 1)
	rw.lock <- 0
	return &rw
}

func (rw *RWMutex) RLock() {
	for {
		rwm := <-rw.lock
		if rwm != -1 {
			rwm++
			rw.lock <- rwm
			break
		}
		rw.lock <- rwm
	}
}

func (rw *RWMutex) RUnlock() {
	rwm := <-rw.lock
	rwm--
	rw.lock <- rwm
}

func (rw *RWMutex) Lock() {
	for {
		rwm := <-rw.lock
		if rwm == 0 {
			rwm = -1
			rw.lock <- rwm
			break
		}
		rw.lock <- rwm
	}
}

func (rw *RWMutex) Unlock() {
	rwm := <-rw.lock
	rwm = 0
	rw.lock <- rwm
}

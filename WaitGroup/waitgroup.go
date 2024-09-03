package waitgroup

type WaitGroup struct {
	lock      chan bool
	waitGroup chan int
}

func New() *WaitGroup {
	var wg WaitGroup
	wg.waitGroup = make(chan int, 1)
	wg.lock = make(chan bool, 1)
	wg.waitGroup <- 0
	return &wg
}

func (wg *WaitGroup) Add(delta int) {
	wg.lock <- true
	waitGroup := <-wg.waitGroup
	waitGroup += delta
	if waitGroup < 0 {
		panic("negative WaitGroup counter")
	}
	wg.waitGroup <- waitGroup
	<-wg.lock
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}
func (wg *WaitGroup) Wait() {
	for {
		wg.lock <- true
		waitGroup := <-wg.waitGroup
		if waitGroup == 0 {
			wg.waitGroup <- waitGroup
			<-wg.lock
			break
		}
		wg.waitGroup <- waitGroup
		<-wg.lock
	}
}

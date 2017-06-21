package goroutine

// Go way
const (
	GoWayRoutine = 1
	GoWayChan    = 2
)

type work struct {
	callback func()
	f        func(callback func())
}

// GoRoutine class
type GoRoutine struct {
	way int
	ch  chan *work
}

// New object
func New(way, chanSize int) *GoRoutine {
	g := &GoRoutine{
		way: way,
	}

	if way == GoWayChan {
		g.ch = make(chan *work, chanSize)
		for i := 0; i < chanSize; i++ {
			go g.worker()
		}
	}

	return g
}

func (g *GoRoutine) worker() {
	for w := range g.ch {
		w.f(w.callback)
	}
}

// Go asynchronous running
func (g *GoRoutine) Go(callback func(), f func(callback func())) {
	switch g.way {
	case GoWayRoutine:
		go f(callback)
	case GoWayChan:
		g.ch <- &work{
			callback: callback,
			f:        f,
		}
	}
}

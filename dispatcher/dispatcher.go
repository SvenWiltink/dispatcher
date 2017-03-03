package dispatcher

type Job interface {
	Execute()
}

type JobDispatcher struct {
	Jobs chan Job
	stop chan bool

	bufferSize  int
	threadCount int
}

func (j *JobDispatcher) Start() {
	j.addRunners()
}

func (j *JobDispatcher) addRunners() {
	for i := 0; i < j.threadCount; i++ {
		runner := NewRunner(i, j)
		go runner.start()
	}
}

func NewJobDispatcher(threadCount int, bufferSize int) *JobDispatcher {
	return &JobDispatcher{
		Jobs:        make(chan Job, bufferSize),
		stop:        make(chan bool),
		threadCount: threadCount,
	}
}

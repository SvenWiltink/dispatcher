package dispatcher

type Runner struct {
	id   int
	jobs chan Job
	stop chan bool
}

func (r *Runner) start() {
	for {
		select {
		case job, more := <-r.jobs:
			if more {
				job.Execute()
			}
		}
	}
}

func NewRunner(id int, listener *JobDispatcher) (runner *Runner) {
	runner = &Runner{
		id:   id,
		jobs: listener.Jobs,
		stop: listener.stop,
	}

	return
}

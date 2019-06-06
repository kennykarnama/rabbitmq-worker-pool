package pool

import "log"

//Job represents the job payload
type Job struct {
	ID        int32
	Resources string
}

//Pool represents the worker pool structure
type Pool struct {
	NumWorkers  int32
	JobChannels chan chan Job
	JobQueue    chan Job
	Stopped     chan bool
}

//Worker represents the actual worker who does the job
type Worker struct {
	ID          int
	JobChannel  chan Job
	JobChannels chan chan Job
	Quit        chan bool
}

//NewPool is a function to construct
//WorkerPool object
func NewPool(numworkers int32) Pool {
	return Pool{
		NumWorkers:  numworkers,
		JobChannels: make(chan chan Job),
		JobQueue:    make(chan Job),
		Stopped:     make(chan bool),
	}
}

//Run is a function to start the pool
func (p *Pool) Run() {
	log.Println("Spawning the workers")
	for i := 0; i < int(p.NumWorkers); i++ {
		worker := Worker{
			ID:          (i + 1),
			JobChannel:  make(chan Job),
			JobChannels: p.JobChannels,
			Quit:        make(chan bool),
		}
		worker.Start()
	}
	p.Allocate()
}

//Allocate is a function to pull
//from the queue and send the job to the channel
func (p *Pool) Allocate() {
	q := p.JobQueue
	s := p.Stopped
	go func(queue chan Job) {
		for {
			select {
			case job := <-q:
				// get from the JobChannels
				availChannel := <-p.JobChannels
				availChannel <- job

			case <-s:
				return
			}
		}
	}(q)
}

//Start the worker
func (w *Worker) Start() {
	log.Printf("Starting Worker ID [%d]", w.ID)
	go func() {
		for {
			w.JobChannels <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				w.work(job)
			case <-w.Quit:
				return
			}

		}
	}()
}

//work is the task performer
func (w *Worker) work(job Job) {
	log.Printf("------")
	log.Printf("Processed by Worker [%d]", w.ID)
	log.Printf("Processed Job With ID [%d] & content: [%s]", job.ID, job.Resources)
	log.Printf("-------")
}

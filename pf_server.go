package main
import("fmt"
				"net"
				"os")
				
const(
	MaxJob = 60
	MaxWorkPool = 10
)
				
type Job struct {
	Cn net.Conn
	Recvbuf []byte
}

type JobQueue chan Job
var jobqueue JobQueue = make(JobQueue,MaxJob)	
type Worker struct {
	WorkerPool chan chan Job
	Jobchan chan Job
	Quit chan bool
}

type Dispatcher struct {
	WorkPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher{
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkPool :pool}
}

func (d *Dispatcher) Run() {
	
	for i := 0 ; i < MaxWorkPool; i++ {
		worker := NewWorker(d.WorkPool)
		worker.Start()
	}
	go d.Dispatch()
}

func (d *Dispatcher)Dispatch() {
	for {
		select {
			case job := <-jobqueue :
					go func(job Job) {
						jobchan := <-d.WorkPool
						jobchan <-job					
					}(job)
		}
	}
} 

func NewWorker(workpools chan chan Job) Worker{
	return Worker{WorkerPool :workpools,
								Jobchan :make(chan Job),
								Quit : make(chan bool)}
}

func respond(job Job) {
	job.Cn.Write([]byte("hi,i am tok,nice to meet you!"))
	
}

/* start work handler process */
func (work Worker) Start() {
	go func() {
		for {
			work.WorkerPool <- work.Jobchan
			select {
			case job := <- work.Jobchan :
				// execute job
				fmt.Println(job)
				respond(job)
			case quit := <- work.Quit :
				// exit routine
				fmt.Println(quit)
				return 
			}
			
		}
	}()
}

/* stop worker */
func (work  Worker) Stop() {
	go func() {
		work.Quit<- true
	}()
}

func main() {
	//read close chan bool, value is false
	//n := make(chan bool)
	//close(n)
	//w := <- n
	//fmt.Println(w)

	var d *Dispatcher ;
	d = NewDispatcher(MaxWorkPool)
	
	listener , err := net.Listen("tcp","10.10.10.58:6688")
	if err != nil {
		fmt.Printf("create listener is error, errno:%v",err)
		os.Exit(-1)
	}
	for {
		cn , err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept is error , errno:%v",err)
			continue
		}
		var j Job
		j.Cn = cn
		j.Recvbuf = make([]byte,4095)
		rcounts , err := j.Cn.Read(j.Recvbuf)
		if err != nil {
			fmt.Printf("read client data is error, errno :%v",err)
		}
		defer j.Cn.Close()
		//fmt.Printf("read count:%d,data:%s",rcounts,j.Recvbuf)
		jobqueue<-j
		d.Run()
		
	}

}



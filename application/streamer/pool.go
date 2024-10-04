package streamer

import "fmt"


type VideoDispatcher struct {

  WorkerPool chan chan VideoProcessingJob
  maxWorkers int
  jobQueue chan VideoProcessingJob
  Processor Processor
}


type videoWorker struct {
  id int
  jobQueue chan VideoProcessingJob
  workerPool chan chan VideoProcessingJob
}


func newVideoWorker(id int, workerPool chan chan VideoProcessingJob) videoWorker {
  fmt.Println("newVideoWorker(): creating worker id:", id)
  return videoWorker{
    id:id,
    jobQueue: make(chan VideoProcessingJob),
    workerPool: workerPool,
    
  }
}

func (vd  *VideoDispatcher) Run()  {
  fmt.Println("vd.Run(): starting worker pool")
  for i := 0; i < vd.maxWorkers; i++{
    fmt.Println("vd.Run(): starting worker:", i)
    worker := newVideoWorker(i, vd.WorkerPool)
    worker.start()
  }

  go vd.dispatch()
}

func (vd  *VideoDispatcher) dispatch()  {
  for {
    job := <-vd.jobQueue
    fmt.Println("vd.dispatch(): sending job ", job.Video.ID," to job queue")
    go func() {
      workerJobQueue := <-vd.WorkerPool
      workerJobQueue <- job
    }()
  }
}

func (w  videoWorker) start()  {
  fmt.Println("w.start(): starting worker:", w.id)
  go func() {
    for {
      w.workerPool <- w.jobQueue
      job := <-w.jobQueue
      w.processVideoJob(job.Video)
    }
  }()
}

func (w  videoWorker) processVideoJob(video Video)  {
  fmt.Println("w.processVideoJob(): start encode for: ", video.ID)
  video.encode()
}

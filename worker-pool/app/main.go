package main

import (
	"fmt"
	"streamer"
)

func main() {

	const numJobs = 4
	const numWorkers = 4

	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	wp := streamer.New(videoQueue, numWorkers)

	wp.Run()

	fmt.Println("started press enter to continue")
	_, _ = fmt.Scanln()


	ops := &streamer.VideoOptions{
		SegmentDuration: 10,
		MaxRate1080p:    "1200k",
		MaxRate720p:     "600k",
		MaxRate480p:     "600k",
		RenameOutput:    true,
	}

	v1 := wp.NewVideo(1, "input/puppy1.mp4", "output", "mp4", notifyChan, nil)
	v2 := wp.NewVideo(2, "input/bad.txt", "output", "mp4", notifyChan, nil)
	v3 := wp.NewVideo(3, "input/puppy2.mp4", "output", "hls", notifyChan, ops)
	v4 := wp.NewVideo(4, "input/puppy2.mp4", "output", "mp4", notifyChan, nil)
  

	videoQueue <- streamer.VideoProcessingJob{Video: v1}
	videoQueue <- streamer.VideoProcessingJob{Video: v2}
	videoQueue <- streamer.VideoProcessingJob{Video: v3}
	videoQueue <- streamer.VideoProcessingJob{Video: v4}

	for i := 0; i < numJobs; i++ {
		msg := <-notifyChan
		fmt.Println("i: ", i, "msg:", msg)
	}

	fmt.Println("done")

}

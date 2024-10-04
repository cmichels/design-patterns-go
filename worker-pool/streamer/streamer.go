package streamer

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/tsawler/toolbox"
)

type ProcessingMessage struct {
	ID         int
	Successful bool
	Message    string
	OutputFile string
}

type VideoProcessingJob struct {
	Video Video
}

type Processor struct {
	Engine Encoder
}

type Video struct {
	ID           int
	InputFile    string
	OutputDir    string
	EncodingType string
	NotifyChan   chan ProcessingMessage
	Options      *VideoOptions
	Encoder      Processor
}

func (v *Video) encode() {

  fmt.Println("v.encode(): encoding to mp4", v.ID)
	var fileName string
	switch v.EncodingType {
	case "mp4":
		name, err := v.encodeToMP4()
		if err != nil {
			v.sendToNotifyChan(false, "", fmt.Sprintf("encode failed for %d: %s", v.ID, err))
			return
		}

		fileName = fmt.Sprintf("%s.mp4", name)

	case "hls":
		name, err := v.encodeToHLS()
		if err != nil {
			v.sendToNotifyChan(false, "", fmt.Sprintf("encode failed for %d: %s", v.ID, err))
			return
		}

		fileName = fmt.Sprintf("%s.m3u8", name)
	default:
    fmt.Println("v.encode(): error:", v.ID)
		v.sendToNotifyChan(false, "", fmt.Sprintf("error processing for %d: invalid encoding type", v.ID))
    return
	}

  fmt.Println("v.encode(): success for id:", v.ID)
	v.sendToNotifyChan(true, fileName, fmt.Sprintf("video %d saved as %s", v.ID, fmt.Sprintf("%s/%s", v.OutputDir, fileName)))
}

func (v *Video) encodeToMP4() (string, error) {
	baseFileName := ""

  fmt.Println("v.encodeToMP4(): encoding:", v.ID)
	if !v.Options.RenameOutput {
		b := path.Base(v.InputFile)
		baseFileName = strings.TrimSuffix(b, filepath.Ext(b))
	} else {
    var t toolbox.Tools
    baseFileName = t.RandomString(10)

	}

	err := v.Encoder.Engine.EncodeToMP4(v, baseFileName)

	if err != nil {
		return "", err
	}


  fmt.Println("v.encodeToMP4(): success for id:", v.ID)

	return baseFileName, nil

}

func (v *Video) encodeToHLS() (string, error) {
 		 
	baseFileName := ""

  fmt.Println("v.encodeToHLS(): encoding:", v.ID)
	if !v.Options.RenameOutput {
		b := path.Base(v.InputFile)
		baseFileName = strings.TrimSuffix(b, filepath.Ext(b))
	} else {
    var t toolbox.Tools
    baseFileName = t.RandomString(10)
	}
  
  err := v.Encoder.Engine.EncodeToHLS(v, baseFileName)


  if err != nil {
    return "", err
  }

  return baseFileName, nil


}

func (v *Video) sendToNotifyChan(succesful bool, fileName, message string) {
  fmt.Println("v.sendToNotifyChan(): sending message for id:", v.ID)
	v.NotifyChan <- ProcessingMessage{
		ID:         v.ID,
		Successful: succesful,
		Message:    message,
		OutputFile: fileName,
	}
}

type VideoOptions struct {
	RenameOutput    bool
	SegmentDuration int
	MaxRate1080p    string
	MaxRate720p     string
	MaxRate480p     string
}

func (vd *VideoDispatcher) NewVideo(id int, input, output, encodingType string, notifyChan chan ProcessingMessage, ops *VideoOptions) Video {
	if ops == nil {
		ops = &VideoOptions{}
	}


  fmt.Println("new video:", id, input)

	return Video{
		ID:           id,
		InputFile:    input,
		OutputDir:    output,
		EncodingType: encodingType,
		NotifyChan:   notifyChan,
		Encoder:      vd.Processor,
		Options:      ops,
	}
}

func New(jobQueue chan VideoProcessingJob, maxWorkers int) *VideoDispatcher {

  fmt.Println("New(): creating worker pool")
	workerPool := make(chan chan VideoProcessingJob, maxWorkers)

	var engine VideoEncoder
	p := Processor{Engine: &engine}

	return &VideoDispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		WorkerPool: workerPool,
		Processor:  p,
	}
}



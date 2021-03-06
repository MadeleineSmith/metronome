package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	//go:embed pop.mp3
	res embed.FS
)

func main() {
	println("Metronome has started!")

	buffer := initializeBuffer()

	bpm, bpb := retrieveBeatsInput()

	initializeMetronome(bpm, bpb, buffer)
}

func initializeBuffer() *beep.Buffer {
	audioFile, _ := res.Open("pop.mp3")

	streamer, format, err := mp3.Decode(audioFile)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	return buffer
}

func retrieveBeatsInput() (float64, int) {
	reader := bufio.NewReader(os.Stdin)

	print("Beats Per Minute (default 60): ")
	bpmInput, _ := reader.ReadString('\n')
	bpm := 60.0
	if bpmInput != "\n" {
		bpm, _ = strconv.ParseFloat(strings.TrimRight(bpmInput, "\n"), 64)
	}

	print("Beats Per Bar (default 4): ")
	bpbInput, _ := reader.ReadString('\n')
	bpb := 4
	if bpbInput != "\n" {
		bpb, _ = strconv.Atoi(strings.TrimRight(bpbInput, "\n"))
	}

	return bpm, bpb
}

func initializeMetronome(bpm float64, bpb int, buffer *beep.Buffer) {
	d := time.Duration(float64(time.Minute) / bpm)
	fmt.Println("Delay:", d)

	t := time.NewTicker(d)
	i := 1

	for _ = range t.C {
		i--
		if i == 0 {
			i = bpb
			fmt.Printf("\nTICK ")

			pop := buffer.Streamer(0, buffer.Len())

			louderPop := &effects.Volume{
				Streamer: pop,
				Base:     1.5,
				Volume:   1,
				Silent:   false,
			}
			speaker.Play(louderPop)

		} else {
			fmt.Printf("tick ")

			pop := buffer.Streamer(0, buffer.Len())
			speaker.Play(pop)
		}
	}
}

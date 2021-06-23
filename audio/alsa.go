//+build linux

package audio

import (
	"errors"
	"github.com/albanseurat/goalsa"
)

type AlsaSteam struct {
	out    []int16
	device *goalsa.PlaybackDevice
}

func NewStream() Stream {
	return &AlsaSteam{}
}

func (s *AlsaSteam) Init() error {
	var err error
	if s.device, err = goalsa.NewPlaybackDevice("pcm.default", 2, goalsa.FormatS16LE, 44100, goalsa.BufferParams{})
		err != nil {
		return err
	}
	return nil
}

func (s *AlsaSteam) Close() error {
	s.device.Close()
	return nil
}

func (s *AlsaSteam) Start() error {
	return nil
}

func (s *AlsaSteam) Stop() error {
	return s.device.Drop()
}

func (s *AlsaSteam) Write(output []int16) error {
	ret , err := s.device.Write(output)
	if err == goalsa.ErrUnderrun {
		// the lib, after an underrun re-prepare the stream, skip the frame
		return nil
	}
	if err != nil {
		return err
	}
	if ret != len(output) {
		return errors.New("number of bytes written is not good")
	}
	return nil
}

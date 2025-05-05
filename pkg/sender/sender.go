package sender

import (
	"time"

	"github.com/Hundemeier/go-sacn/sacn"
)

type Sender struct {
	AddressDst     string
	AddressSrc     string
	Channels       uint16
	Duration       uint64
	Fps            uint32
	Frequency      float64
	IntesintyLower uint8
	IntesintyUpper uint8
	Priority       uint8
	Universes      []uint
}

func Default() Sender {
	return Sender{
		AddressDst:     "",
		AddressSrc:     "",
		Channels:       512,
		Duration:       0,
		Fps:            1,
		Frequency:      1.0,
		IntesintyLower: 0,
		IntesintyUpper: 255,
		Priority:       100,
		Universes:      []uint{1},
	}
}

func (s Sender) Run() error {
	transmitter, err := sacn.NewTransmitter("", [16]byte{1, 2, 3}, "beam CLI utility")
	if err != nil {
		return err
	}

	channel, err := transmitter.Activate(1)
	if err != nil {
		return err
	}
	defer close(channel)

	if s.AddressDst == "" {
		transmitter.SetMulticast(1, true)
	} else {
		transmitter.SetMulticast(1, false)
		transmitter.SetDestinations(1, []string{s.AddressDst})
	}

	for i := range 10 {
		channel <- []byte{byte(s.IntesintyUpper), byte(i & 0xFF)}
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

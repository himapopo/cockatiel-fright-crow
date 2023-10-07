package sound

import "github.com/ebitengine/oto/v3"

type AudioContext struct {
	Ctx *oto.Context
}

func NewAudioContext() *AudioContext {
	op := &oto.NewContextOptions{}

	op.SampleRate = 44100

	op.ChannelCount = 2

	op.Format = oto.FormatSignedInt16LE

	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan

	return &AudioContext{
		Ctx: otoCtx,
	}
}

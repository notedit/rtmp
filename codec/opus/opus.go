package opus

type Codec struct {
	SampleRate int
	Channels   int
}

func NewCodec() *Codec {
	return &Codec{}
}

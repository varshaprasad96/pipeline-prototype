package node

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-logr/logr"
	"github.com/testpipeline/entity"
	channellib "gopkg.in/eapache/channels.v1"
)

type Event struct {
	owner         string
	creationtime  time.Time
	srcId         entity.Identity
	destinationId entity.Identity
	metadata      map[string]string
	eventID       int
	Logger        logr.Logger
}

// =============================== Producer ==============================================

type Producer struct {
	logr.Logger
	Event
	Data           interface{}
	sendingChannel channellib.Channel
}

var _ entity.Node = &Producer{}

func (p *Producer) InjectChannel(ch channellib.Channel, sender bool) error {
	if !sender {
		return fmt.Errorf("cannot inject a non-sending channel into producer.")
	}
	p.sendingChannel = ch
	return nil
}

func (p *Producer) InjectLogger(log logr.Logger) {
	p.Logger = log
}

func (p *Producer) GetEvent(id string) error {
	// TODO
	return nil
}

func (p *Producer) Run() error {
	// test
	for v := 0; v < 4; v++ {
		p.sendingChannel.In() <- p.Data
	}
	p.sendingChannel.Close()
	return nil
}

func (p *Producer) GetBufferedOutputChannel() (channellib.Channel, error) {
	if p == nil || p.sendingChannel == nil {
		return nil, fmt.Errorf("could not fetch the output channel")
	}
	return p.sendingChannel, nil
}

func (p *Producer) Out() []interface{} {
	res := make([]interface{}, 0)
	for val := range p.sendingChannel.Out() {
		res = append(res, val)
	}
	return res
}

// move the event fields into options
func NewProducer[T any](data entity.Contents[T], opts entity.Options) Producer {
	producer := Producer{}

	inChannel := channellib.NewInfiniteChannel()
	producer.sendingChannel = inChannel
	producer.Data = data
	// TODO: fill in rest of the event fields with what the user provides,
	producer.Event = Event{
		owner:         opts.Owner,
		srcId:         entity.Identity(opts.SrcId),
		destinationId: entity.Identity(opts.DestId),
		metadata:      opts.Metadata,
		eventID:       rand.Intn(100),
		creationtime:  time.Now(),
	}
	return producer
}

// =============================== Processor ==============================================

type Processor struct {
	logr.Logger
	Event
	Data             interface{}
	sendingChannel   channellib.Channel
	receivingChannel channellib.Channel
}

var _ entity.Node = &Processor{}

func (p *Processor) GetEvent(id string) error {
	// TODO
	return nil
}

func (p *Processor) InjectLogger(log logr.Logger) {
	p.Logger = log
}

func (p *Processor) Run() error {
	// TODO
	// expose a func that users can pass in to modify the data from
	// sending channel
	channellib.Tee(p.sendingChannel, p.receivingChannel)
	return nil
}

// clean this up, instead of bool use a in/out type
func (p *Processor) InjectChannel(ch channellib.Channel, sender bool) error {
	if sender {
		p.sendingChannel = ch
	} else {
		p.receivingChannel = ch
	}
	return nil
}

func Newprocessor(onwner string) Processor {
	producer := Processor{
		sendingChannel:   channellib.NewInfiniteChannel(),
		receivingChannel: channellib.NewInfiniteChannel(),
		Event: Event{
			creationtime: time.Now(),
		},
	}
	return producer
}

func (p *Processor) Out() []interface{} {
	res := make([]interface{}, 0)
	for val := range p.receivingChannel.Out() {
		res = append(res, val)
	}
	return res
}

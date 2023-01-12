package entity

import (
	"github.com/go-logr/logr"
	channellib "gopkg.in/eapache/channels.v1"
)

// type Producer interface {

// 	// for debugging - should be removed
// 	Out() []interface{}
// 	Node
// }

// This is the node interface which the producer, consumer, and the processor need
// to implement.
type Node interface {
	InjectChannel(ch channellib.Channel, sender bool) error
	InjectLogger(log logr.Logger)
	GetState() State
	Run() error

	// Fix: Does each node contain multiple events?
	// If so, we need to store multiple events
	GetEvent() error
}

// This is the pipeline data which is passed around inbetween nodes through channels
// This could be the pacakage content, deppy constraint or variables etc
// type PipelineData interface {
// 	GetData() Contents
// }

type Contents[T any] struct {
	Data T
}

type Identity string

type Options struct {
	// Current id of the node
	SrcId string
	// Id of the node which it needs to reach to.
	// Need to dig to architecture for this. What triggers a new node?
	// Where do we store all the nodes, to identify them based on id?
	DestId   string
	Owner    string
	Metadata map[string]string
}

// type TransformFunc[T Content] func() (T, error)
// type TransformFunc[T any] func() (T, error)
type State string

const (
	Inactive   State = "inactive"
	Active     State = "active"
	Successful State = "successful"
	Fail       State = "fail"
	Aborted    State = "abort"
)

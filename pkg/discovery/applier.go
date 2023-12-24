package discovery

import "fmt"

type Applier interface {
	ApplyCommand(op string, value []byte) interface{}
}

type CommandApplier struct {
}

func NewCommandApplier() *CommandApplier {
	applier := &CommandApplier{}
	return applier
}

func (c *CommandApplier) ApplyCommand(op string, value []byte) interface{} {
	switch op {
	case "heartbeat":
		return nil
	}
	return fmt.Errorf("command {%s} not found", op)
}

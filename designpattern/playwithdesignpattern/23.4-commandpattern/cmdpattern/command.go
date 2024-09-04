package cmdpattern

type CommandInterface interface {
	executeCommand()
}

type Command struct {
	receiver *Receiver
}

func NewCommand(rver *Receiver) *Command {
	return &Command{
		receiver: rver,
	}
}

func (c *Command) executeCommand() {
	c.receiver.action()
}

type ConcreteCommand struct {
	Command
}

func NewConcreteCommand(rver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{
		Command: *NewCommand(rver),
	}
}

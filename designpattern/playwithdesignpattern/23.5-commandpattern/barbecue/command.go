package barbecue

import "fmt"

type CommandInterface interface {
	executeCommand()
}

type Command struct {
	receiver *Barbecuer
}

func (c *Command) executeCommand() {
	fmt.Printf("embeded executeCommand:Not Implemented")
}

func NewCommand(revier *Barbecuer) *Command {
	return &Command{
		receiver: revier,
	}
}

type BakeMuttonCommand struct {
	Command
}

func NewBakeMuttonCommand(revier *Barbecuer) *BakeMuttonCommand {
	return &BakeMuttonCommand{
		Command: *NewCommand(revier),
	}
}

func (c *BakeMuttonCommand) executeCommand() {
	c.receiver.bakeMutton()
}

type BakeChickenWingCommand struct {
	Command
}

func NewBakeChickenWingCommand(revier *Barbecuer) *BakeChickenWingCommand {
	return &BakeChickenWingCommand{Command: *NewCommand(revier)}
}

func (c *BakeChickenWingCommand) executeCommand() {
	c.receiver.bakeChickenWing()
}

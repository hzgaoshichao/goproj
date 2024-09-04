package cmdpattern

type Invoker struct {
	command CommandInterface
}

func (i *Invoker) SetCommand(cmd CommandInterface) {
	i.command = cmd
}

func (i *Invoker) PerformCommand() {
	i.command.executeCommand()
}

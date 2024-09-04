package barbecue

type Waiter struct {
	command CommandInterface
}

func (w *Waiter) SetOrder(cmd CommandInterface) {
	w.command = cmd
}

func (w *Waiter) NotifyCommand() {
	w.command.executeCommand()
}

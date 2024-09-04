package main

import cmdp "commandpattern/cmdpattern"

func main() {
	receiver := cmdp.Receiver{}
	command := cmdp.NewConcreteCommand(&receiver)
	invoker := cmdp.Invoker{}
	invoker.SetCommand(command)
	invoker.PerformCommand()
}

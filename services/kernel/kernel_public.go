package kernel

func (k Kernel) GetCommands() []KernelCommand {
	return k.commands
}

func (k *Kernel) StartRun() {
	k.run = true
}

func (k *Kernel) StopRun() {
	k.run = false
}

func (k *Kernel) UpdatePrompt(url, scope string) {
	if k.readline == nil {
		return
	}

	k.readline.SetPrompt("(" + url + "):" + scope + "$ ")
}

func (k *Kernel) SetDefault(d KernelCommand) {
	k.def = d
}

func (k *Kernel) SetCommands(cmd []KernelCommand) {
	k.commands = cmd
}

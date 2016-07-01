package keeper

type instruction interface {
	usage() error
	exec() error
}

type instructionQueue interface {
	pullOffQueue() error
}

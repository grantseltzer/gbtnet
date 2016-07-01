package keeper

type queue interface {
	pullOffQueue() error
}

type instructionQueue struct {
	frequency    int
	instructions []instruction
}

func (instructionQueue) pullOffQueue() error {

	return nil
}

type instruction interface {
	usage() error
	timestamp() error
	exec() error
}

package go_concourse_helper

type InCommand struct {
	*Command
}

func NewInCommand(messager *Messager) (*InCommand, error) {
	cmd, err := NewCommand(messager)
	if err != nil {
		return nil, err
	}
	return &InCommand{cmd}, nil
}
func (c InCommand) DestinationFolder() string {
	return c.messager.Directory
}

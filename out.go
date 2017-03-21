package go_concourse_helper

type OutCommand struct {
	*Command
}

func NewOutCommand(messager *Messager) (*OutCommand, error) {
	cmd, err := NewCommand(messager)
	if err != nil {
		return nil, err
	}
	return &OutCommand{cmd}, nil
}

func (c *OutCommand) SourceFolder() string {
	return c.messager.Directory
}

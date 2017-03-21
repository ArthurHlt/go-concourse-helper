package go_concourse_helper

import "errors"

type CheckCommand struct {
	*Command
}

func NewCheckCommand(messager *Messager) (*CheckCommand, error) {
	cmd, err := NewCommand(messager)
	if err != nil {
		return nil, err
	}
	return &CheckCommand{cmd}, nil
}
func (c CheckCommand) Send(versions []Version) {
	c.messager.SendJsonResponse(versions)
}
func (c CheckCommand) Params(v interface{}) error {
	return errors.New("No params in check command")
}

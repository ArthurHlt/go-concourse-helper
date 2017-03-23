package go_concourse_helper

import "encoding/json"

type Command struct {
	messager *Messager
	request  Request
}

func NewCommand(messager *Messager) (*Command, error) {
	command := &Command{
		messager: messager,
	}
	err := command.load()
	if err != nil {
		return nil, err
	}
	return command, nil
}
func (c Command) Messager() *Messager {
	return c.messager
}
func (c Command) Source(v interface{}) error {
	b, _ := json.Marshal(c.request.Source)
	return json.Unmarshal(b, v)
}
func (c Command) Params(v interface{}) error {
	b, _ := json.Marshal(c.request.Params)
	return json.Unmarshal(b, v)
}
func (c *Command) load() error {
	return c.messager.RetrieveJsonRequest(&c.request)
}
func (c Command) Send(metadata []Metadata) {
	c.messager.SendJsonResponse(Response{
		Metadata: metadata,
		Version:  c.request.Version,
	})
}

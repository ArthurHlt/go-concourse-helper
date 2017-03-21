package go_concourse_helper_test

import (
	. "github.com/ArthurHlt/go-concourse-helper"

	"bufio"
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Command", func() {
	var (
		messager        *Messager
		command         *Command
		responseBuffer  bytes.Buffer
		responseWritter *bufio.Writer
		logBuffer       bytes.Buffer
		logWritter      *bufio.Writer
		requestBuffer   bytes.Buffer
		requestReader   *bufio.Reader
	)
	BeforeEach(func() {
		responseBuffer.Reset()
		logBuffer.Reset()
		requestBuffer.Reset()

		logWritter = bufio.NewWriter(&logBuffer)
		responseWritter = bufio.NewWriter(&responseBuffer)
		requestReader = bufio.NewReader(&requestBuffer)
		messager = &Messager{
			LogWriter:      logWritter,
			ResponseWriter: responseWritter,
			RequestReader:  requestReader,
			ExitOnFatal:    false,
		}
		_, err := requestBuffer.Write([]byte(`{
			"source": {"name": "foo", "value": "bar"},
			"version": {"build": "1"},
			"params": {"foo": "bar"}
		}`))
		if err != nil {
			Fail(err.Error())
		}
		command, err = NewCommand(messager)
		if err != nil {
			Fail(err.Error())
		}
	})
	Describe("Source", func() {
		Context("when passing a struct with json template as source", func() {
			It("should give back a struct corresponding to as source", func() {
				type myStruct struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				}
				structToPass := myStruct{}
				_, err := requestBuffer.Write([]byte(`{"source": {"name": "foo", "value": "bar"}}`))
				Expect(err).To(BeNil())
				err = command.Source(&structToPass)
				Expect(err).To(BeNil())
				Expect(structToPass.Name).To(BeEquivalentTo("foo"))
				Expect(structToPass.Value).To(BeEquivalentTo("bar"))
			})
		})
	})
	Describe("Send", func() {
		Context("when passing empty metadata", func() {
			It("should send in output data correct json", func() {
				command.Send([]Metadata{})
				responseWritter.Flush()
				var response Response
				err := json.Unmarshal(responseBuffer.Bytes(), &response)
				Expect(err).To(BeNil())
				Expect(response.Metadata).To(HaveLen(0))
				Expect(response.Version.BuildNumber).To(Equal("1"))
			})
		})
		Context("when passing metadatas", func() {
			It("should send in output data correct json", func() {
				command.Send([]Metadata{
					{
						Name:  "foo",
						Value: "bar",
					},
				})
				responseWritter.Flush()
				var response Response
				err := json.Unmarshal(responseBuffer.Bytes(), &response)
				Expect(err).To(BeNil())
				Expect(response.Metadata).To(HaveLen(1))
				Expect(response.Metadata[0].Name).To(Equal("foo"))
				Expect(response.Metadata[0].Value).To(Equal("bar"))
				Expect(response.Version.BuildNumber).To(Equal("1"))
			})
		})
	})
	Describe("Params", func() {
		It("should feed struct with param", func() {
			param := struct {
				Foo string `json:"foo"`
			}{}
			err := command.Params(&param)
			Expect(err).To(BeNil())
			Expect(param.Foo).To(Equal("bar"))
		})
	})
})

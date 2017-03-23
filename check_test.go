package go_concourse_helper_test

import (
	. "github.com/ArthurHlt/go-concourse-helper"

	"bufio"
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Check", func() {
	var (
		messager        *Messager
		command         *CheckCommand
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
			"version": {"build": "1"}
		}`))
		if err != nil {
			Fail(err.Error())
		}
		command, err = NewCheckCommandWithMessager(messager)
		if err != nil {
			Fail(err.Error())
		}
	})
	Describe("Params", func() {
		It("should return an error", func() {
			param := struct{}{}
			err := command.Params(&param)
			Expect(err).Should(HaveOccurred())
		})
	})
	Describe("Send", func() {
		Context("when passing empty versions", func() {
			It("should send in output data correct json", func() {
				command.Send([]Version{})
				responseWritter.Flush()
				var response []Version
				err := json.Unmarshal(responseBuffer.Bytes(), &response)
				Expect(err).To(BeNil())
				Expect(response).To(HaveLen(0))
			})
		})
		Context("when passing versions", func() {
			It("should send in output data correct json", func() {
				command.Send([]Version{
					{
						BuildNumber: "1",
					},
				})
				responseWritter.Flush()
				var response []Version
				err := json.Unmarshal(responseBuffer.Bytes(), &response)
				Expect(err).To(BeNil())
				Expect(response).To(HaveLen(1))
				Expect(response[0].BuildNumber).To(Equal("1"))
			})
		})
	})
})

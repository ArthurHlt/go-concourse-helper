package go_concourse_helper_test

import (
	. "github.com/ArthurHlt/go-concourse-helper"

	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/colorstring"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Messager", func() {
	var (
		messager        *Messager
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

	})
	Describe("LogIt", func() {
		Context("When passing a simple text", func() {
			Context("without colors", func() {
				It("should output this text", func() {
					myText := "mytext"
					messager.Log(myText)
					logWritter.Flush()
					Expect(logBuffer.String()).To(Equal(myText))
				})
			})
			Context("with colors", func() {
				It("should output this text", func() {
					myText := "[red]mytext"
					messager.Log(myText)
					logWritter.Flush()
					Expect(logBuffer.String()).To(Equal(colorstring.Color(myText)))
				})
			})
		})
		Context("When passing a formatted text", func() {
			Context("without colors", func() {
				It("should output text formatted with good values", func() {
					myText := "mytext %s %d"
					number := 1
					messager.Log(myText, myText, number)
					logWritter.Flush()
					Expect(logBuffer.String()).To(Equal(fmt.Sprintf(myText, myText, number)))
				})
			})
			Context("with colors", func() {
				It("should output text formatted with good values", func() {
					myText := "[red]mytext %s %d"
					number := 1
					messager.Log(myText, myText, number)
					logWritter.Flush()
					Expect(logBuffer.String()).To(Equal(fmt.Sprintf(colorstring.Color(myText), myText, number)))
				})
			})
		})
	})
	Describe("LogItLn", func() {
		Context("When passing a simple text", func() {
			It("should output this text and appending a new line", func() {
				myText := "mytext"
				messager.Logln(myText)
				logWritter.Flush()
				Expect(logBuffer.String()).To(Equal(myText + "\n"))
			})
		})
	})
	Describe("SendJsonResponse", func() {
		Context("when passing a struct with json template", func() {
			It("should output on response writer the formatted json", func() {
				type myStruct struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				}
				structToPass := &myStruct{"foo", "bar"}
				messager.SendJsonResponse(structToPass)
				responseWritter.Flush()
				var reponseJson myStruct
				err := json.Unmarshal(responseBuffer.Bytes(), &reponseJson)
				Expect(err).To(BeNil())
				Expect(reponseJson).To(BeEquivalentTo(reponseJson))
			})
		})
	})
	Describe("RetrieveJsonRequest", func() {
		Context("when passing a struct with json template", func() {
			It("should give back a struct corresponding to entry", func() {
				type myStruct struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				}
				structToPass := myStruct{}
				_, err := requestBuffer.Write([]byte(`{"name": "foo", "value": "bar"}`))
				Expect(err).To(BeNil())
				messager.RetrieveJsonRequest(&structToPass)
				Expect(structToPass.Name).To(BeEquivalentTo("foo"))
				Expect(structToPass.Value).To(BeEquivalentTo("bar"))
			})
		})
	})
	Describe("Fatal", func() {
		Context("when writing an error message", func() {
			It("should write this message on response writer", func() {
				errorMessage := "error"
				messager.Fatal(errorMessage)
				logWritter.Flush()
				Expect(logBuffer.String()).To(ContainSubstring(errorMessage))

				responseWritter.Flush()
				Expect(responseBuffer.String()).To(ContainSubstring(errorMessage))
			})
		})
	})
	Describe("FatalIf", func() {
		Context("when writing an error message and error is nil", func() {
			It("should not write this message on response writer", func() {
				errorMessage := "error"
				messager.FatalIf(errorMessage, nil)
				logWritter.Flush()
				Expect(logBuffer.String()).To(Equal(""))

				responseWritter.Flush()
				Expect(responseBuffer.String()).To(Equal(""))
			})
		})
		Context("when writing an error message and error is not nil", func() {
			It("should not write the message on response writer with error message", func() {
				errorMessage := "error"
				errorDetails := "it's an error"
				messager.FatalIf(errorMessage, errors.New(errorDetails))
				logWritter.Flush()
				Expect(logBuffer.String()).To(ContainSubstring(errorMessage + ": " + errorDetails))

				responseWritter.Flush()
				Expect(responseBuffer.String()).To(ContainSubstring(errorMessage + ": " + errorDetails))
			})
		})
	})
})

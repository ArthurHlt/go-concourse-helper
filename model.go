package go_concourse_helper

// regex semver (v|-|_)?v?((?:0|[1-9]\d*)\.(?:0|[1-9]\d*)\.(?:0|[1-9]\d*)(?:-[\da-z\-]+(?:\.[\da-z\-]+)*)?(?:\+[\da-z\-]+(?:\.[\da-z\-]+)*)?)
type Version struct {
	BuildNumber string `json:"build"`
}

type Metadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Request struct {
	Source  interface{} `json:"source"`
	Version Version     `json:"version"`
	Params  interface{} `json:"params"`
}

type Response struct {
	Metadata []Metadata `json:"metadata"`
	Version  Version    `json:"version"`
}

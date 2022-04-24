package fileszip

import (
	"encoding/json"
	"net/http"
)

type FilesZip struct {
	debug    bool
	client   Client
	userHook UserHook
}
type Sources struct {
	Url   string      `json:"url"`
	extra interface{} `json:"extra"`
}
type (
	Client interface {
		Get(url string) (*http.Response, error)
	}
	UserHook interface {
		TransPath(p Sources) string
	}
)

func (s Sources) String() string {
	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(data)
}

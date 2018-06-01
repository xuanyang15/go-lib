package url

import (
	"bytes"
	"net/url"
	"strings"
)

type URLValue struct {
	Name 	string
	Value   []string
}

type URLValues []URLValue

func (uvs *URLValues) Add(name string, value []string) {
	*uvs = append(*uvs, URLValue{Name: name, Value: value})
}


func (uvs URLValues) Assemble(serverHost string, api string) (URL string) {
	var buf bytes.Buffer
	buf.WriteString(serverHost + api)

	for _, urlValue := range uvs {
		v := url.Values{
			urlValue.Name: urlValue.Value,
		}

		if strings.Contains(buf.String(), "?") {
			buf.WriteByte('&')
		} else {
			buf.WriteByte('?')
		}
		buf.WriteString(v.Encode())
	}

	URL = buf.String()
	return
}

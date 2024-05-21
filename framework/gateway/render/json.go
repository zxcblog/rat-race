package render

import (
	"encoding/json"
	"net/http"
)

type JSON struct {
	Data any
}

func (j JSON) Render(w http.ResponseWriter) error {
	j.WriteContentType(w)
	jsonBytes, err := json.Marshal(j.Data)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

func (j JSON) WriteContentType(w http.ResponseWriter) {
	writerContentType(w, []string{"application/json; charset=utf-8"})
}

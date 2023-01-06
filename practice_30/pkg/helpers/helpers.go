package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseJsonTo(u any, w http.ResponseWriter, req *http.Request) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return err
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	return err
}

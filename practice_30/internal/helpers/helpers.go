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

func DeleteElemFromSlice(s []int, i int) []int {
	s[i] = s[len(s)-1]
	s[len(s)-1] = 0
	s = s[:len(s)-1]
	return s
}

func GetElemIndexFromSlice(s []int, v int) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

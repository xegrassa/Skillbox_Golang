package controls

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	h "skillbox/practice_30/pkg/helpers"
	m "skillbox/practice_30/pkg/models"
	s "skillbox/practice_30/pkg/storage"
	"strings"

	"strconv"

	"github.com/go-chi/chi"
)

func CreateUser(w http.ResponseWriter, req *http.Request) {
	log.Println("API: POST /create")
	var u m.User
	err := h.ParseJsonTo(&u, w, req)
	if err != nil {
		return
	}

	userId := s.UserStorage.AddNewUser(&u)

	response, _ := json.Marshal(map[string]int{"id": userId})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

	fmt.Println(s.UserStorage.ToString())
	fmt.Println("=============================")
}

func MakeFriends(w http.ResponseWriter, req *http.Request) {
	log.Println("API: POST /make_friends")

	var mfr m.MakeFriendsRequest
	err := h.ParseJsonTo(&mfr, w, req)
	if err != nil {
		return
	}

	sourceId := mfr.SourceId
	targetId := mfr.TargetId

	u_1, ok_1 := s.UserStorage.GetUser(sourceId)
	u_2, ok_2 := s.UserStorage.GetUser(targetId)

	if !(ok_1 && ok_2) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка id! В базе нет одного из переданных значений id"))
		return
	}

	u_1.Friends = append(u_1.Friends, targetId)
	u_2.Friends = append(u_2.Friends, sourceId)

	fmt.Println(s.UserStorage.ToString())
	fmt.Println("=============================")
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	log.Println("API: DELETE /user")

	var dur m.DeleteUserRequest
	err := h.ParseJsonTo(&dur, w, req)
	if err != nil {
		return
	}
	u, ok := s.UserStorage.GetUser(dur.TargetId)

	switch ok {
	case true:
		uName := u.Name
		s.UserStorage.DeleteUser(dur.TargetId)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(uName))
	case false:
		w.WriteHeader(http.StatusOK)
	}

	fmt.Println(s.UserStorage.ToString())
	fmt.Println("=============================")
}

func GetFriends(w http.ResponseWriter, req *http.Request) {
	log.Println("API: GET /friends/{userId}")

	uId := chi.URLParam(req, "userId")
	uIdInt, _ := strconv.Atoi(uId)

	u, ok := s.UserStorage.GetUser(uIdInt)
	switch ok {
	case true:
		result := make([]string, 0)
		for _, ufId := range u.Friends {
			uf, _ := s.UserStorage.GetUser(ufId)
			result = append(result, uf.Name)
		}
		rStr := strings.Join(result, " ")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(rStr))
	case false:
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Println(s.UserStorage.ToString())
	fmt.Println("=============================")
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {
	log.Println("API: PUT /{userId}")

	uId := chi.URLParam(req, "userId")
	uIdInt, _ := strconv.Atoi(uId)

	var u m.User
	err := h.ParseJsonTo(&u, w, req)
	if err != nil {
		return
	}

	fmt.Println(u)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "UserId=%d\n", uIdInt)
	w.Write([]byte("возраст пользователя успешно обновлён"))

	fmt.Println(s.UserStorage.ToString())
	fmt.Println("=============================")
}

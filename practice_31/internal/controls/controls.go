package controls

import (
	"encoding/json"
	"log"
	"net/http"
	h "skillbox/practice_30/internal/helpers"
	m "skillbox/practice_30/internal/models"
	s "skillbox/practice_30/internal/storage"

	"strconv"

	"github.com/go-chi/chi"
)

type Controller struct {
	repository s.Storage
}

func CreateUser(w http.ResponseWriter, req *http.Request) {
	log.Println("API: POST /create")

	ctx := req.Context()
	repository := ctx.Value("storage").(s.Storage)

	var u m.User
	err := h.ParseJsonTo(&u, w, req)
	if err != nil {
		return
	}

	userId := repository.AddNewUser(&u)

	r := m.CreateUserResponse{Id: userId}
	response, _ := json.Marshal(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

	log.Println(repository.ToString())
	log.Println("=============================")
}

func MakeFriends(w http.ResponseWriter, req *http.Request) {
	log.Println("API: POST /make_friends")

	ctx := req.Context()
	repository := ctx.Value("storage").(s.Storage)

	var mfr m.MakeFriendsRequest
	err := h.ParseJsonTo(&mfr, w, req)
	if err != nil {
		return
	}

	sourceId := mfr.SourceId
	targetId := mfr.TargetId

	u_1, ok_1 := repository.GetUser(sourceId)
	u_2, ok_2 := repository.GetUser(targetId)

	if !(ok_1 && ok_2) {
		w.WriteHeader(http.StatusBadRequest)
		r := m.MakeFriendsResponse{Message: "Ошибка id! В базе нет одного из переданных значений id"}
		response, _ := json.Marshal(r)
		w.Write(response)
		return
	}

	u_1.Friends = append(u_1.Friends, targetId)
	u_2.Friends = append(u_2.Friends, sourceId)

	log.Println(repository.ToString())
	log.Println("=============================")
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	log.Println("API: DELETE /user")

	ctx := req.Context()
	repository := ctx.Value("storage").(s.Storage)

	var dur m.DeleteUserRequest
	err := h.ParseJsonTo(&dur, w, req)
	if err != nil {
		return
	}
	u, ok := repository.GetUser(dur.TargetId)

	switch ok {
	case true:
		uName := u.Name
		repository.DeleteUser(dur.TargetId)

		r := m.DeleteUserResponse{UserName: uName}
		response, _ := json.Marshal(r)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	case false:
		w.WriteHeader(http.StatusNoContent)
	}

	log.Println(repository.ToString())
	log.Println("=============================")
}

func GetFriends(w http.ResponseWriter, req *http.Request) {
	log.Println("API: GET /friends/{userId}")

	ctx := req.Context()
	repository, _ := ctx.Value("storage").(s.Storage)

	// fmt.Println(repository)
	// fmt.Printf("%T", repository)

	uId := chi.URLParam(req, "userId")
	uIdInt, _ := strconv.Atoi(uId)

	u, ok := repository.GetUser(uIdInt)
	switch ok {
	case true:
		result := make([]string, 0)
		for _, ufId := range u.Friends {
			uf, _ := repository.GetUser(ufId)
			result = append(result, uf.Name)
		}
		r := m.GetFriendsResponse{Friends: result}
		response, _ := json.Marshal(r)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	case false:
		w.WriteHeader(http.StatusBadRequest)
	}

	log.Println(repository.ToString())
	log.Println("=============================")
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {
	log.Println("API: PUT /{userId}")

	ctx := req.Context()
	repository, _ := ctx.Value("storage").(s.Storage)

	uId := chi.URLParam(req, "userId")
	uIdInt, _ := strconv.Atoi(uId)

	var handReq m.UpdateUserAgeRequest
	err := h.ParseJsonTo(&handReq, w, req)
	if err != nil {
		return
	}
	newAgeInt, _ := strconv.Atoi(handReq.Age)

	err = repository.UpdateUserAge(uIdInt, newAgeInt)
	var handResp m.UpdateUserAgeResponse
	if err != nil {
		handResp = m.UpdateUserAgeResponse{Message: err.Error()}
		w.WriteHeader(http.StatusNotFound)
	} else {
		handResp = m.UpdateUserAgeResponse{Message: "Возраст пользователя успешно обновлён"}
		w.WriteHeader(http.StatusOK)
	}
	response, _ := json.Marshal(handResp)

	w.Write(response)

	log.Println(repository.ToString())
	log.Println("=============================")
}

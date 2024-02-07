package handler

import (
	"context"
	"encoding/json"
	"strings"

	"fmt"

	"errors"
	"log"
	"net/http"
	"strconv"

	"test-crud-api/internal/model"
	"test-crud-api/pkg/filter"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// func  createUser add newUsers
func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dsdsdsd")
	var user model.User

	err := render.DecodeJSON(r.Body, &user)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Services.CreateUser(context.TODO(), user)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("createUser() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

//getAllUsersWithFilters output users with filters

func (h *Handler) getAllUsersWithFilters(w http.ResponseWriter, r *http.Request) {
	var filterFields filter.Field
	var operator string
	age := r.URL.Query().Get("age")
	fmt.Println(age)
	if age != "" {
		filterFields.Name = "age"
		operator = "="
		value := age

		if strings.Count(age, ":") != 0 {
			splits := strings.Split(age, ":")
			operator = splits[0]
			value = splits[1]

		}

		filterFields.AddFields(filterFields.Name, value, operator, filter.DataTypeInt)

	}
	recording_dateFrom := r.URL.Query().Get("recordingDateFrom")
	recording_dateTo := r.URL.Query().Get("recordingDateTo")

	if recording_dateTo != "" && recording_dateFrom != "" {

		filterFields.AddFields("recording_dateTo", recording_dateFrom, recording_dateTo, filter.DataTypeInt)

	}

	users, err := h.Services.GetAllUsersWithFilter(context.TODO(), filterFields)
	if err != nil {
		log.Println("getAllUsersWithFilter() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(users)

	response, err := json.Marshal(users)
	if err != nil {
		log.Println("getAllUsersWithFilters() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/text")
	countToByte := strconv.Itoa(len(users))
	w.Write([]byte(countToByte))
	w.Write(response)
}

//func getUserByID output users  by id

func (h *Handler) getUserByID(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	fmt.Println(id)
	user, err := h.Services.GetUserById(context.TODO(), id)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("getUserByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		log.Println("getUserByID() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

//func findAllUsers output all users  */

func (h *Handler) findAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	users, err := h.Services.FindAllUsers(context.TODO())
	if err != nil {
		log.Println("getAllUsersWithFilters() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(users)

	response, err := json.Marshal(users)
	if err != nil {
		log.Println("getAllUsersWithFilters() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)

}

// func deleteUser delete  users by  id
func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	err := h.Services.DeleteUser(context.TODO(), id)
	if err != nil {
		log.Println(" deleteUser() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

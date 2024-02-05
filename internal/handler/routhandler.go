package handler

import (
	"context"
	"encoding/json"
	"strings"

	"fmt"

	"errors"

	"log"
	"net/http"

	"test-crud-api/internal/model"
	"test-crud-api/pkg/filter"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// func  createUser add newUsers
func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dsdsdsd")
	var user model.User
	/*err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("jib,rf", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}*/
	err := render.DecodeJSON(r.Body, &user)
	if err != nil {
		//fmt.Printf("oshibka:", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	err = h.services.CreateUser(context.TODO(), user)
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
	var filterOptions filter.Options
	var operator string
	age := r.URL.Query().Get("age")

	if age != "" {
		operator = "="
		value := age
		if strings.Index(age, ":") != -1 {
			splits := strings.Split(age, ":")
			operator = splits[0]
			value = splits[1]
		}
		err := filterOptions.AddFields("age", operator, value, filter.DataTypeDate)
		if err != nil {
			return
		}
	}
	recording_date := r.URL.Query().Get("recording_date")
	if recording_date != "" {
		value := age
		if strings.Index(recording_date, ":") != -1 {
			operator = filter.OperatorBetween

		} else {
			operator = filter.OperatorEq
			splits := strings.Split(age, ":")
			operator = splits[0]
			value = splits[1]
		}
		filterOptions.AddFields("recording_date", operator, value, filter.DataTypeDate)
	}

	users, err := h.services.GetAllUsersWithFilter(context.TODO(), filterOptions)
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
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

//func getUserByID output users  by id

func (h *Handler) getUserByID(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	fmt.Println(id)
	user, err := h.services.GetUserByID(context.TODO(), id)
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
	users, err := h.services.FindAllUsers(context.TODO())
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
	id := chi.URLParam(r, "Id")
	err := h.services.DeleteUser(context.TODO(), id)
	if err != nil {
		log.Println(" deleteUser() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

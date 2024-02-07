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
	/*err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("jib,rf", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}*/
	err := render.DecodeJSON(r.Body, &user)
	if err != nil {
		//fmt.Printf("Error DecodeJson:", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(user)
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
			fmt.Println(filterFields)
		}
		//fmt.Println(age, operator, value, filter.DataTypeInt)
		filterFields.AddFields(filterFields.Name, value, operator, filter.DataTypeInt)
		fmt.Println(filterFields)

	}

	recording_date := r.URL.Query().Get("recording_date")
	if recording_date != "" {
		filterFields.Name = "recording_date"
		value := recording_date
		if strings.Count(recording_date, ":") != 0 {
			//operator = .Value

		} else {
			operator = filter.OperatorEq
			splits := strings.Split(recording_date, ":")
			operator = splits[0]
			value = splits[1]
		}
		filterFields.AddFields(recording_date, value, operator, filter.DataTypeInt)

	}
	fmt.Println(filterFields)
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

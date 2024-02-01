package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"test-crud-api/internal/model"
	"test-crud-api/pkg/filter"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {

	var user model.User

	err := render.DecodeJSON(r.Body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.services.CreateUser(context.TODO(), user)
	if err != nil {
		if errors.Is(err, errors.New("book not found")) {
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

func (h *Handler) getAllUsersWithFilters(w http.ResponseWriter, r *http.Request) {
	var filterOptions:=r.Context().Value(filter.OptitionsContextKey).(filter.Options)
	age:=r.URL.Query().Get("age")
	
	if(age!="") {
	    operator:="="
	    value:=age
        if strings.Index(age,":")!=-1{
		splits:=strings.Split(age,":")
		operator=splits[0]
		value=splits[1]
		}
		filterOptions.AddField("age",oprator,value,filter.DataTypeInt)
     }
	recording_date:=r.URL.Query().Get("recording_date")
	if(recording_date!=""){
		if strings.Index(recording_date,":")!=-1{

		} else{

		}
		filterOptions.AddField("age",oprator,value,filter.DataTypeDate)
	}
	
	users, err := h.services.GetAllUsersWithFilters(context.TODO(), filteroptions)
	if err != nil {
		log.Println("getAllBooks() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(users)

	response, err := json.Marshal(users)
	if err != nil {
		log.Println("getAllBooks() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) getUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "Id")
	user, err := h.services.GetUserByID(context.TODO(), id)
	if err != nil {
		if errors.Is(err, errors.New("book not found")) {
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
func (h *Handler) findAllUsers(w http.ResponseWriter, r *http.Request) {
	h.services.FindAllUsers(context.TODO())
	users, err := h.services.GetAllUsersWithFilters(context.TODO())
	if err != nil {
		log.Println("getAllBooks() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(users)

	response, err := json.Marshal(users)
	if err != nil {
		log.Println("getAllBooks() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
	
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "Id")
	err := h.services.DeleteUser(context.TODO(), id)
	if err != nil {
		log.Println("deleteBook() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

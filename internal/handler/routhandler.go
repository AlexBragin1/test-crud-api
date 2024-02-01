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
//func  createUser add newUsers
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
//getAllUsersWithFilters output users with filters

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
		err:=filterOptions.AddField("age",oprator,value,filter.DataTypeDate)
		if err != nil{
			return err
		}
     }
	recording_date:=r.URL.Query().Get("recording_date")
	if(recording_date!=""){
		var operator string
   
		if strings.Index(recording_date,":")!=-1{
			operator= filter.OperatorBetween

		} else{
			operator=filter.OperatorEq
            
		}
		filterOptions.AddField("age",oprator,value,filter.DataTypeDate)
	}
	
	users, err := h.services.GetAllUsersWithFilters(context.TODO(), filter.Options)
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

//func getUserByID output users  by id

func (h *Handler) getUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "Id")
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

//func findAllUsers output allusers

func (h *Handler) findAllUsers(w http.ResponseWriter, r *http.Request) {
	
	users, err := h.services.findAllUsers(context.TODO())
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
//func deleteUser delete  users by  id 
func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "Id")
	err := h.services.DeleteUser(context.TODO(), id)
	if err != nil {
		log.Println("getAllUsersWithFilters() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

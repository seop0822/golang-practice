package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

//업데이트 시 빈문자열 구분하려면 이런식으로 한다 일단 패스
//type UpdateUser struct {
//	ID        int       `json:"id"`
//	UpdateFirstName bool 'json:update_first_name'
//	FirstName string    `json:"first_name"`
//	LastName  string    `json:"last_name"`
//	Email     string    `json:"email"`
//	CreatedAt time.Time `json:"created_at"`
//}

var userMap map[int]*User
var lastID int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if len(userMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Users")
		return
	}

	users := []*User{}
	for _, u := range userMap {
		users = append(users, u)
	}
	data, _ := json.Marshal(users)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	user, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", id)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	//client가 보낸 유저정보json 읽어야 해서 User sturct를 만들자
	user := new(User)	//인스턴스 생성
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	// Created User
	lastID++ //유저 하나가 등록될때마다 아이디 증가
	user.ID = lastID
	user.CreatedAt = time.Now()	//현재시간
	userMap[user.ID] = user

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))

}

func deleteUserHandler(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	 _, ok := userMap[id]
	 if !ok {
	 	w.WriteHeader(http.StatusOK)
	 	fmt.Fprint(w, "No User Id:", id)
	 	return
	 }
	 delete(userMap,id)
	 w.WriteHeader(http.StatusOK)
	 fmt.Fprint(w, "Deleted User Id:", id)
}

func updateUserHandler (w http.ResponseWriter, r *http.Request){
	updateUser := new(User)
	err := json.NewDecoder(r.Body).Decode(updateUser)
	if err !=nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w,err)
		return
	}

	user, ok := userMap[updateUser.ID]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,"No User Id:",updateUser.ID)
		return
	}
	if updateUser.FirstName != "" {
		user.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		user.LastName = updateUser.LastName
	}
	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}

	//바꾸고 바궜으면 ok가 나온다
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

// NewHandler
func NewHandler() http.Handler {
	userMap = make(map[int]*User)
	lastID = 0
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler).Methods("GET")
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users", updateUserHandler).Methods("PUT")
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")

	return mux
}

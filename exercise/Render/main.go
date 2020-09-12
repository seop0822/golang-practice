package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:  "Park",
		Email: "Park@naver.com",
	}

	rd.JSON(w, http.StatusOK, user) // 이 한줄로 밑에 4줄 표현
	//w.Header().Add("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//data ,_ := json.Marshal(user)
	//fmt.Fprint(w, string(data))

}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error()) //밑에 2줄표현
		//w.WriteHeader(http.StatusBadRequest)
		//fmt.Fprint(w, err)
		return
	}

	user.CreatedAt = time.Now()
	rd.JSON(w, http.StatusOK, user)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:  "Park",
		Email: "Park@naver.com",
	}
	rd.HTML(w, http.StatusOK, "body", user) //밑에껄 표현, hello는 hello.tmpl을 나타냄

	//tmpl, err := template.New("Hello").ParseFiles("exercise/Render/templates/hello.tmpl")
	//if err != nil {
	//	rd.Text(w, http.StatusBadRequest, err.Error())
	//	return
	//}
	//tmpl.ExecuteTemplate(w, "hello.tmpl", "park")
}

func main() {
	rd = render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello",
	}) //render 패키지 template, json,html등 간단히 표현

	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	n := negroni.Classic()	//public/index.html 포함
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}

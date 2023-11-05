package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type User struct {
	Id   int
	Name string
	Age  int
}

// Ифнормации о пользователях должна загружаться
// Это для примера
var user_list = []User{
	{Id: 0, Name: "John", Age: 28},
	{Id: 1, Name: "Kate", Age: 29},
	{Id: 2, Name: "Pole", Age: 30},
}

func main() {
	router := http.NewServeMux()

	// Главная страница - статическая
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// Страница со списком пользователей
	router.HandleFunc("/users", urersHandler)

	// Страница с инофрмацией об одном пользователе
	router.HandleFunc("/user", urerHandler)

	// Стили, картинки, скрипты
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", noСache(http.StripPrefix("/static/", fs)))

	http.ListenAndServe(":8080", router)
}

func urersHandler(w http.ResponseWriter, r *http.Request) {
	// Будет каждый раз загружать шаблон - не очень хорошо
	tmpl, _ := template.ParseFiles("./templates/users.html")

	// Заполняем шаблон данными и пишем ответ в w (http.ResponseWriter)
	err := tmpl.ExecuteTemplate(w, "users.html", user_list)
	if err != nil {
		fmt.Println(err)
	}
}

func urerHandler(w http.ResponseWriter, r *http.Request) {
	// Достаём из запроса параметр id
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	// Будет каждый раз загружать шаблон - не очень хорошо
	tmpl, _ := template.ParseFiles("./templates/user.html")

	// Заполняем шаблон данными и пишем ответ в w (http.ResponseWriter)
	err := tmpl.ExecuteTemplate(w, "user.html", user_list[id])
	if err != nil {
		fmt.Println(err)
	}
}

// Middleware-чтобы браузер не кэшировал стили и скрипты
// Используется для отладки, чтобы стили обновлялись по F5
func noСache(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		next.ServeHTTP(w, r)
	}
}

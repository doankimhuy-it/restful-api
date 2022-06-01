package api

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

type Task struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    Task   `json:"data"`
	Message string `json:"message"`
}

func ConnectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbmane=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database")
		return nil, err
	}
	return db, nil
}

func Get(w http.ResponseWriter, r *http.Request) {
	taskId := r.FormValue("id")
	if taskId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no id"))
		return
	}
	db, err := ConnectDB()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("db not connected"))
		return
	}
	var task = Task{}
	err = db.QueryRow("SELECT id, title, status FROM todo WHERE id = $1;", taskId).Scan(&task.Id, &task.Title, &task.Status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("task not available"))
		return
	}
	str := fmt.Sprintf("id: %v, title: %s, status: %s", task.Id, task.Title, task.Status)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(str))
}

func Create(w http.ResponseWriter, r *http.Request) {
	taskId := r.FormValue("id")
	taskTitle := r.FormValue("title")
	taskStatus := r.FormValue("status")
	if taskId == "" || taskTitle == "" || taskStatus == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no info"))
		return
	}
	db, err := ConnectDB()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("db not connected"))
		return
	}
	_, err = db.Exec("INSERT INTO todo VALUES ($1, $2, $3);", taskId, taskTitle, taskStatus)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("cannot insert"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	taskId := r.FormValue("id")
	taskTitle := r.FormValue("title")
	taskStatus := r.FormValue("status")
	if taskId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no id"))
		return
	}
	db, err := ConnectDB()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("db not connected"))
		return
	}
	var ok = true
	if taskTitle != "" && taskStatus != "" {
		ok = Exec(db, "UPDATE todo SET title=$1, status=$2 WHERE id=$3;", taskTitle, taskStatus, taskId)
	} else if taskTitle != "" {
		ok = Exec(db, "UPDATE todo SET title=$1 WHERE id=$2;", taskTitle, taskId)
	} else if taskStatus != "" {
		ok = Exec(db, "UPDATE todo SET status=$1 WHERE id=$2;", taskStatus, taskId)
	}
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("cannot update"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func Del(w http.ResponseWriter, r *http.Request) {
	taskId := r.FormValue("id")
	if taskId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no id"))
		return
	}
	db, err := ConnectDB()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("db not connected"))
		return
	}
	_, err = db.Exec("DELETE FROM todo WHERE id=$1;", taskId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("cannot delete"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func Exec(db *sql.DB, query string, args ...interface{}) bool {
	_, err := db.Exec(query, args...)
	return err == nil
}

package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)

	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)

	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := session.GetUser()
		if err != nil {
			log.Println(err)
		}

		todos, err := user.GetTodos()

		if err != nil {
			log.Println(err)
		}

		user.Todos = todos

		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	if _, err := session(w, r); err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	if session, err := session(w, r); err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	} else {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}
		user, err := session.GetUser()
		if err != nil {
			log.Println(err)
		}

		if err := user.CreateTodo(r.PostFormValue("content")); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	if session, err := session(w, r); err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		if _, err := session.GetUser(); err != nil {
			log.Println(err)
		}

		t, err := models.GetTodoById(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	if session, err := session(w, r); err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}

		user, err := session.GetUser()
		if err != nil {
			log.Println(err)
		}

		content := r.PostFormValue("content")

		t := &models.Todo{ID: id, Content: content, UserID: user.ID}

		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	if session, err := session(w, r); err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		if _, err = session.GetUser(); err != nil {
			log.Println(err)
		}

		t, err := models.GetTodoById(id)
		if err != nil {
			log.Println(err)
		}

		if err = t.DeleteTodo(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

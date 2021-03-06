package server

import (
	"blog/models"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
)

func (srv *BlogServer) defineRoutes() {
	srv.router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { srv.GetPage(w, r, "index") })
		r.Get("/blog/", func(w http.ResponseWriter, r *http.Request) { srv.GetPage(w, r, "blog") })
	})
	srv.router.Route("/api/v1/blog/", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) { srv.PostCreate(w, r) })
		r.Put("/{postID}", func(w http.ResponseWriter, r *http.Request) { srv.PostUpdate(w, r) })
		r.Delete("/{postID}", func(w http.ResponseWriter, r *http.Request) { srv.PostDelete(w, r) })
		r.Get("/{postID}", func(w http.ResponseWriter, r *http.Request) { srv.PostGet(w, r) })
	})
}

//GetPage -
func (srv *BlogServer) GetPage(w http.ResponseWriter, r *http.Request, pageName string) {
	file, _ := os.Open("./templates/" + pageName + ".html")
	data, _ := ioutil.ReadAll(file)

	fmt.Println("Подготовка шаблона " + pageName)
	tmp, err := template.New(pageName).Parse(string(data))
	if err != nil {
		fmt.Println("Ошибка!", err)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}
	fmt.Println("Наполнение шаблона " + pageName + " данными")
	tmp.ExecuteTemplate(w, pageName, srv)
}

//PostCreate -
func (srv *BlogServer) PostCreate(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	post := models.BlogPost{}
	_ = json.Unmarshal(data, &post)

	if id, err := srv.Posts.NewBlogPost(srv.DBlink, post.About, post.Text, post.Labels); err != nil {
		fmt.Println(err)
		return
	} else {
		data, _ = json.Marshal(srv.Posts[id])
		w.Write(data)
	}
}

//PostGet -
func (srv *BlogServer) PostGet(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := ioutil.ReadAll(r.Body)
	post := models.BlogPost{}
	_ = json.Unmarshal(data, &post)
	post.ID = int(postID)

	if post2, ok := srv.Posts[post.ID]; !ok {
		fmt.Println("Запрошенный post.ID не найден!")
		return
	} else {
		data, _ = json.Marshal(post2)
		w.Write(data)
	}
}

//PostDelete -
func (srv *BlogServer) PostDelete(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := ioutil.ReadAll(r.Body)
	post := models.BlogPost{}
	_ = json.Unmarshal(data, &post)
	post.ID = int(postID)

	if err := srv.Posts.DeleteBlogPost(srv.DBlink, post.ID); err != nil {
		fmt.Println(err)
		return
	}
}

//PostUpdate -
func (srv *BlogServer) PostUpdate(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.ParseInt(chi.URLParam(r, "postID"), 10, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := ioutil.ReadAll(r.Body)
	post := models.BlogPost{}
	_ = json.Unmarshal(data, &post)
	post.ID = int(postID)

	if err := srv.Posts.UpdateBlogPost(srv.DBlink, post.ID, post.About, post.Text, post.Labels); err != nil {
		fmt.Println(err)
		return
	} else {
		data, _ = json.Marshal(srv.Posts[post.ID])
		w.Write(data)
	}
}

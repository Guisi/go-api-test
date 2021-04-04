package v1

import (
	"encoding/json"
	"fmt"
	"go-api-test/inject"
	"go-api-test/model"
	"go-api-test/service/post"
	"net/http"
)

func GetSystemStatus(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Everything is OK!")
}

func Posts(res http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		GetAllPosts(res)
	case http.MethodPost:
		SavePost(res, req)
	}

}

func GetAllPosts(res http.ResponseWriter) {
	postService := inject.In("postService").(post.PostService)

	allPosts := postService.GetAllPosts()

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(allPosts)
}

func SavePost(res http.ResponseWriter, req *http.Request) {
	postService := inject.In("postService").(post.PostService)

	decoder := json.NewDecoder(req.Body)

	var post *model.Post

	err := decoder.Decode(&post)
	if err != nil {
		panic(err)
	}

	ret := postService.SavePost(post)

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ret)
}

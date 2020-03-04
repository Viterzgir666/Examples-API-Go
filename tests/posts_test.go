package tests

import (
	"net/http"
	"testing"

	"../src/generics"
	"../src/request/built_objects"
	"../src/request/controllers"
	"../src/utils"
)

func TestPosts(t *testing.T) {
	request := new(controllers.Request)
	headers := built_objects.NewHeadersObject()

	t.Run("User is able to add a new Post", func(t *testing.T) {
		postRequestBody := built_objects.NewPostObject(
			generics.RandomUserId,
			generics.RandomId,
			generics.RandomTitle,
			generics.RandomDescription,
		)
		postSchema := utils.LoadSchema(t, "/posts/actions/post")
		resp := request.AddPost(t, headers, postRequestBody).Expect().Status(http.StatusCreated).JSON()
		resp.Schema(postSchema)
	})
}

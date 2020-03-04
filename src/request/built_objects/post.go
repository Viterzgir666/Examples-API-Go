package built_objects

import (
	"encoding/json"
	"fmt"

	"../builders"
)

func NewPostObject(userdId int, id int, title string, desc string) []byte {
	postRequestBodyBuilder := &builders.Post{}
	postRequestBody := postRequestBodyBuilder.
		SetUserId(userdId).
		SetId(id).
		SetTitle(title).
		SetBody(desc).
		Build()
	postRequestBodyJson, err := json.Marshal(postRequestBody)
	if err != nil {
		fmt.Println("Something went wrong. \nObject is: ")
		fmt.Printf("%+v\n", postRequestBody)
		panic(err)
	}
	return postRequestBodyJson
}

package builders

type PostBuildProcess interface {
	SetUserId(userId int) PostBuildProcess
	SetId(id int) PostBuildProcess
	SetTitle(title string) PostBuildProcess
	SetBody(body string) PostBuildProcess
	Build() Post
}

type HeadersBuildProcess interface {
	Init() HeadersBuildProcess
	SetContentType(contentType string) HeadersBuildProcess
	SetAuthorization(id string) HeadersBuildProcess
	SetHeaderKeyValue(key string, value string) HeadersBuildProcess
	Build() map[string]string
}

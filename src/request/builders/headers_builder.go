package builders

type Headers struct {
	headers map[string]string
}

func (h *Headers) Init() HeadersBuildProcess {
	h.headers = make(map[string]string)
	return h
}

func (h *Headers) SetContentType(contentType string) HeadersBuildProcess {
	h.headers["Content-Type"] = contentType
	return h
}

func (h *Headers) SetAuthorization(id string) HeadersBuildProcess {
	h.headers["authorization"] = id
	return h
}

func (h *Headers) SetHeaderKeyValue(key string, value string) HeadersBuildProcess {
	h.headers[key] = value
	return h
}

func (h *Headers) Build() map[string]string {
	return h.headers
}

package builders

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (p *Post) SetUserId(userId int) PostBuildProcess {
	p.UserId = userId
	return p
}

func (p *Post) SetId(id int) PostBuildProcess {
	p.Id = id
	return p
}

func (p *Post) SetTitle(title string) PostBuildProcess {
	p.Title = title
	return p
}

func (p *Post) SetBody(body string) PostBuildProcess {
	p.Body = body
	return p
}

func (p *Post) Build() Post {
	return *p
}

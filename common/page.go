package common

// Page 分页
type Page struct {
	Count   int         `json:"count"`
	Content interface{} `json:"content"`
}

func NewPage() *Page {
	return &Page{}
}

func (p *Page) SetCount(count int) *Page {
	p.Count = count
	return p
}

func (p *Page) SetContent(data interface{}) *Page {
	p.Content = data
	return p
}

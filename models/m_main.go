package models

import (
	"strings"
	"unicode"
)

type (
	Page struct {
		Page			int
		Rows 			int
		Count			bool	`json:"-"`
		Total  			uint 	`form:"-" json:"Total"`
		Sort			string
	}
	/*Model struct {
		ID        uint `gorm:"primary_key"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time `sql:"index"`
	}*/
)

func (p Page) getOrderBy() string {
	if p.Sort == "" {
		return "id desc"
	}
	asc := ""
	splitSort := strings.Split(p.Sort," ")
	if len(splitSort) > 1 && splitSort[1] == "desc" {
		asc = " desc"
	}
	var words []string
	l := 0
	for s := splitSort[0]; s != ""; s = s[l:] {
		l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if l <= 0 {
			l = len(s)
		}
		words = append(words, strings.ToLower(s[:l]))
	}
	return strings.Join(words,"_") + asc
}

func (p *Page) getOffset() (offset int) {
	if p.Page < 1 {
		p.Page = 1
	}
	if(p.Rows < 5){
		p.Rows = 5
	}
	offset = (p.Page - 1) * p.Rows
	return
}

func (p *Page) GetPageInfo() (page Page) {
	page.Page = p.Page
	page.Rows = p.Rows
	page.Total = p.Total
	return
}

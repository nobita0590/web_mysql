package modules

import (
	"html/template"
)

type (
	SEOobject struct {
		Title string
		Description string
		Image string
	}
)

func (s SEOobject) buildSEO() template.HTML {
	return template.HTML(`
		<meta property="og:type"          content="website" />
		<meta property="og:title"         content="` + s.Title + `" />
		<meta property="og:description"   content="` + s.Description + `" />
		<meta property="og:image"         content="http://2study.edu.vn/` + s.Image + `" />
		<title>`+s.Title+`</title>
		<meta name="description" content="` + s.Description + `" />
		<meta name="google" content="notranslate" />
	`)
}

func (s SEOobject) HomePage() template.HTML {
	s.Title = "2Study"
	s.Description = "Học thi online cùng 2Study. Ôn - luyện thi đại học"
	s.Image = "public/frondtend/images/banner-1.png"
	return s.buildSEO()
}

func (s SEOobject) Course() template.HTML {
	s.Title = "2Study - Khóa học"
	s.Description = "Thông tin các khóa học trên hệ thống 2Study"
	s.Image = "public/frondtend/images/banner-1.png"
	return s.buildSEO()
}

func (s SEOobject) News1() template.HTML {
	s.Title = "2Study - Thông tin tuyển sinh"
	s.Description = "2Study - Tin tức - Thông tin các đợt tuyển sinh"
	s.Image = "public/frondtend/images/banner-1.png"
	return s.buildSEO()
}

func (s SEOobject) News2() template.HTML {
	s.Title = "2Study - Bí quyết học thi"
	s.Description = "2Study - Tin tức - Bí quyết học thi - chia sẻ cách học thi hiệu quả nhất"
	s.Image = "public/frondtend/images/banner-1.png"
	return s.buildSEO()
}

func (s SEOobject) Exam() template.HTML {
	s.Title = "2Study - Thi online"
	s.Description = "2Study - Thi online - Kiểm tra kiến thức của bạn. Nhanh chóng và thuận tiện"
	s.Image = "public/frondtend/images/banner-1.png"
	return s.buildSEO()
}
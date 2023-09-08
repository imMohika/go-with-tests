package renderer

import (
	"embed"
	"go-with-tests/blog"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templates *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templates, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templates}, nil
}

func (pr PostRenderer) Render(w io.Writer, post blog.Post) error {
	err := pr.templates.ExecuteTemplate(w, "post.gohtml", post)
	if err != nil {
		return err
	}

	return nil
}

func (pr PostRenderer) RenderIndex(w io.Writer, posts []blog.Post) error {
	err := pr.templates.ExecuteTemplate(w, "index.gohtml", posts)
	if err != nil {
		return err
	}

	return nil
}

package blogrender_test

import (
	blogrender "blogrenderpkg"
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestBlogrender(t *testing.T) {
	var (
		aPost = blogrender.Post{
			Title:       "hello world",
			Body:        "This is the post",
			Description: "This is description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := blogrender.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, filename string) (Post, error) {
	file, err := fileSystem.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()
	post, err := NewPost(file)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeprator         = "Tags: "
	bodySeparator        = "Body: "
)

func NewPost(postfile io.Reader) (Post, error) {
	sc := bufio.NewScanner(postfile)
	readline := func(tagname string) string {
		sc.Scan()
		return strings.TrimPrefix(sc.Text(), tagname)
	}
	title := readline(titleSeparator)
	description := readline(descriptionSeparator)
	tags := strings.Split(readline(tagsSeprator), ",")
	sc.Scan()
	buf := bytes.Buffer{}
	for sc.Scan() {
		fmt.Fprintln(&buf, sc.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")
	post := Post{Title: title, Description: description, Tags: tags, Body: body}
	return post, nil
}

package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

//var posts []Post
//posts = NewPostsFromFS("posts")

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

func getPost(fileSystem fs.FS, name string) (Post, error) {
	file, err := fileSystem.Open(name)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	return readPost(file)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
	arraySeparator       = ", "
)

func readPost(r io.Reader) (Post, error) {
	scanner := bufio.NewScanner(r)

	title, description, tags := readMetadata(scanner)
	body := readBody(scanner)

	post := Post{title, description, body, tags}
	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}

func readMetadata(scanner *bufio.Scanner) (string, string, []string) {
	readMeta := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	title := readMeta(titleSeparator)
	description := readMeta(descriptionSeparator)
	tags := strings.Split(readMeta(tagsSeparator), arraySeparator)
	return title, description, tags
}

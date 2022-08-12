package task

import (
	"bytes"
	"errors"
	"golang.org/x/net/html"
	"io/ioutil"
	"sort"
	"strings"
)

var ErrNotImplemented = errors.New("function not implemented")

// ReadHTMLFromFile should read the file from the current directory, if it exists.
// The file data should be returned as a string.
func ReadHTMLFromFile(fileName string) (string, error) {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bs), nil
	// return "", ErrNotImplemented
}

// CreateBuffer should transfer the contents of a string to a buffer.
func CreateBuffer(data string) bytes.Buffer {
	var b bytes.Buffer
	b.Write([]byte(data))
	return b
	// return bytes.Buffer{}
}

// CreateTree should create the tree representation of HTML represented by the buffer.
func CreateTree(buf bytes.Buffer) (*html.Node, error) {
	return html.Parse(bytes.NewReader(buf.Bytes()))
}

// CountDivTags should return the count of <div> tags in the document tree.
func CountDivTags(node *html.Node) int {
	count := 0
	if node.DataAtom.String() == "div" {
		count++
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		count += CountDivTags(child)
	}
	return count
}

// dfs is a utility function which will help you count the number of unique tags.
func dfs(node *html.Node, tagsCount map[string]int) {
	if node.Type == html.ElementNode {
		tagsCount[node.DataAtom.String()]++
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		dfs(child, tagsCount)
	}
}

// ExtractAllUniqueTagsInSortedOrder should return the unique tags in the document.
// These tags should also be sorted alphabetically.
func ExtractAllUniqueTagsInSortedOrder(node *html.Node) []string {
	count := make(map[string]int)
	dfs(node, count)
	var tags []string
	for tag := range count {
		tags = append(tags, tag)
	}
	sort.Strings(tags)
	return tags
}

// ExtractAllComments returns the list of all comments as they appear in the document.
// You also need to remove all the leading and trailing spaces in the comments.
// HINT: You might need to read about variadic functions.
func ExtractAllComments(node *html.Node) []string {
	var comments []string
	if node.Type == html.CommentNode {
		comments = append(comments, strings.TrimSpace(node.Data))
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		comments = append(comments, ExtractAllComments(child)...)
	}
	return comments
}

// ExtractAllLinks returns all the links in the document, in order of appearance.
func ExtractAllLinks(node *html.Node) []string {
	var links []string
	if node.DataAtom.String() == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		links = append(links, ExtractAllLinks(child)...)
	}
	return links
}

package page

import (
	"io/ioutil"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := "data/" + p.Title + ".txt"

	_, err := ioutil.ReadDir("data")
	if err != nil {
		err = os.Mkdir("data", 0600)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

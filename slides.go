package slides

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jimareed/drawing"
)

// single slide
type Slide struct {
	title    string
	bullets  []string
	hasImage bool
	Drawing  drawing.Drawing
}

type SlideDeck struct {
	Title  string
	Slides []Slide
	text   []string
}

func Read(path string) (SlideDeck, error) {

	deck := SlideDeck{}
	deck.Title = ""

	filename := path + "/README.md"

	f, err := os.Open(filename)
	if err != nil {
		return deck, err
	}
	defer func() {
		if err = f.Close(); err != nil {
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		deck.text = append(deck.text, s.Text())
	}
	err = s.Err()

	if len(deck.text) > 0 {
		title := deck.text[0]
		if strings.HasPrefix(title, "# ") {
			title = string(title[2:])
		}
		deck.Title = title
	}

	filename = path + "/00.draw"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	deck.Slides = append(deck.Slides, Slide{})
	deck.Slides[0].Drawing, err = drawing.FromString(text)

	return deck, err
}

func Write(deck SlideDeck, path string) error {

	filename := path + "/README.md"

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return err
	}

	for _, v := range deck.text {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

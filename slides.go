package slides

import (
	"io/ioutil"
	"log"

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

func Read(path string, name string) (SlideDeck, error) {

	deck := SlideDeck{}
	deck.Title = ""

	filename := path + "/" + name + ".draw"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
		return deck, err
	}
	text := string(content)

	deck.Slides = append(deck.Slides, Slide{})
	deck.Slides[0].Drawing, err = drawing.FromString(text)

	return deck, err
}

func Write(deck SlideDeck, path string) error {

	return nil
}

func ToHtml(deck SlideDeck) (string, error) {

	if len(deck.Slides) > 0 {
		content, err := drawing.ToSvg(deck.Slides[0].Drawing)
		if err == nil {
			return content, nil
		}
	}
	return "", nil
}

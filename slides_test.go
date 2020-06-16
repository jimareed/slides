package slides

import (
	"testing"
)

func TestRead(t *testing.T) {

	sd, err := Read("./test/input")
	if err != nil {
		t.Log("Read error")
		t.Fail()
	}

	if len(sd.Slides) < 1 {
		t.Log("Empty slides error")
		t.Fail()
	}

	d := sd.Slides[0].Drawing

	if d.Width != 900 || d.Height != 600 {
		t.Log("Error, invalid drawing")
		t.Fail()
	}
}

func TestToHtml(t *testing.T) {

	sd, err := Read("./test/input")
	if err != nil {
		t.Log("Read error")
		t.Fail()
	}

	html, err := ToHtml(sd)
	if err != nil {
		t.Log("ToHtml error")
		t.Fail()
	}
	if html == "" {
		t.Log("empty content error")
		t.Fail()
	}
}

package slides

import (
	"testing"
)

func TestRead(t *testing.T) {

	_, err := Read("./test/input")		
	if err != nil {
		t.Log("Error testing stubbed implementation")
		t.Fail()
	}
}

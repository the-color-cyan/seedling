package markdown

import (
	"errors"
	"testing"
)

func TestTopLevelListReturnsTopLevelList(t *testing.T) {
	src := []byte(`
- src/
	- file.go
- docs/

- cat_pics/
	- bingus.png
`)

	doc := Parse(src)
	got, err := topLevelList(doc)
	if err != nil {
		t.Fatal(err)
	}

	if got != doc.root.FirstChild() {
		t.Fatalf("got %T, want first top-level node", got)
	}
}

func TestDecodeRejectsMultipleLists(t *testing.T) {
	src := []byte(`
- src/
	- file.go

separator

- cat_pics/
	- bingus.png
`)

	_, err := Decode(src)
	if !errors.Is(err, ErrListCount) {
		t.Fatal("multiple lists in markdown should be rejected")
	}
}

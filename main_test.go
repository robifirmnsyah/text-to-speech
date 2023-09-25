package main

import (
	"strings"
	"testing"
)

func TestGenerateName(t *testing.T) {
	name := GenerateName()
	ok := strings.Contains(name, ".mp3")
	if !ok {
		t.Fail()
		t.Log("nama bukan termasuk file.mp3 suffix")
	}
}

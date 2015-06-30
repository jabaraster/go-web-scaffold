package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_parse_01(t *testing.T) {
	fmt.Println(filepath.Abs("."))

	f := parse("/css/index.css")
	if f.directory != "/css" {
		t.Errorf("invalid directory -> %s", f.directory)
	}
	if f.name != "index.css" {
		t.Errorf("invalid file name -> %s", f.name)
	}
	if f.extension != "css" {
		t.Errorf("invalid extension -> %s", f.extension)
	}
	if f.nameWithoutExtension != "index" {
		t.Errorf("invalid name without extension -> %s", f.nameWithoutExtension)
	}
}

func Test_parse_02(t *testing.T) {
	f := parse("index.min.css")
	if f.directory != "" {
		t.Errorf("invalid directory -> %s", f.directory)
	}
	if f.name != "index.min.css" {
		t.Errorf("invalid file name -> %s", f.name)
	}
	if f.extension != "css" {
		t.Errorf("invalid extension -> %s", f.extension)
	}
	if f.nameWithoutExtension != "index.min" {
		t.Errorf("invalid name without extension -> %s", f.nameWithoutExtension)
	}
}

func Test_getCssTag_getJsTag(t *testing.T) {
	d, e := filepath.Abs("../../assets")
	if e != nil {
		panic(e)
	}
	o := tagOutput{d}

	css := o.getCssTag("/css/index.css")
	fmt.Println("★ ", css)

	js := o.getJsTag("/js/index.js")
	fmt.Println("★ ", js)
}

func Test_getCssTag_invalidPath(t *testing.T) {
	d, e := filepath.Abs("../../assets")
	if e != nil {
		panic(e)
	}
	o := tagOutput{d}
	o.getCssTag("/../../index.css")
}

package main

import (
	"github.com/goodman-dev/portia/internal/portia"
)

func main() {

	path, err := portia.ChartPath("Ned_I.R._Jennings", "World_War_I")
	if err != nil {
		panic(err)
	}

	for _, page := range path {
		println(page)
	}

}

package main

import (
	"fmt"

	"github.com/snail007/gf/frame/g"

	_ "github.com/snail007/gf/os/gres/testdata"
)

func main() {
	m := g.I18n()
	m.SetLanguage("ja")
	err := m.SetPath("/i18n-dir")
	if err != nil {
		panic(err)
	}
	fmt.Println(m.Translate(`hello`))
	fmt.Println(m.Translate(`{#hello}{#world}!`))
}

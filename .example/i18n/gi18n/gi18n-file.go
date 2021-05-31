package main

import (
	"fmt"

	"github.com/snail007/gf/i18n/gi18n"
)

func main() {
	t := gi18n.New()
	t.SetLanguage("ja")
	err := t.SetPath("./i18n-file")
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Translate(`hello`))
	fmt.Println(t.Translate(`{#hello}{#world}!`))
}

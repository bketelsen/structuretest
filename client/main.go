// +build js,wasm
package main

import (
	"github.com/bketelsen/structuretest/components/slideshow"
	tdl "github.com/bketelsen/structuretest/components/todolist"
	"github.com/factorapp/structure/core"
)

func main() {
	core.RegisterController("todolist", &tdl.TodoList{})
	core.RegisterController("slideshow", &slideshow.SlideshowController{})
	core.Run()
}

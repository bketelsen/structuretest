package todolist

import (
	"fmt"

	"github.com/factorapp/structure/core"
	dom "github.com/gowasm/go-js-dom"
)

type Todo struct {
	Name        string
	Description string
	Done        bool
}

type TodoList struct {
	core.BasicController
	Todos []Todo
}

func (t *TodoList) Add(ev dom.Event) {
	// TODO: how do you pass in form data?
	el := ev.Target()
	fmt.Println("Target:", el)
	fmt.Println("Tag:", el.TagName())
}

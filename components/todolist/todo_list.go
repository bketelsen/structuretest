package todolist

import (
	"fmt"

	"github.com/factorapp/structure/core"
	"github.com/kr/pretty"
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

// func (t *TodoList) Render() {
// 	, err := t.Targets()["todolist"].TextContent()
// }

func (t *TodoList) Add(ctx core.Context) error {
	name, err := ctx.FormInput("name")
	if err != nil {
		return err
	}
	descr, err := ctx.FormInput("description")
	if err != nil {
		return err
	}
	// name := nameElt.Value()
	// description := descriptionElt.Value()

	// these objects have no value in them.
	// don't remember - are the targets kept up to date, or do we need to look them up directly
	// in the DOM?
	pretty.Println("name", name)
	pretty.Println("description", descr)

	str, err := ctx.Templates().Render("new-todo", map[string]interface{}{
		"todo": Todo{
			Name:        name,
			Description: descr,
		},
	})
	if err != nil {
		fmt.Println("error rendering template:", err)
		return err
	}

	fmt.Println("template rendered:", str)

	todoListElts, ok := t.Targets()["todos"]
	if !ok {
		fmt.Println("no todolist element slice")
		return fmt.Errorf("no todolist element slice")
	}
	if len(todoListElts) == 0 {
		fmt.Println("no todo list elts")
		return fmt.Errorf("no todo list elts")
	}
	todoListElt := todoListElts[0]
	if err := todoListElt.AppendHTML(str); err != nil {
		fmt.Println("couldn't append:", err)
		return err
	}
	// newTextContent := oldTextContent + str

	// todoListVal, ok := t.Targets()["todolist"]
	// if !ok {
	// 	fmt.Println("no todolist target")
	// 	return fmt.Errorf("no todolist target")
	// }
	// t.Targets()["todolist"].SetTextContent(newTextContent)
	// t.Targets()["todolist"].AppendTextContent(str)
	// ctx.Element("div#todos").Append(str)

	return nil
	// TODO: how do you pass in form data?
	// name := t.Targets()["name"][0]
	// pretty.Println("name:", name.TagName())

	// name, err := t.Input("name")
	// pretty.Println("error:", err)
	// pretty.Println("name:", name)

	// // TODO: event handlers return errors?
	// name, _ := t.GetFormData("name")
	// pretty.Println("name", name)
	// desc, _ := t.GetFormData("description")
	// pretty.Println("description", desc)
	// name := t.Targets()["name"][0]
	// descr := t.Targets()["description"][0]
	// pretty.Println("name:", name.NodeValue())
	// pretty.Println("description:", descr.NodeValue())
}

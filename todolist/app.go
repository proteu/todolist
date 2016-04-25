package todolist

import "regexp"

type App struct {
	TodoStore Store
}

func NewApp() *App {
	app := &App{TodoStore: NewFileStore()}
	app.TodoStore.Load()
	return app
}

func (a *App) ListTodos(input string) {

	grouped := a.getGroups(input)

	formatter := NewFormatter(grouped)
	formatter.Print()
}

func (a *App) getGroups(input string) *GroupedTodos {
	grouper := &Grouper{}
	contextRegex, _ := regexp.Compile(`by c.*$`)
	projectRegex, _ := regexp.Compile(`by p.*$`)

	var grouped *GroupedTodos

	if contextRegex.MatchString(input) {
		grouped = grouper.GroupByContext(a.TodoStore.Todos())
	} else if projectRegex.MatchString(input) {
		grouped = grouper.GroupByContext(a.TodoStore.Todos())
	} else {
		grouped = grouper.GroupByNothing(a.TodoStore.Todos())
	}
	return grouped
}
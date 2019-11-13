package hserve

func main() {
	ListenAndServe(func(app *Application) {
		app.Name = "test"
		app.Port = 5803
	})
}

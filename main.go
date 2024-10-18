package main

import(
  "html/template"
  "log"
  "net/http"
)

type Todo struct {
  Id int
  Message string
}

func main(){

  // starting a variable with clousures make the variable in memory between requests
  // as long the application is running
  data := map[string][]Todo{
    "Todos": {
      Todo{Id: 1, Message: "Fix LSP"},
    },
  }

  todosHandler := func(w http.ResponseWriter, r *http.Request) {
    templ := template.Must(template.ParseFiles("index.html"))
    templ.Execute(w, data)
  }

  addTodoHandler := func(w http.ResponseWriter, r *http.Request){
    message := r.PostFormValue("message")
    templ := template.Must(template.ParseFiles("index.html"))
    todo := Todo{ Id: len(data["Todos"]) + 1, Message: message}
    
    templ.ExecuteTemplate(w, "todo-list-element", todo)
  }

  http.HandleFunc("/", todosHandler)
  http.HandleFunc("/add-todo", addTodoHandler)
  log.Fatal(http.ListenAndServe(":8000", nil))
}

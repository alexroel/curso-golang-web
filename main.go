package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Funciones de saludo y suma
func greeting() string {
	return "Hola desde la función"
}

func sum(a, b int) int {
	return a + b
}

// Controladores
func home(w http.ResponseWriter, r *http.Request) {
	// Mapa de funciones
	funcs := template.FuncMap{
		"greeting": greeting,
		"sum":      sum,
	}
	// Cargar y analizar el archivo de plantilla HTML
	//template, err := template.ParseFiles("templates/home.html")
	template, err := template.New("home.html").Funcs(funcs).ParseFiles("templates/home.html", "templates/base.html")
	if err != nil {
		log.Println("Error al crear la plantilla:", err)
	}

	// Renderizar la plantilla en la respuesta
	template.Execute(w, nil)
}

type Course struct {
	Name     string
	Duration int
}

type User struct {
	Name     string
	Age      int
	IsActive bool
	IsAdmin  bool
	Skills   []string
	Courses  []Course
}

// Método de tiene permisos de administrador
func (u User) HasAdminPermissions(key string) bool {
	return u.IsActive && u.IsAdmin && key == "admin"
}

func hello(w http.ResponseWriter, r *http.Request) {
	// Cargar y analizar el archivo de plantilla HTML
	template, err := template.New("hello.html").ParseFiles("templates/hello.html", "templates/base.html")
	if err != nil {
		log.Println("Error al crear la plantilla:", err)
	}
	// Lista de habilidades
	skills := []string{"Go", "HTML", "CSS", "JavaScript"}
	// Lista de cursos
	courses := []Course{{"Go", 10}, {"JavaScript", 12}, {"HTM y CSS", 13}}

	user := User{
		Name:     "Alex",
		Age:      28,
		IsActive: true,
		IsAdmin:  true,
		Skills:   skills,
		Courses:  courses,
	}

	// Renderizar la plantilla en la respuesta
	template.Execute(w, user)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Devolver un error personalizado para páginas no encontradas
	http.Error(w, "404 - Página no encontrada", http.StatusNotFound)
}

func main() {
	// Crear un enrutaro
	router := mux.NewRouter()

	// Definir rutas y controladores
	router.HandleFunc("/", home)
	router.HandleFunc("/hello", hello)

	// Manejo personalizado para página no encontrada (404)
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// Creación de un servidor web con Go
	port := ":8080" // Puerto en el que se ejecutará el servidor

	// Asignar el enrutador al servidor
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	//log.Fatal(http.ListenAndServe(port, router))
	log.Fatal(server.ListenAndServe())

}

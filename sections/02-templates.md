# Manejo de plantillas 
1. [Introducción](#introducción)
2. [Rendirizar plantillas](#rendirizar-plantillas)
3. [Enviar datos](#enviar-datos)
4. [Declaración de variable](#declaración-de-variable)
5. [Conposición de plantillas](#conposición-de-plantillas)
6. [Uso de Muts](#uso-de-muts)
7. [Enlaces a páginas](#enlaces-a-páginas)
8. [Resumen](#resumen)

---
## Introducción 
En esta sección del curso de Go web, exploraremos el poderoso manejo de plantillas en el desarrollo de aplicaciones web con Go. Las plantillas son una herramienta fundamental que nos permite separar la lógica de presentación de nuestros datos y generar contenido dinámico en nuestras páginas web.

Durante esta sección, cubriremos los siguientes temas principales:

Rendirizar plantillas: Aprenderemos cómo cargar y rendirizar plantillas HTML utilizando el paquete html/template de Go. Exploraremos cómo integrar nuestras plantillas en nuestro código y enviarlas al cliente para su visualización.

Enviar datos a las plantillas: Descubriremos cómo enviar datos desde nuestro código Go a las plantillas. Exploraremos diferentes formas de pasar variables, estructuras o mapas a las plantillas para que puedan ser utilizados en la generación de contenido dinámico.

Declaración de variables en plantillas: Aprenderemos cómo declarar y utilizar variables en nuestras plantillas. Veremos cómo asignar valores a las variables, realizar operaciones básicas y mostrar su contenido en la página generada.

Composición de plantillas: Exploraremos cómo crear plantillas más complejas utilizando la composición. Aprenderemos a dividir nuestras plantillas en partes reutilizables y combinarlas para construir páginas completas.

Uso de template.Must: Descubriremos cómo utilizar la función template.Must para cargar y analizar nuestras plantillas de forma segura. Aprenderemos a manejar errores de manera más eficiente y garantizar que nuestras plantillas se carguen correctamente.

Enlaces a páginas: Aprenderemos cómo crear enlaces dinámicos a diferentes páginas en nuestra aplicación web. Veremos cómo generar los enlaces en nuestras plantillas y cómo manejar las solicitudes correspondientes en nuestro código Go.

Con estos temas, estarás preparado para aprovechar al máximo el manejo de plantillas en tus proyectos de desarrollo web con Go. ¡Comencemos a crear páginas dinámicas y atractivas!

---
## Rendirizar plantillas 
Para renderizar plantillas en Go, se utilizan bibliotecas como "html/template" o "text/template", que están incluidas en el paquete estándar de Go. Estas bibliotecas proporcionan funcionalidades para generar y renderizar contenido dinámico basado en plantillas.

Aquí hay un ejemplo básico de cómo renderizar una plantilla en Go utilizando la biblioteca "html/template":

1. Primero, asegúrate de importar el paquete `"html/template"` en tu archivo Go:
2. Luego, define una función de controlador para manejar la solicitud HTTP

~~~go
import (
    "html/template"
    "net/http"
)


func Index(w http.ResponseWriter, r *http.Request) {
    // Cargar la plantilla desde un archivo
    tpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    // Renderizar la plantilla con los datos proporcionados
    err = tpl.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
~~~

3. Crea un archivo de plantilla llamado "index.html" en la carpeta "templates". Aquí tienes un ejemplo simple de cómo se podría ver la plantilla:

~~~html
<!DOCTYPE html>
<html>
<head>
    <title>Página de inicio</title>
</head>
<body>
    <h1>¡Bienvenido!</h1>
</body>
</html>
~~~

- La función `ParseFiles` toma una lista de nombres de archivos de plantilla y devuelve un objeto de tipo `*template.Template` que representa la plantilla analizada. En este caso, solo estamos analizando un archivo de plantilla, por lo que solo se proporciona el nombre del archivo "index.html".

La variable `tpl` se declara para almacenar el objeto de plantilla analizada y `err` se declara para capturar cualquier error que pueda ocurrir durante el análisis.

- Si `err` no es nulo, significa que se produjo un error durante el análisis de la plantilla. En ese caso, se utiliza la función `http.Error` para enviar una respuesta de error al cliente con un código de estado HTTP interno del servidor (500) y el mensaje de error como cuerpo de la respuesta. Luego, la función return se utiliza para salir del controlador y evitar que se ejecute el resto del código.

- La función `Execute` se utiliza para renderizar la plantilla y escribir el resultado en el `http.ResponseWriter` proporcionado (w). En este caso, pasamos nil como el segundo argumento, lo que indica que no hay datos adicionales para la plantilla.

---
## Enviar datos 
Para enviar datos a una plantilla en Go, puedes utilizar una estructura, un mapa u otro tipo de datos estructurados. A continuación, te muestro cómo enviar datos a una plantilla utilizando una estructura como ejemplo:

~~~go
func Index(w http.ResponseWriter, r *http.Request) {
    // Cargar la plantilla desde un archivo
    tpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    // Datos para la plantilla (puede ser un struct, map, etc.)
    data := struct {
        Title string
        Message string
    }{
        Title: "Página de inicio",
        Message: "¡Bienvenido!",
    }
    
    // Renderizar la plantilla con los datos proporcionados
    err = tpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
~~~


~~~html
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Message}}</h1>
</body>
</html>
~~~

En este ejemplo, utilizamos la sintaxis `{{.Title}}` y `{{.Message}}` en la plantilla para acceder a las propiedades del objeto data que pasamos al método `Execute`.

---
## Declaración de variable 

En las plantillas de Go, puedes definir y utilizar variables utilizando la sintaxis `{{ $variable := valor }}`. Esto te permite asignar un valor a una variable y luego utilizarla en la plantilla. Aquí tienes un ejemplo de cómo definir y utilizar variables en una plantilla:

~~~go
    data := struct {
        Title string
        Message string
    }{
        Title: "Página de inicio",
        Message: "¡Bienvenido!",
    }
~~~

Dentro de la plantilla "index.html", puedes definir y utilizar variables de la siguiente manera:

~~~html
<!DOCTYPE html>
<html>
<head>
    {{ $data := . }}
    <title>{{ $data.Title }}</title>
</head>
<body>
    <h1>{{ $data.Message }}</h1>
</body>
</html>
~~~

---
## Conposición de plantillas 
En Go, puedes lograr la composición de plantillas utilizando la función `template.ParseFiles` y el método `ExecuteTemplate` de la plantilla. La composición de plantillas te permite crear una plantilla base que define la estructura y el diseño general de tu página, y luego incorporar contenido específico en áreas designadas dentro de esa plantilla base.

Crea una plantilla base que define la estructura y el diseño general de tu página. Por ejemplo, puedes llamarla "base.html":

~~~go
tpl, err := template.ParseFiles("templates/base.html", "templates/index.html")
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

err = tpl.ExecuteTemplate(w, "base", data)
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
~~~
En este ejemplo, estamos renderizando la plantilla base con los datos proporcionados y utilizando el nombre "base" como el nombre de la plantilla.
Cuando se ejecuta ExecuteTemplate, el contenido de la plantilla específica se insertará en el bloque "content" de la plantilla base, y el resultado final se enviará como respuesta HTTP al cliente.

Este enfoque te permite crear múltiples plantillas específicas que extienden la misma plantilla base y reutilizar el diseño y la estructura comunes en tu sitio web. Puedes definir bloques adicionales en la plantilla base y utilizar {{define}} en las plantillas específicas para insertar contenido en esos bloques.

~~~html
{{ define "base" }}
<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <title>{{ template "title" . }}</title>
</head>

<body>
    <header>
        <h2>Encabezado</h2>
    </header>
    
    {{ template "content" . }}

    <footer class="bg-light text-center py-4">
        <p class="mb-0">Creado por ROELCODE &copy; 2023</p>
    </footer>
</body>

</html>
{{ end }}
~~~

Crea una plantilla específica que extienda la plantilla base. Por ejemplo, puedes llamarla "index.html":

~~~html
{{ template "base" . }}

{{ define "title" }}RPS web - Inicio {{ end }}

{{ define "content" }}

<h1>Bienvenido a Piedra, Papel o Tijera</h1>
    
{{ end }}
~~~

---
## Uso de Muts
En Go, `template.Must` es una función que se utiliza para simplificar el manejo de errores al cargar y analizar plantillas. `template.Must` toma como argumento el resultado de la función `template.ParseFiles` u otra función similar que devuelve una plantilla y un posible error, y devuelve la plantilla sin el error.

La función `template.Must` se utiliza para envolver la llamada a la función que carga y analiza las plantillas. Si la función devuelve un error, `template.Must` generará un pánico. Esto es útil cuando estás seguro de que la carga y el análisis de la plantilla deben ser exitosos y no esperas un error en tiempo de ejecución.

quí tienes un ejemplo de cómo utilizar template.Must:

~~~go
tpl := template.Must(template.ParseFiles("template.html"))
~~~

En este ejemplo, `template.ParseFiles("template.html") `carga y analiza la plantilla `"template.html"`. Luego, `template.Must` envuelve la llamada a `template.ParseFiles` y devuelve la plantilla sin el error. Si la carga o el análisis de la plantilla falla, se generará un pánico.

Usar `template.Must` puede ser conveniente cuando estás seguro de que las plantillas están correctamente definidas y no esperas errores en tiempo de ejecución. Sin embargo, debes tener en cuenta que si se produce un error al cargar o analizar la plantilla, se generará un pánico y la aplicación se detendrá.

### Creando un función reutilizable para manejo de plantillas 

~~~go
const (
	templatesDir = "templates/"
	templateBase = templatesDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "index.html", nil)
}

func renderTemplate(w http.ResponseWriter, base, page string, data any) {
	t := template.Must(template.ParseFiles(base, templatesDir+page))

	err := t.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
~~~

---
## Enlaces a páginas 

Para crear enlaces a páginas en una aplicación web en Go, puedes utilizar la función http.ResponseWriter.Write para generar el código HTML correspondiente a los enlaces. Aquí tienes un ejemplo de cómo crear un enlace a una página:

~~~go
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Generar un enlace a otra página
		link := "<a href='/about'>Acerca de</a>"
		
		// Escribir el enlace en la respuesta HTTP
		w.Write([]byte(link))
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Página de Acerca de"))
	})

	http.ListenAndServe(":8000", nil)
}
~~~

En este ejemplo, hemos configurado dos manejadores de ruta: uno para la ruta "/" y otro para la ruta "/about". En el manejador de la ruta "/", generamos un enlace a la página "/about" utilizando la etiqueta HTML `<a>` y el atributo href. Luego, utilizamos http.ResponseWriter.Write para escribir el código HTML del enlace en la respuesta HTTP.

Cuando se accede a la ruta "/", se mostrará el enlace "Acerca de". Si se hace clic en el enlace, se redirigirá a la página "/about", donde se mostrará el texto "Página de Acerca de".

Puedes crear enlaces adicionales a otras páginas siguiendo un enfoque similar. Simplemente genera el código HTML del enlace y utiliza http.ResponseWriter.Write para agregarlo a la respuesta HTTP.

`base.html`
~~~html
{{ define "base" }}
<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8">
    <title>{{ template "title" . }}</title>
</head>

<body>
    <header>
        <div class="logo">
            <a href="/">RPS</a>
        </div>
        <nav>
            <a href="/new">Nuevo</a>
            <a href="/about">About</a>
        </nav>
    </header>
    
    {{ template "content" . }}

    <footer class="bg-light text-center py-4">
        <p class="mb-0">Creado por ROELCODE &copy; 2023</p>
    </footer>
</body>

</html>
{{ end }}
~~~

`/handlers.handlers.go`
~~~go
const (
	templatesDir = "templates/"
	templateBase = templatesDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "new-game.html", nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, templateBase, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, base, page string, data any) {
	t := template.Must(template.ParseFiles(base, templatesDir+page))

	err := t.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
~~~

`main.go`
~~~go
package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {

	router := http.NewServeMux()

	// Manejador para servir los archivos estáticos
	fs := http.FileServer(http.Dir("static"))
	// Ruta para acceder a los archivos estáticos
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080" // Puerto en el que se ejecutará el servidor
	log.Printf("Servidor ejecutándose en http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, router))
}
~~~

---
## Resumen 
En la sección de manejo de plantillas del curso de Go web, exploramos las capacidades de rendirizar plantillas en aplicaciones web. Aprendimos a cargar y rendirizar plantillas HTML utilizando el paquete html/template de Go. Aprendimos a integrar las plantillas en nuestro código y enviarlas al cliente para su visualización.

También aprendimos cómo enviar datos desde el código Go a las plantillas. Exploramos diferentes métodos para pasar variables, estructuras o mapas a las plantillas, permitiendo generar contenido dinámico.

En cuanto a la declaración de variables en las plantillas, aprendimos cómo declarar y utilizar variables, asignarles valores, realizar operaciones básicas y mostrar su contenido en la página generada.

Además, exploramos la composición de plantillas, donde aprendimos a dividir nuestras plantillas en partes reutilizables y combinarlas para construir páginas completas.

En cuanto al uso de template.Must, aprendimos a utilizar esta función para cargar y analizar nuestras plantillas de forma segura. Esto nos permitió manejar errores de manera más eficiente y garantizar que nuestras plantillas se cargaran correctamente.

Finalmente, aprendimos a crear enlaces dinámicos a diferentes páginas en nuestra aplicación web. Aprendimos cómo generar los enlaces en nuestras plantillas y cómo manejar las solicitudes correspondientes en nuestro código Go.

Con estos conocimientos, hemos adquirido habilidades para aprovechar al máximo el manejo de plantillas en nuestros proyectos de desarrollo web con Go. Ahora estamos preparados para crear páginas dinámicas y atractivas en nuestras aplicaciones web.

# Manejo de plantillas 

1. [Introducción](#introducción)
2. [Uso de plantillas](#uso-de-plantillas)
3. [Condicionales](#condicionales)
4. [Ciclos](#ciclos)
5. [Variables](#variables)
6. [Métodos y funciones](#métodos-y-funciones)
7. [Reutilización de plantillas](#reutilización-de-plantillas)
8. [Cargar múltiples plantillas](#cargar-múltiples-plantillas)
9. [Rederizar plantillas](#rederizar-plantillas)
10. [Página de error](#página-de-error)
11. [Resumen](#resumen)


---
## Introducción
Bienvenidos a la sección de "Manejo de plantillas" de nuestro curso de Go. En esta sección, aprenderemos cómo utilizar plantillas para generar contenido dinámico en nuestras aplicaciones web.

Las plantillas son una poderosa herramienta que nos permite separar la lógica de presentación de nuestra aplicación, lo que facilita el mantenimiento y la reutilización del código. A lo largo de esta sección, exploraremos varios aspectos del manejo de plantillas en Go.

Comenzaremos por aprender cómo utilizar condicionales y ciclos en nuestras plantillas. Estas características nos permitirán controlar la ejecución y la apariencia del contenido en función de diferentes condiciones y datos proporcionados.

A continuación, exploraremos cómo trabajar con variables en las plantillas. Las variables nos permiten almacenar y acceder a datos en nuestra plantilla, lo que nos brinda flexibilidad y capacidad para adaptar dinámicamente el contenido generado.

También aprenderemos a utilizar métodos y funciones en nuestras plantillas. Estas funciones nos permiten manipular y transformar datos en tiempo de ejecución, lo que agrega una capa adicional de flexibilidad y personalización a nuestras plantillas.

Además, discutiremos la reutilización de plantillas. Veremos cómo podemos crear plantillas base que contengan elementos comunes y cómo extender y personalizar estas plantillas base para generar diferentes vistas en nuestras aplicaciones.

A medida que avanzamos, también exploraremos cómo cargar múltiples plantillas en nuestras aplicaciones y cómo renderizar estas plantillas para generar el contenido final que se enviará al cliente.

Finalmente, abordaremos el tema de las páginas de error. Aprenderemos a personalizar las respuestas de error en nuestras aplicaciones web, utilizando plantillas para mostrar mensajes de error informativos y amigables para el usuario.

En resumen, en esta sección del curso de Go, adquirirás habilidades fundamentales para trabajar con plantillas en tus aplicaciones web. Te proporcionaremos los conocimientos necesarios para utilizar condicionales, ciclos, variables, métodos y funciones en tus plantillas, así como también para cargar múltiples plantillas y personalizar las páginas de error. ¡Comencemos a explorar el emocionante mundo del manejo de plantillas en Go!

---
## Uso de plantillas
### Nueva plantilla
Si deseas crear una nueva plantilla en Go, puedes hacerlo directamente en tu código utilizando la librería "html/template". Aquí tienes un ejemplo de cómo crear una nueva plantilla en Go sin utilizar un archivo HTML:

~~~go
// Controladores
func home(w http.ResponseWriter, r *http.Request) {

	// Definición de la plantilla
	tmp := `<h1>Hola Mundo</h1>
	<p>Página de inicio </p>`

    //Creación de una instancia de template
	template, err := template.New("hola").Parse(tmp)
	if err != nil {
		log.Println("Error al crear la plantilla:", err)
	}

	template.Execute(w, nil)
	// fmt.Fprintf(w, tmp)
}
~~~

- Definición de la plantilla: Aquí se define una cadena de texto llamada tmp que contiene el contenido HTML de la plantilla. En este caso, la plantilla consiste en un encabezado de nivel 1 y un párrafo que dice "Página de inicio".
- Creación de una instancia de template: Se utiliza template.New("hola") para crear una nueva instancia de template con el nombre "hola". Luego, se utiliza Parse(tmp) para analizar la plantilla y asociarla con la instancia del template. Si ocurre un error durante el análisis, se imprime un mensaje de error.
- Renderización de la plantilla: Se utiliza Execute(w, nil) para renderizar la plantilla y enviarla como respuesta al cliente. w es el objeto http.ResponseWriter que se utiliza para escribir la respuesta HTTP. El segundo argumento nil es opcional y representa los datos que se pueden pasar a la plantilla (en este caso, no se pasa ningún dato).

En resumen, este código crea una plantilla HTML básica utilizando la librería "text/template" de Go. El controlador home se encarga de renderizar la plantilla y enviarla como respuesta al cliente cuando se accede a la página de inicio

### Responder archivos HTML
Si deseas responder con un archivo HTML utilizando la librería html/template y su método ParseFiles, puedes modificar el código anterior de la siguiente manera:

~~~go
// Controladores
func home(w http.ResponseWriter, r *http.Request) {
	// Cargar y analizar el archivo de plantilla HTML
	template, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Println("Error al crear la plantilla:", err)
	}

	// Renderizar la plantilla en la respuesta
	template.Execute(w, nil)
}
~~~

Camabir la ruta: 
~~~go
// Definir rutas y controladores
router.HandleFunc("/", home)
router.HandleFunc("/hello", hello)
~~~

Dentro de de `templates/home.html`: 
~~~html
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Inicio</title>
</head>
<body>
    <h1>Hola Mundo</h1>
    <p>Página de incio.</p>
</body>
</html>
~~~

- Cargar y analizar el archivo de plantilla HTML: En esta sección, utilizamos template.ParseFiles("templates/home.html") para cargar y analizar el archivo de plantilla HTML. La función ParseFiles carga y analiza automáticamente el archivo y devuelve un objeto *template.Template que representa la plantilla.
- Renderizar la plantilla en la respuesta: Utilizamos template.Execute(w, nil) para renderizar la plantilla en la respuesta HTTP. La función Execute toma dos argumentos: el objeto http.ResponseWriter (w) que se utiliza para escribir la respuesta HTTP y nil que representa los datos que se pueden pasar a la plantilla (en este caso, no se pasa ningún dato).

En resumen, este código carga y analiza el archivo de plantilla HTML utilizando template.ParseFiles y luego renderiza la plantilla en la respuesta HTTP utilizando template.Execute. El controlador home se encarga de responder con la página HTML correspondiente cuando se accede a esa ruta en el servidor.

---
## Enviar datos a plantillas
Para enviar datos a una plantilla en Go, puedes seguir los siguientes pasos:

Definir una estructura de datos: Define una estructura en Go que represente los datos que deseas pasar a la plantilla. Por ejemplo:

~~~go
type User struct {
	Name string
	Age  int
}

func hello(w http.ResponseWriter, r *http.Request) {
	// Cargar y analizar el archivo de plantilla HTML
	template, err := template.ParseFiles("templates/hello.html")
	if err != nil {
		log.Println("Error al crear la plantilla:", err)
	}

	user := User{
		Name: "Alex",
		Age:  28,
	}

	// Renderizar la plantilla con los datos
	template.Execute(w, user)
}
~~~

- Preparar los datos que deseas pasar a la plantilla: En este ejemplo, se crea una estructura User con los campos Name y Age y se asignan valores.

- Renderizar la plantilla con los datos: Utilizamos template.Execute(w, user) para renderizar la plantilla en la respuesta HTTP. La función Execute toma dos argumentos: el objeto http.ResponseWriter (w) que se utiliza para escribir la respuesta HTTP y user que representa los datos que se pasan a la plantilla.

Código de `templates/hello.html`: 
~~~html
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hola, {{ .Name}}</title>
</head>
<body>
    <h2>Hola, {{ .Name }} y tu edad es {{ .Age }}</h2>
</body>
</html>
~~~

---
## Condicionales
En las plantillas de Go, puedes utilizar condicionales para mostrar u ocultar contenido según una condición determinada. La sintaxis para los condicionales en las plantillas de Go se basa en la estructura `{{if .Condicion}} ... {{else}} ... {{end}}`. Aquí tienes un ejemplo de cómo puedes usar condicionales en una plantilla de Go:

### Operadores relacionales
En las plantillas de Go, puedes utilizar operadores relacionales para realizar comparaciones y tomar decisiones basadas en ellas. A continuación, se presentan los operadores relacionales disponibles en las plantillas de Go:

- Igualdad `eq`: Comprueba si dos valores son iguales. Sintaxis: `{{if eq .Variable1 .Variable2}} ... {{end}}`
- Desigualdad `ne`: Comprueba si dos valores no son iguales. Sintaxis: `{{if ne .Variable1 .Variable2}} ... {{end}}`
- Menor que `lt`: Comprueba si el primer valor es menor que el segundo valor. Sintaxis: `{{if lt .Variable1 .Variable2}} ... {{end}}`
- Menor o igual que `le`: Comprueba si el primer valor es menor o igual que el segundo valor.Sintaxis: `{{if le .Variable1 .Variable2}} ... {{end}}`
- Mayor que `gt`: Comprueba si el primer valor es mayor que el segundo valor. Sintaxis: `{{if gt .Variable1 .Variable2}} ... {{end}}`
- Mayor o igual que `ge`: Comprueba si el primer valor es mayor o igual que el segundo valor.Sintaxis: `{{if ge .Variable1 .Variable2}} ... {{end}}`

### Operadores lógicos
En las plantillas de Go, también puedes utilizar operadores lógicos para combinar condiciones y controlar el flujo de ejecución. A continuación se presentan los operadores lógicos disponibles:

- Y lógico `and`: Evalúa como verdadero si todas las condiciones son verdaderas.
Sintaxis: `{{if and .Condicion1 .Condicion2}} ... {{end}}`
- O lógico `or`: Evalúa como verdadero si al menos una de las condiciones es verdadera.
Sintaxis: `{{if or .Condicion1 .Condicion2}} ... {{end}}`
- Negación lógica `not`: Niega el resultado de una condición.Sintaxis: `{{if not .Condicion}} ... {{end}}`

### Ejemplo
Estructura de datos y el controlador para este jemplo: 
~~~go
type User struct {
	Name     string
	Age      int
	IsActive bool
	IsAdmin  bool
}

func hello(w http.ResponseWriter, r *http.Request) {
	// Cargar y analizar el archivo de plantilla HTML
	template, err := template.ParseFiles("templates/hello.html")
	if err != nil {
		log.Println("Error al crear la plantilla:", err)
	}

	user := User{
		Name:     "Alex",
		Age:      28,
		IsActive: true,
		IsAdmin:  true,
	}

	// Renderizar la plantilla en la respuesta
	template.Execute(w, user)
}
~~~

Condiciones para en la plantilla:
~~~html
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hola, {{ .Name}}</title>
</head>
<body>
    <h2>Hola, {{ .Name }}</h2>
    <!-- Verifica si esta activo o no -->
    {{ if .IsActive }}
    <p>El usuario esta activo.</p>
    {{ else }}
    <p>El usuario no esta activo.</p>
    {{ end }}

    <!-- Verifica si esta activo y si es administrador  -->
    {{ if (and .IsActive .IsAdmin) }}
    <p>Eres edministrador. </p>
    {{ end }}

    <!-- Si la edad es mayor a 20 -->
    {{ if (gt .Age 22) }}
    <p>El usuario tiene más de 20 años. </p>
    {{ end }}
</body>
</html>
~~~

---
## Ciclos
En las plantillas de Go, también puedes utilizar bucles para repetir bloques de contenido o iterar sobre colecciones de datos. La sintaxis para los bucles en las plantillas de Go se basa en la estructura `{{range .Coleccion}} ... {{end}}`. A continuación te muestro cómo utilizar los bucles en las plantillas de Go:

Ejemplo para usar bucles en plantilla de Go con una lista de elementos. 
~~~go
type User struct {
	Name     string
	Age      int
	IsActive bool
	IsAdmin  bool
	Skills   []string
}

func hello(w http.ResponseWriter, r *http.Request) {
	// Cargar y analizar el archivo de plantilla HTML
	template, err := template.ParseFiles("templates/hello.html")
	if err != nil {
		log.Println("Error al crear la plantilla:", err)
	}

	// Lista de habilidades
	skills := []string{"Go", "HTML", "CSS", "JavaScript"}

	user := User{
		Name:     "Alex",
		Age:      28,
		IsActive: true,
		IsAdmin:  true,
		Skills:   skills,
	}

	// Renderizar la plantilla en la respuesta
	template.Execute(w, user)
}
~~~

En la plantilla usar lo siguiente: 
~~~html
    <!-- Listar habilidades -->
    <h3>Habilidades</h3>
    <ul>
        {{ range .Skills }}
        <li>{{ . }}</li>
        {{ end }}
    </ul>
~~~

- `{{ range .Skills }}`: Esta línea utiliza la estructura {{ range }} de las plantillas de Go para iniciar un bucle sobre la colección de habilidades (Skills). El punto (.) se refiere al contexto actual de la plantilla, que debe ser un objeto que tenga un campo llamado Skills.

- `<li>{{ . }}</li>`: Dentro del bucle, utilizamos `{{ . }}` para acceder a cada elemento de la colección de habilidades y mostrarlo en un elemento de lista (`<li>`).

- `{{ end }}`: Esta línea marca el final del bucle {{ range }}. Todo el código dentro del bucle se repetirá para cada elemento de la colección Skills.

Ejemplo de iteración de objetos: 

~~~go
// Estructura de cursos
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

func hello(w http.ResponseWriter, r *http.Request) {
	// Cargar y analizar el archivo de plantilla HTML
	template, err := template.ParseFiles("templates/hello.html")
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
~~~

Código en la platilla: 
~~~html
    <!-- Lista de cursos -->
    <h3>Cursos</h3>
    <ul>
        {{ range .Courses }}
        <li>Curso de: {{ .Name }} -Duración: {{ .Duration }} horas.</li>
        {{ end }}
    </ul>
~~~

---
## Variables
En las plantillas de Go, puedes utilizar variables para almacenar y acceder a valores dinámicos dentro de la plantilla. Aquí te muestro cómo utilizar variables en las plantillas de Go:

Definición de variables: Para definir una variable en una plantilla, puedes utilizar la siguiente sintaxis:

~~~html
{{ $nombreVariable := valor }}
~~~

Por ejemplo: 
~~~html
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Inicio</title>
</head>
<body>

    {{ $title := "Hola Mundo"}}
    {{ $message := "Página de incio."}}
    <h1>{{ $title }}</h1>
    <p>{{ $message }}</p>

    <!-- Modificar el valor  -->
    {{ $message = "Bienvenido a mi web con Go." }}
    <p>{{ $message }}</p>
    
</body>
</html>
~~~

---
## Métodos y funciones
En las plantillas de Go, puedes utilizar métodos y funciones para realizar operaciones y manipular los datos dentro de la plantilla. Los métodos y funciones en las plantillas de Go se invocan utilizando la sintaxis `{{ .MetodoOFuncion }}` dentro de la plantilla. A continuación, te muestro cómo utilizar métodos y funciones en las plantillas de Go:

### Métodos en platillas 
En Go, puedes definir métodos en una estructura para realizar operaciones específicas en los datos. Aquí tienes un ejemplo:

~~~go
// Método de tiene permisos de administrador
func (u *User) HasAdminPermissions() bool {
	return u.IsActive && u.IsAdmin && "admin"
}
~~~

Usar el método de la siguiente forma:

~~~html
    <!-- Es administrador  -->
    {{ if .HasAdminPermissions }}
    <p>Administrador. </p>
    {{ end }}
~~~

Método con parametros:
~~~go
// Método de tiene permisos de administrador
func (u User) HasAdminPermissions(key string) bool {
	return u.IsActive && u.IsAdmin && key == "admin"
}
~~~

Código en la plantillla: 
~~~html
    <!-- Es administrador  -->
    {{ if .HasAdminPermissions "admin" }}
    <p>Administrador.</p>
    {{ end }}
~~~

### Funciones
Además de los métodos, también puedes utilizar funciones definidas por el usuario en las plantillas. Aquí tienes un ejemplo:

~~~go
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
	template, err := template.New("home.html").Funcs(funcs).ParseFiles("templates/home.html")
	if err != nil {
		log.Println("Error al crear la plantilla:", err)
	}

	// Renderizar la plantilla en la respuesta
	template.Execute(w, nil)
}
~~~

- Se crea un mapa de funciones (template.FuncMap) llamado `funcs`. El mapa de funciones contiene pares clave-valor, donde la clave es el nombre que se utilizará para invocar la función en la plantilla, y el valor es la función en sí. En este ejemplo, se definen dos funciones llamadas greeting y sum.
- Se utiliza template.New("home.html") para crear una nueva instancia de template con el nombre "home.html". Luego, se utiliza el método .Funcs(funcs) para asignar el mapa de funciones `funcs` a la instancia del template. Finalmente, se utiliza .ParseFiles("templates/home.html") para cargar y analizar el archivo de plantilla HTML especificado.

Código en la plantilla: 
~~~html
    <!-- Uso de funciones  -->
    <p>{{ greeting }}</p>
    <p>Resutaldo: {{ sum 10 20 }}</p>
~~~

---
## Reutilización de plantillas
Para reutilizar plantillas en Go, puedes utilizar el bloque `{{define}}` y `{{template}}`. A continuación te explico cómo puedes hacerlo:

**Define una plantilla base utilizando `{{define}}`**: En lugar de definir el contenido completo de la plantilla, puedes utilizar `{{define}}` para declarar un bloque de contenido que se puede reutilizar en diferentes plantillas. Por ejemplo en `base.html`:

~~~html
{{ define "header" }}
<header>
    <ul>
        <li><a href="/">Inicio</a></li>
        <li><a href="/hello">Hola</a></li>
    </ul>
</header>
{{ end }}

{{ define "footer" }}
<hr>
<footer>
    <p>Este es pie de página</p>
</footer>
{{ end }}
~~~

En este ejemplo, definimos un bloque de contenido llamado "header" y "footer" que contiene un un bloque de código. El contenido de este bloque puede variar en diferentes plantillas.

**`Utiliza {{template}}` para incluir la plantilla base en otras plantillas:** Luego de definir la plantilla base, puedes utilizar `{{template}}` para incluir ese bloque de contenido en otras plantillas. Por ejemplo en `home.html` y `hello.html`:

~~~html
{{ template "header" }}

.... 

{{ template "footer" }}
~~~

Modifica los controladores cargando multiples platillas con método `ParseFile`: 

~~~go
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

...

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
~~~

---
## Cargar múltiples plantillas
Para cargar múltiples plantillas con el paquete html/template de Go, puedes utilizar las funciones ParseFiles, ExecuteTemplate y ParseGlob.

### Uso de Must 
En Go, la función Must se utiliza en combinación con otras funciones para manejar errores de una manera más concisa. La función Must asume que el valor de retorno no es nulo y que no produce ningún error. Si ocurre un error, Must generará un pánico.

Aquí tienes un ejemplo de cómo se utiliza la función Must en el contexto del paquete `html/template:`
~~~go
// Controladores
func home(w http.ResponseWriter, r *http.Request) {
	// Mapa de funciones
	funcs := template.FuncMap{
		"greeting": greeting,
		"sum":      sum,
	}
	// Cargar y analizar el archivo de plantilla HTML
	template := template.Must(template.New("home.html").Funcs(funcs).ParseFiles("templates/home.html", "templates/base.html"))

	// Renderizar la plantilla en la respuesta
	template.Execute(w, nil)
}

...

func hello(w http.ResponseWriter, r *http.Request) {
	// Cargar y analizar el archivo de plantilla HTML
	template := template.Must(template.New("hello.html").ParseFiles("templates/hello.html", "templates/base.html"))
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
~~~

### Uso de ExecuteTemplate y ParseGlob
La función `ExecuteTemplate` y `ParseGlob` son funciones del paquete `html/template` de Go que se utilizan para cargar y ejecutar múltiples plantillas utilizando un patrón glob.

La función ParseGlob se utiliza para cargar múltiples plantillas desde el sistema de archivos utilizando un patrón glob y devuelve un objeto `*template.Template` que representa las plantillas cargadas.

Aquí tienes un ejemplo de cómo utilizar `ParseGlob` para cargar múltiples plantillas desde archivos:

~~~go
// Mapa de funciones
var funcs = template.FuncMap{
	"greeting": greeting,
	"sum":      sum,
}

// Gargar platillas
var templates = template.Must(template.New("T").Funcs(funcs).ParseGlob("templates/*.html"))

// Controladores
func home(w http.ResponseWriter, r *http.Request) {

	// Renderizar la plantilla en la respuesta
	err := templates.ExecuteTemplate(w, "home", nil)
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
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
	err := templates.ExecuteTemplate(w, "hello", user)
	if err != nil {
		panic(err)
	}
}
~~~

Definer las platillas en `home.html` y `hello.html`:
~~~html
<!-- home.html -->
{{ define "home"}}
...

{{ end }}

<!-- hello.html -->
{{ define "hello"}}
...

{{ end }}
~~~

---
## Rederizar plantillas
Crearemos una función reutilizable para renderizar plantillas y devolver un error personalizado y encabezado: 
~~~go
func renderTemplate(w http.ResponseWriter, name string, data any) {
	// Encabezado
	w.Header().Set("Content-Type", "text/html")

	// Renderizar la plantilla en la respuesta
	err := templates.ExecuteTemplate(w, "home", nil)
	if err != nil {
		http.Error(w, "No es posible retornar platilla", http.StatusInternalServerError)
	}
}
~~~ 

En esta sección, se establece el encabezado de la respuesta HTTP para indicar que el contenido que se enviará será de tipo HTML. Se utiliza `w.Header().Set("Content-Type", "text/html")` para establecer el encabezado `"Content-Type" en "text/html"`.

### Uso de encabezado
El encabezado en una respuesta HTTP se utiliza para proporcionar información sobre el contenido y la naturaleza de la respuesta. Especifica cómo el cliente (generalmente un navegador web) debe interpretar y procesar la respuesta recibida del servidor.

En el caso específico del encabezado "Content-Type" que se establece en "text/html", se está indicando al cliente que el contenido de la respuesta es un documento HTML. Esto es importante porque permite al cliente saber cómo debe interpretar y mostrar el contenido recibido. El navegador web, por ejemplo, utilizará esta información para renderizar correctamente la página HTML en el navegador del usuario.

El encabezado "Content-Type" es solo uno de los muchos encabezados que se pueden utilizar en una respuesta HTTP para proporcionar información adicional. Otros encabezados comunes incluyen "Content-Length" para indicar la longitud del contenido, "Cache-Control" para controlar la caché del navegador, "Location" para redireccionar el cliente a una nueva ubicación, entre otros.

En resumen, el encabezado en una respuesta HTTP se utiliza para proporcionar metadatos sobre la respuesta y permitir al cliente procesarla adecuadamente. Establecer el encabezado "Content-Type" es especialmente importante para indicar el tipo de contenido que se envía, como HTML, JSON, XML, etc.

Espero que esto aclare el propósito y la importancia del encabezado en una respuesta HTTP. Si tienes más preguntas, ¡no dudes en hacerlas!

---
## Página de error 
El código que proporcionaste muestra cómo puedes utilizar un archivo de plantilla HTML para personalizar la página de error en Go. Para ello, se utiliza el paquete html/template.

~~~go
var errorTemplates = template.Must(template.ParseGlob("templates/**/*.html"))

func handlerError(w http.ResponseWriter, name string, status int) {
	w.WriteHeader(status)
	errorTemplates.ExecuteTemplate(w, name, nil)
}

// código en el renderTemplate
func renderTemplate(w http.ResponseWriter, name string, data any) {
	// Encabezado
	w.Header().Set("Content-Type", "text/html")

	// Renderizar la plantilla en la respuesta
	err := templates.ExecuteTemplate(w, name, nil)
	if err != nil {
		handlerError(w, "500", http.StatusInternalServerError)
	}
}
~~~

La función `handlerError` en el código que proporcionaste es responsable de manejar el envío de una respuesta de error personalizada al cliente en función del estado proporcionado. 

### Error de página no encontrado 

~~~go
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Devolver un error personalizado para páginas no encontradas
	handlerError(w, "404", http.StatusNotFound)
}
~~~


Archivo de html de error dentro de `templates/error/500.html`:
~~~html
{{ define "500" }}
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Error</title>
</head>
<body>
    <h3>No es posible retornar platilla.</h3>
</body>
</html>

{{ end }}
~~~

Archivo de html de error dentro de `templates/error/404.html`:

~~~html
{{ define "404" }}
<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Error</title>
</head>
<body>
    <h3>404 - Página no econtrada</h3>
</body>
</html>

{{ end }}
~~~

---
## Resumen 

En la sección de "Manejo de plantillas" de nuestro curso de Go, aprendimos cómo utilizar plantillas para generar contenido dinámico en nuestras aplicaciones web. Exploramos varios aspectos del manejo de plantillas en Go.

Comenzamos por aprender cómo utilizar condicionales y ciclos en nuestras plantillas. Estas características nos permitieron controlar la ejecución y la apariencia del contenido en función de diferentes condiciones y datos proporcionados.

Luego, exploramos cómo trabajar con variables en las plantillas. Las variables nos permitieron almacenar y acceder a datos en nuestra plantilla, brindándonos flexibilidad y capacidad para adaptar dinámicamente el contenido generado.

También aprendimos a utilizar métodos y funciones en nuestras plantillas. Estas funciones nos permitieron manipular y transformar datos en tiempo de ejecución, agregando una capa adicional de flexibilidad y personalización a nuestras plantillas.

Además, discutimos la reutilización de plantillas. Vimos cómo crear plantillas base que contienen elementos comunes y cómo extender y personalizar estas plantillas base para generar diferentes vistas en nuestras aplicaciones.

A medida que avanzamos, también exploramos cómo cargar múltiples plantillas en nuestras aplicaciones y cómo renderizar estas plantillas para generar el contenido final que se envía al cliente.

Finalmente, abordamos el tema de las páginas de error. Aprendimos a personalizar las respuestas de error en nuestras aplicaciones web, utilizando plantillas para mostrar mensajes de error informativos y amigables para el usuario.

En resumen, en esta sección del curso de Go, adquiriste habilidades fundamentales para trabajar con plantillas en tus aplicaciones web. Te proporcionamos los conocimientos necesarios para utilizar condicionales, ciclos, variables, métodos y funciones en tus plantillas, así como también para cargar múltiples plantillas y personalizar las páginas de error. ¡Exploraste el emocionante mundo del manejo de plantillas en Go!




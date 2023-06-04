# Introducción 

1. [Introducción](#introducción)
2. [Crear proyecto web](#crear-proyecto-web)
3. [Router y handler](#router-y-handler)
4. [Modo debug con air](#modo-debug-con-air)
5. [Resumen](#resumen)

---
## Introducción
¡Bienvenidos a esta sección del curso de Go web! En esta parte del programa, exploraremos los fundamentos esenciales para crear un proyecto web utilizando el lenguaje de programación Go.

En la actualidad, el desarrollo web con Go se ha vuelto extremadamente popular debido a su rendimiento, eficiencia y facilidad de uso. En esta sección, nos centraremos en algunos aspectos clave para comenzar a construir aplicaciones web sólidas y de alto rendimiento utilizando Go.

El primer tema que abordaremos es la creación de un proyecto web. Aprenderemos cómo configurar una estructura básica de carpetas y archivos para nuestro proyecto, y cómo manejar las dependencias utilizando el administrador de paquetes Go Modules. Exploraremos las mejores prácticas para la organización del código y la estructura del proyecto.

A continuación, nos sumergiremos en el concepto de enrutamiento y manejadores (handlers) en Go. Un enrutador es una parte esencial de cualquier aplicación web, ya que dirige las solicitudes entrantes a las funciones de manejo adecuadas. Aprenderemos cómo utilizar el enrutador "mux" de Gorilla, uno de los enrutadores más populares y potentes disponibles para Go.

Una vez que comprendamos los conceptos básicos de enrutamiento y manejo de solicitudes, nos sumergiremos en el modo de depuración. El modo de depuración es crucial durante el desarrollo de una aplicación web, ya que nos permite detectar y solucionar errores de manera eficiente. Presentaremos una herramienta llamada "Air" que simplifica enormemente el proceso de desarrollo y depuración de aplicaciones Go web. Aprenderemos cómo configurar Air en nuestro proyecto y cómo aprovechar sus capacidades para una experiencia de desarrollo fluida y sin interrupciones.

Al final de esta sección, estarás equipado con los conocimientos necesarios para crear un proyecto web en Go, implementar enrutadores y manipular solicitudes entrantes utilizando manejadores eficientes. Además, dominarás el modo de depuración con Air, lo que te permitirá acelerar tu proceso de desarrollo y corregir rápidamente cualquier problema que surja.

¡Comencemos esta emocionante sección y desbloqueemos todo el potencial del desarrollo web con Go!

---
## Crear proyecto web

Para crear un proyecto web "Hola Mundo" en Go llamado "rpsweb". Aquí tienes los pasos para crearlo:

- Crea un directorio para tu proyecto "rpsweb" y navega hasta él en la línea de comandos.
- Dentro del directorio del proyecto, crea un archivo llamado `main.go` y ábrelo en tu editor de código preferido.
- En el archivo `main.go`, escribe el siguiente código:

~~~go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Hola Mundo!")
	})

    port := ":8080" // Puerto en el que se ejecutará el servidor
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

~~~

Guarda el archivo `main.go`.

Desde la línea de comandos, dentro del directorio del proyecto, ejecuta el siguiente comando para compilar y ejecutar el proyecto:

~~~bash
go run main.go
~~~

Ahora puedes abrir tu navegador web y acceder a http://localhost:8080/. Verás el mensaje "¡Hola Mundo!" en la página.
¡Y eso es todo! Has creado un proyecto web "Hola Mundo" en Go llamado "rpsweb". Puedes personalizar el mensaje o agregar más rutas y funcionalidades según tus necesidades.

- Aquí se utiliza `http.HandleFunc()` para registrar una función de controlador para la ruta `"/"` (la ruta raíz del servidor). La función de controlador toma dos argumentos: `w (http.ResponseWriter)` y `r (http.Request)`. El `http.ResponseWriter` se utiliza para enviar la respuesta al cliente, y `http.Request` representa la solicitud recibida del cliente. En este caso, utilizamos `fmt.Fprintf()` para escribir el mensaje `"¡Hola Mundo!"` en el `http.ResponseWriter`, lo que enviará ese mensaje como respuesta al cliente.

- En esta línea, `http.ListenAndServe()` se utiliza para iniciar el servidor web en el puerto `8080`. Toma dos argumentos: el primero es la dirección en la que el servidor debe escuchar, en este caso, `":8080"` indica que se escuchará en el puerto `8080` en todas las interfaces disponibles. El segundo argumento es un `http.Handler` que manejará las solicitudes entrantes. Al pasar `nil`, se utiliza el enrutador predeterminado de Go para manejar las solicitudes. En este caso, como solo se registra una única ruta "/", el enrutador predeterminado se encargará de manejar esa ruta específicamente.

---
## Router y handler

En Go, un router y un handler son dos conceptos clave en el desarrollo de aplicaciones web. Te explicaré brevemente qué son y cómo se utilizan:

- Router (enrutador):
Un router es un componente que se encarga de dirigir las solicitudes HTTP a las funciones de handler correspondientes según la ruta solicitada. En Go, hay varios paquetes populares que proporcionan enrutadores, como "net/http" y "github.com/gorilla/mux". El router se configura para asociar una ruta (URL) con una función de handler específica que se ejecutará cuando se acceda a esa ruta en el servidor.

- Handler (manejador):
Un handler es una función que se encarga de manejar una solicitud HTTP específica. En Go, un handler se representa mediante la interfaz `http.Handler`. Puede ser una función o un objeto que implementa el método `ServeHTTP(ResponseWriter, *Request)`.

En el siguiente ejemplo se utilizan los enrutadores y handlers del paquete `net/http` de Go para definir y manejar diferentes rutas en un servidor web. A continuación, te explico cada línea y qué hace cada handler:

~~~go
	// Crear enrutador
	router := http.NewServeMux()

	// Configurar rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	
    port := ":8080" // Puerto en el que se ejecutará el servidor
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	http.ListenAndServe(port, router)

~~~
`router := http.NewServeMux()`: Aquí se crea un nuevo enrutador utilizando `http.NewServeMux()`. El enrutador es responsable de asignar las rutas a las funciones controladoras correspondientes.

A continuación, te mostraré un ejemplo de cómo se podrían definir los handlers en el paquete `handlers`:

~~~go
package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Página de inicio")
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Crear nuevo juego")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Juego")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Jugar")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Acerca de")
}
~~~

---
## Modo debug con air
Para utilizar el modo de depuración (debug) con la herramienta Air en Go, puedes seguir los siguientes pasos:

Instalar Air: Comienza por instalar Air, una herramienta de reinicio automático similar a Fresh que también admite el modo de depuración. Puedes instalarlo ejecutando el siguiente comando en tu terminal:

A través de go install
Con go 1.18 o superior:
~~~bash
go install github.com/cosmtrek/air@latest
~~~

Puede inicializar el `.air.toml` archivo de configuración en el directorio actual con la configuración predeterminada ejecutando el siguiente comando.

~~~go
air init
~~~

Después de esto, puede ejecutar el aircomando sin argumentos adicionales y usará el `.air.toml` archivo para la configuración.

~~~go
air
~~~

### Error al ejecutar el air 
Agre la siguiente configuración `.profile` o `.basherc` para que puedas usar la instalación de `air`:

~~~bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
~~~

- `export GOPATH=$HOME/go`: Esta línea establece el valor de la variable de entorno GOPATH como $HOME/go. La variable GOPATH define la ubicación de la raíz del árbol de directorios donde Go busca y guarda los paquetes y el código fuente. Al establecer GOPATH como $HOME/go, estás indicando que el directorio go dentro de tu directorio de inicio ($HOME) es el directorio raíz para los paquetes y el código fuente relacionados con Go.

- `export PATH=$PATH:$GOROOT/bin:$GOPATH/bin`: Esta línea agrega los directorios $GOROOT/bin y $GOPATH/bin a la variable PATH. $GOROOT es una variable de entorno que se refiere a la ubicación de la instalación de Go en tu sistema. Al agregar $GOROOT/bin a PATH, permites que los ejecutables específicos de Go que se encuentran en ese directorio estén disponibles desde cualquier ubicación en tu terminal. $GOPATH/bin se agrega a PATH para permitir que los ejecutables de tus proyectos de Go (por ejemplo, aquellos que instalas mediante go install) también estén disponibles en tu entorno.

---
## Resumen 
En esta sección del curso de Go web, exploramos los fundamentos esenciales para crear un proyecto web utilizando el lenguaje de programación Go. Nos centramos en tres temas principales: la creación de un proyecto web, el enrutamiento y los manejadores, y el modo de depuración con Air.

Comenzamos estableciendo una estructura básica de carpetas y archivos para nuestro proyecto web y aprendimos a manejar las dependencias utilizando Go Modules. También nos familiarizamos con las mejores prácticas para la organización del código y la estructura del proyecto.

Luego, nos sumergimos en el enrutamiento y los manejadores en Go. Utilizamos el enrutador "mux" de Gorilla, un enrutador popular y potente, para dirigir las solicitudes entrantes a las funciones de manejo adecuadas. Aprendimos cómo configurar rutas y manipular las solicitudes y respuestas HTTP.

Finalmente, nos adentramos en el modo de depuración utilizando Air. Configuramos esta herramienta en nuestro proyecto y aprovechamos sus capacidades para facilitar el desarrollo y la corrección de errores. El modo de depuración nos permitió detectar y solucionar problemas de manera eficiente durante el desarrollo de nuestra aplicación web.

Al final de la sección, adquirimos los conocimientos necesarios para crear un proyecto web en Go, implementar enrutadores y manejar solicitudes entrantes. Además, aprendimos a utilizar el modo de depuración con Air para acelerar el proceso de desarrollo y solucionar rápidamente cualquier problema que surgiera.

En resumen, esta sección del curso nos proporcionó una base sólida para construir aplicaciones web utilizando Go. Aprendimos a crear proyectos, implementar enrutadores y manejar solicitudes, así como a utilizar el modo de depuración para un desarrollo eficiente. Con estos conocimientos, podemos aprovechar todo el potencial del desarrollo web con Go.
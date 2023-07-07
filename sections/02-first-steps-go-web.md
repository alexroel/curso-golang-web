# Primeros pasos en Go web

1. [Introducció](#introducció)
2. [Arquitectura MVC](#arquitectura-mvc)
3. [Creación de un servidor web con Go](#creación-de-un-servidor-web-con-go)
4. [Controlador y Ruta](#controlador-y-ruta)
5. [Enrutador](#enrutador)
6. [Obtener argumentos de URL](#obtener-argumentos-de-url)
7. [Modo desarrollador con AIR](#modo-desarrollador-con-air)
8. [Resumen](#resumen)

---
## Introducció
En esta sección, te adentrarás en el apasionante mundo del desarrollo web con Go, un lenguaje de programación eficiente y poderoso. Aprenderás los conceptos fundamentales y las herramientas necesarias para construir aplicaciones web utilizando Go.

- Arquitectura MVC: Comenzaremos explorando la arquitectura Modelo-Vista-Controlador (MVC), que es un patrón de diseño ampliamente utilizado en el desarrollo web. Entenderemos cómo se divide la aplicación en capas lógicas, el propósito de cada una y cómo se comunican entre sí. La arquitectura MVC es especialmente adecuada para el desarrollo de aplicaciones web en Go.

- Creación de un servidor web con Go: Aprenderás cómo crear un servidor web básico utilizando el paquete net/http de la biblioteca estándar de Go.

- Controlador y Ruta: Sumergiéndonos más en el desarrollo web, entenderás el concepto de controladores y rutas en Go. Aprenderás cómo escribir funciones de controlador para manejar diferentes rutas y cómo establecer estas rutas en el enrutador. Esto permitirá una navegación eficiente y organizada en tu aplicación web.

- Enrutador: El enrutador es una parte clave del desarrollo web en Go. Exploraremos cómo utilizar el paquete net/http para configurar un enrutador personalizado y dirigir las solicitudes a los controladores adecuados. Con un enrutador bien estructurado, tu aplicación web será fácil de mantener y escalar.

- Obtener argumentos de URL: Las URL a menudo contienen argumentos que son esenciales para el funcionamiento de nuestras aplicaciones web. Aprenderás cómo obtener y utilizar estos argumentos en los controladores de Go. Veremos cómo acceder a los argumentos de la URL y utilizarlos para personalizar las respuestas y realizar acciones específicas.

- Modo desarrollador con AIR: Para agilizar el proceso de desarrollo, te presentaremos AIR, una herramienta que facilita la recarga automática de tu aplicación web en tiempo real mientras realizas cambios en el código fuente. Descubrirás cómo utilizar AIR para acelerar tu flujo de trabajo y ver los cambios reflejados instantáneamente en tu aplicación web.

¡Prepárate para dar tus primeros pasos en el desarrollo web con Go! En la siguiente sección, profundizaremos en cada uno de estos temas, brindándote ejemplos prácticos y ejercicios para fortalecer tus habilidades en el desarrollo web con Go.

---
## Arquitectura MVC
En el desarrollo web con Go, es común seguir una arquitectura de aplicación MVC (Modelo-Vista-Controlador) o una variante de ella. Esta arquitectura ayuda a organizar el código de manera modular y separar las responsabilidades en distintas capas. A continuación, describiré los componentes principales de una arquitectura típica de desarrollo web con Go:

- **Modelo (Model)**: El modelo representa los datos y la lógica de negocio de la aplicación. Aquí se definen las estructuras de datos, las funciones y métodos para acceder y manipular dichos datos. Esto puede incluir interacciones con la base de datos, servicios externos u otras fuentes de datos.

- **Vista (View)**: La vista es la capa encargada de la presentación de los datos al usuario. En Go, la vista generalmente se representa utilizando plantillas HTML o motores de renderizado. Esta capa se encarga de generar el contenido HTML que se enviará al navegador del cliente. Las vistas no deben contener lógica de negocio, solo se enfocan en la presentación de los datos.

- **Controlador (Controller)**: El controlador es el intermediario entre las vistas y los modelos. Recibe las solicitudes del cliente, realiza las validaciones necesarias, interactúa con los modelos para obtener o modificar los datos y finalmente selecciona la vista adecuada para enviar la respuesta al cliente. El controlador es responsable de manejar el flujo de la aplicación y coordinar las acciones entre el modelo y la vista.

- **Enrutador (Router)**: El enrutador es un componente esencial en el desarrollo web con Go. Se encarga de dirigir las solicitudes del cliente a los controladores correspondientes según la ruta solicitada. Puede utilizar bibliotecas populares como Gorilla Mux o el enrutador estándar de Go para manejar el enrutamiento de manera eficiente.

- **Capa de persistencia de datos**: En una arquitectura web típica, se utiliza una capa de persistencia de datos para interactuar con la base de datos. Puedes utilizar un ORM (Object-Relational Mapping) como GORM o interactuar directamente con la base de datos utilizando el paquete database/sql estándar de Go.

- **Capa de servicios**: En ocasiones, es útil introducir una capa de servicios entre los controladores y los modelos. Esta capa se encarga de implementar la lógica de negocio más compleja y coordinar las operaciones entre varios modelos. Los servicios pueden agrupar funcionalidades relacionadas y ayudar a mantener una estructura modular y organizada.

- **Middleware**: El middleware es un concepto importante en Go que permite agregar funcionalidades adicionales a la cadena de manejo de solicitudes antes de que lleguen a los controladores. Puedes utilizar middleware para agregar autenticación, autorización, logging, compresión de respuesta y otros aspectos comunes en el desarrollo web.

Es importante mencionar que esta es una descripción general de la arquitectura de desarrollo web con Go y existen variaciones y enfoques adicionales según los requerimientos de cada proyecto. La modularidad y la escalabilidad son aspectos clave que debes considerar al diseñar la arquitectura de tu aplicación web con Go.


---
## Creación de un servidor web con Go
Para crear un servidor web básico con Go, necesitarás utilizar el paquete net/http que viene incluido en la biblioteca estándar de Go.

~~~go
// Creación de un servidor web con Go
port := ":8080" // Puerto en el que se ejecutará el servidor
log.Printf("Servidor escuchando en http://localhost%s\n", port)
log.Fatal(http.ListenAndServe(port, nil))
~~~
En este fragmento de código:

- `port := ":8080"`: Se declara una variable port que contiene el número de puerto en el que se ejecutará el servidor. En este caso, se ha asignado el puerto 8080, pero puedes cambiarlo según tus necesidades.

- `log.Printf("Servidor escuchando en http://localhost%s\n", port)`: Se utiliza la función `Printf` del paquete `log` para imprimir un mensaje en la consola indicando que el servidor está escuchando en la dirección` http://localhost` seguida del número de puerto. El símbolo %s se utiliza como marcador de posición para el valor de la variable port en la cadena de formato.

- `log.Fatal(http.ListenAndServe(port, nil))`: La función ListenAndServe del paquete net/http se utiliza para iniciar el servidor en el puerto especificado. Esta función toma dos argumentos: el primero es la dirección en la que el servidor escuchará las solicitudes, en este caso, se utiliza la variable port para indicar el puerto. El segundo argumento es un enrutador, que en este caso se ha pasado como nil, lo que significa que se utilizará el enrutador por defecto proporcionado por el paquete net/http.

- El uso de log.Fatal garantiza que si ocurre un error al iniciar el servidor, se mostrará un mensaje de error y el programa se detendrá.

En resumen, este fragmento de código declara el puerto en el que se ejecutará el servidor, imprime un mensaje en la consola indicando la dirección y puerto del servidor, y finalmente inicia el servidor en el puerto especificado.

---
## Controlador y Ruta
En el contexto del desarrollo web con Go, el controlador se refiere a una función que se encarga de manejar una solicitud HTTP específica. Una ruta, por otro lado, es la dirección URL que se utiliza para acceder a un recurso en el servidor.

Aquí tienes un ejemplo de cómo definir un controlador y una ruta en Go utilizando el paquete net/http:
~~~go
// Controlador 
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hola Mundo!")
}

func main() {
	// Rutas 
	http.HandleFunc("/hello", hello)

	// Creación de un servidor web con Go
	port := ":8080" // Puerto en el que se ejecutará el servidor
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
~~~

~~~go
// Controlador 
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hola Mundo!")
}
~~~
En esta parte, se define una función llamada hello que actúa como controlador. Recibe dos parámetros: w http.ResponseWriter, que se utiliza para enviar la respuesta al cliente, y r *http.Request, que proporciona información sobre la solicitud entrante. Dentro del controlador, se utiliza fmt.Fprintf para escribir el mensaje "¡Hola Mundo!" en la respuesta.

~~~go
func main() {
	// Rutas
	http.HandleFunc("/hello", hello)

    ...
}
~~~
En la función `main`, se establecen las rutas y se crea el servidor web utilizando el paquete net/http.

- `http.HandleFunc("/hello", hello)`: Se utiliza la función http.HandleFunc para establecer la ruta "/hello" y asociarla con el controlador hello. Esto significa que cuando un cliente realice una solicitud a "/hello", se ejecutará el controlador hello.

Otra forma de probar usando curl: 
~~~bash
curl get http://localhost:8080/hello
~~~

--- 
## Enrutador 
En el paquete net/http de Go, el enrutador se implementa utilizando el tipo http.ServeMux. Un ServeMux es un multiplexor de servidores HTTP que se utiliza para dirigir las solicitudes a los controladores correspondientes.

Aquí tienes un ejemplo de cómo crear y utilizar un enrutador ServeMux en Go:
~~~go
package main

import (
	"fmt"
	"log"
	"net/http"
)

// Controladores
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Bienvenido a la página de inicio!")
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hola Mundo!")
}

func main() {
	// Crear un enrutaro
	router := http.NewServeMux()

	// Definir rutas y controladores
	router.HandleFunc("/", home)
	router.HandleFunc("/hello", hello)

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
~~~
En este ejemplo:

- Se crea un enrutador utilizando http.NewServeMux(), que devuelve una nueva instancia de ServeMux.

- Se utilizan los métodos `HandleFunc` del enrutador para asociar rutas y controladores. Cada llamada a `HandleFunc` toma una ruta como primer argumento y una función de controlador como segundo argumento. En el ejemplo, se definen dos rutas: "/" y "/hello", y se asocian a las funciones de controlador home y hello, respectivamente.

- Se crea un servidor HTTP utilizando `http.Server` y se le asigna el enrutador como `Handler`.

- Finalmente, se inicia el servidor llamando a `server.ListenAndServe()`.

Utilizando el enrutador ServeMux, puedes definir múltiples rutas y controladores, y dirigir las solicitudes entrantes a los controladores apropiados en función de la ruta especificada. Esto proporciona un enfoque flexible y modular para el enrutamiento en aplicaciones web con Go.

---
## Obtener argumentos de URL
En Go, para obtener los argumentos de una URL, puedes utilizar el paquete net/http y la estructura http.Request. El objeto http.Request contiene información sobre la solicitud HTTP, incluidos los argumentos de la URL.

~~~go
func hello(w http.ResponseWriter, r *http.Request) {
	// Obtener los argumentos de la URL
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	age := queryParams.Get("age")

	// Imprimir los datos y validación
	if name != "" && age != "" {
		fmt.Fprintf(w, "¡Hola %s! Tienes %s años.", name, age)
	} else {
		fmt.Fprintf(w, "¡Hola Mundo!")
	}

}
~~~

- `r.URL.Query()` para obtener un objeto url.Values que contiene los argumentos de la URL. Luego, se utiliza Get() del objeto url.Values para obtener el valor de los argumentos "name" y "age" de la URL.

En esta sección, se realiza una validación para verificar si tanto el nombre (name) como la edad (age) tienen valores no vacíos. Si ambos argumentos tienen valores no vacíos, se imprime un mensaje personalizado utilizando fmt.Fprintf(), donde se muestra el saludo "¡Hola {nombre}! Tienes {edad} años.". En caso contrario, si alguno de los argumentos falta o está vacío, se imprime el mensaje "¡Hola Mundo!".

De esta manera, el controlador hello responde a la solicitud con un saludo personalizado si se proporcionan los argumentos de la URL ("name" y "age"), y responde con un saludo genérico "¡Hola Mundo!" en caso contrario.

Recuerda que este es solo un ejemplo de validación y puedes ajustarlo según tus necesidades y requisitos específicos para los datos de la URL.

---
## Modo desarrollador con AIR
En Go, puedes activar el "modo desarrollador" utilizando la biblioteca Air. Air es una herramienta de desarrollo en tiempo real para Go que recarga automáticamente tu aplicación cada vez que guardas un archivo.

Aquí tienes los pasos para configurar el modo desarrollador con Air en Go:

### Paso 1: Instalar Air
Primero, debes instalar la herramienta Air utilizando el siguiente comando:

~~~go
go install github.com/cosmtrek/air@latest
~~~

### Paso 2: Crear un archivo de configuración

Crea un archivo de configuración llamado `.air.toml` en el directorio raíz de tu proyecto. Aquí tienes un ejemplo de cómo se vería el archivo `.air.toml`. Tambien lo puedes hacer susando el siguiente comando:
~~~bash
air init
~~~

Y tendras el siguiente resultado en el archivo `.air.toml`:

~~~toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ."
  delay = 0
~~~

Para levantar el servidor web jecutas `air`.

~~~bash
air

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ , built with Go

watching .
watching config
watching handlers
watching models
watching templates
watching templates/auth
watching templates/task
!exclude tmp
building...
~~~

---
## Resumen 
En esta sección, nos adentramos en el apasionante mundo del desarrollo web con Go, un lenguaje de programación eficiente y poderoso. Aprendimos los conceptos fundamentales y las herramientas necesarias para construir aplicaciones web utilizando Go.

Arquitectura MVC: Comenzamos explorando la arquitectura Modelo-Vista-Controlador (MVC), un patrón de diseño ampliamente utilizado en el desarrollo web. Entendimos cómo se divide la aplicación en capas lógicas, el propósito de cada una y cómo se comunican entre sí. La arquitectura MVC resultó especialmente adecuada para el desarrollo de aplicaciones web en Go.

Creación de un servidor web con Go: Aprendimos cómo crear un servidor web básico utilizando el paquete net/http de la biblioteca estándar de Go.

Controlador y Ruta: Profundizamos en el concepto de controladores y rutas en Go. Aprendimos cómo escribir funciones de controlador para manejar diferentes rutas y cómo establecer estas rutas en el enrutador. Esto permitirá una navegación eficiente y organizada en nuestra aplicación web.

Enrutador: Exploramos cómo utilizar el paquete net/http para configurar un enrutador personalizado y dirigir las solicitudes a los controladores adecuados. Con un enrutador bien estructurado, nuestra aplicación web será fácil de mantener y escalar.

Obtener argumentos de URL: Descubrimos cómo obtener y utilizar los argumentos de una URL en los controladores de Go. Aprendimos cómo acceder a los argumentos de la URL y utilizarlos para personalizar las respuestas y realizar acciones específicas en nuestra aplicación web.

Modo desarrollador con AIR: Conocimos AIR, una herramienta que facilita la recarga automática de nuestra aplicación web en tiempo real mientras realizamos cambios en el código fuente. Aprendimos cómo utilizar AIR para acelerar nuestro flujo de trabajo y ver los cambios reflejados instantáneamente en nuestra aplicación web.

¡Estamos listos para dar nuestros primeros pasos en el desarrollo web con Go! En la siguiente sección, profundizaremos en cada uno de estos temas, brindándote ejemplos prácticos y ejercicios para fortalecer tus habilidades en el desarrollo web con Go.
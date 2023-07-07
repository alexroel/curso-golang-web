# Enrrutador gorilla/mux

1. [Introducción](#introducción)
2. [Instalación de gorilla/mux](#instalación-de-gorillamux)
3. [Uso de gorilla/mux](#uso-de-gorillamux)
4. [Parametros en rutas](#parametros-en-rutas)
5. [Error personalizado](#error-personalizado)
6. [Resumen](#resumen)

---
## Introducción
¡Bienvenidos a la sección de enrutador Gorilla/mux de nuestro curso de desarrollo web con Go! En esta parte del curso, exploraremos las poderosas capacidades de Gorilla/mux como enrutador para nuestras aplicaciones web.

Comenzaremos con la instalación de Gorilla/mux, una biblioteca popular y altamente flexible para el enrutamiento en aplicaciones web en Go. Aprenderemos cómo configurar y utilizar Gorilla/mux en nuestros proyectos, aprovechando sus numerosas características y funcionalidades.

En esta sección también nos adentraremos en los parámetros en las rutas. Veremos cómo Gorilla/mux nos permite capturar y utilizar de forma eficiente los parámetros presentes en las URL, lo que nos brinda un gran poder y flexibilidad para diseñar rutas dinámicas en nuestras aplicaciones.

Además, exploraremos la capacidad de Gorilla/mux para personalizar los mensajes de error. Aprenderemos cómo manejar y presentar errores de forma más amigable al usuario, brindando información útil y comprensible en caso de que ocurra algún problema en nuestras rutas.

¡Prepárate para sumergirte en el mundo del enrutamiento web con Gorilla/mux! Al final de esta sección, estarás equipado con las habilidades necesarias para crear aplicaciones web robustas y flexibles en Go. ¡Comencemos con la instalación de Gorilla/mux y exploremos las posibilidades infinitas que nos ofrece esta biblioteca poderosa!

---
## Instalación de gorilla/mux
Para instalar Gorilla/mux en tu proyecto de Go, debes seguir los siguientes pasos:

- Asegúrate de tener Go instalado en tu sistema. Puedes verificarlo ejecutando el siguiente comando en la línea de comandos:

~~~bash
go version
~~~

Si Go no está instalado, puedes descargarlo e instalarlo desde el sitio oficial: https://golang.org/dl/

Abre una terminal o línea de comandos y navega hasta el directorio raíz de tu proyecto de Go.

- Utiliza el comando `go get para instalar Gorilla/mux` y sus dependencias. Ejecuta el siguiente comando:

~~~bash
go get -u github.com/gorilla/mux
~~~

Esto descargará e instalará Gorilla/mux en tu proyecto, así como sus dependencias, en la carpeta $GOPATH/src/github.com/gorilla/mux.

- Una vez que la instalación se complete, puedes importar Gorilla/mux en tus archivos Go. Agrega la siguiente línea de importación en los archivos donde desees utilizar la biblioteca:

~~~go
import "github.com/gorilla/mux"
~~~

Con esta importación, podrás utilizar las funcionalidades de Gorilla/mux en tu código.

¡Y eso es todo! Ahora tienes Gorilla/mux instalado en tu proyecto de Go y estás listo para utilizarlo para construir rutas y manejar solicitudes HTTP de manera flexible y poderosa.

---
## Uso de gorilla/mux
Para usar gorilla/mux primero pasas a importar en tu proyecto. 

~~~go
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
~~~

- En la función main(), en lugar de utilizar http.NewServeMux() para crear el enrutador, utiliza mux.NewRouter() de Gorilla/mux para crear el enrutador:

- Reemplaza las llamadas a router.HandleFunc() con las funciones equivalentes proporcionadas por Gorilla/mux. La función HandleFunc() de Gorilla/mux tiene una sintaxis ligeramente diferente y permite un mayor control sobre las rutas y las variables de ruta. Puedes actualizar tu código de esta manera:

~~~go
    // Crear un enrutaro
	router := mux.NewRouter()

	// Definir rutas y controladores
	router.HandleFunc("/", home)
	router.HandleFunc("/hello", hello)
~~~

Con estos cambios, tu código estará utilizando Gorilla/mux para el enrutamiento. Ahora podrás aprovechar las características adicionales que ofrece Gorilla/mux, como el enrutamiento avanzado, el manejo de variables de ruta y la personalización de páginas no encontradas.

---
## Parametros en rutas
Con Gorilla/mux, puedes definir rutas con parámetros utilizando llaves {} en las rutas. Estos parámetros se capturan y se pueden utilizar en tu controlador para procesar la solicitud de manera dinámica. Aquí tienes un ejemplo de cómo puedes trabajar con parámetros de ruta utilizando Gorilla/mux:

~~~go
func hello(w http.ResponseWriter, r *http.Request) {
	// Obtener los argumentos de la URL
	vars := mux.Vars(r)
	name := vars["name"]
	age := vars["age"]

	fmt.Fprintf(w, "¡Hola %s! Tienes %s años.", name, age)
}

func main() {
	// Crear un enrutaro
	router := mux.NewRouter()

	// Definir rutas y controladores
	router.HandleFunc("/", home)
	router.HandleFunc("/hello/{name}/{age:[0-9]+}", hello)

    ...

}
~~~

En el código que proporcionaste, vars := mux.Vars(r) es una línea de código que se utiliza para obtener los valores de los parámetros de ruta capturados en una solicitud HTTP.

Cuando defines una ruta con Gorilla/mux que incluye parámetros de ruta, como "/hello/{name}/{age}", puedes usar mux.Vars(r) para acceder a los valores de esos parámetros. La función mux.Vars() toma como argumento un objeto *http.Request y devuelve un mapa `(map[string]string)` que contiene los nombres de los parámetros de ruta y sus respectivos valores.

En el código proporcionado, se utiliza mux.Vars(r) para obtener el mapa vars que contiene los valores de los parámetros de ruta capturados en la solicitud HTTP. Luego, se utilizan las claves de ese mapa (name y age) para acceder a los valores específicos de los parámetros de ruta.

Por ejemplo, si la ruta capturada es "/hello/john/25", mux.Vars(r) devolverá un mapa con las siguientes asignaciones de clave-valor:

~~~go
{
	"name": "john",
	"age": "25"
}
~~~

Entonces, name := vars["name"] asignará el valor "john" a la variable name, y age := vars["age"] asignará el valor "25" a la variable age. Luego, puedes utilizar estos valores en tu código para realizar acciones específicas según los parámetros de ruta capturados.

En resumen, mux.Vars(r) es una función que te permite acceder a los valores de los parámetros de ruta capturados en una solicitud HTTP utilizando Gorilla/mux. Puedes asignar esos valores a variables individuales para utilizarlos en tu código según sea necesario.

---
## Error personalizado
Para manejar errores personalizados para páginas no existentes (404) utilizando Gorilla/mux, puedes utilizar el enrutador (mux.Router) y establecer un controlador para la ruta no encontrada.

Aquí tienes un ejemplo de cómo implementarlo:

~~~go
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Devolver un error personalizado para páginas no encontradas
	http.Error(w, "404 - Página no encontrada", http.StatusNotFound)
}

func main() {
	// Crear un enrutaro
	router := mux.NewRouter()

    ....

	// Manejo personalizado para página no encontrada (404)
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

    ...
}
~~~


En este ejemplo, http.Error establece automáticamente el código de estado HTTP 404 y escribe el mensaje de error proporcionado en el cuerpo de la respuesta. Al utilizar http.Error, no es necesario llamar explícitamente a w.WriteHeader(http.StatusNotFound).

Puedes personalizar el mensaje de error según tus necesidades, simplemente reemplaza el string "Página no encontrada" con el mensaje que deseas mostrar.

Recuerda que cuando se utiliza http.Error, el controlador automáticamente finaliza el proceso de respuesta, por lo que no es necesario incluir return después de la llamada a http.Error.


---
## Resumen

En esta sección de nuestro curso de desarrollo web con Go, exploramos el enrutador Gorilla/mux y sus capacidades. Comenzamos instalando Gorilla/mux, una biblioteca popular y altamente flexible para el enrutamiento en aplicaciones web en Go. Aprendimos cómo configurar y utilizar Gorilla/mux en nuestros proyectos, aprovechando sus numerosas características y funcionalidades.

Durante el curso, nos adentramos en los parámetros en las rutas y descubrimos cómo Gorilla/mux nos permitía capturar y utilizar de forma eficiente los parámetros presentes en las URL. Esta característica nos brindó un gran poder y flexibilidad para diseñar rutas dinámicas en nuestras aplicaciones.

Además, exploramos la capacidad de Gorilla/mux para personalizar los mensajes de error. Aprendimos cómo manejar y presentar errores de forma más amigable al usuario, brindando información útil y comprensible en caso de que ocurriera algún problema en nuestras rutas.

Al finalizar esta sección, los estudiantes adquirieron las habilidades necesarias para crear aplicaciones web robustas y flexibles en Go utilizando Gorilla/mux.
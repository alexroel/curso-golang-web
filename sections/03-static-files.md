# Archivos estáticos 

1. [Introducción](#introducción)
2. [Configuración](#configuración)
3. [Agregando CSS y Bootstrap 5](#agregando-css-y-bootstrap-5)
4. [Agregando ímagenes](#agregando-ímagenes)
5. [Agregando  JavaScript](#agregando-javascript)
6. [Resumen](#resumen)

---
# Introducción

En esta sección, aprenderemos cómo trabajar con archivos estáticos en una aplicación web desarrollada con el lenguaje de programación Go. Los archivos estáticos, como hojas de estilo CSS, imágenes y scripts de JavaScript, son elementos fundamentales para el diseño y la interactividad de un sitio web. Aprenderemos cómo configurar y servir estos archivos estáticos correctamente en nuestro servidor Go.

Temas que se cubrirán en esta sección:

- Configuración inicial: Veremos cómo organizar la estructura de directorios de nuestro proyecto Go para manejar los archivos estáticos de manera eficiente. Crearemos una carpeta "static" en la raíz del proyecto y aprenderemos a configurar un enrutador para servir estos archivos de manera adecuada.

- Agregando CSS y Bootstrap 5: Aprenderemos cómo utilizar hojas de estilo CSS para dar estilo y diseño a nuestras páginas web. Además, exploraremos cómo agregar Bootstrap 5, uno de los frameworks de CSS más populares, a nuestro proyecto. Veremos cómo enlazar los archivos CSS de Bootstrap correctamente y aplicar estilos predefinidos a nuestros elementos HTML.

- Agregando imágenes: Las imágenes son elementos esenciales en cualquier sitio web. Aprenderemos cómo agregar imágenes estáticas a nuestro proyecto Go y cómo mostrarlas correctamente en nuestras páginas HTML. Veremos cómo organizar las imágenes en una carpeta "images" dentro de la carpeta "static" y cómo enlazar las rutas de las imágenes en nuestro código HTML.

- Agregando JavaScript: El JavaScript es un lenguaje de programación fundamental para agregar interactividad y dinamismo a nuestras aplicaciones web. Aprenderemos cómo agregar archivos de JavaScript estáticos a nuestro proyecto Go y cómo ejecutar código JavaScript en nuestras páginas HTML. Veremos cómo enlazar los archivos JavaScript y cómo utilizar bibliotecas y scripts personalizados para mejorar la funcionalidad de nuestro sitio web.

Al final de esta sección, estarás familiarizado con el manejo de archivos estáticos en Go y podrás crear aplicaciones web completas con estilos atractivos, imágenes llamativas y funciones interactivas utilizando CSS, Bootstrap, imágenes y JavaScript. ¡Comencemos!


---
## Configuración  
En una aplicación web, los archivos estáticos se refieren a archivos como CSS, JavaScript, imágenes u otros recursos que se sirven directamente desde el servidor sin procesamiento adicional. Go proporciona una manera fácil de servir archivos estáticos utilizando el paquete `http.FileServer`.

~~~go
// Manejador para servir los archivos estáticos
	fs := http.FileServer(http.Dir("static"))
	// Ruta para acceder a los archivos estáticos
	router.Handle("/static/", http.StripPrefix("/static/", fs))
~~~

En el código, se utiliza `http.FileServer` y` http.StripPrefix` para servir archivos estáticos en Go.

- `http.FileServer(http.Dir("static"))` crea un servidor de archivos estáticos que sirve archivos desde el directorio "static". El argumento `http.Dir("static") `especifica el directorio desde el cual se deben servir los archivos estáticos.

- `http.StripPrefix("/static/", fs)` se utiliza para eliminar el prefijo "/static/" de las rutas de solicitud antes de pasarlas al servidor de archivos estáticos. Esto es útil cuando deseas que las rutas de tus archivos estáticos comiencen sin el prefijo "/static/".

- `router.Handle("/static/", http.StripPrefix("/static/", fs))` registra el manejador de archivos estáticos en el enrutador. Cuando una solicitud coincide con la ruta "/static/", se pasa al servidor de archivos estáticos después de eliminar el prefijo "/static/" utilizando http.StripPrefix.

---
## Agregando CSS y Bootstrap 5

Para agregar CSS y Bootstrap 5 a tu proyecto Go y utilizarlos en tus archivos estáticos, puedes seguir los siguientes pasos:

Descarga Bootstrap 5: Visita el sitio web oficial de Bootstrap (https://getbootstrap.com/) y descarga la versión más reciente de Bootstrap 5. Extrae los archivos del paquete descargado.

Crea una carpeta "static" en el directorio de tu proyecto Go si aún no existe. Coloca los archivos CSS de Bootstrap (como bootstrap.min.css) en la carpeta "static".

Crea un archivo HTML (por ejemplo, index.html) en el directorio de tu proyecto Go.

En el archivo HTML, agrega las etiquetas HTML necesarias y enlaza los archivos CSS de Bootstrap.

~~~html
<!DOCTYPE html>
<html>
<head>
    <title>Proyecto con Bootstrap 5</title>
    <link rel="stylesheet" href="/static/bootstrap.min.css">
    <link rel="stylesheet" href="static/css/styles.css">
</head>
<body>
    <h1>Hola, mundo!</h1>

    <!-- Contenido adicional utilizando Bootstrap 5 -->

    <script src="/static/bootstrap.bundle.min.js"></script>
</body>
</html>
~~~

En la etiqueta <link rel="stylesheet" href="/static/bootstrap.min.css">, asegúrate de que la ruta coincida con la ruta relativa correcta para acceder a tu archivo CSS de Bootstrap.

Asegúrate de que la carpeta "static" y el archivo HTML estén en el mismo nivel que tu archivo Go principal.

En el código que maneja las rutas en tu servidor Go, agrega el enrutamiento para servir el archivo HTML estático:

~~~go
// Manejador para servir el archivo HTML
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
})
~~~

Ejecuta tu servidor Go y accede al archivo HTML a través de tu navegador web. Deberías poder ver la página con los estilos de Bootstrap aplicados.

---
## Agregando ímagenes 
Para agregar imágenes a tu proyecto Go y mostrarlas en tus archivos estáticos, sigue estos pasos:

Crea una carpeta llamada "images" en el directorio "static" de tu proyecto Go. Puedes utilizar la siguiente estructura de carpetas:

~~~bash
- static
  - images
~~~

Coloca tus imágenes en la carpeta "images". Por ejemplo, si tienes una imagen llamada "logo.png", colócala en "static/images/logo.png".

En tu archivo HTML, puedes utilizar la etiqueta `<img>` para mostrar las imágenes en tu página. Asegúrate de que la ruta de la imagen sea correcta.

~~~html
<!DOCTYPE html>
<html>
<head>
    <title>Proyecto con imágenes</title>
</head>
<body>
    <h1>Hola, mundo!</h1>

    <img src="/static/images/logo.png" alt="Logo">

    <!-- Contenido adicional -->
</body>
</html>
~~~

En la etiqueta `<img src="/static/images/logo.png" alt="Logo">`, asegúrate de que la ruta coincida con la ruta relativa correcta para acceder a tus imágenes.

Asegúrate de que la carpeta "static" (que contiene la carpeta "images") y el archivo HTML estén en el mismo nivel que tu archivo Go principal.

Si ya tienes el código que maneja las rutas en tu servidor Go (como el que mencionaste anteriormente), no es necesario realizar cambios adicionales. El enrutamiento existente ya debe estar sirviendo todos los archivos estáticos, incluidas las imágenes.

Ejecuta tu servidor Go y accede al archivo HTML a través de tu navegador web. Deberías poder ver las imágenes correctamente en tu página.

Recuerda ajustar las rutas y los nombres de archivos según tu estructura de directorios y nombres de archivos específicos.

---
## Agregando  JavaScript 

Para agregar JavaScript a tu proyecto Go y utilizarlo en tus archivos estáticos, sigue estos pasos:

Crea una carpeta llamada "js" en el directorio "static" de tu proyecto Go. Puedes utilizar la siguiente estructura de carpetas:

~~~bash
- static
  - js
~~~

Coloca tus archivos JavaScript en la carpeta "js". Por ejemplo, si tienes un archivo JavaScript llamado "script.js", colócalo en "static/js/script.js".

En tu archivo HTML, puedes utilizar la etiqueta `<script>` para incluir tus archivos JavaScript. Asegúrate de que la ruta del archivo JavaScript sea correcta.

~~~html
<!DOCTYPE html>
<html>
<head>
    <title>Proyecto con JavaScript</title>
</head>
<body>
    <h1>Hola, mundo!</h1>

    <!-- Contenido HTML -->

    <script src="/static/js/script.js"></script>
    <script>
        // Código JavaScript adicional
    </script>

</body>
</html>
~~~

En la etiqueta `<script src="/static/js/script.js"></script>`, asegúrate de que la ruta coincida con la ruta relativa correcta para acceder a tus archivos JavaScript.

Asegúrate de que la carpeta "static" (que contiene la carpeta "js") y el archivo HTML estén en el mismo nivel que tu archivo Go principal.

Si ya tienes el código que maneja las rutas en tu servidor Go (como el que mencionaste anteriormente), no es necesario realizar cambios adicionales. El enrutamiento existente ya debe estar sirviendo todos los archivos estáticos, incluidos los archivos JavaScript.

Ejecuta tu servidor Go y accede al archivo HTML a través de tu navegador web. Deberías poder ejecutar el código JavaScript y ver sus efectos en tu página.

Recuerda ajustar las rutas y los nombres de archivos según tu estructura de directorios y nombres de archivos específicos.

--- 
## Resumen 
En esta sección del curso de Go web, aprendimos cómo trabajar con archivos estáticos en una aplicación web desarrollada con el lenguaje de programación Go. Los archivos estáticos, como hojas de estilo CSS, imágenes y scripts de JavaScript, son elementos fundamentales para el diseño y la interactividad de un sitio web. Aprendimos cómo configurar y servir estos archivos estáticos correctamente en nuestro servidor Go.

En primer lugar, configuramos la estructura de directorios de nuestro proyecto Go para manejar los archivos estáticos de manera eficiente. Creamos una carpeta "static" en la raíz del proyecto y configuramos un enrutador para servir estos archivos adecuadamente.

Luego, agregamos estilos a nuestras páginas web utilizando hojas de estilo CSS. Además, exploramos cómo agregar Bootstrap 5, uno de los frameworks de CSS más populares, a nuestro proyecto. Aprendimos a enlazar los archivos CSS de Bootstrap correctamente y aplicar estilos predefinidos a nuestros elementos HTML.

También aprendimos cómo agregar imágenes estáticas a nuestro proyecto Go y mostrarlas correctamente en nuestras páginas HTML. Organizamos las imágenes en una carpeta "images" dentro de la carpeta "static" y enlazamos las rutas de las imágenes en nuestro código HTML.

Finalmente, agregamos archivos de JavaScript estáticos a nuestro proyecto Go y ejecutamos código JavaScript en nuestras páginas HTML. Aprendimos a enlazar los archivos JavaScript y utilizamos bibliotecas y scripts personalizados para mejorar la funcionalidad de nuestro sitio web.

Al final de esta sección, nos familiarizamos con el manejo de archivos estáticos en Go y pudimos crear aplicaciones web completas con estilos atractivos, imágenes llamativas y funciones interactivas utilizando CSS, Bootstrap, imágenes y JavaScript.

# Base de datos con GORM

1. [Introducción](#introduccións)
2. [Crear la base de datos](#crear-la-base-de-datos)
3. [Conexion a MySQL](#conexion-a-mysql)
4. [Crear modelo](#crear-modelo)
5. [Crear un usuario](#crear-un-usuario)
6. [Obtener un usuario, editar y eliminar](#obtener-un-usuario-editar-y-eliminar)
7. [Resumen](#resumen)


---
## Introducción 
¡Bienvenidos a la sección de Base de Datos con GORM de nuestro curso de Go Web! En esta parte del curso, vamos a sumergirnos en el fascinante mundo de la persistencia de datos y cómo interactuar con una base de datos MySQL utilizando GORM, un poderoso ORM (Object-Relational Mapping) para Go.

A lo largo de esta sección, aprenderemos los fundamentos esenciales para trabajar con bases de datos en Go, desde la creación de la base de datos hasta la manipulación de registros. Aquí están los temas principales que abordaremos:

1. Creación de la base de datos: Comenzaremos aprendiendo cómo crear una base de datos MySQL y diseñar la estructura de nuestras tablas para almacenar la información de nuestros usuarios.

2. Conexión a MySQL: Aprenderemos cómo establecer una conexión segura entre nuestra aplicación Go y la base de datos MySQL, lo que nos permitirá realizar operaciones de lectura y escritura.

3. Creación de modelo: Descubriremos cómo definir y mapear nuestras estructuras Go en tablas de la base de datos utilizando GORM, lo que facilita la sincronización de nuestros objetos con la base de datos.

4. Crear un usuario: Veremos cómo crear nuevos registros en la tabla de usuarios, permitiendo que nuestros usuarios se registren y sean parte de nuestra aplicación web.

5. Obtener un usuario, editar y eliminar: Aprenderemos cómo realizar consultas para obtener información de usuarios específicos, cómo editar sus detalles y cómo eliminar usuarios cuando ya no sean necesarios.

A medida que avancemos, profundizaremos en cada uno de estos temas con ejemplos prácticos y situaciones reales para que puedas entender mejor cómo aplicar estos conceptos en tus propios proyectos web con Go.

¡Prepárate para explorar el poder de GORM y cómo simplifica la interacción con bases de datos MySQL en aplicaciones Go! Estoy emocionado de acompañarte en este viaje de aprendizaje. ¡Vamos a empezar!
¡Bienvenidos a la sección de Base de Datos con GORM de nuestro curso de Go Web! En esta parte del curso, vamos a sumergirnos en el fascinante mundo de la persistencia de datos y cómo interactuar con una base de datos MySQL utilizando GORM, un poderoso ORM (Object-Relational Mapping) para Go.

A lo largo de esta sección, aprenderemos los fundamentos esenciales para trabajar con bases de datos en Go, desde la creación de la base de datos hasta la manipulación de registros. Aquí están los temas principales que abordaremos:

1. Creación de la base de datos: Comenzaremos aprendiendo cómo crear una base de datos MySQL y diseñar la estructura de nuestras tablas para almacenar la información de nuestros usuarios.

2. Conexión a MySQL: Aprenderemos cómo establecer una conexión segura entre nuestra aplicación Go y la base de datos MySQL, lo que nos permitirá realizar operaciones de lectura y escritura.

3. Creación de modelo: Descubriremos cómo definir y mapear nuestras estructuras Go en tablas de la base de datos utilizando GORM, lo que facilita la sincronización de nuestros objetos con la base de datos.

4. Crear un usuario: Veremos cómo crear nuevos registros en la tabla de usuarios, permitiendo que nuestros usuarios se registren y sean parte de nuestra aplicación web.

5. Obtener un usuario, editar y eliminar: Aprenderemos cómo realizar consultas para obtener información de usuarios específicos, cómo editar sus detalles y cómo eliminar usuarios cuando ya no sean necesarios.

A medida que avancemos, profundizaremos en cada uno de estos temas con ejemplos prácticos y situaciones reales para que puedas entender mejor cómo aplicar estos conceptos en tus propios proyectos web con Go.

¡Prepárate para explorar el poder de GORM y cómo simplifica la interacción con bases de datos MySQL en aplicaciones Go! Estoy emocionado de acompañarte en este viaje de aprendizaje. ¡Vamos a empezar!

---
## Crear la base de datos
Para crear una base de datos MySQL desde la terminal en Linux, Asegúrate de tener MySQL instalado en tu sistema. Si no lo tienes, puedes instalarlo ejecutando el siguiente comando: https://github.com/alexroel/tutoriales/blob/master/instalar-mysql.md


Inicia sesión en MySQL como usuario root u otro. Por defecto, el usuario root no tiene contraseña en una instalación nueva. Con  otro usuario sea de esta forma: 

~~~bash
mysql -u alexroel -p
Enter password: ******
~~~


Una vez dentro de MySQL, puedes crear la base de datos con el siguiente comando:

~~~sql
CREATE DATABASE todolist_db;
~~~

Puedes verificar que la base de datos se haya creado correctamente usando el comando:

~~~sql
SHOW DATABASES;
~~~

Para salir del cliente MySQL, simplemente ejecuta el siguiente comando:

~~~sql
exit;
~~~

Con esto, has creado exitosamente una base de datos en MySQL desde la terminal en Linux. Puedes usar esta base de datos para conectarte desde tu aplicación Go usando GORM, como se mostró en la respuesta anterior. Recuerda configurar las credenciales de conexión en tu aplicación para que coincidan con las utilizadas durante la instalación de MySQL y la creación de la base de datos.

---
## Conexion a MySQL

Para conectarte a MySQL con GORM, asegúrate de tener GORM instalado en tu entorno de desarrollo. Puedes hacerlo usando el siguiente comando:

~~~bash
go get -u gorm.io/gorm
~~~
Una vez que hayas instalado GORM, sigue estos pasos para establecer una conexión a tu base de datos MySQL:


Importa las bibliotecas necesarias y define la configuración de conexión a la base de datos:
~~~go
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func main() {

	dsn := "alexroel:123456@tcp(localhost:3306)/todolist_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Reemplaza 'user', 'password', 'host', 'port' y 'database' con los detalles de tu base de datos MySQL
    // charset=utf8mb4 y parseTime=True son opciones recomendadas para asegurarse de que GORM interprete correctamente las fechas.
    // loc=Local establece la ubicación local para las fechas y horas.
  
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
  		panic("Error al conectar a la base de datos: " + err.Error())
	}
}
~~~

Ahora estás conectado a tu base de datos MySQL usando GORM. Puedes realizar consultas, insertar, actualizar y eliminar registros utilizando los métodos proporcionados por GORM para interactuar con la base de datos de manera más fácil y conveniente.

Recuerda que debes reemplazar los valores de dsn con las credenciales y la información adecuada para tu base de datos MySQL.

---
## Crear modelo

Primero configuraremos la conexión dentro de paquete `config/config.go`: 

~~~go
package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	// Reemplaza las credenciales y detalles de la conexión con los tuyos
	dsn := "alexroel:123456@tcp(localhost:3306)/todolist_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Conectarse a la base de datos MySQL utilizando GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
~~~

Para crear el modelo `User` dentro del directorio models utilizando GORM con los atributos id, username y password, sigue estos pasos:

1. Crea el archivo user.go dentro del directorio models con el siguiente contenido:

~~~go
package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
}
~~~

En este modelo, estamos utilizando el tipo `gorm.Model`, que proporciona automáticamente los campos ID, CreatedAt, UpdatedAt y DeletedAt. Luego, definimos los campos adicionales Username y Password.

Si deseas migrar el modelo User sin crear un usuario específico en la base de datos, puedes realizar la migración utilizando GORM sin la necesidad de agregar registros en la tabla. La migración solo creará la tabla correspondiente para el modelo User con los campos definidos en el modelo.

~~~go
package main

import (
	"log"
	"taskmanager/config"
	"taskmanager/models"
)

func main() {
	// Obtener la instancia de la conexión a la base de datos
	db := config.GetDB()

	// Realizar la migración del modelo User
    err = db.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Error al migrar el modelo User:", err)
    }

    log.Println("Migración completada exitosamente.")

}
~~~

---
## Crear un usuario 
Para crear un usuario y guardar sus datos en la base de datos usando GORM en Go, sigue estos pasos:

En el archivo main.go (o cualquier otro archivo donde quieras crear el usuario), importa los paquetes necesarios y crea la lógica para crear y guardar el usuario en la base de datos:

~~~go
package main

import (
	"fmt"
	"log"
	"taskmanager/config"
	"taskmanager/models"
)

func main() {
	// Obtener la instancia de la conexión a la base de datos
	db := config.GetDB()

	// Realizar la migración del modelo User
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error al migrar el modelo User:", err)
	}

	// Crear una nueva instancia del modelo User con los datos del usuario
	newUser := models.User{
		Username: "alexroel",
		Password: "alex123",
	}

	// Crear el registro en la base de datos
	result := db.Create(&newUser)
	if result.Error != nil {
		log.Fatal("Error al crear el usuario:", result.Error)
	}

	fmt.Println("Usuario creado exitosamente con ID:", newUser.ID)

}
~~~

---
## Obtener un usuario, editar y eliminar
Encontrar un usuario existente en la base de datos por su ID

~~~go
    // Encontrar un usuario existente en la base de datos por su ID
    var user models.User
    if err := db.First(&user, 1).Error; err != nil {
        log.Fatal("Error al buscar el usuario:", err)
    }
~~~

Actualizar el usuario: 

~~~go
    // Actualizar los campos del usuario
    user.Username = "alex"
    user.Password = "alex2023"

    // Guardar los cambios en la base de datos
    if err := db.Save(&user).Error; err != nil {
        log.Fatal("Error al actualizar el usuario:", err)
    }

    fmt.Println("Usuario actualizado exitosamente:", user)
~~~

Para eliminar:
~~~go
    // Eliminar el usuario de la base de datos
    if err := db.Delete(&user).Error; err != nil {
        log.Fatal("Error al eliminar el usuario:", err)
    }

    fmt.Println("Usuario eliminado exitosamente:", user)
~~~


---
## Resumen
En la sección de Base de Datos con GORM de nuestro curso de Go Web, exploramos el mundo de la persistencia de datos y la interacción con una base de datos MySQL mediante el uso de GORM, un poderoso ORM para Go.

Comenzamos aprendiendo cómo crear una base de datos MySQL y diseñamos la estructura de nuestras tablas para almacenar la información de los usuarios. Luego, establecimos una conexión segura entre nuestra aplicación Go y la base de datos MySQL para realizar operaciones de lectura y escritura.

Enseguida, descubrimos cómo definir y mapear nuestras estructuras Go en tablas de la base de datos utilizando GORM, lo que facilitó la sincronización de nuestros objetos con la base de datos.

Posteriormente, aprendimos a crear nuevos registros en la tabla de usuarios, permitiendo que nuestros usuarios se registren y sean parte de nuestra aplicación web.

Además, exploramos cómo realizar consultas para obtener información de usuarios específicos, cómo editar sus detalles y cómo eliminar usuarios cuando ya no eran necesarios.

A lo largo del curso, utilizamos ejemplos prácticos y situaciones reales para mejorar la comprensión de los conceptos y aplicarlos de manera efectiva en nuestros proyectos web con Go.

Fue un viaje emocionante para aprender sobre el poder de GORM y cómo simplifica la interacción con bases de datos MySQL en aplicaciones Go. ¡Esperamos que esta sección haya sido útil y te sientas más seguro al trabajar con bases de datos en tus proyectos futuros!





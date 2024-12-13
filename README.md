# AskETSIIT

## Descripción del problema

Como estudiante experimentado de la ETSIIT, sé que es complicado el inicio de un curso nuevo debido a todo el cambio de asignaturas y horarios entre otros. Ésto conlleva cierta desinformación por estár sumidos en las novedades y en la desorganización, sobre todo, si es el primer año. 

Es muy común ver a gente, en cualquier altura del curso, preguntar acerca de:
  - Correos de profesores
  - Horarios de asignaturas
  - Lugar de impartición de una clase
  - Profesor de grupos y subgrupos

Información que se encuentra escondida bajo muchos enlaces o documentos en páginas confusas y que hacen dificultosa la aclaración de las dudas y la consulta constante de las mismas.

## Documentos

[Historias de usuario](./docs/HUs.md)
[Milestones](./docs/MSs.md)

## Gestor de dependencias

Se va a utilizar [Go Modules](https://go.dev/ref/mod), el gestor oficial de los desarrolladores de Golang.

## Gestor de tareas

Se va a utilizar [Just](https://just.systems/) como gestor por su simplicidad.
La primera vez que se use se deberá crear el archivo *"justfile"* el cual editaremos con las configuraciones que necesarias.

## Comprobaciones

El fichero que contiene la declaración de "Clase" será el que implemente la lógica de negocio ya que sobre él recae el procesamiento principal del proyecto.

Para instalar las dependencias usaremos:

```
just install-deps
```

Para comprobar la sintaxis debemos usar:

```
just check
```

## Extra

[Licencia](./LICENSE)
[Configuración](./conf/pasos.txt)
[Checkeo de claves](./conf/claves.png)
[Fuentes](./docs/fuentes.md)

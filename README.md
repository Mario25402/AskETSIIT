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
[Comparativa](./docs/gestor_dependencias.md)

Se va a utilizar [Go Modules](https://go.dev/ref/mod), el gestor oficial de los desarrolladores de Golang.

## Gestor de tareas
[Comparativa](./docs/gestor_tareas.md)

Se va a utilizar [Make](https://www.gnu.org/software/make/manual/make.html) como gestor por su gran capacidad y resultados.
Se deberá crear el archivo *"makefile"* el cual editaremos con las configuraciones que necesarias.

## Ordenes
El fichero que contiene la declaración de "Horario" será el que implemente la lógica de negocio ya que sobre él recae el procesamiento principal del proyecto.

Para instalar las dependencias usaremos:

```
make install
```

Para comprobar la sintaxis debemos usar:

```
make check
```

Respecto a la parte de testing, se ha optado por seguir los estándares del lenguaje utilizando [*Testing*](/docs/test_frameworks.md) en su combinación predeterminada con [*go test*](/docs/test_runner.md).

Para ejecutar el testeo del código se usará:
```
make test
```

## Docker
La aplicación ha sido *"dockerizada"* siguiendo las siguientes [herramientas](/docs/imagenes_docker.md).

Para ejecutar el contenedor usaremos:
```
docker run mario24502/asketsiit:latest
```

## Extra
[Licencia](./LICENSE)
[Configuración](./conf/pasos.txt)
[Checkeo de claves](./conf/claves.png)
[Fuentes](./docs/fuentes.md)

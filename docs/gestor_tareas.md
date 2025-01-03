# Gestores de Tareas
## Requisitos de aceptación:

- **Mantenimiento:** Si un gestor no es mantenido, puede quedar obsoleto y provocar un trabajo adicional motivado por el cambio de una mala elección, debido a ésto, lo recomendable es elegir herramientas que disminuyan la deuda técnica.
  
- **Prestaciones:** Las herramientas deben ser eficientes por lo que será positivo que tengan una buena velocidad de ejecución de las tareas.

## Opciones
### Mage
[Mage](https://magefile.org/) es una herramienta escrita en Go que permite definir tareas utilizando el propio lenguaje.

- Activa pero no tan mantenida últimamene.
- Último *release* hace más de un año, [*commits*](https://github.com/magefile/mage/commits/master/) con frecuencia media.
- Medición de la tarea `check`: 0.288s

### Task
[Task](https://taskfile.dev/) es un gestor que apuesta por la simplicidad y la versatilidad.

- Actualizaciones frecuentes, compatible con proyectos multilenguaje.
- Último *release* hace un mes, [*commits*](https://github.com/go-task/task/commits/main/) frecuentes.
- Medición de la tarea `check`: 0.22s

### Make

[Make](https://www.gnu.org/software/make/) es una herramienta clásica de automatización integrable con todo tipo de lenguajes.

- Estable y fuertemente asentado. Mantenimiento escaso debido a su gran robustez.
- Último *release* hace más de un año.
- Medición de la tarea `check`: 0.16s

### Just
[Just](https://just.systems/) es un gestor ligero similar a ´Make´ pero más simplista.

- Herramienta menos extendida entre la comunidad.
- Último *release* hace unas semanas, [*commits*](https://github.com/casey/just/commits/master/) muy frecuentes.
- Medición de la tarea `check`: 0.19s

### XC

[XC](https://xcfile.dev/) es un gestor similar a `Make` destinado a minimizar la complejidad, se escribe en `Markdown`.

- Proyecto de 3 años con poca comunidad.
- Último *release* hace unos días, sin [*commits*](https://github.com/joerdav/xc/commits/main/) hace pocos días.
- Medición de la tarea `check`: 0.21s

### Taskrunner

[Taskrunner](https://github.com/samsarahq/taskrunner), herramienta escrita en Go que permite crear tareas reutilizables con nombres y dependecias.

- Sin *releases* oficiales, sin [*commits*](https://github.com/samsarahq/taskrunner/commits/master/) desde hace un año.

## Conclusión
La opción de usar `Taskrunner` quedará descartada porque lleva tiempo sin ser actualizada con lo que es posible que se haya abandonado el proyecto y la herramienta quede obsoleta (ya se pueden ir viendo estos detalles por los problemas durante la instalación).

`Mage` es el siguiente descarte debido a que es el que mayor tiempo tarda en realizar la tarea y su ritmo de desarrollo no es tan activo como otras alternativas.

`XC` y `Task` están en una situación similar, proyectos mantenidos, con mejor tiempo que `Mage` pero siguen sin ser los mejores.

Por último entre `Make` y `Just`, que son los dos que mejores tiempos consiguen, elegiremos **Make** porque consigue unas mejores prestaciones.

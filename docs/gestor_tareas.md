# Gestores de Tareas
## Requisitos de aceptación:

- **Mantenimiento:** Un mantenimiento adecuado es importarte a nivel de seguridad para proteger el entorno de desarrollo ya que los gestores de tareas ejecutan comandos y pueden ser objetivos para búsquedas de vulnerabilidades. Por otra parte, si un gestor no es mantenido, puede quedar obsoleto y provocar un trabajo adicional motivado por el cambio de una mala elección, debido a ésto, lo recomendable es elegir herramientas que disminuyan la deuda técnica.
  - La herramienta debe seguir teniendo desarrollo activo y frecuente para maximizar la compatibilidad con el lenguaje, minimizar los errores y las amenazas. 
  - El repositorio de GitHub de cada una será visitado individualmente, revisando la frecuencia y fechas de los últimos commits y versiones.

## Opciones
### Mage
[Mage](https://magefile.org/) es una herramienta escrita en Go que permite definir tareas utilizando el propio lenguaje.

- Activa pero no tan mantenida últimamene.
- Último *release* hace más de un año, [*commits*](https://github.com/magefile/mage) con frecuencia media.

### Task
[Task](https://taskfile.dev/) es un gestor que apuesta por la simplicidad y la versatilidad.

- Actualizaciones frecuentes, compatible con proyectos multilenguaje.
- Último *release* hace unas semana, [*commits*](https://github.com/go-task/task) frecuentes.

### Make

[Make](https://www.gnu.org/software/make/) es una herramienta clásica de automatización integrable con todo tipo de lenguajes.

- Estable y fuertemente asentado. Mantenimiento escaso debido a su gran robustez.
- Último *release* hace más de un año.

### Just
[Just](https://just.systems/) es un gestor ligero similar a ´Make´ pero más simplista.

- Herramienta menos extendida entre la comunidad.
- Último *release* hace unos días, [*commits*](https://github.com/casey/just) muy frecuentes.

### XC

[XC](https://xcfile.dev/) es un gestor similar a `Make` destinado a minimizar la complejidad, se escribe en `Markdown`.

- Proyecto de 3 años con poca comunidad.
- Último *release* hace casi un año, sin [*commits*](https://github.com/joerdav/xc) hace meses.

### Taskrunner

[Taskrunner](https://github.com/samsarahq/taskrunner), herramienta escrita en Go que permite crear tareas reutilizables con nombres y dependecias.

- Sin *releases* oficiales, sin [*commits*](https://github.com/joerdav/xc) desde hace un año.

## Conclusión
Vistas diferentes alternativas para nuestra elección como gestor de tareas, podemos concluir diciendo que no escogeremos `XC` porque su mantenimiento es bajo, al igual que el de `Taskrunner`.
Tampoco elegiremos `Mage` ya que no recibe tanto desarrollo como otras alternativas.

De las herramientas restantes, vamos a seleccionar **Just** por su desarrollo activo frente a *Make* y *Task*.

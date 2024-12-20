# Gestores de Tareas
## Requisitos de aceptación:

- **Madurez:** El nivel de desarrollo de la herramienta debe ser suficientemente completo.
- **Progreso:** La herramienta debe seguir teniendo mantenimiento frecuente para maximizar la compatibilidad con el lenguaje y minimizar los errores.
- El repositorio de cada herramienta será revisado individualmente con el fin de ver el nivel/ritmo de desarollo.

## Opciones
### Mage
[Mage](https://magefile.org/) es una herramienta escrita en Go que permite definir tareas utilizando el propio lenguaje.

- Activa pero no tan mantenida últimamene.
- Último *release* hace más de un año, [*commits*](https://github.com/magefile/mage) con frecuencia media.

### Task
[Task](https://taskfile.dev/) es un gestor que apuesta por la simplicidad y la versatilidad.

- Actualizaciones frecuentes, compatible con proyectos multilenguaje.
- Último *release* hace 1 semana, *commits* frecuentes.

### Make

[Make](https://www.gnu.org/software/make/) es una herramienta clásica de automatización integrable con todo tipo de lenguajes.

- Estable y fuertemente asentado. Mantenimiento escaso debido a su gran robustez.
- Último *release* hace más de un año.

### Just
[Just](https://just.systems/) es un gestor ligero similar a ´Make´ pero más simplista.

- Herramienta menos extendida entre la comunidad.
- Último *release* hace 3 días, [*commits*](https://github.com/casey/just) muy frecuentes.

### Dagger
[Dagger](https://dagger.io/) es una herramienta diseñada para simplificar y unificar flujos de trabajo CI/CD y tareas automatizadas en proyectos de software.

- Es una herramienta relativamente joven pero con mucha interacción de la comunidad.
- Último *release* hace unos días, [*commits*](https://github.com/dagger/dagger) muy frecuentes.

### Goyek
[Goyek](https://github.com/goyek/goyek) es una biblioteca ligera para definir y ejecutar pipelines de tareas directamente en Go.

- Proyecto no tan completo como otras opciones.
- Último *release* hace unos meses, [*commits*](https://github.com/goyek/goyek) frecuentes.

## Conclusión
Vistas diferentes alternativas para nuestra elección como gestor de tareas, podemos concluir diciendo que no escogeremos `Goyek` porque no es una herramienta tan avanzada como las demás, `Dager` tiene dependecias con *docker* lo que añadiría peso innecesario a nuestro proyecto. Tampoco elegiremos `Mage` ya que no recibe tanto desarrollo como otras alternativas.

De las herramientas restantes, vamos a seleccionar **Just** por su desarrollo activo frente a *Make* y *Task*.

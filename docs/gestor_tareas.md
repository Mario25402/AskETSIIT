# Gestores de Tareas
## Requisitos de aceptación:

- **Madurez:** El nivel de desarrollo de la herramienta debe ser suficientemente completo.
- **Progreso:** La herramienta debe seguir teniendo mantenimiento frecuente para maximizar la compatibilidad con el lenguaje y minimizar los errores.

## Opciones
### Mage
[Mage](https://magefile.org/) es una herramienta escrita en Go que permite definir tareas utilizando el propio lenguaje.

- Activa pero no tan mantenida últimamene.
- Último *release* hace más de un año, *commits* con frecuencia media.

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

- Desarrollo menos mantenido y menos extendido entre la comunidad.
- Último *release* hace 3 días, *commits* muy frecuentes.

### Dagger
[Dagger](https://dagger.io/) es una herramienta diseñada para simplificar y unificar flujos de trabajo CI/CD y tareas automatizadas en proyectos de software

- Es una herramienta relativamente joven pero con mucha interacción de la comunidad.
- Último *release* hace unos días, *commits* muy frecuentes.

### Goyek
[Goyek](https://github.com/goyek/goyek) es una biblioteca ligera para definir y ejecutar pipelines de tareas directamente en Go

- Proyecto no tan completo como otras opciones.
- Último *release* hace unos meses, *commits* frecuentes.

## Conclusión
Aunque todas las opciones podrían ser consideradas en diferentes ámbitos debido a sus similitudes, vamos a ir nombrando a la que va a ser nuestra elección:

- **Just**: Su sintaxis única hace que se cuestione su uso además, durante su instalación nos encontramos problemas con los repositorios y de dependecias externas.
- **Make**: Aunque muy buena opción, es preferible usar herramientas más intuitivas y mejor adaptadas a lenguajes actuales.
- **Mage**: Sintaxis un poco más compleja que otras opciones que puede que aumente la curva de aprendizaje al ser escrita en Go.
- **Task**: Configuración con *YAML*, consideramos ésta como nuestra elección.

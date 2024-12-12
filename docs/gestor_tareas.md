# Gestores de Tareas
## Requisitos de aceptación:

- **Madurez y progreso:** El nivel de desarrollo de la herramienta debe ser suficiente y debe seguir teniendo mantenimiento frecuente para maximizar la compatibilidad con el lenguaje y minimizar los errores.
- **Adición de complejidad:** Se valorará que la curva de aprendizaje de la herramienta sea simple o menor que las demás.

## Opciones
### Mage
[Mage](https://magefile.org/) es una herramienta escrita en Go que permite definir tareas utilizando el propio lenguaje.

1. Activa y mantenida, gracias a su integración directa en Go.
2. La complejidad reside en la propia dificultad personal para el desarrollo de estos archivos en Go.

### Task
[Task](https://taskfile.dev/) es un gestor que apuesta por la simplicidad y la versatilidad.

1. Actualizaciones frecuentes, compatible con proyectos multilenguaje.
2. Sintaxis sencilla por el uso de *YAML*.

### Make

[Make](https://www.gnu.org/software/make/) es una herramienta clásica de automatización integrable con todo tipo de lenguajes.

1. Estable y fuertemente asentado. Mantenimiento escaso debido a su gran robustez.
2. No requiere de grandes conocimientos aunque utiliza una sintaxis específica no muy compleja.

### Just
[Just](https://just.systems/) es un gestor ligero similar a ´Make´ pero más simplista.

1. Desarrollo menos mantenido y menos extendido.
2. Sintaxis única y limpia, similar a la de `Make`.

## Conclusión
Aunque todas las opciones podrían ser consideradas en diferentes ámbitos debido a sus similitudes, vamos a ir nombrando a la que va a ser nuestra elección:

- **Just**: Su sintaxis única hace que se cuestione su uso además, durante su instalación nos encontramos problemas con los repositorios y de dependecias externas.
- **Make**: Aunque muy buena opción, es preferible usar herramientas más intuitivas y mejor adaptadas a lenguajes actuales.
- **Mage**: Sintaxis un poco más compleja que otras opciones que aumentará la curva de aprendizaje.
- **Task**: Gracias a la simple configuración con *YAML* consideramos ésta como nuestra elección.

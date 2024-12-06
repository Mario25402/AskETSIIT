# Gestores de Tareas
## Requisitos de aceptación:

- **Madurez y robustez:** El nivel de desarrollo de la herramienta debe ser suficiente y debe seguir teniendo mantenimiento frecuente para maximizar la compatibilidad con el lenguaje y minimizar los errores.
- **Adición de complejidad:** Se valorará la simplicidad de inclusión y uso del gestor, de manera que no aumente la dificultad del proyecto.

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
2. Integración secilla pero su sintaxis es menos moderna que las herramientas actuales.

### Just
[Just](https://just.systems/) es un gestor ligero similar a Mage pero más simplista.

1. Desarrollo activo aunque menos utilizado que otras herramientas.
2. Sintaxis limpia y ligera.

## Conclusión
Aunque todas las opciones podrían ser consideradas en diferentes ámbitos debido a sus similitudes, vamos a ir nombrando a la que va a ser nuestra elección:

- **Just**: Durante su instalación nos encontramos problemas con los repositorios y de dependecias externas.
- **Make**: Aunque muy buena opción, es preferible usar herramientas más modernas.
- **Mage**: Sintaxis un poco más compleja que otras opciones que aumentará la curva de aprendizaje.
- **Task**: Gracias a la simple configuración con *YAML* consideramos ésta como nuestra elección.

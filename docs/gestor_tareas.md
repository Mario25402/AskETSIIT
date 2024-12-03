# Gestores de Tareas
## Requisitos de aceptación:

- Deuda técnica.
- Versatilidad.
- Soporte y comunidad activa.
- Características y funcionalidades.

## Comparativa
### Mage
	- Mage es una herramienta escrita y diseñada para Go.
	- Permite la definición de tareas con el propio lenguaje.
	- Es escalable gracias a su integración directa con el lenguaje.
	- Rápida por no tener que interpretar archivos externos.
 	- La última versión se lanzó hace tiempo, por lo que es posible que nos encontremos con problemas.
	- Recomendado para proyectos exclusivos de Go.

### Task
	- Un gestor ligero y moderno que utiliza configuraciones en YAML.
	- Facilita la ejecución paralela y las dependencias entre tareas.
	- Deuda técina baja por su buen soporte y comunidad activa.
	- Buen rendimiento para tareas en paralelo.
	- Versátil e ideal para proyectos multilenguaje.
	- Susceptible a errores en archivos de configuración.
	- Recomendado para proyectos hibridos.

### Make
	- Task runner genérico y ampliamente utilizado, soporta el uso de Go.
	- Permite definir reglas personalizadas.
	- Deuda técnica media debido a su sintaxis y posible complejidad ante grandes proyectos.
	- Versátil y portable.
	- Más "rústico" que las otras opciones.
	- Sin soporte para funciones complejas.
	- Riesgos de seguridad al ejecutar comandos de shell.
	- Recomendado para proyectos pequeños que busquen estabilidad.

### GoReleaser
	- Herramienta especializada para proyectos Go.
	- Útil con tareas de despliegue continuo.
	- Deuda técnica baja por su optimización con Go.
	- Susceptible a errores en archivos de configuración.
	- Menor flexibilidad en tareas genéricas.
	- Recomendado para proyectos que requieren gestión automatizada de lanzamientos.

## Conclusión

El criterio de la deuda técnica es muy importante ya que no queremos que conforme avance el proyecto nos vayamos poniendo más "piedras en el camino" por lo que descartaremos *Make*.

De entre los demás tampoco utilizaremos *Mage* porque aunque está mejor integrado, no es común que los proyectos tenga exclusivamente un solo lenguaje.

Entre los dos restantes, nos quedaremos con *Task* porque aunque ambos utilizan configuraciones en *"yml"* y *"yaml"*, *Task* permite definir tareas personalizadas de manera más simple que *GoReleaser*.

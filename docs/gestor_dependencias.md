
# Gestores de Dependencias

Exceptuando gestores como [**Glide**](https://github.com/Masterminds/glide), en la que los propios autores no recomiendan su uso ya que actualmente la mayoría del proyecto no se encuentra en mantenimiento, la mayoría de gestores de dependencias utilizadas para Go están obsoletas.

La razón principal es la siguiente opción:

## Go Modules

El gestor oficial de los desarrolladores del lenguaje, nacido en la versión 1.11 y obligatorio desde la versión 1.16.

Es el gestor de dependencias por antonomasia de *Golang* debido a que viene incluido de forma nativa en el ecosistema y no necesita de configuración externa, lo que facilita mucho la elección de su uso por los consumidores.

El modo de funcionamiento es simple a partir de la creación de un fichero llamado ***go.mod*** que actúa de director. Éste incluye las versiones de cada paquete necesitado en el proyecto.

Ejecutando una orden, el fichero se actualizará con las dependencias creando además un fichero que indicará las versiones y el *hash* de integridad de cada una.

## Configuración manual

Existe la posibilidad de gestionar manualmente las dependencias, a través de la carpeta *"vendor"*, la cuál no funciona como un gestor sino que la gestión se hace de forma explicita por la/s persona/s encargada/s. Ésta contendrá el código fuente de los paquetes utilizados.

Algunas de las ventajas que ofrece éste método es la posibilidad de programar de manera *offline* (en caso de tener las dependencias descargadas), buscar una versión específica, acelera la compilación sobretodo en entornos de trabajo en equipo.

Aunque es una opción interesante de considerar en ciertas situaciones, nuestra elección será el gestor estándar gracias a su simplicidad, eficacia e integración.

# Conclusión

Aunque el uso de la gestión manual es una opción interesante de considerar en ciertas situaciones, nuestra elección será el gestor estándar gracias a su simplicidad, eficacia e integración.

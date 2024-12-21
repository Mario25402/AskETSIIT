
# Gestores de Dependencias

Exceptuando gestores como [**Glide**](https://github.com/Masterminds/glide), en la que los propios autores no recomiendan su uso ya que actualmente la mayoría del proyecto no se encuentra en mantenimiento, la mayoría de gestores de dependencias utilizadas para Go están obsoletas.

La razón principal es la siguiente opción:

## Go Modules

El [gestor oficial](https://go.dev/blog/using-go-modules) de los desarrolladores del lenguaje, nacido en la versión 1.11 y obligatorio desde la versión 1.16.

Es el gestor de dependencias por antonomasia de *Golang* debido a que viene incluido de forma nativa en el ecosistema y no necesita de configuración externa, lo que facilita mucho la elección de su uso por los consumidores.

El modo de funcionamiento es simple a partir de la creación de un fichero llamado ***go.mod*** que actúa de director. Éste incluye las versiones de cada paquete necesitado en el proyecto.

Ejecutando una orden, el fichero se actualizará con las dependencias creando además un fichero que indicará las versiones y el *hash* de integridad de cada una.

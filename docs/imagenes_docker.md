# Imágenes Docker
## Requisitos de aceptación
- **Mantenimiento**: Que una imagen se mantenga actualizada es importante para estar a la última de las novedades y sin riesgo de producir una deuda técnica futura, además de reducir posibles brechas de seguridad.
    - Se valorará que tenga actualizaciones recientes.

- **Tamaño**: Por norma general una imagen pequeña estará más enfocada y será más eficiente que una de mayor tamaño.
    - Se priorizarán imágenes con menor tamaño.

- **Insignias**: En DockerHub exiten unas insignias llamadas "Trusted Content" que indican si la imagen es *Oficial*, si es de un *Creador Verificado* o si es de *Código Abierto*. La concesión de alguna de estas insignias indican contenido de calidad por lo que:
    - Se seleccionarán imágenes con insignias.

## Opciones
### **[Golang](https://hub.docker.com/_/golang)**
Su versión BookWorm basada en Debian 11, cuenta con 58 vulnerabilidades registradas. La versión BullsEye no se escogió debido que esta basada en Debian 10 por lo que es preferible usar una versión más reciente.
- Última actualización hace horas.
- Tamaño base de 288.23MB.
- Make instalado: *Y*.
- Versión oficial.

### **[Bitnami](https://hub.docker.com/r/bitnami/golang)**
- Última actualización hace días.
- Tamaño base de 230.74MB.
- Make instalado: *Y*.
- Desarrollador verificado.

### **[Debian](https://hub.docker.com/_/debian)**
- Última actualización hace días.
- Tamaño base de 47.17MB.
- Make instalado: *N*.
- Versión oficial.

### **[Alpine](https://hub.docker.com/_/alpine)**
- Última actualización hace días.
- Tamaño base de 3.3MB.
- Make instalado: *N*.
- Versión oficial.

## Conclusión

Las dos imágenes de *Go* van a quedar descartadas ya que aunque ambas traen `Make` instalado por defecto, ámbas tienen insignias de DockerHub, ámbas tienen actualizaciones recientes, el peso en comparación con otras opciones es muy elevado.

Entre selecciónar Debian y Linux Alpine, vamos a elegir usar **Alpine**, porque aunque en ambas deberemos instalar `Go` y `Make`, la versión reducidad de Linux cuenta con un menor peso frente a `Debian`, y muy pocas funcionalidades instaladas por lo que mejorará la seguridad y la eficiencia de la imagen.
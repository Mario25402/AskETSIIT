# Frameworks
## Criterios de aceptación
 - Mantenimiento y actualización del paquete, que cumpla este requisito es importante debido a que si una herramienta no se corrige o no se actualiza conforme se renuevan versiones del lenguaje, puede quedar obsoleta o que parte de código antiguo deje de funcionar.
    - Se valorará que haya commits recientes y que solucionen issues o pull requests abiertos por la comunidad.

## Opciones
### 1. [Testing](https://pkg.go.dev/testing)
- Inlcuido en Go de forma nativa.
- Alineado con las mejores prácticas ya que es desarollado por los mismos creadores.
- Mantenido y actualizado en cada versión lanzada si fuera necesario.

### 2. [Ginkgo](https://github.com/onsi/ginkgo)
- Framework que sigue un estilo BDD con un DSL no convencional.
- Repositorio con 8.5k estrellas, con [commits](https://github.com/onsi/ginkgo/commits/master/) recientes.

### 3. [GoConvey](https://github.com/smartystreets/goconvey)
- Herramienta conocida por su interfaz web en tiempo real con un DSL sencillo.
- Repositorio con 8.3k estrellas, sin [commits](https://github.com/smartystreets/goconvey/commits/master/) desde hace casi un año.

## Conclusión
`GoConvey` es una opción interesante y llamativa debido a su característica gráfica pero su casi nulo mantenimiento puede resultar en problemas con las últimas novedades.

Aunque el uso de `Ginkgo` (en combinación con *Gomega*, generalmente) es notable, nos adheriremos a las mejores prácticas establecidas para el lenguaje en cuyo caso la mejor opción es elegir `Testing`, el estándar desarollado por Google en Go.
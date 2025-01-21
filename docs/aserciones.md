# Biblioteca de aserciones
## Criterios de aceptación
 - Mantenimiento y actualización del paquete, que cumpla este requisito es importante debido a que si una herramienta no se corrige o no se actualiza conforme se renuevan versiones del lenguaje, puede quedar obsoleta o que parte de código antiguo deje de funcionar.
    - Se valorará que haya commits recientes y que solucionen issues o pull requests abiertos por la comunidad.
    - Será preferible una herramienta que cuente con mayor apoyo por parte de la comunidad, de manera que obtenga más estrellas en su repositorio de GitHub.
- Sería muy adecuado que la elección no añada dependecias externas. El no utilizar paquetes no incluidos "de fábrica" es aconsejable porque es posible que aumenten la deuda técnica en caso de quedar deprecados.

## Opciones
### 1. [Testify](https://github.com/stretchr/testify)
- Repositorio con casi 24k estrellas, y con [commits](https://github.com/stretchr/testify/commits/master/) recientes de hace unas semanas.

### 2. [Gomega](https://github.com/onsi/gomega)
- Repositorio con 2.2k estrellas, [commits](https://github.com/onsi/gomega/commits/master/) muy recientes.

### 3. [Check](https://github.com/go-check/check)
- Repositorio con alrededor de 60 estrellas, sin [commits](https://github.com/go-check/check/commits/master/) desde hace varios años.

### 4. [Testing](https://pkg.go.dev/testing)
- Testing es un paquete nativo de *Go* que engloba funcionalidades de testing, incluidas las de aserción y matching, útilies para este apartado.


## Conclusión
`Check` es una opción completa pero sin mantenimiento desde hace 5 años por lo que la catalogaremos como obsoleta y quedará descartada.

Entre `Testify` y `Gomega` hay una clara diferencia de apoyo por parte de la comunidad, sin embargo, ambas necesitan de descargar paquetes externos para su uso, por lo que `Testing` será nuestra opción.
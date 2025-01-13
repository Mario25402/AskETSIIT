# Biblioteca de aserciones
## Criterios de aceptación
 - Compatibilidad con el estándar *go test*.
 - Mantenimiento y actualización del paquete, que cumpla este requisito es importante debido a que si una herramienta no se corrige o no se actualiza conforme se renuevan versiones del lenguaje, puede quedar obsoleta o que parte de código antiguo deje de funcionar.
    - Se valorará que haya commits recientes y que solucionen issues o pull requests abiertos por la comunidad.

## Opciones
### 1. [Testify](https://github.com/stretchr/testify)
- Aserciones avanzadas e inclusión de mocks (componentes que imitan el comportamiento real en un sistema).
- Compatible con *go test*.
- Repositorio con casi 24k estrellas, y con [commits](https://github.com/stretchr/testify/commits/master/) recientes de hace unas semanas.

### 2. [Gomega](https://github.com/onsi/gomega)
- Enfocada en BDD y con posibilidad de crear aserciones personalizadas.
- Compatible en cierta medida pero mayormente pensada para utilizar con el framework **Ginkgo**.
- Repositorio con 2.2k estrellas, [commits](https://github.com/onsi/gomega/commits/master/) muy recientes.

### 3. [Check](https://github.com/go-check/check)
- Conjunto amplio de aserciones, con funciones avanzadas.
- Compatible con *go test*.
- Repositorio con alrededor de 60 estrellas, sin [commits](https://github.com/go-check/check/commits/master/) desde hace varios años.

## Conclusión
`Check` es una opción completa pero sin mantenimiento desde hace 5 años por lo que la catalogaremos como obsoleta y quedará descartada.

En igualdad de condiciones ante la compatibilidad con *go test*, nuestra elección será *Testify* por la mera diferencia de apoyo por parte de la comunidad basandonos en que el repositorio cuenta con casi 22 mil estrellas más de puntuación.
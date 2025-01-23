# Biblioteca de aserciones
## Criterios de aceptación
 - Mantenimiento y actualización del paquete, que cumpla este requisito es importante debido a que si una herramienta no se corrige o no se actualiza conforme se renuevan versiones del lenguaje, puede quedar obsoleta o que parte de código antiguo deje de funcionar.
    - Es adecuado que haya commits en un periodo de tiempo relativamente corto que indique que el paquete sigue supervisado.
    - Será preferible una herramienta que cuente con mayor apoyo por parte de la comunidad, de manera que obtenga más estrellas en su repositorio de GitHub.
- Sería imperante que la elección no añada dependecias externas. El no utilizar paquetes no incluidos "de fábrica" es aconsejable porque favorece el aumento de la deuda técnica en caso de quedar deprecados.

## Opciones
### 1. [Testify](https://github.com/stretchr/testify)
- Repositorio con casi 24k estrellas, útlimos [commits](https://github.com/stretchr/testify/commits/master/) de hace un mes.

### 2. [Gomega](https://github.com/onsi/gomega)
- Repositorio con 2.2k estrellas, con [commits](https://github.com/onsi/gomega/commits/master/) hace unas semanas.

### 3. [Check](https://github.com/go-check/check)
- Repositorio con alrededor de 60 estrellas, sin [commits](https://github.com/go-check/check/commits/master/) desde hace varios años.

### 4. [Testing](https://pkg.go.dev/testing)
- Testing es un paquete nativo de *Go* que engloba funcionalidades de testing, incluidas las de aserción y matching, útilies para este apartado.
- No contamos con información sobre el repositorio específico por la misma razón que el apartado anterior, viene includio por defecto.

## Conclusión
Contamos con 3 opciones que necesitan de la inclusión de un paquete externo, aunque `Check` cuenta sin commits desde hace años, `Testify` y `Gomega` siguen con mantenimiento pero `Testify` recoge más apoyo por parte de la comunidad. Sin embargo la opción restante, `Testing` está incorporada de manera nativa, con lo cuál será nuestra opción.

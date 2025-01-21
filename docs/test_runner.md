# Herramienta CLI de ejucción

## Criterios
- Es preferible que la opción elegida venga incluida de forma predeterminada en el entorno del lenguaje, manteniendo las buenas prácticas como se menciona en el [guión](https://jj.github.io/IV/documentos/proyecto/4.Tests).

## Opciones
El ecosistema de *Go* incluye una herrmienta integrada de ejecución de tests invoada por el comando: **go test**, integrada de forma de forma nativa desde la versión *1.0*.

Existen herramientas que sustituyen a *go test*, la mayoría incluídas en un framework como [**GoConvey**](https://smartystreets.github.io/goconvey/) y [**Ginkgo**](https://onsi.github.io/ginkgo/) mencionadas en el dicho [documento](./test_frameworks.md).

Otras simulan el sistema de testing de Go como pueden ser `GoTestSum` y `RichGo`. La carácterística principal de estas herramientas es que añaden a la salida un *formateo* de la salida, añadiendo colores y una sintaxis más legible.

<img src="./images/test_runners.png" alt="Imagen"/>

## Conclusión
"*Go test*" será nuestra elección como herramienta de ejecución de test ya que viene incluido de manera predeterminada.


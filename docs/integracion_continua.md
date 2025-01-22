# Herramientas de Integración Continua
## Criterios de elección
- **Docker**: Acercandonos un poco más al despligue en la nube, el sistema deberá permitir la ejecución de imágenes de *Docker*, en particular la creada en el objetivo anterior.

- **GitHub**: Las herramientas seleccionadas deben ser compatibles en GitHub para aprovechar que el desarrollo se está realizando en esta funcional plataforma.

- **Costo**: No se seleccionarán herramientas que sean completamente de pago, deben ser gratis o tener opciones gratuitas.

## Opciones
### [GitHub Actions](https://docs.github.com/es/actions)
- Permite la ejecución de contenedores Docker.
- Integración nativa con GitHub, configurado en YAML.
- 2000 minutos al mes con el plan gratuito en repositorios privados, ilimitado en repositorios públicos.

### [Circle CI](https://circleci.com/)
- Permite la ejecución de contenedores Docker.
- Compatible con GitHub mediante un archivo YAML.
- 2500 minutos al mes gratis.

### [Semaphore CI](https://semaphoreci.com/)
- Permite la ejecución de contenedores Docker.
- Compatible con GitHub mediante un archivo YAML.
- Uso gratuiro de 1300 minutos al mes.

### [Travis CI](https://www.travis-ci.com/)
- Permite la ejecución de contenedores Docker.
- Compatible con GitHub mediante un archivo YAML.
- Usos limitados, se descuentan en base a los segundos de ejecución.

## Conclusión
Teniendo en cuenta que todas son compatibles con la ejecución de imágenes de Docker y que además pueden ser utilizadas en GitHub, hay una que destaca por el hecho de estar incluida de forma nativa, lo que hace que nos evitemos procesos de configuración externos: las `GitHub Actions`.

Para el otro sistema vamos a escoger `Circle CI` simplemente porque nos permite más tiempo de ejecución que las opciones restantes.
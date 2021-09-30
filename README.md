# Go2021
## Integrantes:
- Bartolo, Gisele
- Gauna, Santiago
- Villegas, Matias
## Consigna:
Crear una funcion que dada una cadena con un formato determinado genere una instancia de una estructura con los valores de los campos correspondientes al formato de la cadena
## Consideraciones:
- Agregar test de unidad
- Se espera una cobertura del 80% o mas
- Tener en cuenta distintos casos para los test, por ejemplo cadenas inválidas.
- Usar go modules (go mod)
- Entregar un repositorio de github publico con el README.md que explique la solucion que hicieron
- Se pueden hacer grupos de 2 o 3 integrantes.
### Solución:
Dada una cadena de strings se toman los dos primeros caracteres y se valida que sean letras. A continuación se toman los próximos dos para verificar que sean números ya que determina la longitud de la próxima subcadena. Para la próxima subcadena se verifica que la longitud de la misma coincida con el número especificado. En caso de que alguna condición no se cumpla se retorna un resultado vacio con un error de "Estructura inválida".

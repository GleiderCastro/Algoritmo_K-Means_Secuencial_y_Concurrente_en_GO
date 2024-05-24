# 1. Introducción 

K-means es un método de agrupación muy utilizado en la minería de datos, tiene como objetivo dividir los datos en grupos o clústeres de tal manera que los puntos en el mismo clúster sean más similares entre sí.
Nuestro objetivo es aplicar este algoritmo con el uso de goroutines en el lenguaje GO, utilizando así la programación concurrente de manera eficiente.
Utilizaremos datos aleatorios creados en el mismo código (1000000 registros) y nuestro número de clústeres será de 3 (K=3).
# 2. Algoritmo Secuencial(K-Means) secuencial y de manera concurrente.
## 2.1. Explicación del Algoritmo Secuencial(K-Means). 
El código implementa el algoritmo K-means para realizar clustering en un conjunto de puntos. Aquí un desglose de su funcionamiento:
### 2.1.1. Generación de puntos aleatorios:
La función generatePoints crea un conjunto de n puntos con coordenadas X e Y aleatorias dentro del rango (0, 100).
### 2.1.2.Inicialización de Clusters:
Se define el número de clusters deseado (k) en la función main.
La función kmeans inicializa k clústeres con centros aleatorios dentro del mismo rango (0, 100).
### 2.1.3. Loop principal de K-means:
El loop principal itera hasta que los centros de los clústeres convergen:
#### Asignación de puntos a clústeres:
- Para cada punto, se calcula la distancia a cada centro de clúster.
- El punto se asigna al clúster con el centro más cercano.
#### Actualización de centros de clústeres:
* Se calcula la media de las coordenadas X e Y de los puntos pertenecientes a cada clúster.
* Los centros de los clústeres se actualizan con las nuevas medias.
* Se comprueba si los nuevos centros son iguales a los anteriores (convergencia).
### 2.1.4. Convergencia y resultado:
* El loop termina cuando los centros de los clusters dejan de cambiar, indicando que se ha encontrado una configuración estable.
* Finalmente, se imprimen los centros de cada cluster y los puntos asignados a cada uno.
## 2.2. Explicación del algoritmo K-means concurrente.
Nuestro código proporcionado implementa el algoritmo K-means de manera concurrente para agrupar un conjunto de datos en clústeres. Aquí un desglose de su funcionamiento: 
### 2.2.1. Inicialización:	
* Se genera un conjunto de datos aleatorios (data) con nSamples puntos y nFeatures características. 
* Se define el número de clústeres deseados (k). 
* Se inicializan los centroides (centroids) aleatoriamente a partir de los datos. 
* Se crean variables para almacenar las asignaciones de puntos a clústeres (assignments) y los vectores para un procesamiento eficiente (vectors).
### 2.2.2. Loop principal de K-means:
El algoritmo itera por un número máximo de iteraciones (maxIterations):
#### Asignación de puntos a clusters concurrente:
* Se utiliza un grupo de goroutines (wg) para asignar puntos a los centroides más cercanos en paralelo.
* Cada goroutine itera sobre los puntos y calcula la distancia a cada centroide.
* El punto se asigna al cluster con el centroide más cercano y la asignación se almacena en assignments.
#### Actualización de centroides concurrente:
* Se utilizan goroutines para actualizar los centroides en paralelo. 
* Se crea un mutex (mu) para evitar condiciones de carrera durante la actualización de cada centroide. 
* Se crean grupos de clusters (clusters) para almacenar los puntos asignados a cada cluster. 
* Cada goroutine calcula la media de los puntos pertenecientes a su cluster asignado y actualiza el correspondiente centroide.

### 2.2.3 Resultado e impresión:
* Se imprimen los centroides finales.
* Se imprime un número definido de asignaciones de puntos a clusters (printAssignments).
* Se muestra el tiempo de ejecución del algoritmo.

# 3. Justificación del uso de los mecanismos de paralelización y sincronización utilizados.
## 3.1. Paralelización:
La asignación de puntos a clusters es una tarea independiente para cada punto, por lo que se presta naturalmente a la paralelización. Utilizar goroutines permite aprovechar múltiples núcleos de la CPU para acelerar esta fase del algoritmo, mejorando significativamente el rendimiento.

## 3.2. Justificación de la sincronización:
La actualización de cada centroide requiere sumar los puntos asignados a ese cluster y calcular la media. El mutex (mu) garantiza que solo un goroutine acceda y modifique un centroide a la vez, evitando condiciones de carrera y resultados inconsistentes. Esto asegura la integridad de los datos y la convergencia del algoritmo.
# 4. Explicación de las pruebas realizadas y resultados.

## 4.1. Algoritmo K-means en Go de manera secuencial
Este código implementa el algoritmo K-means en Go de manera secuencial. Inicializa aleatoriamente los centroides, asigna puntos de datos a los centroides más cercanos en cada iteración y actualiza los centroides con la media de los puntos asignados. Finalmente, imprime los centroides finales y las asignaciones de los puntos de datos.
![image](https://github.com/GleiderCastro/Algoritmo_K-Means_Secuencial_y_Concurrente_en_GO/assets/81375850/229949b1-fe6d-46c2-82b8-659757b7ad47)

## 4.2. Algoritmo K-means en Go de manera concurrente
La diferencia significativa entre esta versión concurrente y la implementación secuencial anterior radica en cómo maneja la asignación de puntos de datos y las actualizaciones de centroides. Al utilizar goroutines y primitivas de sincronización, el código realiza estas tareas de manera concurrente. Este enfoque puede mejorar potencialmente la velocidad de ejecución general aprovechando los múltiples núcleos o subprocesos disponibles en el sistema.

![image](https://github.com/GleiderCastro/Algoritmo_K-Means_Secuencial_y_Concurrente_en_GO/assets/81375850/37398fa9-39bc-41b4-a5d2-fe118859db42)

# 5. Explicación de la simulación realizada con promela, pegar las imágenes de evidencia. 

En la simulación con promela, se definen 3 canales(data_chanel,centroid_chanel y convergence_chanel) estos canales sirven para la comunicación entre los diferentes procesos. Después los procesos de “datageneratos” y “centroidinitializer” son los procesos donde se generan los datos que necesitamos y los centroides.
Después se verifican si los procesos convergieron con el proceso “convergencechecker” y envía una señal a través del canal correspondiente y por último el proceso main inicializa todos los procesos.

![image](https://github.com/GleiderCastro/Algoritmo_K-Means_Secuencial_y_Concurrente_en_GO/assets/81375850/cc8e48c1-5c60-4f33-8dbb-f6614ae220d1)

![image](https://github.com/GleiderCastro/Algoritmo_K-Means_Secuencial_y_Concurrente_en_GO/assets/81375850/367d5359-48c4-4237-9907-1c4b40fb193d)

# 6. Explicación del análisis usando spin.

Una vez ejecutado el código con el comando “spin -a” se crea en nuestra carpeta varios archivos, el más importante es el archivo “pan.c”, luego ejecutamos el código “gcc -o pan pan.c” para que se pueda crear el archivo “pan”(Que aparece en verde en la primera imagen). luego para poder ejecutar el archivo se utiliza el comando “./pan”.

![image](https://github.com/GleiderCastro/Algoritmo_K-Means_Secuencial_y_Concurrente_en_GO/assets/81375850/b7c3a9bb-0baa-4c48-b3f6-74d0add23440)


Esto es lo que se obtiene cuando se ejecuta el comando “./pan” donde nos muestra información importante como:
* La versión del spin.
* El tamaño del vector indicado por el state-vector (988 bytes).
* La profundidad alcanzada que es lo que el spin ha recorrido hasta encontrar el Estado final (226).
* El número de errores que hay en la simulación.
* El número de estados unicos almacenados por el spin durante la simulación (227).
* El número total de transiciones exploradas en la simulación (227).
* El número de conflictos de hash resueltos (0).
* Para terminar en la parte final proporciona información estadística sobre el uso de la memoria en la simulación.

![image](https://github.com/GleiderCastro/Algoritmo_K-Means_Secuencial_y_Concurrente_en_GO/assets/81375850/99deaa96-97dd-4727-b502-9590e6dac327)


# 7. Conclusiones.
El uso de goroutines nos permite asignar puntos a los centroides de manera concurrente, acelerando así el proceso de asignación en comparación con un enfoque secuencial.

Al utilizar concurrencia y goroutines se puede reducir significativamente el tiempo de ejecución del algoritmo, incluso cuando se trabaja con datos grandes como es el caso.
Se concluye que el algoritmo K-means concurrente ofrece un mayor potencial de rendimiento y escalabilidad en sistemas con múltiples núcleos. Sin embargo, es importante considerar las necesidades específicas de la aplicación, la complejidad de la implementación y los recursos disponibles al elegir entre el algoritmo secuencial y el concurrente.
# 8. Link del repositorio de Github y del video.
* A continuación se muestra el link del repositorio: https://github.com/GleiderCastro/Aplicacion-en-GO---PCYD.git.
* Link del video: https://youtu.be/O0qQkNbNE7k 
# 9. Bibliografía
Likebupt. (2023, 1 junio). Agrupación en clústeres K-means: referencia del componente - Azure Machine Learning. Microsoft Learn. https://learn.microsoft.com/es-es/azure/machine-learning/component-reference/k-means-clustering?view=azureml-api-2 
Okereke, G. E., Bali, M. C., Okwueze, C. N., Ukekwe, E. C., Echezona, S. C., & Ugwu, C. I. (2023). K-means clustering of electricity consumers using time-domain features from smart meter data. Journal of Electrical Systems and Information Technology, 10(1), 2. doi:https://doi.org/10.1186/s43067-023-00068-3 

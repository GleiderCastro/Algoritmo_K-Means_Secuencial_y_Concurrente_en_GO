package main

//librerias utilizadas
import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Struc de los vectores
type Vector struct {
	Data []float64
}

// la funcion squaredistance sirve para poder calcular la distancia euclidiana entre
// los vectores
func squaredDistance(p1, p2 *Vector) float64 {
	var sum float64
	for i := 0; i < len(p1.Data); i++ {
		sum += (p1.Data[i] - p2.Data[i]) * (p1.Data[i] - p2.Data[i])
	}
	return sum
}

// la funcion initializecentroids sirve para poder inicializar los centroides de manera
// aleatoria
func initializeCentroids(data [][]float64, k int) [][]float64 {
	nSamples := len(data)
	centroids := make([][]float64, k)
	for i := 0; i < k; i++ {
		idx := rand.Intn(nSamples)
		centroids[i] = make([]float64, len(data[0]))
		copy(centroids[i], data[idx])
	}
	return centroids
}

// la funcion mean sirve para poder calcular la media de los vectores
func mean(vectors []*Vector) *Vector {
	n := len(vectors)
	meanVector := make([]float64, len(vectors[0].Data))
	for _, v := range vectors {
		for i := range v.Data {
			meanVector[i] += v.Data[i] / float64(n)
		}
	}
	return &Vector{meanVector}
}

// la funcion Kmerans es la que contienen el algoritmo de K-means
func kMeans(data [][]float64, k int, maxIterations int) ([][]float64, []int) {
	nSamples := len(data)
	centroids := initializeCentroids(data, k)
	assignments := make([]int, nSamples)
	vectors := make([]*Vector, nSamples)
	for i := range data {
		vectors[i] = &Vector{data[i]}
	}
	for iter := 0; iter < maxIterations; iter++ {
		var wg sync.WaitGroup

		// Asignar puntos a los centroides más cercanos en paralelo
		//se utiliza goroutines para hacer una asignacion paralela
		wg.Add(nSamples)
		for i := 0; i < nSamples; i++ {
			go func(i int) {
				defer wg.Done()
				localCentroids := make([][]float64, k)
				for j := range centroids {
					localCentroids[j] = make([]float64, len(centroids[j]))
					copy(localCentroids[j], centroids[j])
				}
				minDist := squaredDistance(vectors[i], &Vector{localCentroids[0]})
				assignments[i] = 0
				for j := 1; j < k; j++ {
					dist := squaredDistance(vectors[i], &Vector{localCentroids[j]})
					if dist < minDist {
						minDist = dist
						assignments[i] = j
					}
				}
			}(i)
		}
		wg.Wait()

		// Actualizar centroides en paralelo
		//se utilizan goroutines para poder calcular la media de los puntos
		//se utiliza Mutex para poder evitar las condiciones de carrera
		//haciendo que solo un goroutine pueda actulizar un centroide
		//evitando que varios quieran actualizar al mismo tiempo
		//ademas se paraleliza la actualizacion de cada centroide
		//haciendo que cada goroutine actualiza un centroide en especifico
		var mu sync.Mutex
		clusters := make([][]*Vector, k) // Variable clusters
		for i := range clusters {
			clusters[i] = make([]*Vector, 0)
		}
		for i, idx := range assignments {
			clusters[idx] = append(clusters[idx], vectors[i])
		}
		var wgCentroids sync.WaitGroup
		for j := 0; j < k; j++ {
			wgCentroids.Add(1)
			go func(j int) {
				defer wgCentroids.Done()
				mu.Lock()
				defer mu.Unlock()
				centroids[j] = mean(clusters[j]).Data
			}(j)
		}
		wgCentroids.Wait()
	}
	return centroids, assignments
}

// con la funcion printAssignments se imprime las asignaciones de los clusteres
func printAssignments(assignments []int, numAssignments int) {
	if numAssignments > len(assignments) {
		numAssignments = len(assignments)
	}
	for i := 0; i < numAssignments; i++ {
		fmt.Printf("%d: %d\n", i+1, assignments[i])
	}
	fmt.Println()
}

// esta es la funcion main
func main() {
	rand.Seed(time.Now().UnixNano())

	nSamples := 1000000 //es numero de datos que tenemos como pide el enunciado
	nFeatures := 2
	k := 3                //el numero de clusteres
	maxIterations := 1000 //el numero de iteraciones como pide el enunciado

	data := make([][]float64, nSamples)
	for i := range data {
		data[i] = make([]float64, nFeatures)
		for j := range data[i] {
			data[i][j] = rand.NormFloat64()
		}
	}
	//se piden la cantidad de asiganciones que desea observar
	var numAssignments int
	fmt.Print("Ingrese la cantidad de asignaciones que desea imprimir: ")
	fmt.Scan(&numAssignments)
	//aplicacion de K-means
	start := time.Now()
	centroids, assignments := kMeans(data, k, maxIterations)
	elapsed := time.Since(start)
	//se imrpimen los centroides obtenidos
	fmt.Println("Centroids:")
	for _, c := range centroids {
		fmt.Println(c)
	}
	//se imprime las asignaciones a los clusteres
	fmt.Println("Assignments:")
	printAssignments(assignments, numAssignments)
	//se imprime el tiempo transcurrido en el algoritmo
	fmt.Printf("Tiempo transcurrido: %s\n", elapsed)
	//escribir en archivo por ahora
	file, err := os.OpenFile("datos.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()
	dato := elapsed.String() + "\n"
	_, err = file.WriteString(dato)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}
}

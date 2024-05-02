package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Vector struct {
	Data []float64
}

func squaredDistance(p1, p2 *Vector) float64 {
	var sum float64
	for i := 0; i < len(p1.Data); i++ {
		sum += (p1.Data[i] - p2.Data[i]) * (p1.Data[i] - p2.Data[i])
	}
	return sum
}

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

func kMeans(data [][]float64, k int, maxIterations int) ([][]float64, []int) {
	nSamples := len(data)
	centroids := initializeCentroids(data, k)
	assignments := make([]int, nSamples)
	vectors := make([]*Vector, nSamples)
	for i := range data {
		vectors[i] = &Vector{data[i]}
	}

	for iter := 0; iter < maxIterations; iter++ {
		// Asignar puntos a los centroides mas cercanos
		for i := 0; i < nSamples; i++ {
			minDist := squaredDistance(vectors[i], &Vector{centroids[0]})
			assignments[i] = 0
			for j := 1; j < k; j++ {
				dist := squaredDistance(vectors[i], &Vector{centroids[j]})
				if dist < minDist {
					minDist = dist
					assignments[i] = j
				}
			}
		}

		// Actualizar centroides
		clusters := make([][]*Vector, k)
		for i := range clusters {
			clusters[i] = make([]*Vector, 0)
		}
		for i, idx := range assignments {
			clusters[idx] = append(clusters[idx], vectors[i])
		}
		for j := 0; j < k; j++ {
			centroids[j] = mean(clusters[j]).Data
		}
	}

	return centroids, assignments
}

func printAssignments(assignments []int, numAssignments int) {
	if numAssignments > len(assignments) {
		numAssignments = len(assignments)
	}
	for i := 0; i < numAssignments; i++ {
		fmt.Printf("%d: %d\n", i+1, assignments[i])
	}
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	nSamples := 100000    // Numero de muestras
	nFeatures := 2        // Numero de caracteristicas
	k := 3                // Numero de clusters
	maxIterations := 1000 // Numero maximo de iteraciones

	data := make([][]float64, nSamples)
	for i := range data {
		data[i] = make([]float64, nFeatures)
		for j := range data[i] {
			data[i][j] = rand.NormFloat64()
		}
	}

	var numAssignments int
	fmt.Print("Ingrese la cantidad de asignaciones que desea imprimir: ")
	fmt.Scan(&numAssignments)

	start := time.Now()
	centroids, assignments := kMeans(data, k, maxIterations)
	elapsed := time.Since(start)

	fmt.Println("Centroids:")
	for _, c := range centroids {
		fmt.Println(c)
	}

	fmt.Println("Assignments:")
	printAssignments(assignments, numAssignments)

	fmt.Printf("Tiempo transcurrido: %s\n", elapsed)
}

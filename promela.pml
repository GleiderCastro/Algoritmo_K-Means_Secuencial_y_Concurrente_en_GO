mtype = {INIT, DATA, CENTROID, CONVERGED}

chan data_channel = [100] of {mtype, int};  // Canal para datos
chan centroid_channel = [10] of {mtype, int};  // Canal para centroides
chan convergence_channel = [1] of {mtype, int}; // Canal para señalar convergencia

active proctype DataGenerator() {
    data_channel ! INIT, 0; // Iniciar la transmisión de datos
    int i = 0;
    do
    :: data_channel ! DATA, i; // Enviar dato
       i = i + 1;
    od
}

active proctype CentroidInitializer() {
    centroid_channel ! INIT, 0; // Iniciar la transmisión de centroides
    int i = 0;
    do
    :: centroid_channel ! CENTROID, i; // Enviar centroide
       i = i + 1;
    od
}

proctype Centroid() {
    int centroid_value;
    bool converged = false;
    do
    :: centroid_channel ? CENTROID, centroid_value; // Recibir centroide
       printf("Received centroid: %d\n", centroid_value);
       // Lógica de actualización de centroides
    :: convergence_channel ? CONVERGED, _; // Verificar convergencia
       converged = true;
    od
}

proctype Data() {
    int data_value;
    do
    :: data_channel ? DATA, data_value; // Recibir dato
       printf("Received data: %d\n", data_value);
       // Lógica de asignación de datos a centroides
    od
}

proctype ConvergenceChecker() {
    convergence_channel ! INIT, 0; // Iniciar la señal de convergencia
    do
    :: // Lógica para verificar la convergencia y enviar señal
       convergence_channel ! CONVERGED, 0;
    od
}

active proctype Main() {
    run DataGenerator();
    run CentroidInitializer();
    run Data();
    run Centroid();
    run ConvergenceChecker();
}
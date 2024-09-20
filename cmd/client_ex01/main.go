package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/zkhrg/go_team00/pkg/detector"
)

const freqCountToCalculateParameters = 100

func main() {
	var anomalyCoefficient float64
	flag.Float64Var(&anomalyCoefficient, "k", 0, "Standard deviation anomaly coefficient")
	flag.Parse()

	if anomalyCoefficient <= 0 {
		flag.Usage()
		os.Exit(1)
	}

	freqChan := make(chan float64)
	defer close(freqChan)

	anomaliesDetector := detector.NewAnomaliesDetector(anomalyCoefficient, freqCountToCalculateParameters)
	anomaliesDetector.DetectAnomalies(freqChan)

	var readValue float64

	for {
		_, err := fmt.Fscan(os.Stdin, &readValue)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading value from stdin: %s", err)
		}
		freqChan <- readValue
	}
}

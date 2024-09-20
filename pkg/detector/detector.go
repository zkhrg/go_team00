package detector

import (
	"fmt"
	"math"
	"sync"
)

type AnomaliesDetector struct {
	SessionID                      string
	Mean                           float64
	Sd                             float64
	mu                             sync.Mutex
	CountRecords                   uint64
	CountAnomalies                 uint64
	sumForMean                     float64
	sumForSD                       float64
	anomalyCoefficient             float64
	initValues                     []float64
	freqCountToCalculateParameters uint64
}

func (a *AnomaliesDetector) InitCalculateMeanSD(freq float64) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.initValues = append(a.initValues, freq)
	a.CountRecords++

	if a.CountRecords < a.freqCountToCalculateParameters {
		return false
	}

	for _, value := range a.initValues {
		a.sumForMean += value
	}

	a.Mean = a.sumForMean / float64(a.CountRecords)

	for _, value := range a.initValues {
		a.sumForSD += math.Pow(value-a.Mean, 2)
	}

	a.Sd = math.Sqrt(a.sumForSD / float64(a.CountRecords))
	a.initValues = nil // Clear the slice

	return true
}

func (a *AnomaliesDetector) ProcessNextFrequency(freq float64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.CountRecords++
	previousMean := a.Mean
	a.sumForMean += freq
	a.Mean = a.sumForMean / float64(a.CountRecords)

	a.sumForSD += math.Pow(freq-previousMean, 2)
	a.Sd = math.Sqrt(a.sumForSD / float64(a.CountRecords))

	if math.Abs(freq-a.Mean) > a.anomalyCoefficient*a.Sd {
		fmt.Printf("Found anomaly: %.4f\n", freq)
		a.CountAnomalies++
	}
}

func (a *AnomaliesDetector) DetectAnomalies(freqChan chan float64) {
	go func() {
		fmt.Printf("Compute parameters stage (need %d frequencies to calculate).\n", a.freqCountToCalculateParameters)

		for freq := range freqChan {
			if a.InitCalculateMeanSD(freq) {
				break
			}
		}

		fmt.Printf("Anomaly detection stage. Parameters - mean: %.4f, sd: %.4f, k*sd: %.4f\n",
			a.Mean, a.Sd, a.Sd*a.anomalyCoefficient)

		for freq := range freqChan {
			a.ProcessNextFrequency(freq)
		}
	}()
}

func NewAnomaliesDetector(anomalyCoefficient float64, freqCountToCalculateParameters uint64) *AnomaliesDetector {
	return &AnomaliesDetector{
		anomalyCoefficient:             anomalyCoefficient,
		freqCountToCalculateParameters: freqCountToCalculateParameters,
	}
}

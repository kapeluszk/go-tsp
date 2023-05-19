package base

import (
	"log"
	"math"
)

type Population struct {
	tours []Tour
}

func (a *Population) InitEmpty(pSize int) {
	a.tours = make([]Tour, pSize)
}

func (a *Population) InitPopulation(pSize int, tm, tr TourManager) {
	a.tours = make([]Tour, pSize)
	tmp := float64(pSize) * 0.99

	for i := 0; i < pSize; i++ {
		nT := Tour{}
		if pSize >= int(math.Round(tmp)) {
			nT.InitTourCities(tm)
			a.SaveTour(i, nT)
		} else {
			tr.destCities = ShuffleCities(tr.destCities)
			nT.InitTourCities(tr)
			a.SaveTour(i, nT)
		}

	}
}

func (a *Population) SaveTour(i int, t Tour) {
	a.tours[i] = t
}

func (a *Population) GetTour(i int) *Tour {
	return &a.tours[i]
}

func (a *Population) PopulationSize() int {
	return len(a.tours)
}

func (a *Population) GetFittest() *Tour {
	fittest := a.tours[0]
	// Loop through all tours taken by population and determine the fittest
	for i := 0; i < a.PopulationSize(); i++ {
		log.Println("Current Tour: ", i)
		if fittest.Fitness() <= a.GetTour(i).Fitness() {
			fittest = *a.GetTour(i)
		}
	}
	return &fittest
}
func (a *Population) GetTop(num int) []Tour {
	tours := make([]Tour, num)
	// Inicjalizacja tours slice z pierwszymi num trasami
	for i := 0; i < num; i++ {
		tours[i] = *a.GetTour(i)
	}

	// Pętla przez wszystkie trasy w populacji i wyznaczenie najfitalniejszej
	for i := num; i < a.PopulationSize(); i++ {
		log.Println("Current Tour: ", i)
		currentTour := a.GetTour(i)
		// Znajdź indeks najmniej fit trasy wśród wybranych tras
		minIndex := 0
		minFitness := tours[minIndex].Fitness()
		for j := 1; j < num; j++ {
			if tours[j].Fitness() < minFitness {
				minIndex = j
				minFitness = tours[j].Fitness()
			}
		}
		// Zamień najmniej fit trasę na bieżącą trasę, jeśli jest bardziej fit
		if currentTour.Fitness() > minFitness {
			tours[minIndex] = *currentTour
		}
	}

	return tours
}

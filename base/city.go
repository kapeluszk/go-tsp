package base

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// City : coordinates of city
type City struct {
	number int
	x      int
	y      int
}

// GenerateRandomCity : Generate city with random coordinates
func GenerateRandomCity() City {
	c := City{}
	c.x = rand.Intn(100) * 100
	c.y = rand.Intn(100) * 100
	return c
}

// GenerateCity : Generate city with user defined coordinates
func GenerateCity(x int, y int) City {
	c := City{}
	c.x = x
	c.y = y
	return c
}

// SetLocation : User defined coordinates for a city
func (a *City) SetLocation(x int, y int) {
	a.x = x
	a.y = y
}

// DistanceTo : distance of current city to target city
func (a *City) DistanceTo(c City) float64 {
	idx := a.x - c.x
	idy := a.y - c.y

	if idx < 0 {
		idx = -idx
	}
	if idy < 0 {
		idy = -idy
	}

	fdx := float64(idx)
	fdy := float64(idy)

	fd := math.Sqrt((fdx * fdx) + (fdy * fdy))
	return fd
}

func (a *City) X() int {
	return a.x
}

func (a *City) Y() int {
	return a.y
}

func (a City) String() string {
	return fmt.Sprintf("{x%d y%d}", a.x, a.y)
}

// ShuffleCities : return a shuffled []City given input []City
func ShuffleCities(in []City) []City {
	out := make([]City, len(in), cap(in))
	perm := rand.Perm(len(in))
	for i, v := range perm {
		out[v] = in[i]
	}
	return out
}

func ReadCitiesFromFile(filename string) ([]City, error) {

	//otwieramy plik
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//otwieramy skaner i tworzymy tablicę miast
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var cities []City
	// num, err := strconv.Atoi(scanner.Text())
	// if err != nil {
	// 	fmt.Println("first line is not an amount of cities")
	// }
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 3 {
			return nil, fmt.Errorf("incorrect format of input file")
		}

		// Konwersja wartości na odpowiednie typy
		num, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}
		xArg, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}
		yArg, err := strconv.Atoi(fields[2])
		if err != nil {
			return nil, err
		}

		// Dodanie miasta do tablicy
		cities = append(cities, City{
			number: num,
			x:      xArg,
			y:      yArg,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cities, nil
}

func GenerateTxtInstance(citiesAmount int) (string, error) {
	var newFileName string
	fmt.Println("Podaj nazwę nowego pliku")
	fmt.Scanln(&newFileName)

	fileNew, err := os.Create(newFileName)
	if err != nil {
		log.Fatal()
	}
	defer fileNew.Close()

	writer := bufio.NewWriter(fileNew)

	fmt.Fprintf(fileNew, "%d\n", citiesAmount)
	writer.Flush()

	for i := 0; i < citiesAmount; i++ {
		x := rand.Intn(2000)
		y := rand.Intn(2000)
		fmt.Fprintf(writer, "%d %d %d\n", i+1, x, y)
		writer.Flush()
	}
	return newFileName, err
}

func Distance(city1, city2 City) float64 {
	xDist := city1.x - city2.x
	yDist := city1.y - city2.y
	return math.Sqrt(math.Pow(float64(xDist), 2) + math.Pow(float64(yDist), 2))
}

func NearestNeighbor(cities []City) []City {
	visited := make(map[int]bool) // mapa odwiedzonych miast
	tour := make([]City, len(cities))

	// Wybieramy losowe miasto jako punkt początkowy
	currentCity := cities[0]
	tour[0] = currentCity
	visited[currentCity.number] = true

	// Przeglądamy każde miasto w kolejności "najbliższego sąsiada"
	for i := 1; i < len(cities); i++ {
		nearestDistance := math.MaxFloat64
		var nearestCity City

		// Szukamy najbliższego miasta
		for _, city := range cities {
			if !visited[city.number] {
				dist := Distance(currentCity, city)
				if dist < nearestDistance {
					nearestDistance = dist
					nearestCity = city
				}
			}
		}

		// Dodajemy najbliższe miasto do trasy i oznaczamy jako odwiedzone
		tour[i] = nearestCity
		visited[nearestCity.number] = true
		currentCity = nearestCity // aktualizujemy aktualne miasto
	}

	// Dodajemy punkt początkowy na końcu trasy
	tour = append(tour, tour[0])

	return tour
}

func CalculateTotalDistance(tour []City) float64 {
	totalDistance := 0.0
	for i := 0; i < len(tour)-1; i++ {
		currentCity := tour[i]
		nextCity := tour[i+1]
		dist := Distance(currentCity, nextCity)
		totalDistance += dist
	}
	return totalDistance
}

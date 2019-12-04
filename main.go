package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func getFuelForMass(mass int) int {
	fuel := 0
	for mass > 6 {
		temp := (mass / 3) - 2
		fuel += temp
		fmt.Println("Masse:", mass, "Added", temp, "Fuel :", fuel)
		mass = temp
	}
	return fuel
}

func day1() {
	csvFile, _ := os.Open("day1.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var total = 0
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		value, _ := (strconv.Atoi(line[0]))
		wiegth := getFuelForMass(value)
		total += wiegth
		fmt.Println(wiegth, total)

	}
	fmt.Println(total)

}

func intCodeEngine(verb int, noun int) int {
	csvFile, _ := os.Open("day2")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	ret := 0
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		var t2 = []int{}

		for _, i := range line {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			t2 = append(t2, j)
		}
		t2[1] = verb
		t2[2] = noun

		for i := 0; i < len(t2); i += 4 {
			switch opcode := line[i]; opcode {
			case "1":
				t2[t2[i+3]] = t2[t2[i+1]] + t2[t2[i+2]]
			case "2":
				t2[t2[i+3]] = t2[t2[i+1]] * t2[t2[i+2]]
			case "99":
				fmt.Println("END ", t2[0])
				ret = t2[0]
				break
			default:
				break
			}

		}

	}

	return ret
}

func day2() {
	value := 0
	verb := 0
	noun := 0
	value = intCodeEngine(verb, noun)
	for value != 19690720 {

		if verb < 99 {
			verb++
		} else {
			verb = 0
			noun++
		}
		value = intCodeEngine(verb, noun)
	}
	fmt.Println("FOUND ", verb, noun, value, (verb*100)+noun)
}

type Point struct {
	x     int
	y     int
	steps float64
}

type Wire struct {
	Points []Point
}

func orientation(p Point, q Point, r Point) int {
	val := (q.y-p.y)*(r.x-q.x) - (q.x-p.x)*(r.y-q.y)

	if val == 0 {
		return 0
	} else if val > 0 {
		return 1
	} else {
		return 2
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func day3() {
	csvFile, _ := os.Open("day3")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	start := time.Now()
	var Wires []Wire
	for {
		posY := 0
		posX := 0
		line, error := reader.Read()
		if error == io.EOF {
			break
		}
		Wire := Wire{Points: []Point{Point{x: posX, y: posY}}}
		s := 0.
		for _, i := range line {
			direction := i[0]
			val, _ := strconv.Atoi(i[1:])
			for i := 0; i < val; i++ {
				switch d := direction; d {
				case 'D':
					posY--
				case 'U':
					posY++
				case 'R':
					posX++
				case 'L':
					posX--
				}
			}
			s += float64(val)
			Wire.Points = append(Wire.Points, Point{x: posX, y: posY, steps: s})
		}
		Wires = append(Wires, Wire)
	}
	step1 := 0.
	step2 := 0.
	var steps float64
	var mans float64
	for i, w1 := range Wires {
		for _, w2 := range Wires[i+1:] {

			step1 = 0.
			for j, pw1 := range w1.Points[:len(w1.Points)-1] {
				step2 = 0.
				w1p1 := pw1
				w1p2 := w1.Points[j+1]

				for h, pw2 := range w2.Points[:len(w2.Points)-1] {
					w2p1 := pw2
					w2p2 := w2.Points[h+1]
					if j+h != 0 && (orientation(w1p1, w1p2, w2p1) != orientation(w1p1, w1p2, w2p2)) && (orientation(w2p1, w2p2, w1p1) != orientation(w2p1, w2p2, w1p2)) {
						var intersect Point
						step1 := w1p1.steps
						step2 := w2p1.steps
						if w1p1.x == w1p2.x {
							intersect = Point{x: w1p1.x, y: w2p1.y}
							step1 += absoluteDiff(w1p1.y, intersect.y)
							step2 += absoluteDiff(w2p1.x, intersect.x)
						} else {
							intersect = Point{x: w2p1.x, y: w1p1.y}
							step1 += absoluteDiff(w1p1.x, intersect.x)
							step2 += absoluteDiff(w2p1.y, intersect.y)

						}
						man := Manhattan([]int{intersect.x, intersect.y}, []int{0, 0})
						step := step1 + step2

						if mans == 0 || man < mans {
							mans = man
						}
						if steps == 0 || step < steps {
							steps = step
						}
					}
				}
			}
		}
	}
	fmt.Println("M1", mans, steps)
	fmt.Println("AAA", step1, step2, step1+step2)
	elapsed := time.Since(start)
	fmt.Printf("M1 took %s", elapsed)

}

func absoluteDiff(a, b int) float64 {
	return math.Max((float64(a)), (float64(b))) - math.Min((float64(a)), (float64(b)))
}

func day3Sol1() {
	csvFile, _ := os.Open("day3")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	start := time.Now()
	// r := new(big.Int)
	// fmt.Println(r.Binomial(1000, 10))

	var Wires []Wire
	for {
		posY := 0
		posX := 0
		line, error := reader.Read()
		if error == io.EOF {
			break
		}
		Wire := Wire{Points: []Point{Point{x: posX, y: posY}}}
		for _, i := range line {
			direction := i[0]
			val, _ := strconv.Atoi(i[1:])
			for i := 0; i < val; i++ {
				switch d := direction; d {
				case 'D':
					posY--
				case 'U':
					posY++
				case 'R':
					posX++
				case 'L':
					posX--
				}
				Wire.Points = append(Wire.Points, Point{x: posX, y: posY})
			}
		}
		Wires = append(Wires, Wire)
	}
	var steps int
	var mans float64
	for i, w1 := range Wires {
		for _, w2 := range Wires[i+1:] {
			for g, pw1 := range w1.Points {
				for h, pw2 := range w2.Points[i+1:] {
					if pw1.x == pw2.x && pw1.y == pw2.y {
						man := Manhattan([]int{pw1.x, pw1.y}, []int{0, 0})
						step := g + h + 1
						if mans == 0 || man < mans {
							mans = man
						}
						if steps == 0 || step < steps {
							steps = step
						}
					}
				}
			}
		}
	}
	fmt.Println(steps, mans)
	elapsed := time.Since(start)
	fmt.Printf("M2 took %s", elapsed)

}
func Manhattan(vector1, vector2 []int) float64 {
	var distance float64 = 0
	for i, val1 := range vector1 {
		distance += math.Abs(float64(val1) - float64(vector2[i]))
	}

	return distance
}

func main() {
	day3()
	day3Sol1()

}

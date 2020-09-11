package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Box represents a package to be wrapped
type Box struct {
	length int
	width  int
	height int
}

func minFromSlice(candidates []int) int {
	var min int
	for i, candidate := range candidates {
		if i == 0 || candidate < min {
			min = candidate
		}
	}
	return min
}

func specToBox(spec string) (Box, error) {
	values := strings.Split(spec, "x")
	length, err := strconv.Atoi(values[0])
	if err != nil {
		return Box{}, err
	}
	width, err := strconv.Atoi(values[1])
	if err != nil {
		return Box{}, err
	}
	height, err := strconv.Atoi(values[2])
	if err != nil {
		return Box{}, err
	}
	return Box{
		length: length,
		width:  width,
		height: height,
	}, nil
}

func calculateNeededPaper(box Box) int {
	var sides []int
	sides = append(sides, box.height*box.length)
	sides = append(sides, box.length*box.width)
	sides = append(sides, box.width*box.height)
	smallestSide := minFromSlice(sides)

	paper := smallestSide
	for _, side := range sides {
		paper = paper + 2*side
	}

	return paper
}

func calculateNeededRibbon(box Box) int {
	var perimeters []int
	perimeters = append(perimeters, box.height*2+box.length*2)
	perimeters = append(perimeters, box.length*2+box.width*2)
	perimeters = append(perimeters, box.width*2+box.height*2)
	smallestPerimeter := minFromSlice(perimeters)
	return smallestPerimeter + box.length*box.height*box.width
}

func partOne() {
	runningTotal := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		spec := scanner.Text()
		box, err := specToBox(spec)
		if err != nil {
			panic(err)
		}
		paper := calculateNeededPaper(box)
		runningTotal = runningTotal + paper
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(runningTotal)
}

func partTwo() {
	runningTotal := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		spec := scanner.Text()
		box, err := specToBox(spec)
		if err != nil {
			panic(err)
		}
		ribbon := calculateNeededRibbon(box)
		runningTotal = runningTotal + ribbon
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(runningTotal)
}

func main() {
	partTwo()
}

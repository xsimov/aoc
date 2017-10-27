// Triangler solves the puzzle for AOC day3: count how many of the input triplets of values could make a valid triangle.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	triangles, err := getTriangles()
	if err != nil {
		log.Fatalf("could not import triangles: %v", err)
	}
	fmt.Println(triangles)
}

func getTriangles() ([][]int, error) {
	var path = flag.String("file", "assets/triangles.txt", "Select a file containing space-separated triangle values' lines")
	var example = flag.String("example", "", "Pass in an example that will ignore the file path")

	flag.Parse()

	if *example != "" {
		return extractTriangles([]string{*example})
	}

	if fileContents, err := ioutil.ReadFile(*path); err == nil {
		return extractTriangles(strings.Split(string(fileContents), "\n"))
	}

	return nil, fmt.Errorf("error opening file: %v", *path)
}

func extractTriangles(lines []string) (triangles [][]int, err error) {
	var triangle []int
	for _, line := range lines {
		splitted := strings.Split(line, " ")
		for _, number := range splitted {
			if number != "" {
				num, err := strconv.Atoi(number)
				if err != nil {
					return nil, fmt.Errorf("could not convert %q from %q into an int", number, line)
				}
				triangle = append(triangle, num)
			}
		}
		triangles = append(triangles, triangle)
		triangle = nil
	}
	return
}

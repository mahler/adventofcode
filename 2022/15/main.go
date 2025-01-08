package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Point struct {
	x, y int
}

type Pair struct {
	sensor Point
	beacon Point
}

type Line struct {
	start Point
	end   Point
}

type Range struct {
	start int
	end   int
}

func parsePair(line string) (Pair, error) {
	var sx, sy, bx, by int
	_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
	if err != nil {
		return Pair{}, err
	}
	return Pair{
		sensor: Point{x: sx, y: sy},
		beacon: Point{x: bx, y: by},
	}, nil
}

func parseInput(filename string) ([]Pair, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pairs []Pair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		pair, err := parsePair(line)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, pair)
	}
	return pairs, scanner.Err()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p Point) manhattanDistance(other Point) int {
	return abs(p.x-other.x) + abs(p.y-other.y)
}

func (p Pair) coverDistance() int {
	return p.sensor.manhattanDistance(p.beacon)
}

func (p Pair) topLeft() Line {
	dist := p.coverDistance() + 1
	return Line{
		start: Point{x: p.sensor.x - dist, y: p.sensor.y},
		end:   Point{x: p.sensor.x, y: p.sensor.y - dist},
	}
}

func (p Pair) topRight() Line {
	dist := p.coverDistance() + 1
	return Line{
		start: Point{x: p.sensor.x, y: p.sensor.y - dist},
		end:   Point{x: p.sensor.x + dist, y: p.sensor.y},
	}
}

func (p Pair) bottomLeft() Line {
	dist := p.coverDistance() + 1
	return Line{
		start: Point{x: p.sensor.x - dist, y: p.sensor.y},
		end:   Point{x: p.sensor.x, y: p.sensor.y + dist},
	}
}

func (p Pair) bottomRight() Line {
	dist := p.coverDistance() + 1
	return Line{
		start: Point{x: p.sensor.x, y: p.sensor.y + dist},
		end:   Point{x: p.sensor.x + dist, y: p.sensor.y},
	}
}

func (p Pair) covers(point Point) bool {
	return p.sensor.manhattanDistance(point) <= p.coverDistance()
}

func (p Pair) coveredXs(row int) *Range {
	xOffset := p.coverDistance() - abs(p.sensor.y-row)
	if xOffset < 0 {
		return nil
	}
	return &Range{
		start: p.sensor.x - xOffset,
		end:   p.sensor.x + xOffset + 1,
	}
}

func (l Line) slope() int {
	if l.end.y-l.start.y > 0 {
		return 1
	}
	return -1
}

func (l Line) yIntercept() int {
	return l.start.y - l.slope()*l.start.x
}

func (l Line) yAt(x int) *int {
	if l.start.x > x || l.end.x < x {
		return nil
	}
	y := l.slope()*x + l.yIntercept()
	return &y
}

func (l Line) overlap(other Line) *Line {
	if l.slope() != other.slope() || l.yIntercept() != other.yIntercept() {
		return nil
	}
	x := max(l.start.x, other.start.x)
	y := l.yAt(x)
	if y == nil {
		return nil
	}
	start := Point{x: x, y: *y}

	x = min(l.end.x, other.end.x)
	y = l.yAt(x)
	if y == nil {
		return nil
	}
	end := Point{x: x, y: *y}
	return &Line{start: start, end: end}
}

func (l Line) interception(other Line) *Point {
	if l.slope() == other.slope() {
		return nil
	}
	yInterceptDiff := other.yIntercept() - l.yIntercept()
	slopeDiff := l.slope() - other.slope()
	x := yInterceptDiff / slopeDiff
	y := l.yAt(x)
	if y == nil {
		return nil
	}
	otherY := other.yAt(x)
	if otherY == nil || *y != *otherY {
		return nil
	}
	return &Point{x: x, y: *y}
}

func solvePart1(pairs []Pair, exampleRow int) int {
	var ranges []Range
	for _, pair := range pairs {
		if r := pair.coveredXs(exampleRow); r != nil {
			ranges = append(ranges, *r)
		}
	}

	// Sort ranges by start position
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	// Merge overlapping ranges
	var merged []Range
	if len(ranges) > 0 {
		merged = append(merged, ranges[0])
		for i := 1; i < len(ranges); i++ {
			last := &merged[len(merged)-1]
			if ranges[i].start <= last.end {
				if ranges[i].end > last.end {
					last.end = ranges[i].end
				}
			} else {
				merged = append(merged, ranges[i])
			}
		}
	}

	// Calculate total coverage
	covered := 0
	for _, r := range merged {
		covered += r.end - r.start
	}

	// Count beacons and sensors on the row
	seen := make(map[Point]bool)
	for _, pair := range pairs {
		if pair.sensor.y == exampleRow {
			seen[pair.sensor] = true
		}
		if pair.beacon.y == exampleRow {
			seen[pair.beacon] = true
		}
	}

	return covered - len(seen)
}

func solvePart2(pairs []Pair, maxCoord int) int64 {
	// Create maps to store lines by y-intercept
	topRight := make(map[int][]Line)
	topLeft := make(map[int][]Line)
	for _, pair := range pairs {
		tr := pair.topRight()
		yi := tr.yIntercept()
		topRight[yi] = append(topRight[yi], tr)

		tl := pair.topLeft()
		yi = tl.yIntercept()
		topLeft[yi] = append(topLeft[yi], tl)
	}

	// Find overlapping lines
	var positiveSlopes []Line
	var negativeSlopes []Line

	for _, pair := range pairs {
		bl := pair.bottomLeft()
		yi := bl.yIntercept()
		for _, tr := range topRight[yi] {
			if ol := bl.overlap(tr); ol != nil {
				positiveSlopes = append(positiveSlopes, *ol)
			}
		}

		br := pair.bottomRight()
		yi = br.yIntercept()
		for _, tl := range topLeft[yi] {
			if ol := br.overlap(tl); ol != nil {
				negativeSlopes = append(negativeSlopes, *ol)
			}
		}
	}

	// Find intersection point
	for _, pos := range positiveSlopes {
		for _, neg := range negativeSlopes {
			if p := pos.interception(neg); p != nil {
				if p.x >= 0 && p.x <= maxCoord && p.y >= 0 && p.y <= maxCoord {
					covered := false
					for _, pair := range pairs {
						if pair.covers(*p) {
							covered = true
							break
						}
					}
					if !covered {
						return int64(p.x)*4000000 + int64(p.y)
					}
				}
			}
		}
	}

	return -1
}

func main() {
	pairs, err := parseInput("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Part 1: In the row where y=2000000, how many positions cannot contain a beacon?")
	fmt.Println(solvePart1(pairs, 2000000))

	fmt.Println()
	fmt.Println("Part 2: Find the only possible position for the distress beacon. What is its tuning frequency?")
	fmt.Println(solvePart2(pairs, 4000000))
}

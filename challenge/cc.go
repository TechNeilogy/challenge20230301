package challenge

// Coding Challenge Solution Spring 2023

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Maze struct {
	xSize, ySize   int
	xStart, yStart int
	xGoal, yGoal   int
}

type Dir struct {
	path   string
	dx, dy int
}

var Dirs []Dir = []Dir{
	{"U", 0, -1},
	{"D", 0, 1},
	{"L", -1, 0},
	{"R", 1, 0},
}

func GetOpenDirs(path string) (rtn []*Dir) {
	hash := md5.Sum([]byte(path))
	four := hex.EncodeToString(hash[:])[0:4]
	for i, c := range four {
		if c >= 'b' {
			rtn = append(rtn, &Dirs[i])
		}
	}
	return rtn
}

type Path struct {
	path string
	x, y int
}

func (m *Maze) Move(x, y, xd, yd int) (int, int) {
	x += xd
	if x < 0 || x >= m.xSize {
		return -1, -1
	}
	y += yd
	if y < 0 || y >= m.ySize {
		return -1, -1
	}
	return x, y
}

func (m *Maze) BreadthFirstSearch(
	paths []Path,
) *Path {
	for len(paths) > 0 {

		h := paths[0]

		if h.x == m.xGoal && h.y == m.yGoal {
			return &h
		}

		paths = paths[1:]

		for _, d := range GetOpenDirs(h.path) {
			xn, yn := m.Move(h.x, h.y, d.dx, d.dy)
			if xn >= 0 {
				paths = append(
					paths,
					Path{
						h.path + d.path,
						xn,
						yn,
					},
				)
			}
		}
	}

	return nil
}

// Setting findAll to true results in an exhaustive search.
func (m *Maze) depthFirstSearch(
	path string,
	x int,
	y int,
	pathAcc *[]string,
	findAll bool,
) bool {
	if x == m.xGoal && y == m.yGoal {
		*pathAcc = append(*pathAcc, path)
		return true
	}

	for _, d := range GetOpenDirs(path) {
		xn, yn := m.Move(x, y, d.dx, d.dy)
		if xn >= 0 {
			if m.depthFirstSearch(
				path+d.path,
				xn,
				yn,
				pathAcc,
				findAll,
			) && !findAll {
				return true
			}
		}
	}

	return false
}

func (m *Maze) DepthFirstSearch(
	key string,
) []string {
	pathAcc := []string{}
	m.depthFirstSearch(key, m.xStart, m.yStart, &pathAcc, true)
	return pathAcc
}

func RunBreadthFirst(key string) {
	maze := &Maze{
		4, 4,
		0, 0,
		3, 3,
	}

	path := maze.BreadthFirstSearch(
		[]Path{
			{
				key,
				maze.xStart,
				maze.yStart,
			},
		},
	)

	fmt.Printf("Breadth First\n")
	if path != nil {
		shortest := path.path[len(key):]
		fmt.Printf("    Shortest: (%v) %s\n", len(shortest), shortest)
	} else {
		fmt.Println("    No valid path.")
	}
}

func RunDepthFirst(key string) {
	maze := &Maze{
		4, 4,
		0, 0,
		3, 3,
	}

	paths := maze.DepthFirstSearch(
		key,
	)

	fmt.Printf("Depth First\n")
	if len(paths) > 0 {
		shortest := paths[0]
		longest := paths[0]
		for _, p := range paths[1:] {
			if len(p) < len(shortest) {
				shortest = p
			}
			if len(p) > len(longest) {
				longest = p
			}
		}
		shortest = shortest[len(key):]
		longest = longest[len(key):]
		fmt.Printf("    Shortest: (%v) %s\n", len(shortest), shortest)
		fmt.Printf("    Longest: (%v) %s\n", len(longest), longest)
	} else {
		fmt.Println("    No valid path.")
	}
}

func Run(key string) {
	RunBreadthFirst(key)
	RunDepthFirst(key)
}

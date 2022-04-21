package shape

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dannyroes/raytrace/data"
)

type OBJDetails struct {
	Ignored   int
	Vertices  []data.Tuple
	Normals   []data.Tuple
	Triangles int
	Polys     int
	Groups    map[string]*GroupType
}

func ParseObj(file string) (res OBJDetails) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	r := &res
	r.Groups = make(map[string]*GroupType)
	currentGroup := "default_group"

	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), " ")
		parts := strings.Split(l, " ")
		switch parts[0] {
		case "v":
			r.AddVertex(parts[1:])
		case "f":
			r.Polys++
			r.AddFace(parts[1:], currentGroup)
		case "g":
			currentGroup = parts[1]
		case "vn":
			r.AddNormal(parts[1:])
		default:
			r.Ignored++
		}
	}

	return res
}

func (o *OBJDetails) AddVertex(loc []string) {
	for loc[0] == "" {
		loc = loc[1:]
	}
	o.Vertices = append(o.Vertices, parsePoint(loc))
}

func (o *OBJDetails) AddFace(i []string, currentGroup string) {
	for i[0] == "" {
		i = i[1:]
	}
	indexes, normals, hasNormals := parseIndexes(i)
	if len(indexes) < 3 {
		panic("Not enough indexes")
	}

	if o.Groups[currentGroup] == nil {
		o.Groups[currentGroup] = Group()
	}

	for x := 1; x < len(indexes)-1; x++ {
		o.Triangles++
		if !hasNormals {
			o.Groups[currentGroup].AddChild(Triangle(o.Vertices[indexes[0]-1], o.Vertices[indexes[x]-1], o.Vertices[indexes[x+1]-1]))
		} else {
			o.Groups[currentGroup].AddChild(SmoothTriangle(o.Vertices[indexes[0]-1], o.Vertices[indexes[x]-1], o.Vertices[indexes[x+1]-1],
				o.Normals[normals[0]-1], o.Normals[normals[x]-1], o.Normals[normals[x+1]-1]))

		}
	}
}

func (o *OBJDetails) GetGroup() *GroupType {
	g := Group()
	for _, group := range o.Groups {
		g.AddChild(group)
	}
	return g
}

func (o *OBJDetails) AddNormal(loc []string) {
	for loc[0] == "" {
		loc = loc[1:]
	}
	o.Normals = append(o.Normals, parseVector(loc))
}

func parseIndexes(raw []string) ([]int, []int, bool) {
	indexes := make([]int, len(raw))
	normals := make([]int, len(raw))
	hasNormals := false

	for i, v := range raw {
		ind := strings.Split(v, "/")
		index, err := strconv.ParseInt(ind[0], 10, 32)
		if err != nil {
			fmt.Printf("%+v\n", raw)
			panic(err)
		}
		indexes[i] = int(index)

		if len(ind) > 1 {
			hasNormals = true
			index, err := strconv.ParseInt(ind[2], 10, 32)
			if err != nil {
				fmt.Printf("%+v\n", raw)
				panic(err)
			}
			normals[i] = int(index)
		}
	}

	return indexes, normals, hasNormals
}

func parsePoint(loc []string) data.Tuple {
	x, y, z := parseValues(loc)
	return data.Point(x, y, z)
}

func parseVector(loc []string) data.Tuple {
	x, y, z := parseValues(loc)
	return data.Vector(x, y, z)
}

func parseValues(loc []string) (float64, float64, float64) {
	x, err := strconv.ParseFloat(loc[0], 64)
	if err != nil {
		fmt.Printf("%+v\n", loc)
		panic(err)
	}

	y, err := strconv.ParseFloat(loc[1], 64)
	if err != nil {
		panic(err)
	}

	z, err := strconv.ParseFloat(loc[2], 64)
	if err != nil {
		panic(err)
	}

	return x, y, z
}

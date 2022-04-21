package shape

import (
	"fmt"
	"testing"

	"github.com/dannyroes/raytrace/data"
)

func TestObjParse(t *testing.T) {
	cases := []struct {
		file     string
		expected OBJDetails
	}{
		{file: "tests/invalid.obj",
			expected: OBJDetails{Ignored: 5, Groups: map[string]*GroupType{}},
		},
		{
			file: "tests/vertices.obj",
			expected: OBJDetails{Vertices: []data.Tuple{
				data.Point(-1, 1, 0),
				data.Point(-1, 0.5, 0),
				data.Point(1, 0, 0),
				data.Point(1, 1, 0),
			}, Groups: map[string]*GroupType{},
			},
		},
		{
			file: "tests/triangles.obj",
			expected: OBJDetails{Ignored: 1, Vertices: []data.Tuple{
				data.Point(-1, 1, 0),
				data.Point(-1, 0, 0),
				data.Point(1, 0, 0),
				data.Point(1, 1, 0),
			},
				Groups: map[string]*GroupType{
					"default_group": {
						Children: []Shape{
							Triangle(data.Point(-1, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0)),
							Triangle(data.Point(-1, 1, 0), data.Point(1, 0, 0), data.Point(1, 1, 0)),
						},
					},
				},
			},
		},
		{
			file: "tests/polygon.obj",
			expected: OBJDetails{Ignored: 1, Vertices: []data.Tuple{
				data.Point(-1, 1, 0),
				data.Point(-1, 0, 0),
				data.Point(1, 0, 0),
				data.Point(1, 1, 0),
				data.Point(0, 2, 0),
			},
				Groups: map[string]*GroupType{
					"default_group": {
						Children: []Shape{
							Triangle(data.Point(-1, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0)),
							Triangle(data.Point(-1, 1, 0), data.Point(1, 0, 0), data.Point(1, 1, 0)),
							Triangle(data.Point(-1, 1, 0), data.Point(1, 1, 0), data.Point(0, 2, 0)),
						},
					},
				},
			},
		},
		{
			file: "tests/groups.obj",
			expected: OBJDetails{Ignored: 1, Vertices: []data.Tuple{
				data.Point(-1, 1, 0),
				data.Point(-1, 0, 0),
				data.Point(1, 0, 0),
				data.Point(1, 1, 0),
			},
				Groups: map[string]*GroupType{
					"FirstGroup": {
						Children: []Shape{
							Triangle(data.Point(-1, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0)),
						},
					},
					"SecondGroup": {
						Children: []Shape{
							Triangle(data.Point(-1, 1, 0), data.Point(1, 0, 0), data.Point(1, 1, 0)),
						},
					},
				},
			},
		},
		{
			file: "tests/normals.obj",
			expected: OBJDetails{Normals: []data.Tuple{
				data.Vector(0, 0, 1),
				data.Vector(0.707, 0, -0.707),
				data.Vector(1, 2, 3),
			}, Groups: map[string]*GroupType{},
			},
		},
		{
			file: "tests/normal_faces.obj",
			expected: OBJDetails{Ignored: 2, Vertices: []data.Tuple{
				data.Point(0, 1, 0),
				data.Point(-1, 0, 0),
				data.Point(1, 0, 0),
			},
				Normals: []data.Tuple{
					data.Vector(-1, 0, 0),
					data.Vector(1, 0, 0),
					data.Vector(0, 1, 0),
				},
				Groups: map[string]*GroupType{
					"default_group": {
						Children: []Shape{
							SmoothTriangle(data.Point(0, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0),
								data.Vector(0, 1, 0), data.Vector(-1, 0, 0), data.Vector(1, 0, 0),
							),
							SmoothTriangle(data.Point(0, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0),
								data.Vector(0, 1, 0), data.Vector(-1, 0, 0), data.Vector(1, 0, 0),
							),
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		res := ParseObj(tc.file)
		if !OBJDetailsEqual(res, tc.expected) {
			t.Errorf("Result mismatch expected %+v received %+v", tc.expected, res)
		}
	}
}

func TestObjGroups(t *testing.T) {
	o := ParseObj("tests/groups.obj")

	group := o.GetGroup()

	expect := o.Groups["FirstGroup"]
	found := false
	for _, g := range group.Children {
		if g == expect {
			found = true
			break
		}
	}

	if !found {
		t.Error("Group did not contain FirstGroup")
	}

	expect = o.Groups["SecondGroup"]
	found = false
	for _, g := range group.Children {
		if g == expect {
			found = true
			break
		}
	}

	if !found {
		t.Error("Group did not contain SecondGroup")
	}
}

func OBJDetailsEqual(a, b OBJDetails) bool {
	if a.Ignored != b.Ignored {
		return false
	}

	if len(a.Vertices) != len(b.Vertices) {
		return false
	}

	for i := range a.Vertices {
		if !data.TupleEqual(a.Vertices[i], b.Vertices[i]) {
			return false
		}
	}

	for i := range a.Normals {
		if !data.TupleEqual(a.Normals[i], b.Normals[i]) {
			return false
		}
	}

	if len(a.Groups) != len(b.Groups) {
		return false
	}

	for name := range a.Groups {
		if b.Groups[name] == nil {
			fmt.Printf("Missing group %s, %+v", name, b.Groups)
			return false
		}

		if len(a.Groups[name].Children) != len(b.Groups[name].Children) {
			return false
		}
		for i := range a.Groups[name].Children {
			if !triangleEqual(a.Groups[name].Children[i].(*TriangleType), b.Groups[name].Children[i].(*TriangleType)) {
				fmt.Println(a.Groups[name].Children[i], b.Groups[name].Children[i])
				return false
			}
		}
	}
	return true
}

func triangleEqual(a, b *TriangleType) bool {
	return data.TupleEqual(a.p1, b.p1) && data.TupleEqual(a.p2, b.p2) && data.TupleEqual(a.p3, b.p3) &&
		data.TupleEqual(a.n1, b.n1) && data.TupleEqual(a.n2, b.n2) && data.TupleEqual(a.n3, b.n3) &&
		a.smooth == b.smooth
}

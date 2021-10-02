package world

import (
	"fmt"
	"os"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
	"github.com/dannyroes/raytrace/shape"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

type SceneCamera struct {
	Width       int
	Height      int
	FieldOfView float64 `mapstructure:"field-of-view"`
	Supersample int
	From        []float64
	To          []float64
	Up          []float64
}

type SceneObject struct {
	Type      string `mapstructure:"add"`
	Material  *SceneMaterial
	Transform [][]interface{}
}

type SceneMaterial struct {
	Colour          *[]float64
	Diffuse         *float64
	Ambient         *float64
	Specular        *float64
	Shininess       *float64
	Reflective      *float64
	Transparency    *float64
	RefractiveIndex *float64 `mapstructure:"refractive-index"`
	Pattern         *ScenePattern
}

type ScenePattern struct {
	Type string
	A    []float64
	B    []float64
}

type SceneLight struct {
	Type      string
	At        []float64
	Intensity []float64
}

func LoadScene(filename string) (*CameraType, WorldType, error) {
	w := World()
	c := &CameraType{}
	definitions := map[string]interface{}{}

	yamlScene, err := os.ReadFile(filename)
	if err != nil {
		return c, w, err
	}
	items := []map[string]interface{}{}

	err = yaml.Unmarshal(yamlScene, &items)
	if err != nil {
		return c, w, err
	}

	for _, item := range items {
		if t, exists := item["add"]; exists {
			item = addDefinitions(item, definitions)
			switch t {
			case "camera":
				c = processCamera(item)
			case "sphere", "cube", "plane":
				w.Objects = append(w.Objects, processObject(item))
			case "light":
				w.Light = processLight(item)
			}
		} else if name, exists := item["define"]; exists {
			var result interface{}

			switch v := item["value"].(type) {
			case map[interface{}]interface{}:
				if ext, exists := item["extend"]; exists {
					for key, value := range definitions[ext.(string)].(map[interface{}]interface{}) {
						if _, exists := v[key]; !exists {
							v[key] = value
						}
					}
				}
				result = v
			case []interface{}:
				result = mergeDefinitions(v, definitions)
			}

			definitions[name.(string)] = result
		}
	}

	return c, w, nil
}

func addDefinitions(item map[string]interface{}, definitions map[string]interface{}) map[string]interface{} {
	for key, val := range item {
		switch v := val.(type) {
		case string:
			if def, exists := definitions[v]; exists {
				item[key] = def
			}
		case []interface{}:
			item[key] = mergeDefinitions(v, definitions)
		}
	}

	return item
}

func mergeDefinitions(item []interface{}, definitions map[string]interface{}) []interface{} {
	type merge struct {
		index int
		val   []interface{}
	}
	merges := []merge{}

	for i, val := range item {
		switch v := val.(type) {
		case string:
			if def, exists := definitions[v]; exists {
				switch d := def.(type) {
				case []interface{}:
					merges = append(merges, merge{i, d})
				default:
					merges = append(merges, merge{i, []interface{}{def}})
				}
			}
		}
	}

	for _, m := range merges {
		end := append(m.val, item[m.index+1:]...)
		item = append(item[:m.index], end...)
	}

	return item
}

func processCamera(item map[string]interface{}) *CameraType {
	var result SceneCamera

	err := mapstructure.Decode(item, &result)
	if err != nil {
		fmt.Println(err)
	}

	c := Camera(result.Width, result.Height, result.FieldOfView)
	c.Transform = data.ViewTransform(sliceToPoint(result.From), sliceToPoint(result.To), sliceToVector(result.Up))

	if result.Supersample > 0 {
		c.Supersample = result.Supersample
	}

	return c
}

func processObject(item map[string]interface{}) shape.Shape {
	var result SceneObject

	err := mapstructure.Decode(item, &result)
	if err != nil {
		fmt.Println(err)
	}

	var obj shape.Shape

	switch result.Type {
	case "sphere":
		obj = shape.Sphere()
	case "cube":
		obj = shape.Cube()
	case "plane":
		obj = shape.Plane()
	}

	mat := material.Material()
	if result.Material != nil {
		mat = processMaterial(*result.Material)
	}

	obj.SetMaterial(mat)

	if len(result.Transform) > 0 {
		obj.SetTransform(processTransform(result.Transform))
	}

	return obj
}

func processMaterial(mat SceneMaterial) material.MaterialType {
	m := material.Material()
	if mat.Colour != nil {
		m.Colour = sliceToColour(*mat.Colour)
	}
	if mat.Diffuse != nil {
		m.Diffuse = *mat.Diffuse
	}
	if mat.Ambient != nil {
		m.Ambient = *mat.Ambient
	}
	if mat.Specular != nil {
		m.Specular = *mat.Specular
	}
	if mat.Shininess != nil {
		m.Shininess = *mat.Shininess
	}
	if mat.Reflective != nil {
		m.Reflective = *mat.Reflective
	}
	if mat.Transparency != nil {
		m.Transparency = *mat.Transparency
	}
	if mat.RefractiveIndex != nil {
		m.RefractiveIndex = *mat.RefractiveIndex
	}

	return m
}

func processLight(item map[string]interface{}) Light {
	var result SceneLight

	err := mapstructure.Decode(item, &result)
	if err != nil {
		fmt.Println(err)
	}

	light := PointLight(sliceToPoint(result.At), sliceToColour(result.Intensity))

	return light
}

func processTransform(item [][]interface{}) data.Matrix {
	transform := data.IdentityMatrix()
	for _, t := range item {
		v := sliceToFloats(t[1:])
		switch t[0].(string) {
		case "scale":
			transform = transform.Scale(v[0], v[1], v[2])
		case "translate":
			transform = transform.Translate(v[0], v[1], v[2])
		case "rotate-x":
			transform = transform.RotateX(v[0])
		case "rotate-y":
			transform = transform.RotateY(v[0])
		case "rotate-z":
			transform = transform.RotateZ(v[0])
		case "shear":
			transform = transform.Shear(v[0], v[1], v[2], v[3], v[4], v[5])
		}

	}

	return transform
}

func sliceToPoint(s []float64) data.Tuple {
	return data.Point(s[0], s[1], s[2])
}

func sliceToVector(s []float64) data.Tuple {
	return data.Vector(s[0], s[1], s[2])
}

func sliceToColour(s []float64) material.ColourTuple {
	return material.Colour(s[0], s[1], s[2])
}

func sliceToFloats(s []interface{}) []float64 {
	result := make([]float64, len(s))

	for i, val := range s {
		switch v := val.(type) {
		case int:
			result[i] = float64(v)
		case float64:
			result[i] = v
		}
	}
	return result
}

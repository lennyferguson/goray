package src

import (
	"math"
	"sync"
	"time"
)

type (
	Pixel struct {
		X     int
		Y     int
		Color Vec4
	}
	Image struct {
		data []Pixel
	}

	// Ray is a type that contains a Src and Dir vec types
	Ray struct {
		Src Vec
		Dir Vec
	}

	//Material describes the components of an object computed using a BDRF function
	Material struct {
		Ambient    Vec4
		Specular   Vec4
		Diffuse    Vec4
		Reflective float32
	}

	//Camera contains the components necessary to describ the properties of the camera / view
	Camera struct {
		Eye         Vec3
		LookAt      Vec3
		AspectRatio float32
		Width       int
		Height      int
	}

	//Object describes the various renderable entities in a scene
	Object interface {
		Hit(Ray) (*Object, float32)
		CalculateColor(float32, Ray, *Scene) Vec4
	}

	/*LightSource types are objects that generate light
	The Light function returns a Vec4 of the form
	[R,G,B,Brightness]
	*/
	LightSource interface {
		Light() Vec4
		Sample() Vec3
	}

	//Renderer describes a set of functions that an entity can be invoked with to
	//update and render an image
	Renderer interface {
		Update(float64)
		Render() Image
	}

	//Scene is a type
	Scene struct {
		Camera       Camera
		Objects      []Object
		lightSources *[]LightSource
	}

	Sphere struct {
		Origin Vec3
		Radius float32
		Material
	}

	Rectangle struct {
		Planes [6]Plane
	}

	Plane struct {
		Verts [4]Vec3
		Material
	}

	RenderPoperties struct {
		frames int
	}

	Hit struct {
		Point  Vec3
		Normal Vec3
		Ray
	}
)

var BkgColor = NewVec4(0.0, 0.0, 0.0, 0.0)
var Red = Material{}
var Blue = Material{}
var Green = Material{}
var Mirror = Material{}
var Floor = Material{}

//CalculateColor calculates the color from a meterial hiting a particular hit location
//that contains the information necessary to calculate the color using a BDRF function
func (m Material) CalculateColor(h Hit, s Scene) Vec4 {
	panic("not implemented")
}

func (s *Scene) Update(dt float64) {
	panic("not implemented")
}

func (s Scene) LightSources() []LightSource {
	if s.lightSources == nil {
		ans := []LightSource{}
		for _, v := range s.Objects {
			if object, is := v.(LightSource); is {
				ans = append(ans, object)
			}
		}
		s.lightSources = &ans
	}
	return *s.lightSources
}

func (s *Scene) Render() Image {
	size := s.Camera.Width & s.Camera.Height
	ans := Image{
		data: make([]Pixel, size)[:],
	}
	width := s.Camera.Width
	var wg sync.WaitGroup
	for y := 0; y < s.Camera.Height; y++ {
		for x := 0; x < s.Camera.Width; x++ {
			go func(lx int, ly int) {
				wg.Add(1)
				defer func() {
					wg.Done()
				}()
				ray := s.Camera.CalculateViewRay(lx, ly)
				pixel := Pixel{
					X: lx,
					Y: ly,
				}
				nearest := float32(math.MaxFloat32)
				var hit *Object
				for _, object := range s.Objects {
					if _, isLightsource := object.(LightSource); !isLightsource {
						if o, t := object.Hit(ray); t < nearest {
							nearest = t
							hit = o
						}
					}
				}
				if hit != nil {
					pixel.Color = (*hit).CalculateColor(nearest, ray, s)
				} else {
					pixel.Color = BkgColor
				}
				ans.data[ly*width+lx] = pixel
			}(x, y)
		}
	}
	// Wait for each goroutine calculate the pixel color to finish
	wg.Wait()
	return ans
}

func (r Rectangle) Hit(ray Ray) (*Object, float32) {
	for _, plane := range r.Planes {
		if p, t := plane.Hit(ray); p != nil {
			return p, t
		}
	}
	return nil, 0.0
}

func (c Camera) CalculateViewRay(x int, y int) (r Ray) {
	panic("not implemented")
}
func (r Rectangle) CalculateColor(t float32, ray Ray) Vec4 {
	return NewVec4(0.0, 0.0, 0.0, 0.0)
}

func (p Plane) Hit(r Ray) (*Object, float32) {
	panic("not implemented")
}

func (p Plane) CalculateColor(t float32, r Ray, scene *Scene) Vec4 {
	panic("not implemented")
}

func (s Sphere) Hit(r Ray) (*Object, float32) {
	panic("not implemented")
}

func (s Sphere) CalculateColor(t float32, r Ray, scene *Scene) Vec4 {
	panic("not implemented")
}

func (r Ray) Reflect(h Hit) Ray {
	panic("not implemented")
}

func Render(r Renderer, p RenderPoperties) []Image {
	ans := make([]Image, p.frames)
	t := time.Now()
	for i := 0; i < p.frames; i++ {
		current := time.Now()
		dt := current.Sub(t).Seconds()
		r.Update(dt)
		ans[i] = r.Render()
		t = current
	}
	return ans
}

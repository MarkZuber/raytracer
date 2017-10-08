package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/MarkZuber/raytrace"
	"github.com/MarkZuber/raytrace/materials"
	"github.com/MarkZuber/raytrace/shapes"
)

func loadPng(filePath string) (*image.RGBA, error) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	pngImage, err := png.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	pngRect := pngImage.Bounds()

	rgbaImage := image.NewRGBA(image.Rect(0, 0, pngRect.Dx(), pngRect.Dy()))
	draw.Draw(rgbaImage, rgbaImage.Bounds(), pngImage, image.ZP, draw.Over)

	return rgbaImage, nil
}

func loadJpg(filePath string) (*image.RGBA, error) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	pngImage, err := jpeg.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	pngRect := pngImage.Bounds()

	rgbaImage := image.NewRGBA(image.Rect(0, 0, pngRect.Dx(), pngRect.Dy()))
	draw.Draw(rgbaImage, rgbaImage.Bounds(), pngImage, image.ZP, draw.Over)

	return rgbaImage, nil
}

func createTestScene() *raytrace.Scene {
	scene := raytrace.CreateScene()
	scene.SetCamera(raytrace.CreateCamera(
		raytrace.CreateVector(10, -.2, 7),
		raytrace.CreateVector(-1, .3, -1),
		raytrace.CreateVector(0, 0, 1))) // z is up

	// setup the chessboard floor
	scene.AddShape(shapes.CreatePlaneShape(
		raytrace.CreateVector(0, 0, -1),
		materials.CreateChessboardMaterial(
			raytrace.CreateDoubleColor(1, 1, 1),
			raytrace.CreateDoubleColor(0, 0, 0),
			0.7,
			1.0,
			0.0,
			0.2),
		1.2))

	// create a back wall

	scene.AddShape(shapes.CreatePlaneShape(
		raytrace.CreateVector(-1, 0, 0),
		materials.CreateChessboardMaterial(
			raytrace.CreateDoubleColor(1, 1, 1),
			raytrace.CreateDoubleColor(0, 0, 0),
			0.7,
			1.0,
			0.0,
			0.2),
		1.2))

	marbleImage, err := loadPng("/home/mzuber/go/src/github.com/MarkZuber/raytracer/marble1.png")
	if err != nil {
		os.Exit(2)
	}

	scene.AddShape(shapes.CreateSphereShape(
		raytrace.CreateVector(0, 0, 3),
		materials.CreateSolidMaterial(
			raytrace.CreateDoubleColor(0.85, 0.0, 0.0),
			0.1,
			0.0,
			0.2,
			0.0),
		1.5))

	scene.AddShape(shapes.CreateSphereShape(
		raytrace.CreateVector(2, 6, 4),
		materials.CreateTextureMaterial(
			marbleImage,
			1,
			0.2,
			0.0,
			0.2,
			0.5),
		1.6))

	scene.AddShape(shapes.CreateSphereShape(
		raytrace.CreateVector(0, -2, 3),
		materials.CreateTextureMaterial(
			marbleImage,
			1,
			0.2,
			0.0,
			0.2,
			0.5),
		1.4)) // todo: weird bug where can't go much smaller in radius or the object disappears

	/*
		signImage, err := loadPng("/Users/zube/go/src/github.com/MarkZuber/raytracer/stop_sign_page.png")
		if err != nil {
			os.Exit(2)
		}
	*/

	// boxImage, err := loadJpg("/Users/zube/go/src/github.com/MarkZuber/raytracer/rainbowcolor.jpg")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(2)
	// }

	// scene.AddShape(shapes.CreateBoxShape(
	// 	materials.CreateTextureMaterial(
	// 		boxImage,
	// 		1,
	// 		0.2,
	// 		0.0,
	// 		0.2,
	// 		0.5),
	// 	raytrace.CreateVector(1, 4, 1),
	// 	raytrace.CreateVector(3, 2.0, 3)))

	scene.AddLight(raytrace.CreateLight(
		raytrace.CreateVector(5, 5, 10),
		raytrace.CreateDoubleColor(.5, .5, .5)))

	/*
		scene.AddLight(raytrace.CreateLight(
			raytrace.CreateVector(15, -10, 10),
			raytrace.CreateDoubleColor(.5, .5, .5)))
	*/
	/*
		scene.AddLight(raytrace.CreateLight(
			raytrace.CreateVector(1, 4, 10),
			raytrace.CreateDoubleColor(.9, .9, .9)))
	*/

	return scene
}

func createBallScene() *raytrace.Scene {
	scene := raytrace.CreateScene()
	scene.SetCamera(raytrace.CreateCamera(
		raytrace.CreateVector(0, 0, -15),
		raytrace.CreateVector(-0.2, 0, 5),
		raytrace.CreateVector(0, 1, 0)))

	// setup a solid reflecting sphere
	scene.AddShape(shapes.CreateSphereShape(
		raytrace.CreateVector(-1.5, 0.5, 0),
		materials.CreateSolidMaterial(
			raytrace.CreateDoubleColor(0.5, 0.0, 0.0),
			0.1,
			0.0,
			0.2,
			0.0),
		0.5))

	marbleImage, err := loadPng("/home/mzuber/go/src/github.com/MarkZuber/raytracer/marble1.png")
	if err != nil {
		os.Exit(2)
	}

	// signImage, err := loadPng("/Users/zube/go/src/github.com/MarkZuber/raytracer/stop_sign_page.png")
	// if err != nil {
	// 	os.Exit(2)
	// }

	scene.AddShape(shapes.CreateSphereShape(
		raytrace.CreateVector(0, 0, 0),
		materials.CreateTextureMaterial(
			marbleImage,
			1,
			0.2,
			0.0,
			0.2,
			0.5),
		0.9))

	// scene.AddShape(shapes.CreateBoxShape(
	// 	materials.CreateTextureMaterial(
	// 		signImage,
	// 		1,
	// 		0.2,
	// 		0.0,
	// 		0.2,
	// 		0.5),
	// 	raytrace.CreateVector(-3, 4, 0),
	// 	raytrace.CreateVector(-1, 2.0, 4)))

	// setup the chessboard floor
	scene.AddShape(shapes.CreatePlaneShape(
		raytrace.CreateVector(0.1, 0.9, -0.5).Normalize(),
		materials.CreateChessboardMaterial(
			raytrace.CreateDoubleColor(1, 1, 1),
			raytrace.CreateDoubleColor(0, 0, 0),
			0.7,
			1.0,
			0.0,
			0.2),
		1.2))

	//add two lights for better lighting effects
	scene.AddLight(raytrace.CreateLight(
		raytrace.CreateVector(5, 10, -1),
		raytrace.CreateDoubleColor(0.8, 0.8, 0.8)))
	scene.AddLight(raytrace.CreateLight(
		raytrace.CreateVector(-3, 5, -15),
		raytrace.CreateDoubleColor(0.8, 0.8, 0.8)))

	return scene
}

func savePng(filePath string, img image.Image) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// initialize parameters
	outPath := "/home/mzuber/raytrace.png"
	sizeRect := image.Rect(0, 0, 1000, 1000)
	// scene := createBallScene()
	scene := createTestScene()

	// generate the image via the raytracer
	img := image.NewRGBA(sizeRect)
	rt := raytrace.CreateRayTracer(img.Rect, scene)
	rt.SimpleRender(img)

	// save image to disk
	err := savePng(outPath, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("Finished image is at: %s\n", outPath)
}

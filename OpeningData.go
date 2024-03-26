package main

import rl "github.com/gen2brain/raylib-go/raylib"

func MakeOpeningData() Scene {
	var testScene Scene

	var dotGroup1 DotGroup
	dotGroup1.Dots = MakeRectangularGroupOfCircles(0.5, 0.5, 0.05, 0.05, 1, 1, rl.White, 0.02)
	dotGroup1.StartTime = 0.0
	dotGroup1.EndTime = 1.0
	testScene.Dots = append(testScene.Dots, dotGroup1)

	var dotGroup2 DotGroup
	dotGroup2.Dots = MakeRectangularGroupOfCircles(0.3, 0.3, 0.02, 0.02, 10, 10, rl.Red, 0.2)
	dotGroup2.StartTime = 2.0
	dotGroup2.EndTime = 4.0
	testScene.Dots = append(testScene.Dots, dotGroup2)

	var dotGroup3 DotGroup
	dotGroup3.Dots = MakeRectangularGroupOfCircles(0.4, 0.4, 0.0, 0.0, 4, 4, rl.Green, 0.2)
	dotGroup3.StartTime = 5.0
	dotGroup3.EndTime = 10.0
	testScene.Dots = append(testScene.Dots, dotGroup3)
	return testScene

}

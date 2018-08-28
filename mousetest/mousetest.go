package mousetest

import (
"fmt"
"github.com/go-vgo/robotgo"
)

func Screen() {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
	color := robotgo.GetPixelColor(100, 200)
	fmt.Println("color----", color)
}
func Mouse() {
	//robotgo.ScrollMouse(3, "up")
	//robotgo.MouseClick("left", true)
	robotgo.MoveMouseSmooth(540, 960, 1.0, 5.0)
}
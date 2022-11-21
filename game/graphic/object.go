package graphic

import (
	"time"

	"github.com/nsf/termbox-go"
)

type Coordinates struct{ X, Y int }

type DrawableAtom struct {
	Char rune
	Fg   termbox.Attribute
	Bg   termbox.Attribute
}

type DrawableObject struct {
	Atoms [][]DrawableAtom
}

func (obj DrawableObject) width() int {
	return len(obj.Atoms)
}

func (obj DrawableObject) DrawObject(coord Coordinates) {
	for ly := 0; ly < obj.width(); ly++ {
		for lx := 0; lx < len(obj.Atoms[ly]); lx++ {
			atom := obj.Atoms[ly][lx]
			termbox.SetCell(coord.X+lx, coord.Y+ly, atom.Char, atom.Fg, atom.Bg)
		}
	}
}

func (obj DrawableObject) ClearObject(coord Coordinates) {
	for ly := 0; ly < obj.width(); ly++ {
		for lx := 0; lx < len(obj.Atoms[ly]); lx++ {
			termbox.SetCell(coord.X+lx, coord.Y+ly, ' ', termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func (obj DrawableObject) MoveObjectVertically(origin Coordinates, steps int, ch chan bool) {
	ticker := time.NewTicker(time.Millisecond * 50)
	defer ticker.Stop()

	stepsTaken := 0

	for {
		select {
		case <-ch:
			return
		default:
			if stepsTaken == 0 {
				obj.DrawObject(origin)
				stepsTaken += 1
			} else if origin.Y+stepsTaken <= steps {
				lastCoordinates := Coordinates{X: origin.X, Y: origin.Y + stepsTaken - 1}
				newCoordinates := Coordinates{X: origin.X, Y: origin.Y + stepsTaken}
				obj.ClearObject(lastCoordinates)
				obj.DrawObject(newCoordinates)
				stepsTaken += 1
			} else if origin.Y+stepsTaken > steps {
				return
			}
		}
		termbox.Flush()
		<-ticker.C
	}
}

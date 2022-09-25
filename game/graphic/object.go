package graphic

import (
	"time"

	"github.com/nsf/termbox-go"
)

type Cordinates struct{ X, Y int }

type DrawableAtom struct {
	Char rune
	Fg   termbox.Attribute
	Bg   termbox.Attribute
}

type DrawableObject struct {
	Atoms    [][]DrawableAtom
	isMoving bool
}

func (obj DrawableObject) width() int {
	return len(obj.Atoms)
}

func (obj DrawableObject) IsMoving() bool {
	return obj.isMoving
}

func (obj DrawableObject) DrawObject(coord Cordinates) {
	for ly := 0; ly < obj.width(); ly++ {
		for lx := 0; lx < len(obj.Atoms[ly]); lx++ {
			atom := obj.Atoms[ly][lx]
			termbox.SetCell(coord.X+lx, coord.Y+ly, atom.Char, atom.Fg, atom.Bg)
		}
	}
}

func (obj DrawableObject) ClearObject(coord Cordinates) {
	for ly := 0; ly < obj.width(); ly++ {
		for lx := 0; lx < len(obj.Atoms[ly]); lx++ {
			termbox.SetCell(coord.X+lx, coord.Y+ly, ' ', termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

func (obj DrawableObject) MoveObjectVertically(origin Cordinates, steps int, ch chan bool) {
	obj.isMoving = true
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
				lastCoordinates := Cordinates{X: origin.X, Y: origin.Y + stepsTaken - 1}
				newCoordinates := Cordinates{X: origin.X, Y: origin.Y + stepsTaken}
				obj.ClearObject(lastCoordinates)
				obj.DrawObject(newCoordinates)
				stepsTaken += 1

			} else if origin.Y+stepsTaken > steps {
				obj.isMoving = false
				return
			}
		}
		termbox.Flush()
		<-ticker.C
	}
}

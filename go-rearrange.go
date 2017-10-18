// Package rearrange (gorearrange.go) :
// This is a Golang library to interactively rearrange a text data by users.
package rearrange

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

// reaData : Main data for rearranging.
type reaData struct {
	X                    int
	Y                    int
	FTop                 bool
	FBot                 bool
	Row                  int
	SelectedValue        string
	SelectedValueHistory []valuesHistory
	Data                 []string
	DataWithInf          []dataWithInfEle
	Pos                  int
	CPos                 int
	Width                int
	Height               int
	Pgupdn               int
	Onflag               bool
	SelectMode           bool
	IndexMode            bool
}

// valuesHistory : History data.
type valuesHistory struct {
	Index int
	Value string
}

type dataWithInfEle struct {
	Index   int
	DataEle string
}

// initRearrange : Set initial values.
func initRearrange(data []string, pgupdn int, selectmode bool, index bool) *reaData {
	r := &reaData{}
	r.Data = data
	r.Pgupdn = pgupdn
	r.Width, r.Height = termbox.Size()
	r.FTop = true
	r.FBot = false
	r.X = 0
	r.Y = 0
	r.Row = 0
	r.Onflag = false
	r.SelectMode = selectmode
	r.IndexMode = index
	r.inDat(data)
	return r
}

// inDat : Import data to DataWithInf.
func (r *reaData) inDat(data []string) {
	for i, e := range data {
		d := &dataWithInfEle{
			Index:   i,
			DataEle: e,
		}
		r.DataWithInf = append(r.DataWithInf, *d)
	}
}

// drawDef : Drawing data to buffer with default colors.
func (r *reaData) drawDef(x, y int, str []rune, color, backgroundColor termbox.Attribute) {
	for _, e := range str {
		termbox.SetCell(x, y, e, color, backgroundColor)
		x += runewidth.RuneWidth(e)
	}
}

// drawC : Drawing data to buffer with colors.
func (r *reaData) drawC(x, y int, str []rune, color, backgroundColor termbox.Attribute) {
	for i := 0; i < len(str); i += 1 {
		termbox.SetCell(x+i, y, str[i], color, backgroundColor)
	}
}

// drawLineDef : Set one line to buffer as a default color. This is used for static data.
func (r *reaData) drawLineDef(x, y int, str []rune) {
	r.drawDef(x, y, str, termbox.ColorDefault, termbox.ColorDefault)
}

// drawLineMove : Set one line with a special color to buffer. This is used for dynamic data.
func (r *reaData) drawLineMove(x, y int, str []rune) {
	r.drawC(x, y, str, termbox.ColorBlack, termbox.ColorWhite)
}

// drawLineSelect : Set one line with a special color to buffer. This is used for selected data.
func (r *reaData) drawLineSelect(x, y int, str []rune) {
	r.drawC(x, y, str, termbox.ColorBlack, termbox.ColorCyan)
}

// inBufD : Display data as a default color. This is used for static data.
func (r *reaData) inBufD(val []string) {
	for i, e := range val {
		r.drawLineDef(0, i, []rune(e))
	}
}

// getDatFromBuf : Get data from buffer
func (r *reaData) getDatFromBuf(x, y int) []rune {
	tex0 := []rune{}
	for i := 0; i < r.Width; i += 1 {
		tex0 = append(tex0, termbox.CellBuffer()[(r.Width*y)+i].Ch)
	}
	return tex0
}

// inBufC : Display data as a special color. This is used for dynamic data.
func (r *reaData) inBufC(x, y int) {
	r.drawLineMove(x, y, r.getDatFromBuf(x, y))
}

// inBufS : Display data as a special color. This is used for selected data.
func (r *reaData) inBufS(x, y int) {
	r.drawLineSelect(x, y, r.getDatFromBuf(x, y))
}

// moveEleArb : Create data for moving cursor.
func (r *reaData) moveEleArb() {
	temp := r.Data[r.CPos]
	r.Data = append(r.Data[:r.CPos], r.Data[r.CPos+1:len(r.Data)]...)
	ar1 := r.Data[:r.Pos]
	ar2 := r.Data[r.Pos:]
	ar2 = append(ar2[:1], ar2[0:]...)
	ar2[0] = temp
	r.Data = append(ar1, ar2...)
	r.moveEleArbBase()
}

// moveEleArbBase : Create data for moving cursor. (json)
func (r *reaData) moveEleArbBase() {
	temp := r.DataWithInf[r.CPos]
	r.DataWithInf = append(r.DataWithInf[:r.CPos], r.DataWithInf[r.CPos+1:len(r.DataWithInf)]...)
	ar1 := r.DataWithInf[:r.Pos]
	ar2 := r.DataWithInf[r.Pos:]
	ar2 = append(ar2[:1], ar2[0:]...)
	ar2[0] = temp
	r.DataWithInf = append(ar1, ar2...)
}

// firstDraw : Draw data. This is used for the first time after it launches this.
func (r *reaData) firstDraw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	r.inBufD(r.Data)
	r.inBufC(0, 0)
	termbox.SetCursor(0, 0)
	termbox.Flush()
}

// moveCursorUp : Move cursor to upper direction.
func (r *reaData) moveCursorUp(mv int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if r.Height > len(r.Data) {
		if r.Onflag && !r.SelectMode {
			r.CPos = r.Y
			if r.Y-mv < 0 {
				r.Y = 0
			} else {
				r.Y -= mv
			}
			r.Pos = r.Y
			r.moveEleArb()
			r.inBufD(r.Data)
			r.inBufS(r.X, r.Y)
		} else {
			if r.SelectMode {
				r.Onflag = false
			}
			if r.Y-mv < 0 {
				r.Y = 0
			} else {
				r.Y -= mv
			}
			r.inBufD(r.Data)
			r.inBufC(r.X, r.Y)
		}
	} else {
		if r.Onflag && !r.SelectMode {
			r.CPos = r.Row + r.Y
			switch {
			case (r.Y - mv) > 0:
				r.Y -= mv
				r.Pos -= mv
				r.FTop = false
				r.FBot = false
			case (r.Y - mv) == 0:
				r.Y -= mv
				r.Pos -= mv
				r.FTop = true
				r.FBot = false
			case (r.Y - mv) < 0:
				if !r.FTop {
					if (r.Row - mv) > 0 {
						r.Row -= mv - r.Y
						r.Pos -= mv
						r.Y = 0
						r.FTop = true
						r.FBot = false
					} else {
						r.Y = 0
						r.Row = 0
						r.Pos = 0
						r.FTop = true
						r.FBot = false
					}
				} else {
					if (r.Row - mv) > 0 {
						r.Row -= mv
						r.Pos -= mv
					} else {
						r.Row = 0
						r.Pos = 0
					}
				}
			}
			r.moveEleArb()
			r.inBufD(r.Data[r.Row : r.Row+r.Height])
			r.inBufS(r.X, r.Y)
		} else {
			if r.SelectMode {
				r.Onflag = false
			}
			switch {
			case (r.Y - mv) > 0:
				r.Y -= mv
				r.FTop = false
				r.FBot = false
			case (r.Y - mv) == 0:
				r.Y -= mv
				r.FTop = true
				r.FBot = false
			case (r.Y - mv) < 0:
				if !r.FTop {
					if (r.Row - mv) > 0 {
						r.Row -= mv - r.Y
						r.Y = 0
						r.FTop = true
						r.FBot = false
					} else {
						r.Y = 0
						r.Row = 0
						r.FTop = true
						r.FBot = false
					}
				} else {
					if (r.Row - mv) > 0 {
						r.Row -= mv
					} else {
						r.Row = 0
					}
				}
			}
			r.inBufD(r.Data[r.Row : r.Row+r.Height])
			r.inBufC(r.X, r.Y)
		}
	}
	termbox.Flush()
}

// moveCursorDn : Move cursor to lower direction.
func (r *reaData) moveCursorDn(mv int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if r.Height > len(r.Data) {
		if r.Onflag && !r.SelectMode {
			r.CPos = r.Y
			if r.Y+mv > len(r.Data)-1 {
				r.Y = len(r.Data) - 1
			} else {
				r.Y += mv
			}
			r.Pos = r.Y
			r.moveEleArb()
			r.inBufD(r.Data)
			r.inBufS(r.X, r.Y)
		} else {
			if r.SelectMode {
				r.Onflag = false
			}
			if r.Y+mv > len(r.Data)-1 {
				r.Y = len(r.Data) - 1
			} else {
				r.Y += mv
			}
			r.inBufD(r.Data)
			r.inBufC(r.X, r.Y)
		}
	} else {
		if r.Onflag && !r.SelectMode {
			r.CPos = r.Row + r.Y
			switch {
			case (r.Y + mv) < r.Height-1:
				r.Y += mv
				r.Pos += mv
				r.FTop = false
				r.FBot = false
			case (r.Y + mv) == r.Height-1:
				r.Y += mv
				r.Pos += mv
				r.FTop = false
				r.FBot = true
			case (r.Y + mv) > r.Height-1:
				if !r.FBot {
					if (r.Row + mv) < (len(r.Data) - r.Height) {
						r.Row += (mv - ((r.Height - 1) - r.Y))
						r.Pos += mv
						r.Y = r.Height - 1
						r.FTop = false
						r.FBot = true
					} else {
						r.Y = r.Height - 1
						r.Row = len(r.Data) - r.Height
						r.Pos = len(r.Data) - 1
						r.FTop = false
						r.FBot = true
					}
				} else {
					if (r.Row + mv) < (len(r.Data) - r.Height) {
						r.Row += mv
						r.Pos += mv
					} else {
						r.Row = len(r.Data) - r.Height
						r.Pos = len(r.Data) - 1
					}
				}
			}
			r.moveEleArb()
			r.inBufD(r.Data[r.Row : r.Row+r.Height])
			r.inBufS(r.X, r.Y)
		} else {
			if r.SelectMode {
				r.Onflag = false
			}
			switch {
			case (r.Y + mv) < r.Height-1:
				r.Y += mv
				r.FTop = false
				r.FBot = false
			case (r.Y + mv) == r.Height-1:
				r.Y += mv
				r.FTop = false
				r.FBot = true
			case (r.Y + mv) > r.Height-1:
				if !r.FBot {
					if (r.Row + r.Y + mv) < len(r.Data)-1 {
						r.Row += (mv - ((r.Height - 1) - r.Y))
						r.Y = r.Height - 1
						r.FTop = false
						r.FBot = true
					} else {
						r.Y = r.Height - 1
						r.Row = len(r.Data) - r.Height
						r.FTop = false
						r.FBot = true
					}
				} else {
					if (r.Row + r.Y + mv) < len(r.Data)-1 {
						r.Row += mv
					} else {
						r.Row = len(r.Data) - r.Height
					}
				}
			}
			r.inBufD(r.Data[r.Row : r.Row+r.Height])
			r.inBufC(r.X, r.Y)
		}
	}
	termbox.Flush()
}

// grabData : Grab data by enter key.
func (r *reaData) grabData() {
	if r.Onflag {
		r.Onflag = false
		r.SelectedValue = ""
		r.Pos = 0
		r.inBufC(r.X, r.Y)
	} else {
		r.Onflag = true
		r.Pos = r.Row + r.Y
		r.SelectedValue = r.Data[r.Pos]
		vh := &valuesHistory{}
		vh.Index = r.Pos
		vh.Value = r.SelectedValue
		r.SelectedValueHistory = append(r.SelectedValueHistory, *vh)
		r.inBufS(r.X, r.Y)
	}
	termbox.Flush()
}

// resetDat : Reset rearranged data.
func (r *reaData) resetDat(backupDat []string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	baseDat := make([]string, len(backupDat))
	r.Data = baseDat
	_ = copy(r.Data, backupDat)
	if r.Height > len(r.Data) {
		r.inBufD(r.Data)
	} else {
		r.inBufD(r.Data[r.Row : r.Row+r.Height])
	}
	r.inBufC(r.X, r.Y)
	if r.Onflag {
		r.Onflag = false
		r.SelectedValue = ""
		r.Pos = 0
	}
	var temp reaData
	r.DataWithInf = temp.DataWithInf
	r.inDat(baseDat)
	termbox.Flush()
}

// setResult : When data index is output, this is used.
func (r *reaData) setResult(backupDat []string) *reaData {
	if r.IndexMode {
		var temp []string
		for _, d := range r.DataWithInf {
			temp = append(temp, strconv.Itoa(d.Index))
		}
		r.Data = temp
	}
	return r
}

// output : Output results.
func (r *reaData) output() ([]string, []valuesHistory, error) {
	return r.Data, r.SelectedValueHistory, nil
}

// rearrange : Main method of this library.
func (r *reaData) rearrange() *reaData {
	backupDat := make([]string, len(r.Data))
	_ = copy(backupDat, r.Data)
	r.firstDraw()
	termbox.HideCursor()
	q := make(chan termbox.Event, 1)
	defer close(q)
	go func() {
		for {
			q <- termbox.PollEvent()
		}
	}()
	for {
		e := <-q
		switch e.Type {
		case termbox.EventKey:
			switch e.Key {
			case termbox.KeyCtrlC, termbox.KeyEsc:
				return r.setResult(backupDat)
			case termbox.KeyHome:
				r.moveCursorUp(r.Y + r.Row)
			case termbox.KeyEnd:
				r.moveCursorDn(len(r.Data) - (r.Y + r.Row))
			case termbox.KeyPgup:
				r.moveCursorUp(r.Pgupdn)
			case termbox.KeyPgdn:
				r.moveCursorDn(r.Pgupdn)
			case termbox.KeyArrowUp:
				r.moveCursorUp(1)
			case termbox.KeyArrowDown:
				r.moveCursorDn(1)
			case termbox.KeyEnter:
				r.grabData()
			case termbox.KeySpace, termbox.KeyBackspace:
				r.resetDat(backupDat)
			}
		}
	}
}

// Do : Method called from users.
func Do(data []string, pgupdn int, selectmode bool, index bool) ([]string, []valuesHistory, error) {
	defer termbox.Close()
	if len(data) == 0 || reflect.TypeOf(data).String() != "[]string" {
		return nil, nil, errors.New("Error: No data or wrong data.")
	}
	if err := termbox.Init(); err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Error: %v\n", err))
	}
	return initRearrange(data, pgupdn, selectmode, index).rearrange().output()
}

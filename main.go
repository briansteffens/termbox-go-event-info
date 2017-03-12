package main

import (
	"fmt"
	"os"
	"github.com/nsf/termbox-go"
)

func eventTypeToString(t termbox.EventType) string {
	switch t {
	case termbox.EventKey:
		return "EventKey"
	case termbox.EventResize:
		return "EventResize"
	case termbox.EventMouse:
		return "EventMouse"
	case termbox.EventError:
		return "EventError"
	case termbox.EventInterrupt:
		return "EventInterrupt"
	case termbox.EventRaw:
		return "EventRaw"
	case termbox.EventNone:
		return "EventNone"
	default:
		panic("Unrecognized termbox.EventType value")
	}
}

func keyToString(k termbox.Key) string {
	switch k {
	case termbox.KeyF1:
		return "KeyF1"
	case termbox.KeyF2:
		return "KeyF2"
	case termbox.KeyF3:
		return "KeyF3"
	case termbox.KeyF4:
		return "KeyF4"
	case termbox.KeyF5:
		return "KeyF5"
	case termbox.KeyF6:
		return "KeyF6"
	case termbox.KeyF7:
		return "KeyF7"
	case termbox.KeyF8:
		return "KeyF8"
	case termbox.KeyF9:
		return "KeyF9"
	case termbox.KeyF10:
		return "KeyF10"
	case termbox.KeyF11:
		return "KeyF11"
	case termbox.KeyF12:
		return "KeyF12"
	case termbox.KeyInsert:
		return "KeyInsert"
	case termbox.KeyDelete:
		return "KeyDelete"
	case termbox.KeyHome:
		return "KeyHome"
	case termbox.KeyEnd:
		return "KeyEnd"
	case termbox.KeyPgup:
		return "KeyPgup"
	case termbox.KeyPgdn:
		return "KeyPgdn"
	case termbox.KeyArrowUp:
		return "KeyArrowUp"
	case termbox.KeyArrowDown:
		return "KeyArrowDown"
	case termbox.KeyArrowLeft:
		return "KeyArrowLeft"
	case termbox.KeyArrowRight:
		return "KeyArrowRight"
	case termbox.MouseLeft:
		return "MouseLeft"
	case termbox.MouseMiddle:
		return "MouseMiddle"
	case termbox.MouseRight:
		return "MouseRight"
	case termbox.MouseRelease:
		return "MouseRelease"
	case termbox.MouseWheelUp:
		return "MouseWheelUp"
	case termbox.MouseWheelDown:
		return "MouseWheelDown"
	case termbox.KeyCtrlTilde:
		return "KeyCtrlTilde | KeyCtrl2 | KeyCtrlSpace"
	case termbox.KeyCtrlA:
		return "KeyCtrlA"
	case termbox.KeyCtrlB:
		return "KeyCtrlB"
	case termbox.KeyCtrlC:
		return "KeyCtrlC"
	case termbox.KeyCtrlD:
		return "KeyCtrlD"
	case termbox.KeyCtrlE:
		return "KeyCtrlE"
	case termbox.KeyCtrlF:
		return "KeyCtrlF"
	case termbox.KeyCtrlG:
		return "KeyCtrlG"
	case termbox.KeyBackspace:
		return "KeyBackspace | KeyCtrlH"
	case termbox.KeyTab:
		return "KeyTab | KeyCtrlI"
	case termbox.KeyCtrlJ:
		return "KeyCtrlJ"
	case termbox.KeyCtrlK:
		return "KeyCtrlK"
	case termbox.KeyCtrlL:
		return "KeyCtrlL"
	case termbox.KeyEnter:
		return "KeyEnter | KeyCtrlM"
	case termbox.KeyCtrlN:
		return "KeyCtrlN"
	case termbox.KeyCtrlO:
		return "KeyCtrlO"
	case termbox.KeyCtrlP:
		return "KeyCtrlP"
	case termbox.KeyCtrlQ:
		return "KeyCtrlQ"
	case termbox.KeyCtrlR:
		return "KeyCtrlR"
	case termbox.KeyCtrlS:
		return "KeyCtrlS"
	case termbox.KeyCtrlT:
		return "KeyCtrlT"
	case termbox.KeyCtrlU:
		return "KeyCtrlU"
	case termbox.KeyCtrlV:
		return "KeyCtrlV"
	case termbox.KeyCtrlW:
		return "KeyCtrlW"
	case termbox.KeyCtrlX:
		return "KeyCtrlX"
	case termbox.KeyCtrlY:
		return "KeyCtrlY"
	case termbox.KeyCtrlZ:
		return "KeyCtrlZ"
	case termbox.KeyEsc:
		return "KeyEsc | KeyCtrlLsqBracket | KeyCtrl3"
	case termbox.KeyCtrl4:
		return "KeyCtrl4 | KeyCtrlBackslash"
	case termbox.KeyCtrl5:
		return "KeyCtrl5 | KeyCtrlRsqBracket"
	case termbox.KeyCtrl6:
		return "KeyCtrl6"
	case termbox.KeyCtrl7:
		return "KeyCtrl7 | KeyCtrlSlash | KeyCtrlUnderscore"
	case termbox.KeySpace:
		return "KeySpace"
	case termbox.KeyBackspace2:
		return "KeyBackspace2 | KeyCtrl8"
	default:
		return "?"
	}
}

func termPrintf(x, y int, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	for i, c := range s {
		termbox.SetCell(x + i, y, c, termbox.ColorWhite,
				termbox.ColorBlack)
	}
}

func pad(code, comment string) string {
	for len(code) < 15 {
		code = code + " "
	}

	return "\t" + code + " // " + comment + "\n"
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	log, err := os.Create("log")
	if err != nil {
		panic(err)
	}
	defer log.Close()

	termPrintf(5, 5, "Logging events to ./log")
	termPrintf(5, 6, "Try 'tail -f log' from this directory")

	termbox.Flush()
	termbox.Sync()

	for {
		ev := termbox.PollEvent()

		log.WriteString("termbox.Event {\n")

		log.WriteString(pad(fmt.Sprintf("Type: %d,", ev.Type),
				    eventTypeToString(ev.Type)))

		log.WriteString(fmt.Sprintf("\tMod: %d,\n",
				ev.Mod))

		log.WriteString(pad(fmt.Sprintf("Key: %d,", ev.Key),
				    keyToString(ev.Key)))

		log.WriteString(pad(fmt.Sprintf("Ch: %d,", ev.Ch),
				    "'" + string(ev.Ch) + "'"))

		log.WriteString(fmt.Sprintf("\tWidth: %d,\n",
				ev.Width))

		log.WriteString(fmt.Sprintf("\tHeight: %d,\n",
				ev.Height))

		errStr := "nil"
		if ev.Err != nil {
			errStr = fmt.Sprintf("%s", ev.Err)
		}

		log.WriteString(fmt.Sprintf("\tErr: %s,\n",
				errStr))

		log.WriteString(fmt.Sprintf("\tMouseX: %d,\n",
				ev.MouseX))

		log.WriteString(fmt.Sprintf("\tMouseY: %d,\n",
				ev.MouseY))

		log.WriteString(fmt.Sprintf("\tN: %d,\n",
				ev.N))

		log.WriteString("}\n")

		if ev.Type == termbox.EventKey && ev.Key == termbox.KeyCtrlC {
			break
		}
	}
}

termbox-go-event-info
=====================

This is a little program to log all the termbox.Event structs to a file to help
figure out which events are being detected and find non-standard escape
sequences.

# Usage

Download and run it:

```bash
git clone https://github.com/briansteffens/termbox-go-event-info
cd termbox-go-event-info
go get github.com/nsf/termbox-go
go run main.go
```

Press keys or resize the window, and event information will be written to a
file called ```log``` in the same directory.

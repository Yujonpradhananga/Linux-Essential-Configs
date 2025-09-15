package main

import (
	"os"
	"syscall"
	"unsafe"
)

func (d *DocumentViewer) getTerminalSize() (int, int) {
	width, height := 80, 24
	type winsize struct {
		Row, Col, Xpixel, Ypixel uint16
	}
	ws := &winsize{}
	ret, _, _ := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(os.Stdout.Fd()),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)
	if ret == 0 {
		width = int(ws.Col)
		height = int(ws.Row)
	}
	return width, height
}

func (d *DocumentViewer) setRawMode() (*termios, error) {
	fd := int(os.Stdin.Fd())
	var old termios
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TCGETS, uintptr(unsafe.Pointer(&old)))
	if err != 0 {
		return nil, err
	}
	new := old
	new.Lflag &^= syscall.ECHO | syscall.ICANON
	new.Cc[syscall.VMIN] = 1
	new.Cc[syscall.VTIME] = 0
	_, _, err = syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TCSETS, uintptr(unsafe.Pointer(&new)))
	if err != 0 {
		return nil, err
	}
	d.oldState = &old
	return &old, nil
}

func (d *DocumentViewer) restoreTerminal(old *termios) {
	if old != nil {
		fd := int(os.Stdin.Fd())
		syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TCSETS, uintptr(unsafe.Pointer(old)))
	}
}

func (d *DocumentViewer) readSingleChar() byte {
	buf := make([]byte, 1)
	os.Stdin.Read(buf)
	return buf[0]
}

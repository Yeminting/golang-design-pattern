package bridge

import "fmt"

type IFileType interface {
	PlayFile()
}

type MP4 struct {
}

func (*MP4) PlayFile() {
	fmt.Print("play mp4 file")
}

type AVI struct {
}

func (*AVI) PlayFile() {
	fmt.Print("play avi file")
}

type IOSType interface {
	OSPlayer()
}

type WindowsPlayer struct {
}

func (*WindowsPlayer) OSPlayer() {
	fmt.Print(" with windows player")
}

type MacPlayer struct {
}

func (*MacPlayer) OSPlayer() {
	fmt.Print(" with mac player")
}

type LinuxPlayer struct {
}

func (*LinuxPlayer) OSPlayer() {
	fmt.Print(" with linux player")
}

type ViderPlayer struct {
	file IFileType
	os   IOSType
}

func (p *ViderPlayer) Play() {
	p.file.PlayFile()
	p.os.OSPlayer()
	fmt.Print("\n")
	//fmt.Println(p.file.PlayFile(), p.player.OSPlayer())
}

func (p *ViderPlayer) SetFile(f IFileType) {
	p.file = f
}
func (p *ViderPlayer) SetOS(o IOSType) {
	p.os = o
}

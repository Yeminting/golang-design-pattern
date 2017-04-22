package bridge

import "testing"

func TestBridge(t *testing.T) {
	var os IOSType
	var file IFileType

	p := ViderPlayer{}

	os = new(WindowsPlayer)
	file = new(MP4)
	p.SetFile(file)
	p.SetOS(os)
	p.Play()
	//out
	//play mp4 file with windows player

	os = new(MacPlayer)
	file = new(AVI)
	p.SetFile(file)
	p.SetOS(os)
	p.Play()
	//out
	//play avi file with mac player
}

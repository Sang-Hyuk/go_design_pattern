package main

type Command interface {
	Execute()
	Undo()
}

type EmptyCommand struct {

}

func (e EmptyCommand) Execute() {

}

func (e EmptyCommand) Undo() {

}

type LightOnCommand struct {
	light Light
}

func NewLightOnCommand(light Light) Command {
	return &LightOnCommand{light: light}
}

func (l LightOnCommand) Execute() {
	l.light.On()
}

func (l LightOnCommand) Undo() {
	l.light.Off()
}

type LightOffCommand struct {
	light Light
}

func NewLightOffCommand(light Light) Command {
	return &LightOffCommand{light: light}
}

func (l LightOffCommand) Execute() {
	l.light.Off()
}

func (l LightOffCommand) Undo() {
	l.light.On()
}

type StereoOnWithCDCommand struct {
    stereo Stereo
}

func NewStereoOnWithCDCommand(stereo Stereo) Command {
	return &StereoOnWithCDCommand{stereo: stereo}
}

func (s StereoOnWithCDCommand) Execute() {
	s.stereo.On()
	s.stereo.SetCD()
	s.stereo.SetVolume(11)
}

func (s StereoOnWithCDCommand) Undo() {
	s.stereo.Off()
}

type StereoOffWithCDCommand struct {
	stereo Stereo
}

func NewStereoOffWithCDCommand(stereo Stereo) Command {
	return &StereoOffWithCDCommand{stereo: stereo}
}

func (s StereoOffWithCDCommand) Execute() {
	s.stereo.Off()
}

func (s StereoOffWithCDCommand) Undo() {
	s.stereo.On()
	s.stereo.SetCD()
	s.stereo.SetVolume(11)
}



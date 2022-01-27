package main

func main() {
	//simpleVersion()

	upgradeVersion()
}

func simpleVersion()  {
	remote := SimpleRemoteControl{}
	light := NewKitchenLight()
	lightOn := NewLightOnCommand(light)

	remote.SetCommand(lightOn)
	remote.PressBtn()
}

func upgradeVersion() {
	remote := NewRemoteControl()

	kitchenLight := NewKitchenLight()
	stereo := NewDefaultStereo()

	kitchenLightOn := NewLightOnCommand(kitchenLight)
	kitchenLightOff := NewLightOffCommand(kitchenLight)
	stereoOn := NewStereoOnWithCDCommand(stereo)
	stereoOff := NewStereoOffWithCDCommand(stereo)

	remote.SetCommand(Slot1, kitchenLightOn, kitchenLightOff)
	remote.SetCommand(Slot2, stereoOn, stereoOff)

	remote.PressOnBtn(Slot1)
	remote.PressOffBtn(Slot1)
	remote.PressUndoBtn()

	remote.PressOnBtn(Slot2)
	remote.PressOffBtn(Slot2)
	remote.PressUndoBtn()

	remote.PressOnBtn(Slot3)
	remote.PressOffBtn(Slot3)
	remote.PressUndoBtn()
}

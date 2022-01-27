package main

type SimpleRemoteControl struct {
	slot Command
}

func (c *SimpleRemoteControl) SetCommand(command Command) {
	c.slot = command
}

func (c SimpleRemoteControl) PressBtn() {
	c.slot.Execute()
}

const MaxSlot = 7

type Slot int

const (
	Slot1 Slot  = iota
	Slot2
	Slot3
	Slot4
	Slot5
	Slot6
	Slot7
)

type RemoteControl struct {
	onCommands []Command
	offCommands []Command
	undoCommand Command
}

func NewRemoteControl() *RemoteControl {
	rc := RemoteControl{}
	rc.onCommands = make([]Command, MaxSlot)
	rc.offCommands = make([]Command, MaxSlot)
	rc.undoCommand = EmptyCommand{}
	for i := 0; i < MaxSlot; i++ {
		rc.onCommands[i] = EmptyCommand{}
		rc.offCommands[i] = EmptyCommand{}
	}
	return &rc
}

func (r *RemoteControl) SetCommand(slot Slot, onCommand, offCommand Command) {
	if onCommand == nil {
		onCommand = EmptyCommand{}
	}
	if offCommand == nil{
		offCommand = EmptyCommand{}
	}

	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

func (r *RemoteControl) PressOnBtn(slot Slot) {
	r.onCommands[slot].Execute()
	r.undoCommand = r.onCommands[slot]
}

func (r *RemoteControl) PressOffBtn(slot Slot) {
	r.offCommands[slot].Execute()
	r.undoCommand = r.offCommands[slot]
}

func (r RemoteControl) PressUndoBtn() {
	r.undoCommand.Undo()
}

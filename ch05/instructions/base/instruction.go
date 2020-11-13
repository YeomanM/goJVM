package base

import "goJVM/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (noi NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

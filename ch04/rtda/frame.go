package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) SetOperandStack(operandStack *OperandStack) {
	f.operandStack = operandStack
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) SetLocalVars(localVars LocalVars) {
	f.localVars = localVars
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

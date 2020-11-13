package rtda

import "math"

type OperandStack struct {
	size uint
	slot []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slot: make([]Slot, maxStack),
		}
	}
	return nil
}

func (os *OperandStack) PushInt(value int32) {
	os.slot[os.size].num = value
	os.size++
}

func (os *OperandStack) PopInt() int32 {
	os.size--
	return os.slot[os.size].num
}

func (os *OperandStack) PushFloat(value float32) {
	bits := math.Float32bits(value)
	os.slot[os.size].num = int32(bits)
	os.size++
}

func (os *OperandStack) PopFloat() float32 {
	os.size--
	bits := os.slot[os.size].num
	return math.Float32frombits(uint32(bits))
}

func (os *OperandStack) PushLong(value int64) {
	os.slot[os.size].num = int32(value)
	os.slot[os.size+1].num = int32(value >> 32)
	os.size += 2
}

func (os *OperandStack) PopLong() int64 {
	os.size -= 2
	low := os.slot[os.size].num
	high := os.slot[os.size+1].num
	return int64(uint32(high))<<32 | int64(uint32(low))
}

func (os *OperandStack) PushDouble(value float64) {
	bits := math.Float64bits(value)
	os.PushLong(int64(bits))
}

func (os *OperandStack) PopDouble() float64 {
	bits := os.PopLong()
	return math.Float64frombits(uint64(bits))
}

func (os *OperandStack) PushRef(ref *Object) {
	os.slot[os.size].ref = ref
	os.size++
}

func (os *OperandStack) PopRef() *Object {
	os.size--
	ref := os.slot[os.size].ref
	os.slot[os.size].ref = nil
	return ref
}

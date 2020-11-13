package rtda

import "math"

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}

	return nil
}

func (localVars LocalVars) SetInt(index uint, value int32) {
	localVars[index].num = value
}

func (localVars LocalVars) GetInt(index uint) int32 {
	return localVars[index].num
}

func (localVars LocalVars) SetFloat(index uint, value float32) {
	bits := math.Float32bits(value)
	localVars[index].num = int32(bits)
}

func (localVars LocalVars) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(localVars[index].num))
}

func (localVars LocalVars) SetLong(index uint, value int64) {
	localVars[index].num = int32(value)
	localVars[index+1].num = int32(value >> 32)
}

func (localVars LocalVars) GetLong(index uint) int64 {
	lower := localVars[index].num
	higher := localVars[index+1].num
	return int64(uint32(higher))<<32 | int64(uint32(lower))
}

func (localVars LocalVars) SetDouble(index uint, value float64) {
	bits := math.Float64bits(value)
	localVars.SetLong(index, int64(bits))
}

func (localVars LocalVars) GetDouble(index uint) float64 {
	bits := localVars.GetLong(index)
	return math.Float64frombits(uint64(bits))
}

func (localVars LocalVars) SetRef(index uint, ref *Object) {
	localVars[index].ref = ref
}

func (localVars LocalVars) GetRef(index uint) *Object {
	return localVars[index].ref
}

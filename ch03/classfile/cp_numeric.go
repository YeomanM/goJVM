package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	self.val = int32(reader.readUint32())
}

type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	self.val = math.Float32frombits(reader.readUint32())
}

type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	self.val = int64(reader.readUint64())
}

type COnstantDoubleInfo struct {
	val float64
}

func (self *COnstantDoubleInfo) readInfo(reader *ClassReader) {
	self.val = math.Float64frombits(reader.readUint64())
}

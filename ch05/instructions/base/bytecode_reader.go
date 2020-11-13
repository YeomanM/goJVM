package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func (br *BytecodeReader) Reset(code []byte, pc int) {
	br.code = code
	br.pc = pc
}

func (br *BytecodeReader) ReadUint8() uint8 {
	i := br.code[br.pc]
	br.pc++
	return i
}

func (br *BytecodeReader) ReadInt8() int8 {
	return int8(br.ReadUint8())
}

func (br *BytecodeReader) ReadUint16() {

}

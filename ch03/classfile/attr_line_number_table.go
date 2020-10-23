package classfile

type LineNumberTableAttribute struct {
	lineNumberTableEntry []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	count := reader.readUint16()
	entries := make([]*LineNumberTableEntry, count)

	for i := range entries {
		entries[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}

}

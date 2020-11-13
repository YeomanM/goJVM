package classfile

type LocalVariableTableAttribute struct {
	localVariableTableEntry []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	count := reader.readUint16()
	entries := make([]*LocalVariableTableEntry, count)

	for i := range entries {
		entries[i] = &LocalVariableTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}

}

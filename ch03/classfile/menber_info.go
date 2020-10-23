package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributeCount  uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	count := reader.readUint16()
	members := make([]*MemberInfo, count)

	for i := range members {
		members[i] = readMember(reader, cp)
	}

	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	memberInfo := &MemberInfo{
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributeCount:  reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
	return memberInfo
}

func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

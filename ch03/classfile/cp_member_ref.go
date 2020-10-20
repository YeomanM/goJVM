package classfile

type ConstantMemberRefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberRefInfo) ClassName() string {

}

func (self *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {

}

type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantInterfacesRefInfo struct {
	ConstantMemberRefInfo
}

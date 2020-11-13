package classfile

import "fmt"

type ClassFile struct {
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) SetAccessFlags(accessFlags uint16) {
	self.accessFlags = accessFlags
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)

			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(cr *ClassReader) {
	self.readAndCheckMagic(cr)
	self.readAndCheckVersion(cr)
	self.constantPool = readConstantPool(cr)
	self.accessFlags = cr.readUint16()
	self.thisClass = cr.readUint16()
	self.superClass = cr.readUint16()
	self.interfaces = cr.readUint16s()
	self.fields = readMembers(cr, self.constantPool)
	self.methods = readMembers(cr, self.constantPool)
	self.attributes = readAttributes(cr, self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(cr *ClassReader) {
	magic := cr.readUint32()

	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}

}

func (self *ClassFile) readAndCheckVersion(cr *ClassReader) {
	self.minorVersion = cr.readUint16()
	self.majorVersion = cr.readUint16()

	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) Attributes() []AttributeInfo {
	return self.attributes
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) Interfaces() []uint16 {
	return self.interfaces
}

func (self *ClassFile) SuperClass() uint16 {
	return self.superClass
}

func (self *ClassFile) ThisClass() uint16 {
	return self.thisClass
}

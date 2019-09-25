package stores

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Store into reference array
type AASTORE struct{ base.NoOperandsInstruction }

func (instr *AASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Refs()[index] = ref
	}
}

// Store into byte or boolean array
type BASTORE struct{ base.NoOperandsInstruction }

func (instr *BASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Bytes()[index] = int8(val)
	}
}

// Store into char array
type CASTORE struct{ base.NoOperandsInstruction }

func (instr *CASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Chars()[index] = uint16(val)
	}
}

// Store into double array
type DASTORE struct{ base.NoOperandsInstruction }

func (instr *DASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Doubles()[index] = val
	}
}

// Store into float array
type FASTORE struct{ base.NoOperandsInstruction }

func (instr *FASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Floats()[index] = val
	}
}

// Store into int array
type IASTORE struct{ base.NoOperandsInstruction }

func (instr *IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Ints()[index] = val
	}
}

// Store into long array
type LASTORE struct{ base.NoOperandsInstruction }

func (instr *LASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Longs()[index] = val
	}
}

// Store into short array
type SASTORE struct{ base.NoOperandsInstruction }

func (instr *SASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Shorts()[index] = int16(val)
	}
}

func _checkArrayAndIndex(frame *rtda.Frame, arrRef *heap.Object, index int32) bool {
	if arrRef == nil {
		frame.Thread().ThrowNPE()
		return false
	}
	if index < 0 || index >= heap.ArrayLength(arrRef) {
		frame.Thread().ThrowArrayIndexOutOfBoundsException(index)
		return false
	}
	return true
}

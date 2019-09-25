package stack

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }

func (instr *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	stack.PushSlot(val1)
	stack.PushSlot(val2)
}

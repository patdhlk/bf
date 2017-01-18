package compile

import (
    
	"github.com/patdhlk/bf/lang"
)

type Compiler struct {
	code       string
	codeLength int
	position   int

	instructions []*lang.Instruction
}

func NewCompiler(code string) *Compiler {
	return &Compiler{
		code:         code,
		codeLength:   len(code),
		instructions: []*lang.Instruction{},
	}
}

func (c *Compiler) Compile() []*lang.Instruction {
	loopStack := []int{}

	for c.position < c.codeLength {
		current := c.code[c.position]

		switch current {
		case '[':
			insPos := c.EmitWithArg(lang.JumpIfZero, 0)
			loopStack = append(loopStack, insPos)
		case ']':
			// Pop position of last JumpIfZero ("[") instruction off stack
			openInstruction := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]
			// Emit the new JumpIfNotZero ("]") instruction, with correct position as argument
			closeInstructionPos := c.EmitWithArg(lang.JumpIfNotZero, openInstruction)
			// Patch the old JumpIfZero ("[") instruction with new position
			c.instructions[openInstruction].Argument = closeInstructionPos

		case '+':
			c.CompileFoldableInstruction('+', lang.Plus)
		case '-':
			c.CompileFoldableInstruction('-', lang.Minus)
		case '<':
			c.CompileFoldableInstruction('<', lang.Left)
		case '>':
			c.CompileFoldableInstruction('>', lang.Right)
		case '.':
			c.CompileFoldableInstruction('.', lang.PutChar)
		case ',':
			c.CompileFoldableInstruction(',', lang.ReadChar)
		}

		c.position++
	}

	return c.instructions
}

func (c *Compiler) CompileFoldableInstruction(char byte, insType lang.InsType) {
	count := 1

	for c.position < c.codeLength-1 && c.code[c.position+1] == char {
		count++
		c.position++
	}

	c.EmitWithArg(insType, count)
}

func (c *Compiler) EmitWithArg(insType lang.InsType, arg int) int {
	ins := &lang.Instruction{Type: insType, Argument: arg}
	c.instructions = append(c.instructions, ins)
	return len(c.instructions) - 1
}

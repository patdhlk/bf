package compile

import (
	"testing"

	"github.com/patdhlk/bf/lang"
)

func TestCompile(t *testing.T) {
	input := `
	+++++
	-----
	+++++
	>>>>>
	<<<<<
	`
	expected := []*lang.Instruction{
		&lang.Instruction{lang.Plus, 5},
		&lang.Instruction{lang.Minus, 5},
		&lang.Instruction{lang.Plus, 5},
		&lang.Instruction{lang.Right, 5},
		&lang.Instruction{lang.Left, 5},
	}

	compiler := NewCompiler(input)
	bytecode := compiler.Compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("wrong op. want=%+v, got=%+v", op, bytecode[i])
		}
	}
}

func TestCompileLoops(t *testing.T) {
	input := `+[+[+]+]+`
	expected := []*lang.Instruction{
		&lang.Instruction{lang.Plus, 1},
		&lang.Instruction{lang.JumpIfZero, 7},
		&lang.Instruction{lang.Plus, 1},
		&lang.Instruction{lang.JumpIfZero, 5},
		&lang.Instruction{lang.Plus, 1},
		&lang.Instruction{lang.JumpIfNotZero, 3},
		&lang.Instruction{lang.Plus, 1},
		&lang.Instruction{lang.JumpIfNotZero, 1},
		&lang.Instruction{lang.Plus, 1},
	}

	compiler := NewCompiler(input)
	bytecode := compiler.Compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("wrong op. want=%+v, got=%+v", op, bytecode[i])
		}
	}
}

func TestCompileEverything(t *testing.T) {
	input := `+++[---[+]>>>]<<<`
	expected := []*lang.Instruction{
		&lang.Instruction{lang.Plus, 3},
		&lang.Instruction{lang.JumpIfZero, 7},
		&lang.Instruction{lang.Minus, 3},
		&lang.Instruction{lang.JumpIfZero, 5},
		&lang.Instruction{lang.Plus, 1},
		&lang.Instruction{lang.JumpIfNotZero, 3},
		&lang.Instruction{lang.Right, 3},
		&lang.Instruction{lang.JumpIfNotZero, 1},
		&lang.Instruction{lang.Left, 3},
	}

	compiler := NewCompiler(input)
	bytecode := compiler.Compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("wrong bytecode length. want=%+v, got=%+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("wrong op. want=%+v, got=%+v", op, bytecode[i])
		}
	}
}

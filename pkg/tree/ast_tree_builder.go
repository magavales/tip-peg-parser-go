package tree

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"strings"
	"tip-peg-parser-go/pkg/statement"
)

type ASTreeBuilder struct {
	Root           statement.RootStatement
	levelStack     *stack.Stack
	statementStack *stack.Stack
}

func (b *ASTreeBuilder) Build() ASTree {
	var (
		sbuilder strings.Builder
	)
	sbuilder.WriteString("\n")
	b.levelStack = new(stack.Stack)
	b.statementStack = new(stack.Stack)

	// push in stack in reverse order
	for i := len(b.Root.GetBodyStatement()) - 1; i >= 0; i-- {
		b.statementStack.Push(*(b.Root.GetBodyStatement()[i]))
	}
	b.levelStack.Push(len(b.Root.GetBodyStatement()))

	for b.statementStack.Len() > 0 {
		if b.levelStack.Peek() == 0 {
			b.levelStack.Pop()
		}
		stm := b.statementStack.Pop()
		switch t := (stm).(type) {
		case *statement.OutputStatement:
			sbuilder.WriteString(b.getOffset())
			sbuilder.WriteString("output: ")
			sbuilder.WriteString(fmt.Sprintf("%s", t.Expression))
			sbuilder.WriteString("\n")
			b.reduceLevelReminderOrPop()
		case *statement.EqualStatement:
			sbuilder.WriteString(b.getOffset())
			sbuilder.WriteString(t.Id)
			sbuilder.WriteString(" = ")
			sbuilder.WriteString(fmt.Sprintf("%s", t.Expression))
			sbuilder.WriteString("\n")
			b.reduceLevelReminderOrPop()
		case *statement.IfStatement:
			sbuilder.WriteString(b.getOffset())
			sbuilder.WriteString("if ")
			sbuilder.WriteString(fmt.Sprintf("%s", t.Condition))
			sbuilder.WriteString(" is true ↓\n")
			b.reduceLevelReminderOrPop()
			b.pushToStatementStack(t)
		case *statement.ElseStatement:
			sbuilder.WriteString(b.getOffset())
			sbuilder.WriteString("else ↓\n")
			b.reduceLevelReminder()
			b.pushToStatementStack(t)
		case *statement.WhileStatement:
			sbuilder.WriteString(b.getOffset())
			sbuilder.WriteString("while ")
			sbuilder.WriteString(fmt.Sprintf("%s", t.Condition))
			sbuilder.WriteString(" is true ↺\n")
			b.reduceLevelReminder()
			b.pushToStatementStack(t)
		default:
			panic("")
		}
	}

	fmt.Println(sbuilder.String())
	return ASTree{Str: sbuilder.String()}
}

func (b *ASTreeBuilder) getOffset() string {
	return strings.Repeat("    ", b.levelStack.Len())
}

func (b *ASTreeBuilder) reduceLevelReminderOrPop() {
	reminder := b.levelStack.Pop().(int)
	if (reminder - 1) > 0 {
		b.levelStack.Push(reminder - 1)
	}
}

func (b *ASTreeBuilder) reduceLevelReminder() {
	reminder := b.levelStack.Pop().(int)
	b.levelStack.Push(reminder - 1)
}

func (b *ASTreeBuilder) pushToStatementStack(container statement.Container) {
	for i := len(container.GetBodyStatement()) - 1; i >= 0; i-- {
		b.statementStack.Push(*(container.GetBodyStatement())[i])
	}
	b.levelStack.Push(len(container.GetBodyStatement()))
}

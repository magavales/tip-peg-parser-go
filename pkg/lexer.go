package pkg

import (
	"errors"
	"fmt"
	"github.com/golang-collections/collections/stack"
	"strings"
	"tip-peg-parser-go/pkg/expression"
	"tip-peg-parser-go/pkg/expression/atomic"
	"tip-peg-parser-go/pkg/statement"
)

var (
	stmTokens = statement.GetAllStatementTokens()
	expTokens = expression.GetAllExpressionTokens()
	atmTokens = expression.GetAllAtomicTokens()
)

func Parse(lines []string) statement.RootStatement {
	flatted := lineFlatter(lines)
	statements, _ := statementLexer(flatted)
	fmt.Printf("Statements: ")
	for _, s := range statements {
		fmt.Printf("%s", *s)
	}
	fmt.Printf("\n")
	rootStatement, err := statementGrouper(statements)
	if err != nil {
		panic(err)
	}
	return rootStatement
}

func statementGrouper(stments []*statement.Statement) (statement.RootStatement, error) {
	headStatement := new(stack.Stack)
	headStatement.Push(new(statement.RootStatement))

	for _, stm := range stments {
		t := (*stm).(interface{})
		switch t.(type) {
		case *statement.IfStatement:
			headStatement.Push(*stm)
		case *statement.WhileStatement:
			headStatement.Push(*stm)
		case *statement.ElseStatement:
			headStatement.Push(*stm)
		case *statement.EndBodyStatement:
			currentHead := headStatement.Pop()
			root := headStatement.Pop()
			root.(statement.Container).Put(currentHead.(statement.Statement))
			headStatement.Push(root)
		default:
			root := headStatement.Pop()
			root.(statement.Container).Put(*stm)
			headStatement.Push(root)
		}
	}

	if headStatement.Len() != 1 {
		return statement.RootStatement{}, errors.New(fmt.Sprintf("Cannot to resolve: %s", headStatement.Peek()))
	}

	returnValue := headStatement.Pop()
	v := returnValue.(statement.Container)
	return statement.RootStatement{Body: statement.BodyStatement{Statements: v.GetBodyStatement()}}, nil
}

func statementLexer(lines []string) ([]*statement.Statement, error) {
	var (
		output []*statement.Statement
	)

	prev := ""
	for _, line := range lines {
		next := prev + " " + line

		statement, err := tryToGuessStatement(next)
		if err != nil {
			fmt.Printf("Cannot recognize token: %s\n", prev)
			prev = next
			continue
		}

		output = append(output, &statement)

		fmt.Printf("Resolve token: %s\n", next)
		pattern, _ := statement.GetToken().GetPattern()
		replaced := pattern.ReplaceAllString(next, "")
		prev = strings.TrimSpace(replaced)
	}

	if strings.TrimSpace(prev) != "" {
		fmt.Printf("Cannot recognize statement: %s\n", prev)
		return []*statement.Statement{}, errors.New("cannot recognize statement: " + prev)
	}

	return output, nil
}

func tryToGuessStatement(line string) (statement.Statement, error) {
	for _, token := range stmTokens {
		pattern, _ := token.GetPattern()
		matches := pattern.FindAllStringSubmatch(line, -1)
		if len(matches) != 0 {
			switch token.Pattern {
			case statement.IfExp:
				guessExpression, err := tryToGuessExpression(matches[0][1])
				return &statement.IfStatement{Condition: guessExpression}, err
			case statement.WhileExp:
				guessExpression, err := tryToGuessExpression(matches[0][1])
				return &statement.WhileStatement{Condition: guessExpression}, err
			case statement.IfElse:
				return &statement.ElseStatement{}, nil
			case statement.OutputExp:
				exp, err := tryToGuessExpression(matches[0][2])
				return &statement.OutputStatement{Operation: matches[0][1], Expression: exp}, err
			case statement.EndStm:
				return &statement.EndBodyStatement{}, nil
			case statement.IdEqExp:
				exp, err := tryToGuessExpression(matches[0][2])
				return &statement.EqualStatement{Id: matches[0][1], Expression: exp}, err
			default:
				return nil, errors.New("switch case mismatch")
			}
		}
	}
	return nil, errors.New("cannot recognize token")
}

func tryToGuessExpression(line string) (expression.Expression, error) {
	for _, token := range expTokens {
		pattern, _ := token.GetPattern()
		matches := pattern.FindAllStringSubmatch(line, -1)
		if len(matches) != 0 {
			fmt.Printf("Expression: %s\n", matches)
			switch token.Pattern {
			case expression.ExpPlusExp:
				left, err := tryToGuessExpression(matches[0][1])
				right, err := tryToGuessExpression(matches[0][2])
				return expression.PlusExpression{BiExpression: expression.BiExpression{Left: left, Right: right}}, err
			case expression.ExpMinusExp:
				left, err := tryToGuessExpression(matches[0][1])
				right, err := tryToGuessExpression(matches[0][2])
				return expression.MinusExpression{BiExpression: expression.BiExpression{Left: left, Right: right}}, err
			case expression.ExpMultiExp:
				left, err := tryToGuessExpression(matches[0][1])
				right, err := tryToGuessExpression(matches[0][2])
				return expression.MultiExpression{BiExpression: expression.BiExpression{Left: left, Right: right}}, err
			case expression.ExpDivExp:
				left, err := tryToGuessExpression(matches[0][1])
				right, err := tryToGuessExpression(matches[0][2])
				return expression.DivExpression{BiExpression: expression.BiExpression{Left: left, Right: right}}, err
			case expression.ExpBox:
				exp, err := tryToGuessExpression(matches[0][1])
				return expression.BoxExpression{Expression: exp}, err
			}
		}
	}

	return tryToGuessAtomic(line)
}

func tryToGuessAtomic(line string) (expression.Expression, error) {
	for _, token := range atmTokens {
		pattern, _ := token.GetPattern()
		matches := pattern.FindAllStringSubmatch(line, -1)
		if len(matches) != 0 {
			fmt.Printf("Atomic: %s\n", matches)
			switch token.Pattern {
			case expression.Id:
				return atomic.IdExpression{Id: matches[0][1]}, nil
			case expression.Int:
				return atomic.IntExpression{Value: matches[0][1]}, nil
			case expression.Input:
				return atomic.InputExpression{}, nil
			}
		}
	}

	return nil, errors.New("cannot recognize token")
}

func lineFlatter(lines []string) []string {
	var output []string
	for _, line := range lines {
		line = strings.ReplaceAll(line, "(", " ( ")
		line = strings.ReplaceAll(line, ")", " ) ")
		line = strings.ReplaceAll(line, ";", " ; ")
		for _, split := range strings.Split(line, " ") {
			if split != "" {
				output = append(output, split)
			}
		}
	}

	fmt.Printf("Flatter: %s\n", output)
	return output
}

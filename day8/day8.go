package day8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FrankBro/AdventOfCode2020/util"
)

const (
	errReadLines = "failed to read input as lines: %w"
)

type OpType int

const (
	OpTypeNop OpType = iota
	OpTypeAcc
	OpTypeJmp
)

type Op struct {
	Type  OpType
	Value int
}

func parseLines(lines []string) []Op {
	ops := make([]Op, len(lines))
	for i, line := range lines {
		ops[i] = parseLine(line)
	}
	return ops
}

func parseLine(line string) Op {
	spaceIndex := strings.Index(line, " ")
	op := line[:spaceIndex]
	value := line[spaceIndex+1:]
	var opType OpType
	switch op {
	case "nop":
		opType = OpTypeNop
	case "acc":
		opType = OpTypeAcc
	case "jmp":
		opType = OpTypeJmp
	default:
		panic("wat")
	}
	isMinus := value[0] == '-'
	i, err := strconv.Atoi(value[1:])
	if err != nil {
		panic(err)
	}
	if isMinus {
		i *= -1
	}
	return Op{
		Type:  opType,
		Value: i,
	}
}

func exec(ops []Op) (int, bool) {
	var acc int
	var pc int
	visited := make(map[int]struct{})
	for {
		if _, ok := visited[pc]; ok {
			return acc, false
		}
		visited[pc] = struct{}{}
		op := ops[pc]
		switch op.Type {
		case OpTypeNop:
			pc++
		case OpTypeAcc:
			acc += op.Value
			pc++
		case OpTypeJmp:
			pc += op.Value
		}
		if pc == len(ops) {
			return acc, true
		}
	}
}

func Part1() (int, error) {
	lines, err := util.ReadLinesNoEmpty()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	ops := parseLines(lines)
	acc, _ := exec(ops)
	return acc, nil
}

func Part2() (int, error) {
	lines, err := util.ReadLinesNoEmpty()
	if err != nil {
		return 0, fmt.Errorf(errReadLines, err)
	}
	ops := parseLines(lines)
	for i, op := range ops {
		if op.Type == OpTypeNop || op.Type == OpTypeJmp {
			copiedOps := make([]Op, len(ops))
			copy(copiedOps, ops)
			if op.Type == OpTypeNop {
				op.Type = OpTypeJmp
			} else if op.Type == OpTypeJmp {
				op.Type = OpTypeNop
			}
			copiedOps[i] = op
			acc, good := exec(copiedOps)
			if good {
				return acc, nil
			}
		}
	}
	return 0, nil
}

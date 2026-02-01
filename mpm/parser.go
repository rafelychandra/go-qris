package mpm

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	IDWordCount          = 2
	ValueLengthWordCount = 2
)

type (
	Parser struct {
		current int64
		max     int64
		source  []rune
		err     error
	}

	ParserError struct {
		Func string
		Err  error
	}
)

func NewParser(tlv string) *Parser {
	if len(strings.TrimSpace(tlv)) == 0 {
		return &Parser{err: fmt.Errorf("invalid tlv: tlv is empty")}
	}

	return &Parser{
		current: -1,
		max:     int64(utf8.RuneCountInString(tlv)),
		source:  []rune(tlv),
		err:     nil,
	}
}

func (p *Parser) Next() bool {
	if p.err != nil {
		return false
	}
	if p.current < 0 {
		p.current = 0
	} else {
		valueLength := p.ValueLength()
		if p.err != nil {
			return false
		}
		p.current += valueLength + IDWordCount + ValueLengthWordCount
	}
	if p.current >= p.max {
		return false
	}
	return true
}

func (p *Parser) Tag() Tag {
	const fnID = "Tag"
	start := p.current
	end := start + IDWordCount
	if p.current < 0 {
		p.err = notCallError(fnID)
		return ""
	}
	if p.max < end {
		p.err = outOfRangeError(fnID, p.current, p.max, start, end)
		return ""
	}
	id := Tag(p.source[start:end])
	return id
}

func (p *Parser) ValueLength() int64 {
	const fnValueLength = "ValueLength"
	start := p.current + IDWordCount
	end := start + ValueLengthWordCount
	if p.current < 0 {
		p.err = notCallError(fnValueLength)
		return 0
	}
	if p.max < end {
		p.err = outOfRangeError(fnValueLength, p.current, p.max, start, end)
		return 0
	}
	strValueLength := string(p.source[start:end])
	valueLength, err := strconv.ParseInt(strValueLength, 10, 64)
	if err != nil {
		p.err = syntaxError(fnValueLength, strValueLength)
		return 0
	}
	return valueLength
}

func (p *Parser) Value() string {
	const fnValue = "Value"
	start := p.current + IDWordCount + ValueLengthWordCount
	end := start + p.ValueLength()
	if p.current < 0 {
		p.err = notCallError(fnValue)
		return ""
	}
	if p.max < end {
		p.err = outOfRangeError(fnValue, p.current, p.max, start, end)
		return ""
	}
	return string(p.source[start:end])
}

func (p *Parser) Err() error {
	return p.err
}

func (e *ParserError) Error() string {
	return "parser." + e.Func + ": " + e.Err.Error()
}

func notCallError(fn string) *ParserError {
	return &ParserError{
		Func: fn,
		Err:  fmt.Errorf("not call Next()"),
	}
}

func outOfRangeError(fn string, current, max, start, end int64) *ParserError {
	return &ParserError{
		Func: fn,
		Err:  fmt.Errorf("bounds out of range. current: %d, max: %d, start: %d, end: %d", current, max, start, end),
	}
}

func syntaxError(fn, str string) *ParserError {
	return &ParserError{
		Func: fn,
		Err:  fmt.Errorf("parsing %s: %s", strconv.Quote(str), strconv.ErrSyntax.Error()),
	}
}

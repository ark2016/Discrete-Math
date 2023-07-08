package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type token struct {
	teg   byte
	value interface{}
}

type pair struct {
	head int
	tail int
}

type elem struct {
	numPer   int
	position int
}

type lexemes struct {
	tail        int
	graph       Pairs
	numFunc     int
	mapFunc     MSE
	arr         tokens
	index       int
	currentArgs strings
}

type Vertex struct {
	name   int
	timeIn int
	low    int
	comp   int
}

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value *Vertex
		prev  *node
	}
)

type MSE map[string]elem
type Pairs []pair
type tokens []token
type strings []string

func (lex lexemes) check(s string) {
	for i := 0; i < len(lex.currentArgs); i++ {
		if lex.currentArgs[i] == s {
			return
		}
	}
	lex.err("Not in current args")
}

func (lex lexemes) err(s string) {
	//fmt.Printf("error: " + s)
	//os.Exit(1)
	fmt.Printf("error")
	os.Exit(0)
}
func (lex lexemes) Teg() byte {
	return lex.arr[lex.index].teg
}
func (lex lexemes) Val() interface{} {
	return lex.arr[lex.index].value
}
func (lex *lexemes) ExpectTeg(a byte) {
	if a == lex.Teg() {
		lex.index++
	} else {
		lex.err("Not expected lex.Teg")
	}
	return
}
func (lex *lexemes) ExpectTegAndVal(a byte, v interface{}) {
	if a == lex.Teg() && v == lex.Val() {
		lex.index++
	} else {
		lex.err("Not expected lex.Teg & lex.Val")
	}
	return
}
func (lex *lexemes) Programm() {
	lex.numFunc++
	lex.Function()

	if 5 == lex.Teg() {
		if lex.numFunc != len(lex.mapFunc) {
			lex.err("number of functions")
		}
		return
	} else {
		lex.Programm()
	}
}

func (lex *lexemes) Function() {
	lex.ExpectTeg(1)
	nameIfFunction := getBeforeVal(lex)
	lex.ExpectTeg(2)
	n, s := lex.FormalArgsList()
	lex.currentArgs = s

	if num, ok := lex.mapFunc[nameIfFunction]; ok {
		if n != num.numPer {
			lex.err("number of variables")
		}
		lex.tail = num.position
	} else {
		lex.tail = len(lex.mapFunc)
		lex.mapFunc[nameIfFunction] = elem{numPer: n, position: lex.tail}
	}

	lex.ExpectTeg(3)
	lex.ExpectTegAndVal(4, ":=")
	lex.Expr()
	lex.ExpectTegAndVal(4, ";")
}

func getBeforeVal(lex *lexemes) string {
	nameIfFunction := lex.arr[lex.index-1].value.(string)
	return nameIfFunction
}

func (lex *lexemes) FormalArgsList() (l int, s []string) {
	if 3 != lex.Teg() {
		lex.Ident_List(&s)
	}

	l = len(s)
	return
}

func (lex *lexemes) Ident_List(s *[]string) {
	lex.ExpectTeg(1)
	*s = append(*s, getBeforeVal(lex))

	if 4 == lex.Teg() && "," == lex.Val() {
		lex.index++
		lex.Ident_List(s)
	}
}

func (lex *lexemes) Expr() {
	lex.Comparison_Expr()

	if 4 == lex.Teg() && "?" == lex.Val() {
		lex.index++
		lex.Comparison_Expr()
		lex.ExpectTegAndVal(4, ":")
		lex.Expr()
	}
}

func (lex *lexemes) Comparison_Expr() {
	lex.ArithExpr()

	if lex.Comparison_Op() {
		lex.index++
		lex.ArithExpr()
	}
}

func (lex lexemes) Comparison_Op() bool {
	return 4 == lex.Teg() && ("=" == lex.Val() || "<>" == lex.Val() ||
		"<" == lex.Val() || ">" == lex.Val() ||
		"<=" == lex.Val() || ">=" == lex.Val())
}

func (lex *lexemes) ArithExpr() {
	lex.Term()
	lex.ArithExprTagAndVal()
}

func (lex *lexemes) ArithExprTagAndVal() {
	if lex.Teg() == 4 && (lex.Val() == "+" || lex.Val() == "-") {
		lex.index++
		lex.Term()
		lex.ArithExprTagAndVal()
	}
}

func (lex *lexemes) Term() {
	lex.Factor()
	lex.TermTagAndVal()
}

func (lex *lexemes) TermTagAndVal() {
	if 4 == lex.Teg() && ("*" == lex.Val() || "/" == lex.Val()) {
		lex.index++
		lex.Factor()
		lex.TermTagAndVal()
	}
}

func (lex *lexemes) Factor() {
	switch lex.Teg() {
	case 0:
		lex.index++
	case 1:
		lex.index++
		if 2 == lex.Teg() {
			nameOfFunction := getBeforeVal(lex)
			lex.index++
			n := lex.ActualArgsList()
			if num, ok := lex.mapFunc[nameOfFunction]; ok {
				if num.numPer != n {
					lex.err("number of variables")
				}
				lex.graph = append(lex.graph, pair{tail: lex.tail, head: num.position})
			} else {
				lex.graph = append(lex.graph, pair{tail: lex.tail, head: len(lex.mapFunc)})
				lex.mapFunc[nameOfFunction] = elem{numPer: n, position: len(lex.mapFunc)}
			}
			lex.ExpectTeg(3)
		} else {
			lex.check(getBeforeVal(lex))
		}
	case 2:
		lex.index++
		lex.Expr()
		lex.ExpectTeg(3)
	case 4:
		if "-" == lex.Val() {
			lex.index++
			lex.Factor()
		} else {
			lex.err("val of variables")
		}
	default:
		lex.err("Tag of variables")
	}
}

func (lex *lexemes) ActualArgsList() int {
	k := 0

	if 3 != lex.Teg() {
		k++
		lex.ExprList(&k)
	}
	return k
}

func (lex *lexemes) ExprList(k *int) {
	lex.Expr()

	if 4 == lex.Teg() && "," == lex.Val() {
		lex.index++
		*k++
		lex.ExprList(k)
	}
}

func Lexer(text string) (arr tokens) {
	var flnum, flval bool
	var c byte
	var start, num, i int

	for i = 0; i < len(text); i++ {
		c = text[i]

		if flnum {
			if '9' >= c && '0' <= c {
				continue
			}
			num, _ = strconv.Atoi(text[start:i])
			arr = append(arr, token{teg: 0, value: num})
			flnum = false
		}

		if flval {
			if 'a' <= c && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' {
				continue
			}
			arr = append(arr, token{teg: 1, value: text[start:i]})
			flval = false
		}

		switch c {
		case '=', ',', '*', '/', '+', '-', '?', ';':
			arr = append(arr, token{teg: 4, value: text[i : i+1]})
		case ':', '>', '<':
			if i+1 < len(text) && ('=' == text[i+1] || '<' == c && '>' == text[i+1]) {
				arr = append(arr, token{teg: 4, value: text[i : i+2]})
				i++
			} else {
				arr = append(arr, token{teg: 4, value: text[i : i+1]})
			}
		case '(':
			arr = append(arr, token{teg: 2, value: nil})
		case ')':
			arr = append(arr, token{teg: 3, value: nil})
		case ' ', '\n', '\t':
			continue
		default:
			if 'a' <= c && 'z' >= c || 'A' <= c && 'Z' >= c {
				start = i
				flval = true
			} else if '0' <= c && '9' >= c {
				start = i
				flnum = true
			} else {
				fmt.Printf("error")
				os.Exit(0)
			}
		}
	}

	if flnum {
		num, _ = strconv.Atoi(text[start:i])
		arr = append(arr, token{teg: 0, value: num})
		flnum = false
	}

	if flval {
		arr = append(arr, token{teg: 1, value: text[start:i]})
		flval = false
	}

	arr = append(arr, token{teg: 5, value: nil})

	return
}

func Parser(text string) lexemes {
	var lex lexemes

	setGraph(lex)
	lex.arr = Lexer(text)
	lex.mapFunc = make(MSE)
	lex.Programm()
	return lex
}

func setGraph(lex lexemes) {
	lex.graph = make(Pairs, 0)
}

func (s *Stack) Pop() *Vertex {
	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

func (s *Stack) Push(value *Vertex) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}

func Tarjan(G [][]*Vertex, V []*Vertex) int {
	var v *Vertex
	var count, time int
	var S Stack

	time++

	for _, v = range V {
		if v.timeIn == 0 {
			VisitvertexTarjan(G, v, &S, &time, &count)
		}
	}
	return count
}
func VisitvertexTarjan(G [][]*Vertex, v *Vertex, S *Stack, time, count *int) {
	var u *Vertex

	v.timeIn, v.low = *time, *time
	*time++
	S.Push(v)

	for _, u = range G[v.name] {
		if u.timeIn == 0 {
			VisitvertexTarjan(G, u, S, time, count)
		}
		if u.comp == -1 && v.low > u.low {
			v.low = u.low
		}
	}

	if v.timeIn == v.low {
		for {
			u = S.Pop()
			u.comp = *count
			if u.name == v.name {
				break
			}
		}
		*count++
	}
}

func main() {
	//buf, _ := io.ReadAll(os.Stdin)
	//buf := "f(x, y) := x+y; 10"
	var buf string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buf += scanner.Text()
	}

	lex := Parser(buf)

	var m, n, a, b, i int

	n = lex.numFunc
	m = len(lex.graph)
	V := make([]*Vertex, n)

	for i = 0; i < n; i++ {
		V[i] = &Vertex{
			name: i,
			comp: -1,
		}
	}

	G := make([][]*Vertex, n)

	for i = 0; i < m; i++ {
		a = lex.graph[i].tail
		b = lex.graph[i].head
		G[a] = append(G[a], V[b])
	}

	count := Tarjan(G, V)
	fmt.Println(count)
}

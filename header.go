package main

const (
	TTYPE_IDENT int = iota
	TTYPE_PUNCT
	TTYPE_INT
	TTYPE_CHAR
	TTYPE_STRING
)

type Token struct {
	typ int
	v   struct { // wanna be Union
		ival  int
		sval  []byte
		punct byte
		c     byte
	}
}

const (
	AST_LITERAL byte = iota
	AST_STRING
	AST_LVAR
	AST_LREF
	AST_GVAR
	AST_GREF
	AST_FUNC
	AST_FUNCALL
	AST_DECL
	AST_ARRAY_INIT
	AST_ADDR
	AST_DEREF
	AST_IF
)

const (
	CTYPE_VOID int = iota
	CTYPE_INT
	CTYPE_CHAR
	CTYPE_ARRAY
	CTYPE_PTR
)

type Ctype struct {
	typ  int
	ptr  *Ctype
	size int
}

type Ast struct {
	typ   byte
	ctype *Ctype
	// want to be "union"
	// Integer
	ival int
	// Char
	c byte
	// String
	str struct {
		val    []byte
		slabel []byte
	}
	// Local variable
	variable struct {
		lname []byte
		loff  int
	}
	// Global variable
	gvar struct {
		gname  []byte
		glabel []byte
	}
	// Local reference
	lref struct {
		ref *Ast
		off int
	}
	// Global reference
	gref struct {
		ref *Ast
		off int
	}
	// Binary operator
	binop struct {
		left  *Ast
		right *Ast
	}
	// Unary operator
	unary struct {
		operand *Ast
	}
	// Function call or function declaration
	fnc struct {
		fname []byte
		args  []*Ast
		params []*Ast
		locals []*Ast
		body []*Ast
	}
	// Declaration
	decl struct {
		declvar  *Ast
		declinit *Ast
	}
	// Array initializer
	array_initializer struct {
		csize      int
		arrayinit []*Ast
	}
	// If statement
	_if struct {
		cond *Ast
		then []*Ast
		els  []*Ast
	}
}

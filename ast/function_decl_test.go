package ast

import (
	"testing"
)

func TestFunctionDecl(t *testing.T) {
	nodes := map[string]Node{
		`0x7fb5a90e60d0 <line:231:1, col:22> col:7 clearerr 'void (FILE *)'`: &FunctionDecl{
			Addr:         0x7fb5a90e60d0,
			Pos:          NewPositionFromString("line:231:1, col:22"),
			Prev:         "",
			Position2:    "col:7",
			Name:         "clearerr",
			Type:         "void (FILE *)",
			Type2:        "",
			IsExtern:     false,
			IsImplicit:   false,
			IsUsed:       false,
			IsReferenced: false,
			IsStatic:     false,
			IsInline:     false,
			ChildNodes:   []Node{},
		},
		`0x7fb5a90e2a50 </usr/include/sys/stdio.h:39:1, /usr/include/AvailabilityInternal.h:21697:126> /usr/include/sys/stdio.h:39:5 renameat 'int (int, const char *, int, const char *)'`: &FunctionDecl{
			Addr:         0x7fb5a90e2a50,
			Pos:          NewPositionFromString("/usr/include/sys/stdio.h:39:1, /usr/include/AvailabilityInternal.h:21697:126"),
			Prev:         "",
			Position2:    "/usr/include/sys/stdio.h:39:5",
			Name:         "renameat",
			Type:         "int (int, const char *, int, const char *)",
			Type2:        "",
			IsExtern:     false,
			IsImplicit:   false,
			IsUsed:       false,
			IsReferenced: false,
			IsStatic:     false,
			IsInline:     false,
			ChildNodes:   []Node{},
		},
		`0x7fb5a90e9b70 </usr/include/stdio.h:244:6> col:6 implicit fprintf 'int (FILE *, const char *, ...)' extern`: &FunctionDecl{
			Addr:         0x7fb5a90e9b70,
			Pos:          NewPositionFromString("/usr/include/stdio.h:244:6"),
			Prev:         "",
			Position2:    "col:6",
			Name:         "fprintf",
			Type:         "int (FILE *, const char *, ...)",
			Type2:        "",
			IsExtern:     true,
			IsImplicit:   true,
			IsUsed:       false,
			IsReferenced: false,
			IsStatic:     false,
			IsInline:     false,
			ChildNodes:   []Node{},
		},
		`0x7fb5a90e9d40 prev 0x7fb5a90e9b70 <col:1, /usr/include/sys/cdefs.h:351:63> /usr/include/stdio.h:244:6 fprintf 'int (FILE *, const char *, ...)'`: &FunctionDecl{
			Addr:         0x7fb5a90e9d40,
			Pos:          NewPositionFromString("col:1, /usr/include/sys/cdefs.h:351:63"),
			Prev:         "0x7fb5a90e9b70",
			Position2:    "/usr/include/stdio.h:244:6",
			Name:         "fprintf",
			Type:         "int (FILE *, const char *, ...)",
			Type2:        "",
			IsExtern:     false,
			IsImplicit:   false,
			IsUsed:       false,
			IsReferenced: false,
			IsStatic:     false,
			IsInline:     false,
			ChildNodes:   []Node{},
		},
		`0x7fb5a90ec210 <line:259:6> col:6 implicit used printf 'int (const char *, ...)' extern`: &FunctionDecl{
			Addr:         0x7fb5a90ec210,
			Pos:          NewPositionFromString("line:259:6"),
			Prev:         "",
			Position2:    "col:6",
			Name:         "printf",
			Type:         "int (const char *, ...)",
			Type2:        "",
			IsExtern:     true,
			IsImplicit:   true,
			IsUsed:       true,
			IsReferenced: false,
			IsStatic:     false,
			IsInline:     false,
			ChildNodes:   []Node{},
		},
		`0x2ae30d8 </usr/include/math.h:65:3, /usr/include/x86_64-linux-gnu/sys/cdefs.h:57:54> <scratch space>:17:1 __acos 'double (double)' extern`: &FunctionDecl{
			Addr:         0x2ae30d8,
			Pos:          NewPositionFromString("/usr/include/math.h:65:3, /usr/include/x86_64-linux-gnu/sys/cdefs.h:57:54"),
			Prev:         "",
			Position2:    "<scratch space>:17:1",
			Name:         "__acos",
			Type:         "double (double)",
			Type2:        "",
			IsExtern:     true,
			IsImplicit:   false,
			IsUsed:       false,
			IsReferenced: false,
			IsStatic:     false,
			IsInline:     false,
			ChildNodes:   []Node{},
		},
		`0x7fc595071500 <line:26:1, line:69:1> line:26:5 referenced main 'int (int, char **)'`: &FunctionDecl{
			Addr:         0x7fc595071500,
			Pos:          NewPositionFromString("line:26:1, line:69:1"),
			Prev:         "",
			Position2:    "line:26:5",
			Name:         "main",
			Type:         "int (int, char **)",
			Type2:        "",
			IsExtern:     false,
			IsImplicit:   false,
			IsUsed:       false,
			IsReferenced: true,
			IsStatic:     false,
			IsInline:     false,
			ChildNodes:   []Node{},
		},
		`0x55973a008cb0 <line:93619:1, line:93630:1> line:93619:12 used exprIsConst 'int (Expr *, int, int)' static`: &FunctionDecl{
			Addr:         0x55973a008cb0,
			Pos:          NewPositionFromString("line:93619:1, line:93630:1"),
			Prev:         "",
			Position2:    "line:93619:12",
			Name:         "exprIsConst",
			Type:         "int (Expr *, int, int)",
			Type2:        "",
			IsExtern:     false,
			IsImplicit:   false,
			IsUsed:       true,
			IsReferenced: false,
			IsStatic:     true,
			IsInline:     false,
			ChildNodes:   []Node{},
		},
		`0x563ade547cb8 <safe_math.h:13:1, line:25:1> line:14:2 safe_unary_minus_func_int8_t_s 'int8_t (int8_t)':'int8_t (int8_t)' static`: &FunctionDecl{
			Addr:         0x563ade547cb8,
			Pos:          NewPositionFromString("safe_math.h:13:1, line:25:1"),
			Prev:         "",
			Position2:    "line:14:2",
			Name:         "safe_unary_minus_func_int8_t_s",
			Type:         "int8_t (int8_t)",
			Type2:        "int8_t (int8_t)",
			IsExtern:     false,
			IsImplicit:   false,
			IsUsed:       false,
			IsReferenced: false,
			IsStatic:     true,
			IsInline:     false,
			ChildNodes:   []Node{},
		},
		`0x556cac571be0 <tests/asm.c:9:1, line:13:1> line:9:26 sqlite3Hwtime1 'unsigned long (void)' inline`: &FunctionDecl{
			Addr:         0x556cac571be0,
			Pos:          NewPositionFromString("tests/asm.c:9:1, line:13:1"),
			Prev:         "",
			Position2:    "line:9:26",
			Name:         "sqlite3Hwtime1",
			Type:         "unsigned long (void)",
			Type2:        "",
			IsExtern:     false,
			IsImplicit:   false,
			IsUsed:       false,
			IsReferenced: false,
			IsStatic:     false,
			IsInline:     true,
			ChildNodes:   []Node{},
		},
		`0x21c3da0 <line:8201:1, line:8786:1> line:8201:25 used insertvertex 'enum insertvertexresult (struct mesh *, struct behavior *, vertex, struct otri *, struct osub *, int, int)'`: &FunctionDecl{
			Addr:         0x21c3da0,
			Pos:          NewPositionFromString("line:8201:1, line:8786:1"),
			Prev:         "",
			Position2:    "line:8201:25",
			Name:         "insertvertex",
			Type:         "enum insertvertexresult (struct mesh *, struct behavior *, vertex, struct otri *, struct osub *, int, int)",
			Type2:        "",
			IsExtern:     false,
			IsImplicit:   false,
			IsUsed:       true,
			IsReferenced: false,
			IsStatic:     false,
			IsInline:     false,
			ChildNodes:   []Node{},
		},
	}

	runNodeTests(t, nodes)
}

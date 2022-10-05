package types

import "github.com/google/go-jsonnet/ast"

func prepareStdlib(g *typeGraph) {
	g.newPlaceholder()

	arrayOfString := anyArrayType
	stringOrArray := anyType
	stringOrNumber := anyType
	jsonType := anyType // It actually cannot functions anywhere

	required := func(name string) ast.Parameter {
		return ast.Parameter{Name: ast.NewIdentifier(name)}
	}

	dummyDefaultArg := &ast.LiteralNull{}
	optional := func(name string) ast.Parameter {
		return ast.Parameter{Name: ast.NewIdentifier(name), DefaultArg: dummyDefaultArg}
	}

	fields := map[string]placeholderID{

		// External variables
		"extVar": g.newSimpleFuncType(anyType, ast.NewIdentifier("x")),

		// Types and reflection
		"thisFile":        stringType,
		"type":            g.newSimpleFuncType(stringType, ast.NewIdentifier("x")),
		"length":          g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),
		"objectHas":       g.newSimpleFuncType(boolType, ast.NewIdentifier("o"), ast.NewIdentifier("f")),
		"objectFields":    g.newSimpleFuncType(arrayOfString, ast.NewIdentifier("o")),
		"objectValues":    g.newSimpleFuncType(anyArrayType, ast.NewIdentifier("o")),
		"objectHasAll":    g.newSimpleFuncType(boolType, ast.NewIdentifier("o"), ast.NewIdentifier("f")),
		"objectFieldsAll": g.newSimpleFuncType(arrayOfString, ast.NewIdentifier("o")),
		"objectValuesAll": g.newSimpleFuncType(anyArrayType, ast.NewIdentifier("o")),
		"prune":           g.newSimpleFuncType(anyObjectType, ast.NewIdentifier("a")),
		"mapWithKey":      g.newSimpleFuncType(anyObjectType, ast.NewIdentifier("func"), ast.NewIdentifier("obj")),
		"get":             g.newFuncType(anyType, []ast.Parameter{required("o"), required("f"), optional("default"), optional("inc_hidden")}),

		// isSomething
		"isArray":    g.newSimpleFuncType(boolType, ast.NewIdentifier("v")),
		"isBoolean":  g.newSimpleFuncType(boolType, ast.NewIdentifier("v")),
		"isFunction": g.newSimpleFuncType(boolType, ast.NewIdentifier("v")),
		"isNumber":   g.newSimpleFuncType(boolType, ast.NewIdentifier("v")),
		"isObject":   g.newSimpleFuncType(boolType, ast.NewIdentifier("v")),
		"isString":   g.newSimpleFuncType(boolType, ast.NewIdentifier("v")),

		// Mathematical utilities
		"abs":      g.newSimpleFuncType(numberType, ast.NewIdentifier("n")),
		"sign":     g.newSimpleFuncType(numberType, ast.NewIdentifier("n")),
		"max":      g.newSimpleFuncType(numberType, ast.NewIdentifier("a"), ast.NewIdentifier("b")),
		"min":      g.newSimpleFuncType(numberType, ast.NewIdentifier("a"), ast.NewIdentifier("b")),
		"pow":      g.newSimpleFuncType(numberType, ast.NewIdentifier("x"), ast.NewIdentifier("n")),
		"exp":      g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),
		"log":      g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),
		"exponent": g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),
		"mantissa": g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),
		"floor":    g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),
		"ceil":     g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),
		"sqrt":     g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),
		"sin":      g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),
		"cos":      g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),
		"tan":      g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),
		"asin":     g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),
		"acos":     g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),
		"atan":     g.newSimpleFuncType(numberType, ast.NewIdentifier("x")),

		// Assertions and debugging
		"assertEqual": g.newSimpleFuncType(boolType, ast.NewIdentifier("a"), ast.NewIdentifier("b")),

		// String Manipulation

		"toString":    g.newSimpleFuncType(stringType, ast.NewIdentifier("a")),
		"codepoint":   g.newSimpleFuncType(numberType, ast.NewIdentifier("str")),
		"char":        g.newSimpleFuncType(stringType, ast.NewIdentifier("n")),
		"substr":      g.newSimpleFuncType(stringType, ast.NewIdentifier("str"), ast.NewIdentifier("from"), ast.NewIdentifier("len")),
		"findSubstr":  g.newSimpleFuncType(numberArrayType, ast.NewIdentifier("pat"), ast.NewIdentifier("str")),
		"startsWith":  g.newSimpleFuncType(boolType, ast.NewIdentifier("a"), ast.NewIdentifier("b")),
		"endsWith":    g.newSimpleFuncType(boolType, ast.NewIdentifier("a"), ast.NewIdentifier("b")),
		"stripChars":  g.newSimpleFuncType(stringType, ast.NewIdentifier("str"), ast.NewIdentifier("chars")),
		"lstripChars": g.newSimpleFuncType(stringType, ast.NewIdentifier("str"), ast.NewIdentifier("chars")),
		"rstripChars": g.newSimpleFuncType(stringType, ast.NewIdentifier("str"), ast.NewIdentifier("chars")),
		"split":       g.newSimpleFuncType(arrayOfString, ast.NewIdentifier("str"), ast.NewIdentifier("c")),
		"splitLimit":  g.newSimpleFuncType(arrayOfString, ast.NewIdentifier("str"), ast.NewIdentifier("c"), ast.NewIdentifier("maxsplits")),
		"strReplace":  g.newSimpleFuncType(stringType, ast.NewIdentifier("str"), ast.NewIdentifier("from"), ast.NewIdentifier("to")),
		"asciiUpper":  g.newSimpleFuncType(stringType, ast.NewIdentifier("str")),
		"asciiLower":  g.newSimpleFuncType(stringType, ast.NewIdentifier("str")),
		"stringChars": g.newSimpleFuncType(stringType, ast.NewIdentifier("str")),
		"format":      g.newSimpleFuncType(stringType, ast.NewIdentifier("str"), ast.NewIdentifier("vals")),
		// TODO(sbarzowski) Fix when they match the documentation
		"escapeStringBash":    g.newSimpleFuncType(stringType, ast.NewIdentifier("str_")),
		"escapeStringDollars": g.newSimpleFuncType(stringType, ast.NewIdentifier("str_")),
		"escapeStringJson":    g.newSimpleFuncType(stringType, ast.NewIdentifier("str_")),
		"escapeStringPython":  g.newSimpleFuncType(stringType, ast.NewIdentifier("str")),

		// Parsing

		"parseInt":   g.newSimpleFuncType(numberType, ast.NewIdentifier("str")),
		"parseOctal": g.newSimpleFuncType(numberType, ast.NewIdentifier("str")),
		"parseHex":   g.newSimpleFuncType(numberType, ast.NewIdentifier("str")),
		"parseJson":  g.newSimpleFuncType(jsonType, ast.NewIdentifier("str")),
		"parseYaml":  g.newSimpleFuncType(jsonType, ast.NewIdentifier("str")),
		"encodeUTF8": g.newSimpleFuncType(numberArrayType, ast.NewIdentifier("str")),
		"decodeUTF8": g.newSimpleFuncType(stringType, "arr"),

		// Manifestation

		"manifestIni":          g.newSimpleFuncType(stringType, "ini"),
		"manifestPython":       g.newSimpleFuncType(stringType, "v"),
		"manifestPythonVars":   g.newSimpleFuncType(stringType, "conf"),
		"manifestTomlEx":       g.newSimpleFuncType(stringType, "value", "indent"),
		"manifestJsonEx":       g.newSimpleFuncType(stringType, "value", "indent"),
		"manifestJsonMinified": g.newSimpleFuncType(stringType, "value"),
		"manifestYamlDoc":      g.newSimpleFuncType(stringType, "value"),
		"manifestYamlStream":   g.newSimpleFuncType(stringType, "value"),
		"manifestXmlJsonml":    g.newSimpleFuncType(stringType, "value"),

		// Arrays

		"makeArray":     g.newSimpleFuncType(anyArrayType, "sz", "func"),
		"count":         g.newSimpleFuncType(numberType, "arr", "x"),
		"member":        g.newSimpleFuncType(boolType, "arr", "x"),
		"find":          g.newSimpleFuncType(numberArrayType, "value", "arr"),
		"map":           g.newSimpleFuncType(anyArrayType, "func", "arr"),
		"mapWithIndex":  g.newSimpleFuncType(anyArrayType, "func", "arr"),
		"filterMap":     g.newSimpleFuncType(anyArrayType, "filter_func", "map_func", "arr"),
		"flatMap":       g.newSimpleFuncType(anyArrayType, "func", "arr"),
		"filter":        g.newSimpleFuncType(anyArrayType, "func", "arr"),
		"foldl":         g.newSimpleFuncType(anyType, "func", "arr", "init"),
		"foldr":         g.newSimpleFuncType(anyType, "func", "arr", "init"),
		"repeat":        g.newSimpleFuncType(anyArrayType, "what", "count"),
		"slice":         g.newSimpleFuncType(arrayOfString, "indexable", "index", "end", "step"),
		"range":         g.newSimpleFuncType(numberArrayType, "from", "to"),
		"join":          g.newSimpleFuncType(stringOrArray, "sep", "arr"),
		"lines":         g.newSimpleFuncType(arrayOfString, "arr"),
		"flattenArrays": g.newSimpleFuncType(anyArrayType, "arrs"),
		"sort":          g.newFuncType(anyArrayType, []ast.Parameter{required("arr"), optional("keyF")}),
		"uniq":          g.newFuncType(anyArrayType, []ast.Parameter{required("arr"), optional("keyF")}),

		// Sets

		"set":       g.newFuncType(anyArrayType, []ast.Parameter{required("arr"), optional("keyF")}),
		"setInter":  g.newFuncType(anyArrayType, []ast.Parameter{required("a"), required("b"), optional("keyF")}),
		"setUnion":  g.newFuncType(anyArrayType, []ast.Parameter{required("a"), required("b"), optional("keyF")}),
		"setDiff":   g.newFuncType(anyArrayType, []ast.Parameter{required("a"), required("b"), optional("keyF")}),
		"setMember": g.newFuncType(boolType, []ast.Parameter{required("x"), required("arr"), optional("keyF")}),

		// Encoding

		"base64":            g.newSimpleFuncType(stringType, "input"),
		"base64DecodeBytes": g.newSimpleFuncType(numberType, ast.NewIdentifier("str")),
		"base64Decode":      g.newSimpleFuncType(stringType, ast.NewIdentifier("str")),
		"md5":               g.newSimpleFuncType(stringType, "s"),

		// JSON Merge Patch

		"mergePatch": g.newSimpleFuncType(anyType, "target", "patch"),

		// Debugging

		"trace": g.newSimpleFuncType(anyType, "str", "rest"),

		// Undocumented
		"manifestJson":     g.newSimpleFuncType(stringType, "value"),
		"objectHasEx":      g.newSimpleFuncType(boolType, "obj", "fname", "hidden"),
		"objectFieldsEx":   g.newSimpleFuncType(arrayOfString, "obj", "hidden"),
		"modulo":           g.newSimpleFuncType(numberType, "x", "y"),
		"primitiveEquals":  g.newSimpleFuncType(boolType, "x", "y"),
		"mod":              g.newSimpleFuncType(stringOrNumber, "a", "b"),
		"native":           g.newSimpleFuncType(anyFunctionType, "x"),
		"$objectFlatMerge": g.newSimpleFuncType(anyObjectType, "x"),
	}

	fieldContains := map[string][]placeholderID{}
	for name, t := range fields {
		fieldContains[name] = []placeholderID{t}
	}

	g._placeholders[stdlibType] = concreteTP(TypeDesc{
		ObjectDesc: &objectDesc{
			allFieldsKnown: true,
			unknownContain: nil,
			fieldContains:  fieldContains,
		},
	})
}

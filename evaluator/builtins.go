package evaluator

import (
	"fmt"

	"github.com/kitagry/monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			array, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `first` not supported, got %s", args[0].Type())
			}

			if len(array.Elements) > 0 {
				return array.Elements[0]
			}
			return NULL
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			array, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `last` not supported, got %s", args[0].Type())
			}

			if len(array.Elements) > 0 {
				return array.Elements[len(array.Elements)-1]
			}
			return NULL
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			array, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `rest` not supported, got %s", args[0].Type())
			}

			length := len(array.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, array.Elements[1:length])
				return &object.Array{Elements: newElements}
			}
			return NULL
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			array, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			length := len(array.Elements)
			newElements := make([]object.Object, length+1)
			copy(newElements, array.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
}

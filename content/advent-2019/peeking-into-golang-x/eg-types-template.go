package example

import (
	"context"

	"./types"
)

func before(a types.After, b types.Before, s string) { b.Execute(s) }
func after(a types.After, b types.Before, s string)  { a.Exec(context.Background(), s) }

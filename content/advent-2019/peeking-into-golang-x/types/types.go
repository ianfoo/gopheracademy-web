package types

import "context"

type Before struct{}

func (Before) Execute(string) {}

type After struct{}

func (After) Exec(context.Context, string) {}

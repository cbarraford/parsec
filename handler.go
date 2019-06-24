package parsec

import "github.com/cosmos/cosmos-sdk/types"

type Handler func(ctx Context, msg types.Msg) types.Result

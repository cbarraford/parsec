package parsec

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type Parsec struct{}

var _ = Suite(&Parsec{})

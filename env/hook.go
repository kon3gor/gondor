package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/kon3gor/gondor"
)

type EnvHook struct{}

func NewEnvHook() gondor.Hook[string] {
	return EnvHook{}
}

func (h EnvHook) Apply(v string) string {
	fmt.Println(v)
	name, ok := strings.CutPrefix(v, "env:")
	if !ok {
		fmt.Println("value has no env: prefix")
		return v
	}

	return os.Getenv(name)
}

package gondor

type Hook[T any] interface {
	Apply(T) T
}

type stringUnmarshaler struct {
	hooks []Hook[string]
}

func (u *stringUnmarshaler) unmarshall(v *string, data []byte) error {
	raw := string(data[1 : len(data)-1])
	for _, hook := range u.hooks {
		raw = hook.Apply(raw)
	}

	*v = raw

	return nil
}

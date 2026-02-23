package types

type Command interface {
	GetName() string
	Execute(params []byte) (response any, err error)
}

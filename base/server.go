package base

type IServer interface {
	Start()
	Stop() error
}

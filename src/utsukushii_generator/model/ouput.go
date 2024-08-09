package model

type OutputPrinter interface {
	Print() ([]byte, error)
}

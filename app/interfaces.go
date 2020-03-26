package app
type Processor interface {
	Process(source string) (string, error)
}


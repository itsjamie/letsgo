package main

type PublicType struct {
	Exported   string
	unexported int64
}

func (this *PublicType) GetInt() int64 {
	return this.unexported
}

func PublicApi() ([]byte, error) {}

func privateImpl() ([]byte, example.PublicType, error) {}

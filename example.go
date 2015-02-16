package example

type PublicType struct {
	Exported   string
	unexported int64
}

func (this *PublicType) GetInt() int64 {
	return this.unexported
}

func (this *PublicType) privateFunc() {}

func PublicApi() ([]byte, error) {}

func privateImpl() ([]byte, example.PublicType, error) {}

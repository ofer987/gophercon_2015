package speakers

type Speaker interface {
	SayHello() string
}

type English struct {
	Name string
}

type Chinese struct {
	Name string
}

func (sp English) SayHello() string {
	return "Hello World " + sp.Name
}

func (sp Chinese) SayHello() string {
	return "你好世界 " + sp.Name
}

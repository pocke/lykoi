package lykoi

type Component interface {
	Render() error
	UpdateRenderer(*Renderer)
}

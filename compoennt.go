package lykoi

type Component interface {
	Render(w, h int)
}

type TextArea struct {
}

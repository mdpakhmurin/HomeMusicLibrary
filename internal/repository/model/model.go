package model

type Models struct {
	Model interface{}
}

// Возвращает все используемые модели
func GetModels() []Models {
	return []Models{
		{Model: Song{}},
	}
}

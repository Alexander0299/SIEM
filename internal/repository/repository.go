package repository

type Repository struct {
	dataSource string
}

func NewRepository(dataSource string) *Repository {
	return &Repository{dataSource: dataSource}
}

func (r *Repository) GetAllItems() (string, error) {
	// Тут можешь подключить чтение из файла items.csv
	return `["item1", "item2", "item3"]`, nil
}

package models

type BookInventory struct {
	ISBN            int    `json:"isbn" gorm:"primary_key"`
	LibID           int    `json:"libid"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Publisher       string `json:"publisher"`
	Version         string `json:"version"`
	TotalCopies     uint64 `json:"totalcopies"`
	AvailableCopies uint64 `json:"availablecopies"`
}

func (b *BookInventory) CheckBook() bool {
	var bk BookInventory
	if err := DB.Where("ISBN = ?", b.ISBN).Find(&bk).Error; err != nil {
		return false
	}
	return true
}

func (b *BookInventory) NewBook() (*BookInventory, error) {
	if err := DB.Create(&b).Error; err != nil {
		return &BookInventory{}, err
	}
	return b, nil
}
func (b *BookInventory) IncreaseBook() (*BookInventory, error) {
	var bk BookInventory
	if err := DB.Where("ISBN = ?", b.ISBN).Find(&bk).Error; err != nil {
		return &BookInventory{}, err
	}
	bk.TotalCopies = bk.TotalCopies + b.TotalCopies
	bk.AvailableCopies = bk.AvailableCopies + b.AvailableCopies

	if err := DB.Where("ISBN = ?", b.ISBN).Save(&bk).Error; err != nil {
		return &BookInventory{}, err
	}
	return b, nil
}

func (b *BookInventory) DecreaseBook(no uint64) (*BookInventory, error) {
	n := no
	var bk BookInventory
	if err := DB.Where("ISBN = ?", b.ISBN).Find(&bk).Error; err != nil {
		return &BookInventory{}, err
	}
	if bk.AvailableCopies < n {
		n = n - (n - bk.AvailableCopies)
	}
	bk.TotalCopies = bk.TotalCopies - n
	bk.AvailableCopies = bk.AvailableCopies - n

	if err := DB.Where("ISBN = ?", b.ISBN).Save(&bk).Error; err != nil {
		return &BookInventory{}, err
	}
	return b, nil
}

func (b *BookInventory) UpdateDetails() (*BookInventory, error) {
	var bk BookInventory
	if err := DB.Where("ISBN = ?", b.ISBN).Find(&bk).Error; err != nil {
		return &BookInventory{}, err
	}
	bk.ISBN = b.ISBN
	bk.LibID = b.LibID
	bk.Title = b.Title
	bk.Author = b.Author
	bk.Publisher = b.Publisher
	bk.Version = b.Version
	if b.TotalCopies > 0 {
		bk.TotalCopies = bk.TotalCopies + b.TotalCopies
		bk.AvailableCopies = bk.AvailableCopies + b.AvailableCopies
	}

	if err := DB.Where("ISBN = ?", b.ISBN).Save(&bk).Error; err != nil {
		return &BookInventory{}, err
	}
	return b, nil
}

type Query struct {
	Query string `json:"query"`
}

func FindBook(query Query) ([]BookInventory, error) {
	var b []BookInventory
	if err := DB.Where("Title LIKE ? or Author LIKE ? or Publisher LIKE ?",
		"%"+query.Query+"%", "%"+query.Query+"%", "%"+query.Query+"%").Find(&b).Error; err != nil {
		return []BookInventory{}, err
	}
	return b, nil
}

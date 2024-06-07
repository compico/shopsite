package products

import "time"

type Category interface {
	GetId() int
	GetName() string
	SetName(string)
	IsSubcategory() bool
	GetParent() Category
	SetParentId(*int)
	GetChildren() Categories
	AddChildIds(...int)
}

type Categories []Category

type CategoryImpl struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	ParentId  int        `json:"parent_id"`
	Parent    Category   `json:"-"`
	Child     Categories `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (category *CategoryImpl) GetName() string {
	return category.Name
}

func (category *CategoryImpl) SetName(name string) {
	category.Name = name
}

func (category *CategoryImpl) GetId() int {
	return category.Id
}

func (category *CategoryImpl) GetParent() Category {
	return category.Parent
}

func (category *CategoryImpl) SetParentId(parentId int) {
	category.ParentId = parentId
}

func (category *CategoryImpl) GetChildren() Categories {
	return category.Child
}

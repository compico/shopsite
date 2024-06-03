package products

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

func NewCategory(name string, parentId *int, childIds ...int) Category {
	category := &CategoryImpl{}

	category.SetName(name)
	category.AddChildIds(childIds...)

	if parentId != nil {
		category.SetParentId(parentId)
		category.IsChild = true
	}

	return category
}

type CategoryImpl struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	IsChild  bool       `json:"is_child"`
	ChildIds []int      `json:"child_ids"`
	ParentId *int       `json:"parentId,omitempty"`
	Parent   Category   `json:"-"`
	Child    Categories `json:"-"`
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

func (category *CategoryImpl) IsSubcategory() bool {
	return category.IsChild
}

func (category *CategoryImpl) GetParent() Category {
	return category.Parent
}

func (category *CategoryImpl) SetParentId(parentId *int) {
	category.ParentId = parentId
}

func (category *CategoryImpl) GetChildren() Categories {
	return category.Child
}

func (category *CategoryImpl) AddChildIds(childIds ...int) {
	category.ChildIds = append(category.ChildIds, childIds...)
}

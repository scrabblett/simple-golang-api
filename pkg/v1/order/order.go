package order

type Order struct {
	BookTitle string `json:"book_title" validate:"required,min=1,max=100"`
	Count     int64  `json:"count" validate:"required,min=1,max=1000"`
}

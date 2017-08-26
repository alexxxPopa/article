package article

type Connection interface{
	isDuplicateTitle(title string) bool
}
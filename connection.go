package article

type Connection interface{
	isTitleUnique(title string) bool
}
package article



type API struct {
	conn Connection
}

func (api API) isDuplicateTitle(title string) string {

	if api.conn.isTitleUnique(title) {
		return "No good"
	}
	return "good to go"
}
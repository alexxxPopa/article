package article



type API struct {
	conn Connection
}

func (api API) findTitle() string {
	return api.conn.findTitle()
}
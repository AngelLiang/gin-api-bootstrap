package serializer

type PaginationOut struct {
	Current int `json:"current"`
	Size int    `json:"size"`
	Count int64  `json:"count"`
}

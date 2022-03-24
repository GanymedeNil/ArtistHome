package request

type UriID struct {
	ID uint64 `uri:"id" binding:"required"`
}

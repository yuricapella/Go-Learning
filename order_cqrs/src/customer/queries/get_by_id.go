package queries

type GetByID struct {
	ID int64 `uri:"id" binding:"required"`
}

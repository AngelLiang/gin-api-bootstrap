package util

type Pagination struct {
    Page      int     `form:"current"`
    PerPage   int     `form:"size"`
}

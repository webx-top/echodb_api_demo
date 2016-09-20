package handler

import (
	"fmt"
	"math"

	"github.com/webx-top/echo"
)

func Paging(c echo.Context) (page int, size int) {
	page = c.Formx(`page`).Int()
	size = c.Formx(`size`).Int()
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 1000 {
		size = 10
	}
	return
}

func JSON(c echo.Context, data interface{}, err error) error {
	a := &API{Status: 1}
	if err != nil {
		a.Status = 0
		a.Message = err.Error()
	} else if data != nil {
		a.Data = data
	}
	return c.JSON(200, a)
}

func JSONList(c echo.Context, data interface{}, cnt func() int64, page int, size int) error {
	a := &ListAPI{}
	if fmt.Sprintf(`%v`, data) == `[]` {
		data = []string{}
	}
	a.List = data
	a.Total = c.Formx(`total`).Int64()
	if a.Total < 1 {
		a.Total = cnt()
	}
	a.Page = page
	a.Pages = Pages(a.Total, size)
	if a.Pages > page {
		a.NextURL = GenNextURL(c, c.Request().URL().Path())
		t := `?`
		for key, values := range c.Queries() {
			if key == `page` || key == `size` || key == `total` || len(values) < 1 {
				continue
			}
			a.NextURL += t + key + `=` + values[0]
			t = `&`
		}
		a.NextURL += t + fmt.Sprintf(`page=%d&size=%d&total=%d`, page+1, size, a.Total)
	}
	return JSON(c, a, nil)
}

func SiteURL(c echo.Context) string {
	scheme := c.Request().Scheme()
	if len(scheme) == 0 {
		scheme = "http"
	}
	return scheme + "://" + c.Request().Host() + "/"
}

func GenNextURL(c echo.Context, urlPath string, params ...interface{}) string {
	fullURL := SiteURL(c) + urlPath
	if len(params) == 0 {
		return fullURL
	}
	return fmt.Sprintf(fullURL, params...)
}

func PageOffset(page, size int) int {
	if page < 1 {
		page = 1
	}
	return (page - 1) * size
}

func Pages(total int64, size int) int {
	return int(math.Ceil(float64(total) / float64(size)))
}

type API struct {
	Status  int         `json:"status" xml:"status"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

type ListAPI struct {
	List    interface{} `json:"list" xml:"list"`
	Pages   int         `json:"pages" xml:"pages"`
	Page    int         `json:"page" xml:"page"`
	Total   int64       `json:"total" xml:"total"`
	NextURL string      `json:"nextURL,omitempty" xml:"nextURL,omitempty"`
}

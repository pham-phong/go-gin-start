package database

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Paginator struct {
	// TotalPage   int         `json:"total_page"`
	Total        int         `json:"total"`
	Data         interface{} `json:"data"`
	Current_page int         `json:"current_page"`
	Last_page    int         `json:"last_page"`
	Next         int         `json:"next"`
	Prev         int         `json:"prev"`
	Links        struct {
		Next int `json:"next"`
		Prev int `json:"prev"`
	} `json:"link"`
	// PrevPage    int         `json:"prev_page"`
	// NextPage    int         `json:"next_page"`
}

func Paginate(c *gin.Context, result interface{}) *Paginator {
	db := c.MustGet("db").(*gorm.DB)
	var total int
	var paginator Paginator

	//Count all Record
	db.Model(result).Count(&total)
	//Get page
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}
	offset := (page - 1) * limit

	data := db.Offset(offset).Limit(limit).Find(result)

	paginator.Data = data.Value
	paginator.Total = total
	paginator.Current_page = page
	paginator.Last_page = int(math.Ceil(float64(total) / float64(limit)))

	if paginator.Current_page > 1 {
		paginator.Prev = paginator.Current_page - 1
	} else {
		paginator.Prev = paginator.Current_page
	}

	if paginator.Current_page == paginator.Total {
		paginator.Next = paginator.Current_page
	} else {
		paginator.Next = paginator.Current_page + 1
	}
	paginator.Links.Next = paginator.Next
	paginator.Links.Prev = paginator.Prev

	return &paginator
}

func Paginate1(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

		if page < 1 {
			page = 1
		}
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit

		result := db.Offset(offset).Limit(limit)
		return result
	}
}

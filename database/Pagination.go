package database

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Paginator struct {
	Data interface{} `json:"data"`
	// Total        int         `json:"total"`
	// Current_page int         `json:"current_page"`
	// Last_page    int         `json:"last_page"`
	// Next         int         `json:"next"`
	// Prev         int         `json:"prev"`
	Links struct {
		Next int `json:"next"`
		Prev int `json:"prev"`
	} `json:"links"`
	Meta struct {
		Total        int `json:"total"`
		Current_page int `json:"current_page"`
		Last_page    int `json:"last_page"`
	} `json:"meta"`
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
	paginator.Meta.Total = total
	paginator.Meta.Current_page = page
	paginator.Meta.Last_page = int(math.Ceil(float64(total) / float64(limit)))

	if paginator.Meta.Current_page > 1 {
		paginator.Links.Prev = paginator.Meta.Current_page - 1
	} else {
		paginator.Links.Prev = paginator.Meta.Current_page
	}

	if paginator.Meta.Current_page == paginator.Meta.Last_page {
		paginator.Links.Next = paginator.Meta.Current_page
	} else {
		paginator.Links.Next = paginator.Meta.Current_page + 1
	}

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

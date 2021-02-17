package models

// import (
// 	"net/http"

// 	"github.com/jinzhu/gorm"
// )

// func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.ModelStruct.DB) *gorm.DB {
// 		page, _ := strconv.Atoi(c.Param("page"))
// 		if page == 0 {
// 			page = 1
// 		}

// 		pageSize, _ := strconv.Atoi(c.Param("page_size"))
// 		switch {
// 		case pageSize > 100:
// 			pageSize = 100
// 		case pageSize <= 0:
// 			pageSize = 10
// 		}

// 		offset := (page - 1) * pageSize
// 		return db.Offset(offset).Limit(pageSize)
// 	}
// }

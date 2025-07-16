package dao

import "go-mall/repository/db/model"

func migrate() (err error) {
	//err = db.Set("gorm:table_options", "charset=utf8mb4").
	//	AutoMigrate(&model.User{}, &model.Favorite{},
	//		&model.Order{}, &model.Admin{}, &model.Address{},
	//		&model.Cart{}, &model.Category{}, &model.Carousel{},
	//		&model.Notice{}, &model.Notice{}, &model.Product{},
	//		&model.ProductImg{}, &model.SkillProduct{},
	//	)
	err = db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{})

	return
}

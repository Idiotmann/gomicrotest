package model

type Category struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	//指定在json序列化之后的key值 而不希望用大写开始的Name  根据json tag进行结构体赋值
	CategoryName        string `gorm:"unique_index,not_null" json:"category_name"`
	CategoryLevel       uint32 `json:"category_level"`
	CategoryParent      int64  `json:"category_parent"`
	CategoryImage       string `json:"category_image"`
	CategoryDescription string `json:"category_description"`
}

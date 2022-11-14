package repository

import (
	"github.com/Idiotmann/gomicrotest/domain/model"
	"github.com/jinzhu/gorm"
)

// ICategoryRepository 实现与数据库进行交互的接口
type ICategoryRepository interface {
	InitTable() error
	FindCategoryByID(int64) (*model.Category, error)
	CreateCategory(*model.Category) (int64, error)
	DeleteCategoryByID(int64) error
	UpdateCategory(*model.Category) error
	FindAll() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

// NewCategoryRepository 实例化
func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb: db}
}

type CategoryRepository struct {
	mysqlDb *gorm.DB //mysql数据库
}

// InitTable 初始化表
func (u *CategoryRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Category{}).Error
}

// FindCategoryByID 根据ID查找Category信息
func (u *CategoryRepository) FindCategoryByID(categoryID int64) (category *model.Category, err error) {
	category = &model.Category{}
	return category, u.mysqlDb.First(category, categoryID).Error
}

// CreateCategory 创建Category信息
func (u *CategoryRepository) CreateCategory(category *model.Category) (int64, error) {
	return category.ID, u.mysqlDb.Create(category).Error
}

// DeleteCategoryByID 根据ID删除Category信息
func (u *CategoryRepository) DeleteCategoryByID(categoryID int64) error {
	return u.mysqlDb.Where("id = ?", categoryID).Delete(&model.Category{}).Error
}

// UpdateCategory 更新Category信息
func (u *CategoryRepository) UpdateCategory(category *model.Category) error {
	return u.mysqlDb.Model(category).Update(category).Error
}

// FindAll 获取结果集
func (u *CategoryRepository) FindAll() (categoryAll []model.Category, err error) {
	return categoryAll, u.mysqlDb.Find(&categoryAll).Error
}

// FindCategoryByName 实现接口
func (u *CategoryRepository) FindCategoryByName(categoryName string) (category *model.Category, err error) {
	return category, u.mysqlDb.Where("name = ?", categoryName).Find(category).Error
}
func (u *CategoryRepository) FindCategoryByLevel(categoryLevel uint32) (categorySlice []model.Category, err error) {
	return categorySlice, u.mysqlDb.Where("level = ?", categoryLevel).Find(categorySlice).Error
}
func (u *CategoryRepository) FindCategoryByParent(categoryParent int64) (categorySlice []model.Category, err error) {
	return categorySlice, u.mysqlDb.Where("parent = ?", categoryParent).Find(categorySlice).Error
}

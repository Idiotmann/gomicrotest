package service

//实现数据库增删改查
import (
	"github.com/Idiotmann/gomicrotest/domain/model"
	"github.com/Idiotmann/gomicrotest/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(*model.Category) (int64, error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(int64) (*model.Category, error)
	FindAllCategory() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

// NewCategoryDataService 实例化
func NewCategoryDataService(categoryRepository repository.ICategoryRepository) ICategoryDataService {
	return &CategoryDataService{categoryRepository}
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}

// AddCategory 插入
func (u *CategoryDataService) AddCategory(category *model.Category) (int64, error) {
	return u.CategoryRepository.CreateCategory(category)
}

// DeleteCategory 删除
func (u *CategoryDataService) DeleteCategory(categoryID int64) error {
	return u.CategoryRepository.DeleteCategoryByID(categoryID)
}

// UpdateCategory 更新
func (u *CategoryDataService) UpdateCategory(category *model.Category) error {
	return u.CategoryRepository.UpdateCategory(category)
}

// FindCategoryByID 查找
func (u *CategoryDataService) FindCategoryByID(categoryID int64) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByID(categoryID)
}

// FindAllCategory 查找
func (u *CategoryDataService) FindAllCategory() ([]model.Category, error) {
	return u.CategoryRepository.FindAll()
}

func (u *CategoryDataService) FindCategoryByName(categoryName string) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByName(categoryName)
}
func (u *CategoryDataService) FindCategoryByLevel(categoryLevel uint32) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByLevel(categoryLevel)
}
func (u *CategoryDataService) FindCategoryByParent(categoryParent int64) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByParent(categoryParent)
}

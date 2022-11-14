package handler

import (
	"context"
	"github.com/Idiotmann/gomicrotest/common"
	"github.com/Idiotmann/gomicrotest/domain/model"
	"github.com/Idiotmann/gomicrotest/domain/service"
	pb "github.com/Idiotmann/gomicrotest/proto/category"
	"github.com/prometheus/common/log"
)

//实现proto文件中定义的rpc服务接口

type Category struct {
	CategoryDataService service.ICategoryDataService
}

// CreateCategory 提供创建分类的服务
func (c *Category) CreateCategory(ctx context.Context, request *pb.CategoryRequest, response *pb.CreateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category) //通过json tag进行结构体赋值
	if err != nil {
		return err
	}
	categoryID, err := c.CategoryDataService.AddCategory(category)
	if err != nil {
		return err
	}
	response.Message = "Create category success"
	response.CategoryId = categoryID
	return nil
}
func (c *Category) UpdateCategory(ctx context.Context, request *pb.CategoryRequest, response *pb.UpdateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category) //通过json tag进行结构体赋值，不然要一个一个赋值
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(category)
	if err != nil {
		return err
	}
	response.Message = "Update category success"
	return nil
}
func (c *Category) DeleteCategory(ctx context.Context, request *pb.DeleteCategoryRequest, response *pb.DeleteCategoryResponse) error {

	err := c.CategoryDataService.DeleteCategory(request.CategoryId)
	if err != nil {
		return err
	}
	response.Message = "Delete category success"
	return nil
}
func (c *Category) FindCategoryByName(ctx context.Context, request *pb.FindByNameRequest, response *pb.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)

}
func (c *Category) FindCategoryByID(ctx context.Context, request *pb.FindByIdRequest, response *pb.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByID(request.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)
}

func (c *Category) FindCategoryByLevel(ctx context.Context, request *pb.FindByLevelRequest, response *pb.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByLevel(request.Level)
	if err != nil {
		return err
	}
	CategorySliceToResponse(categorySlice, response)
	return nil
}

func (c *Category) FindCategoryByParent(ctx context.Context, request *pb.FindByParentRequest, response *pb.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByParent(request.ParentId)
	if err != nil {
		return err
	}
	CategorySliceToResponse(categorySlice, response)
	return nil
}

// FindAllCategory 提供查询所有分类的服务
func (c *Category) FindAllCategory(ctx context.Context, request *pb.FindAllRequest, response *pb.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	CategorySliceToResponse(categorySlice, response)
	return nil
}

// CategorySliceToResponse 返回值categorySlice是切片，不能直接common.SwapTo 需要转化一下pb使用的类型
func CategorySliceToResponse(categorySlice []model.Category, response *pb.FindAllResponse) {
	for _, cg := range categorySlice {
		cr := &pb.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Fatal(err)
			break
		}
		response.Category = append(response.Category, cr)
	}
}

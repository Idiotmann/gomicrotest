package main

import (
	"github.com/Idiotmann/gomicrotest/common"
	"github.com/Idiotmann/gomicrotest/domain/repository"
	service2 "github.com/Idiotmann/gomicrotest/domain/service"
	"github.com/Idiotmann/gomicrotest/handler"
	category "github.com/Idiotmann/gomicrotest/proto/category"
	"github.com/go-micro/plugins/v4/registry/consul"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"
)

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "micro/config")
	if err != nil {
		log.Fatal(err)
	}
	//注册中心
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8082"), //服务启动的地址
		micro.Registry(consulReg),       //注册中心
	)
	//获取mysql配置,路径中不带前缀
	//mysql需要手动加载数据库驱动
	mysqlConfig, err := common.GetMysqlFromConsul(consulConfig, "mysql")
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open("mysql", mysqlConfig.User+":"+mysqlConfig.Password+"@/"+mysqlConfig.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SingularTable(true) //禁用表名复数
	// Initialise service
	service.Init()
	//因为有service服务，所以给自己写的起了别名
	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))
	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{CategoryDataService: categoryDataService})
	if err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

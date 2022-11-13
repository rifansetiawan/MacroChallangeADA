// package brick_auth

// import (
// 	"fmt"

// 	"github.com/sirupsen/logrus"
// 	"gorm.io/gorm"
// )

// type Repository interface {
// 	FindAll() ([]Stock, error)
// 	FindStockByCode(code string) (Stock, error)
// 	FindStocksName() ([]StockNameFormatter, error)
// 	FindStocksCode() ([]StockCodeFormatter, error)
// }

// type repository struct {
// 	db *gorm.DB
// }

// func NewRepository(db *gorm.DB) *repository {
// 	return &repository{db}
// }

// func (r *repository) FindAll() ([]Stock, error) {
// 	var stocks []Stock
// 	err := r.db.Find(&stocks).Error
// 	if err != nil {
// 		logrus.Info("something went wrong")
// 		return stocks, err
// 	}
// 	return stocks, nil
// }

// func (r *repository) FindStockByCode(code string) (Stock, error) {
// 	var stock Stock
// 	err := r.db.Where("code = ?", code).Find(&stock).Error
// 	if err != nil {
// 		logrus.Info("something went wrong")
// 		return stock, err
// 	}
// 	return stock, nil
// }

// func (r *repository) FindStocksName() ([]StockNameFormatter, error) {
// 	var stocks []StockNameFormatter
// 	err := r.db.Table("stocks").Select("uuid,name").Find(&stocks).Error
// 	if err != nil {
// 		logrus.Info("something went wrong")
// 		return stocks, err
// 	}
// 	return stocks, nil
// }

// func (r *repository) FindStocksCode() ([]StockCodeFormatter, error) {
// 	var stocks []StockCodeFormatter
// 	err := r.db.Table("stocks").Select("uuid,code,name").Find(&stocks).Error
// 	fmt.Println(stocks[1])
// 	if err != nil {
// 		logrus.Info("something went wrong")
// 		return stocks, err
// 	}
// 	return stocks, nil
// }

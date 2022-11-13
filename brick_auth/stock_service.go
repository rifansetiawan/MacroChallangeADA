package brick_auth

import "fmt"

type Service interface {
	FindStocks() ([]Stock, error)
	FindStocksName() ([]StockNameFormatter, error)
	FindStocksCode() ([]StockCodeFormatter, error)
}

func (s *service) FindStocks() ([]Stock, error) {
	stocks, err := s.repository.FindAll()
	if err != nil {
		return stocks, err
	}

	return stocks, nil
}

func (s *service) FindStocksName() ([]StockNameFormatter, error) {
	stocks, err := s.repository.FindStocksName()
	if err != nil {
		return stocks, err
	}

	return stocks, nil
}

func (s *service) FindStocksCode() ([]StockCodeFormatter, error) {
	stocks, err := s.repository.FindStocksCode()
	fmt.Println(stocks[1])
	if err != nil {
		return stocks, err
	}

	return stocks, nil
}

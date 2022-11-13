package brick_auth

type ResponseFormatter struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    DataResponse `json:"data"`
}

type DataResponseFormatter struct {
	AccessToken  string `json:"access_token"`
	PrimaryColor string `json:"primary_color"`
}

func FormatStock(stock Stock) StockFormatter {
	stockFormatter := StockFormatter{}
	stockFormatter.UUID = stock.UUID
	stockFormatter.Name = stock.Name
	stockFormatter.Code = stock.Code
	stockFormatter.HeadOffice = stock.HeadOffice
	stockFormatter.Phone = stock.Phone
	stockFormatter.RepresentativeName = stock.RepresentativeName
	stockFormatter.WebsiteURL = stock.WebsiteURL
	stockFormatter.Address = stock.Address
	stockFormatter.TotalEmployees = stock.TotalEmployees
	stockFormatter.IsActive = stock.IsActive
	stockFormatter.StockSubSectorID = stock.StockSubSectorID
	stockFormatter.LastStockDataID = stock.LastStockDataID
	stockFormatter.ExchangeAdministration = stock.ExchangeAdministration
	stockFormatter.NPWP = stock.NPWP
	stockFormatter.NPKP = stock.NPKP
	stockFormatter.ListingDate = stock.ListingDate
	stockFormatter.AnnualDividend = stock.AnnualDividend
	stockFormatter.BoardRecording = stock.BoardRecording
	stockFormatter.GeneralInformation = stock.GeneralInformation
	stockFormatter.Fax = stock.Fax
	stockFormatter.FoundingDate = stock.FoundingDate
	stockFormatter.CompanyEmail = stock.CompanyEmail
	stockFormatter.ListedShare = stock.ListedShare
	stockFormatter.NewSubIndustryId = stock.NewSubIndustryId
	stockFormatter.CreatedAt = stock.CreatedAt
	stockFormatter.UpdatedAt = stock.UpdatedAt

	return stockFormatter
}

func FormatStocks(stocks []Stock) []StockFormatter {
	var stocksFormatter []StockFormatter
	for _, stock := range stocks {
		stockFormatter := FormatStock(stock)
		stocksFormatter = append(stocksFormatter, stockFormatter)
	}
	return stocksFormatter
}

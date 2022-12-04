package user

import "time"

type User struct {
	UUID           string
	UserName       string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	RegistrationId string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Session struct {
	Username  string `json:"username"`
	UniqueID  string `json:"uniqueId"`
	SessionID string `json:"sessionId"`
	OtpToken  string `json:"otpToken"`
}

type SessionPayload struct {
	Username  string `json:"username"`
	UniqueID  string `json:"uniqueId"`
	SessionID string `json:"sessionId"`
	OtpToken  string `json:"otpToken"`
	Otp       int    `json:"otp"`
}

type AccountListTransactionResponsePerAccessToken struct {
	Status        int                          `json:"status"`
	Message       string                       `json:"message"`
	LastUpdatedAt string                       `json:"lastUpdateAt"`
	Session       string                       `json:"session"`
	Data          []DataAccountListTransaction `json:"data"`
}
type DataAccountListTransaction struct {
	AccountId     string                         `json:"accountId"`
	AccountHolder string                         `json:"accountHolder"`
	AccountNumber string                         `json:"accountNumber"`
	Balances      BalancesAccountListTransaction `json:"balances"`
	Currency      string                         `json:"currency"`
	Type          string                         `json:"type"`
	InstitutionID int                            `json:"institution_id"`
	Transactions  []ListTransactionOnly          `json:"transactions"`
}

type BalancesAccountListTransaction struct {
	Available float64 `json:"available"`
	Current   float64 `json:"current"`
	Limit     float64 `json:"limit"`
}

type ErrorResponseAccountListTransaction struct {
	Status        int    `json:"status"`
	Message       string `json:"message"`
	Data          string `json:"data"`
	InstitutionID int    `json:"institution_id"`
}

type MergedAccountListTransaction struct {
	Status    int                                   `json:"status"`
	Message   string                                `json:"message"`
	Data      []DataAccountListTransaction          `json:"data"`
	DataError []ErrorResponseAccountListTransaction `json:"errorData"`
}

type ListTransactionFromTo struct {
	Status        int                   `json:"status"`
	Message       string                `json:"message"`
	LastUpdatedAt string                `json:"lastUpdateAt"`
	Session       string                `json:"session"`
	Data          []ListTransactionOnly `json:"data"`
}

type ListTransactionOnly struct {
	DateTimestamp     string   `json:"dateTimestamp"`
	Id                int      `json:"id"`
	AccountId         string   `json:"account_id"`
	AccountNumber     string   `json:"account_number"`
	AccountCurrencty  string   `json:"account_currency"`
	InstitutionID     int      `json:"institution_id"`
	MerchantId        int      `json:"merchant_id"`
	OutletOutletId    int      `json:"outlet_outlet_id"`
	LocationCityID    int      `json:"location_city_id"`
	LocationCounterId int      `json:"location_country_id"`
	Date              string   `json:"date"`
	Amount            float64  `json:"amount"`
	Description       string   `json:"description"`
	Status            string   `json:"status"`
	Direction         string   `json:"direction"`
	ReferenceId       string   `json:"reference_id"`
	Category          Category `json:"category"`
	TransactionType   string   `json:"transaction_type"`
}

type Category struct {
	CategoryId               int    `json:"category_id"`
	CategoryName             string `json:"category_name"`
	ClassificationGroupId    int    `json:"classification_group_id"`
	ClassificationGroup      string `json:"classification_group"`
	ClassificationSubgroupId int    `json:"classification_subgroup_id"`
	ClassificationSubgroup   string `json:"classification_subgroup"`
}

type AccountListTransactionResponsePerAccessTokenScheduler struct {
	Status        int                   `json:"status"`
	Message       string                `json:"message"`
	LastUpdatedAt string                `json:"lastUpdateAt"`
	Session       string                `json:"session"`
	Data          []ListTransactionOnly `json:"data"`
}

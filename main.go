package main

import (
	"encoding/json"
	"fmt"
	"kaia/auth"
	"kaia/handler"
	"kaia/helper"
	"kaia/user"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func scheduler() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/kaia?charset=utf8mb4&parseTime=True&loc=Local"
	dateTomorrow := time.Now().AddDate(0, 0, 10).Format("2006-01-02")
	dateYesterday := time.Now().AddDate(0, 0, -10).Format("2006-01-02")
	fmt.Println(dateTomorrow)
	fmt.Println(dateYesterday)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	accessTokens, _ := userRepository.GetAllAccessTokens()

	for _, accTokenOne := range accessTokens {
		var responseAccountListTransactionInterface user.AccountListTransactionResponsePerAccessTokenScheduler

		clientAccountListTransactions := http.Client{}
		url := fmt.Sprintf("https://api.onebrick.io/v1/transaction/list?from=%s&to=%s", dateYesterday, dateTomorrow)
		reqAccountListTransaction, err := http.NewRequest("GET", url, nil)
		if err != nil {
			//Handle Error
		}

		reqAccountListTransaction.Header = http.Header{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer " + accTokenOne.AccessToken},
		}

		responseAccountListTransaction, err := clientAccountListTransactions.Do(reqAccountListTransaction)

		defer responseAccountListTransaction.Body.Close()

		// if responseAccountListTransaction.StatusCode = 200 {
		// 	err = json.NewDecoder(resGetInstList.Body).Decode(&respDataSessionError)
		// 	return respDataSessionError, nil
		// }
		json.NewDecoder(responseAccountListTransaction.Body).Decode(&responseAccountListTransactionInterface)
		// fmt.Println(responseAccountListTransactionInterface)

		// if transactions != nil {
		// 	transactions = append(transactions, responseAccountListTransactionInterface.Data[panjangData-1])
		// 	fmt.Println("aku sudah terisiiii : ", transactions)
		// }
		lastTransactions, err := userRepository.GetLastTransactions(accTokenOne.UserID, accTokenOne.UserEmail)
		fmt.Println(lastTransactions.Amount)
		panjangData := len(responseAccountListTransactionInterface.Data)
		fmt.Println("panjangData : ", panjangData)
		// fmt.Println(responseAccountListTransactionInterface.Data[panjangData-1].Amount)
		if responseAccountListTransactionInterface.Data != nil {
			if lastTransactions.Amount != responseAccountListTransactionInterface.Data[panjangData-1].Amount &&
				lastTransactions.Description != responseAccountListTransactionInterface.Data[panjangData-1].Description {
				fmt.Println("DONT notify meee PLEASEEE")
				fmt.Println("lastTransactions.Amount : ", lastTransactions.Amount)
				fmt.Println("responseAccountListTransactionInterface.Data[panjangData-1].Amount : ", responseAccountListTransactionInterface.Data[panjangData-1].Amount)

				var response interface{}
				var userSaya, err = userRepository.FindByUUID(accTokenOne.UserID)
				payloadString := fmt.Sprintf(`{
					"registration_ids": ["%s"],
					"notification": {
						"body": "New Transaction out %d",
						"title": "There is new transaction out!"
					},
					"data": {
						"url": "https://www.google.com",
						"dl": "kaia://"
					}
				}`, userSaya.RegistrationId, int(responseAccountListTransactionInterface.Data[panjangData-1].Amount))
				fmt.Println("regis ID : ", userSaya)
				payload := strings.NewReader(payloadString)
				clientFirebase := http.Client{}
				reqClientFirebase, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", payload)
				if err != nil {
					//Handle Error
				}

				reqClientFirebase.Header = http.Header{
					"Content-Type":  {"application/json"},
					"Authorization": {"key=AAAAunUoSKc:APA91bHwKFpoTCHkBqDgbtrHYV1VR2yG35tQ_cx8s_R6zzK-0q59THtnoljJZkXJwbDyx2ncTgyqPyr0H690P26LC9ZR-I9HP9OkBTd2mbj52a--vrf-qxY4PGgmmmyNoJGjOZ-6IBcd"},
				}
				fmt.Println(payload)
				resGetInstList, err := clientFirebase.Do(reqClientFirebase)
				json.NewDecoder(resGetInstList.Body).Decode(&response)
				fmt.Println("this is response firebase : ", response)
				userRepository.SaveLastTransactions(userSaya.UUID, userSaya.Email, userSaya.UserName, responseAccountListTransactionInterface.Data[panjangData-1].Amount, responseAccountListTransactionInterface.Data[panjangData-1].Description)

			}
		}

		// if len(transactions) != 0 && responseAccountListTransactionInterface.Data[panjangData-1]
	}

	/*
		1. Get Access Tokens
		2. Loop transaction in it
		3. buat array X kosong untuk diappend transaction [0]
		4. ambil transaction 0 dan append ke Array X
		5. []

		contoh :
		trx[]
		trxZero = 10000


		append trx[1000]----->diappend jika trx[] nya itu kosong, kalau ada isinya , diganti dengan yang baru
		jika trx[0].value == trxZerro -> dont notify

	*/
}

func task() {
	fmt.Println("I am running task.")
}
func main() {
	// c := cron.New()
	// c.AddFunc("@every 1s", scheduler)
	// c.Start()
	dsn := "root:1234@tcp(127.0.0.1:3306)/kaia?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn server
	// dsn := "rifan:1234@tcp(135.148.157.241:3306)/pasardanamobile?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	gocron.Every(1).Second().Do(task)

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	gocron.Every(1).Second().Do(task)

	router := gin.Default()
	addr := ":3333"
	router.Use(cors.Default())
	router.Static("/images", "./images")

	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))
	router.Use(sessions.Sessions("pasardanamobile", cookieStore))

	api := router.Group("/api/v1")

	//Users Route
	api.POST("/register-user", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	//Auth to get public token
	api.GET("/auth/token", authMiddleware(authService, userService), userHandler.AuthToken)

	//Auth to connect Bank Account or any other financials account
	api.POST("/auth", authMiddleware(authService, userService), userHandler.AuthTokenToAccessToken)

	//Auth to connect Gojek
	api.POST("/auth/gopay", authMiddleware(authService, userService), userHandler.AuthTokenToAccessTokenGopay)
	api.POST("/auth/gopay/token", authMiddleware(authService, userService), userHandler.OTPSessionToToken)

	//Account List
	api.GET("/account/list/:start/:end", authMiddleware(authService, userService), userHandler.AccountListTransactions)

	api.POST("/device-token", authMiddleware(authService, userService), userHandler.SaveDeviceToken)

	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	router.Run(addr)

	//CHRON JOBS

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := claim["user_id"].(string)

		user, err := userService.GetUserByUUID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

func authAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userIDSession := session.Get("userID")

		if userIDSession == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

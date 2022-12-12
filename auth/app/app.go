package app

func Start() {
	sanityCheck()

	router := gin.Default()

	authRepository := domain.NewAuthRepository(getDBClient())

	ah := AuthHandler{service: service.NewLoginService(authRepository, domain.GetRolePermissions()}

	router.POST("/auth/login", ah.Login)
	router.POST("/auth/register", ah.NotImplementedHanlder)
	router.GET("/auth/verify", ah.Verify)

	router.Run(":8011")
}

func getDBClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/ibank")
	if err != nil {
		panic(err)
	}
	// See Important settings section
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}

func sanityCheck() {
	
}

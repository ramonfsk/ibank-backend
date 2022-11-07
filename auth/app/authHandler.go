package app

func (ah *AuthHandler) NotImplementedHanlder(c *gin.Context) {
	fmt.Fprint(c, "Hanlder not implemented...")
}

func (ah *AuthHandler) Login(c *gin.Context) {

}

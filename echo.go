package main 

func main() {
	//instanciar echo 
	e := echo.New()

	//ruta
	e.GET("/", func(c echo.Context)) error {
		return c.String(http.StatusOK, "Hola a tod@s")   

	})
	e.Logger.Fatal(e.Start(":1323")) 

}
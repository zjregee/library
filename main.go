package main

import (
	"github.com/labstack/echo/v4"
	"library/controller"
)

func main() {
	e := echo.New()

	login := e.Group("/login")
	login.POST("/login", controller.Login)
	login.POST("/register", controller.Register)
	login.POST("/password", controller.ChangePassword)

	personal := e.Group("/personal")
	personal.POST("/information", controller.GetPersonalInformation)
	personal.POST("/setinformation", controller.SetPersonalInformation)
	personal.POST("/getborrow", controller.GetPersonalBorrow)
	personal.POST("/getfavoritenum", controller.GetPersonalFavoriteNum)
	personal.POST("/getfavorite", controller.GetPersonalFavorite)

	book := e.Group("/book")
	book.POST("/borrow", controller.BorrowBook)
	book.POST("/return", controller.ReturnBook)
	book.POST("/borrowtime", controller.BorrowBookTime)
	book.POST("/borrownum", controller.BorrowBookNum)
	book.POST("/information", controller.GetBookBasicInformation)
	book.POST("/specificinformation", controller.GetBookSpecificInformation)

	favorite := e.Group("/favorite")
	favorite.POST("/isfavorite", controller.BookIsFavorite)
	favorite.POST("/changefavorite", controller.BookChangeFavorite)

	activity := e.Group("/activity")
	activity.GET("/hot", controller.GerHotActivity)
	activity.GET("/new", controller.GetNewActivity)
	activity.POST("/information", controller.GetActivityInformation)
	activity.POST("/apply", controller.ApplyActivity)

	home := e.Group("/home")
	home.GET("/recommend", controller.GetHotBook)
	home.GET("/new", controller.GetNewBook)

	search := e.Group("/search")
	search.POST("/books", controller.SearchBooks)
	search.POST("/lost", controller.SearchLost)
	search.POST("/activity", controller.SearchActivity)

	lost := e.Group("/lost")
	lost.POST("/getall", controller.GetAllLost)
	lost.POST("/information", controller.GetLostInformation)

	e.Logger.Fatal(e.Start(":8080"))
}
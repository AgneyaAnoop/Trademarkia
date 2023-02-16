package main

import (
	"Trademarkia/pkg/Storage"
	"Trademarkia/pkg/Utils"
	"Trademarkia/pkg/Models"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"github.com/Valgard/godotenv"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (r *Repository) GetProceedingInfo(c *fiber.Ctx) error {
	var proceedingInfo []Models.ProceedingInfo
	r.DB.Find(&proceedingInfo)
	return c.JSON(proceedingInfo)
}

func (r *Repository) GetProceedingInfoById(c *fiber.Ctx) error {
	id := c.Params("id")
	var proceedingInfo Models.ProceedingInfo
	r.DB.First(&proceedingInfo, id)
	return c.JSON(proceedingInfo)
}


type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/ProceedingInfo", r.GetProceedingInfo)
	
	api.Get("/ProceedingInfo/:id", r.GetProceedingInfo)
	

}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	file, err := ioutil.ReadFile("tt230101.xml") //Reading the File
	if err != nil {
		fmt.Println("File reading error", err)
		panic(err)
	}
	Utils.Converter(file)

	config := &Storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := Storage.NewConnection(config)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Models.ProceedingInfo{})


	if err != nil {
		log.Fatal("could not migrate")
	}



	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":3000")

}

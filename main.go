package main

import (
	"backend-go/models"
	"backend-go/storage"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type BadanUsaha struct {
	NamaBadanUsaha string `json:"nama_badan_usaha"`
	NoTelp         string `json:"no_telp"`
	Email          string `json:"email"`
}

func (r *Repository) CreateBadanUsaha(context *fiber.Ctx) error {
	book := BadanUsaha{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create badan usaha"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "badan usaha has been added"})
	return nil
}

func (r *Repository) GetBadanUsaha(context *fiber.Ctx) error {
	badanUsahaModels := &[]models.BadanUsaha{}

	err := r.DB.Find(badanUsahaModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get badan usaha"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "badan usaha fetched successfully",
		"data":    badanUsahaModels,
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_badan_usaha", r.CreateBadanUsaha)
	api.Get("/badan_usaha", r.GetBadanUsaha)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}
	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")

}

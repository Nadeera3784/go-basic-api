package controllers

import (
	"github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v2"
	"application/src/models"
	"application/src/services"
)

func GetAll(context *fiber.Ctx) error {
	books, err := services.GetAll();
	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    books,
		"message": "Books returned successfully!",
	});
}


func GetById(context *fiber.Ctx) error {
	// Get the book ID from the URL parameters
	bookID := context.Params("id")

	// Call the FindById function from the services package
	book, findErr := services.GetById(bookID)
	if findErr != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": findErr.Error()})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{"data": book, "message": "Book retrieved successfully"})
}

func Create(context *fiber.Ctx,) error {
	var book models.Book;
	if err := context.BodyParser(&book); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var validate = validator.New()

	if err := validate.Struct(&book); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			// If it's not a validation error, return a generic error response
			return context.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
		}

		// Map validation errors to a more human-readable format
		errorDetails := make(map[string]string)
		for _, e := range validationErrors {
			errorDetails[e.Field()] = e.Tag()
		}

		return context.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": "Validation failed", "errors": errorDetails})
	}

	data, createErr := services.Create(book)

	if createErr != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": createErr.Error()})
	}
	return context.Status(fiber.StatusCreated).JSON(fiber.Map{"data": data, "message": "Book created successfully"})
}


func Update(context *fiber.Ctx) error {
	// Get the book ID from the URL parameters
	bookID := context.Params("id")

	var updatedBook models.Book
	if err := context.BodyParser(&updatedBook); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	validate := validator.New()

	err := validate.Struct(updatedBook);
	if err != nil {
		return context.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": "Validation failed", "data": err.Error()})
	}

	// Call the Update function from the services package, passing the book ID
	updatedBook, updateErr := services.Update(bookID, updatedBook)
	if updateErr != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": updateErr.Error()})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book updated successfully", "data": updatedBook})
}

func Delete(context *fiber.Ctx) error {
	// Get the book ID from the URL parameters
	bookID := context.Params("id")

	// Call the DeleteById function from the services package
	deleteErr := services.Delete(bookID)
	if deleteErr != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": deleteErr.Error()})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully"})
}
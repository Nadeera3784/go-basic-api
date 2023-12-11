package services

import (
	"application/src/config"
	"application/src/models"
)

func GetAll() ([]models.Book, error) {
	var books []models.Book

	result := config.Instance.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func GetById(bookID string) (models.Book, error) {
	var book models.Book
	result := config.Instance.First(&book, bookID)
	if result.Error != nil {
		return models.Book{}, result.Error
	}

	return book, nil
}

func Create(book models.Book) (models.Book, error) {
	result := config.Instance.Create(&book)
	if result.Error != nil {
		return models.Book{}, result.Error
	}
	return book, nil
}

func Update(bookID string, updatedBook models.Book) (models.Book, error) {
	// Implement your logic to update the book in the database using the config.Instance
	// Fetch the existing book from the database using the provided ID
	var existingBook models.Book
	result := config.Instance.First(&existingBook, bookID)
	if result.Error != nil {
		return models.Book{}, result.Error
	}

	// Update the existing book with the new data
	existingBook.Title = updatedBook.Title
	existingBook.Description = updatedBook.Description

	// Save the updated book to the database
	result = config.Instance.Save(&existingBook)
	if result.Error != nil {
		return models.Book{}, result.Error
	}

	return existingBook, nil
}

func Delete(bookID string) error {

	existingBook, err := GetById(bookID);
	if err != nil {
		return err
	}

	result := config.Instance.Unscoped().Delete(&existingBook, bookID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
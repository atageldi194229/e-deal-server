package controllers

import (
	"errors"

	"github.com/atageldi194229/e-deal-server/database"
	"github.com/atageldi194229/e-deal-server/models"
	"github.com/gofiber/fiber/v2"
)

// Serializer
type Category struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	ParentID *uint  `json:"parent_id"`
}

func CreateResponseCategory(categoryModel models.Category) Category {
	return Category{ID: categoryModel.ID, Name: categoryModel.Name, ParentID: categoryModel.ParentID}
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&category)
	responseCategory := CreateResponseCategory(category)

	return c.Status(200).JSON(responseCategory)
}

func GetCategories(c *fiber.Ctx) error {
	categories := []models.Category{}

	database.Database.Db.Find(&categories)

	responseCategories := []Category{}
	for _, category := range categories {
		responseCategory := CreateResponseCategory(category)
		responseCategories = append(responseCategories, responseCategory)
	}

	return c.Status(200).JSON(responseCategories)
}

func findUser(id int, category *models.Category) error {
	database.Database.Db.Find(&category, "id = ?", id)

	if category.ID == 0 {
		return errors.New("Category does not exist")
	}

	return nil
}

func GetCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var category models.Category

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &category); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseCategory := CreateResponseCategory(category)

	return c.Status(200).JSON(responseCategory)
}

func UpdateCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var category models.Category

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &category); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateCategory struct {
		Name     string `json:"name"`
		ParentID *uint  `json:"parent_id"`
	}

	var updateCategory UpdateCategory

	if err := c.BodyParser(&updateCategory); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	category.Name = updateCategory.Name
	category.ParentID = updateCategory.ParentID

	database.Database.Db.Save(&category)

	responseCategory := CreateResponseCategory(category)

	return c.Status(200).JSON(responseCategory)
}

func DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var category models.Category

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &category); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&category).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Successfully deleted category")
}

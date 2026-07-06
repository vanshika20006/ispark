package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iips-oss/ispark/api/config"
	"github.com/iips-oss/ispark/api/models"
	"github.com/iips-oss/ispark/api/utils"
)

// UpdateProfile updates student's editable profile info
func UpdateProfile(c *fiber.Ctx) error {
	rollNo := c.Locals("roll_no").(string)

	type ProfileUpdate struct {
		EmailID   string `json:"email_id"`
		ContactNo string `json:"contact_no"`
		DOB       string `json:"dob"`
		Gender    string `json:"gender"`
	}

	var input ProfileUpdate
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	var student models.Student
	if err := config.DB.Where("roll_no = ?", rollNo).First(&student).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Student not found",
		})
	}

	// Update fields if provided
	if input.EmailID != "" {
		student.EmailID = input.EmailID
	}
	if input.ContactNo != "" {
		student.ContactNo = input.ContactNo
	}
	if input.DOB != "" {
		student.DOB = input.DOB
	}
	if input.Gender != "" {
		student.Gender = input.Gender
	}

	if err := config.DB.Save(&student).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update profile",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Profile updated successfully",
		"student": fiber.Map{
			"roll_no":       student.RollNo,
			"name":          student.Name,
			"email_id":      student.EmailID,
			"course_name":   student.CourseName,
			"semester":      student.Semester,
			"contact_no":    student.ContactNo,
			"dob":           student.DOB,
			"gender":        student.Gender,
			"enrollment_no": student.EnrollmentNo,
			"is_verified":   student.IsVerified,
		},
	})
}

type ChangePasswordInput struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

// ChangePassword changes the authenticated student's password
func ChangePassword(c *fiber.Ctx) error {
	rollNo := c.Locals("roll_no").(string)

	var input ChangePasswordInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON input",
		})
	}

	if len(input.NewPassword) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "New password must be at least 6 characters long",
		})
	}

	var student models.Student
	if err := config.DB.Where("roll_no = ?", rollNo).First(&student).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Student not found",
		})
	}

	// Verify current password
	if !utils.CheckPasswordHash(input.CurrentPassword, student.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Incorrect current password",
		})
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	student.Password = hashedPassword
	if err := config.DB.Save(&student).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update password",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Password changed successfully",
	})
}

package config

import (
	"log"

	"github.com/iips-oss/ispark/api/models"
	"github.com/iips-oss/ispark/api/utils"
)

func SeedDefaultAdmin() {
	hashed, _ := utils.HashPassword("admin123")

	var admin models.Admin
	if err := DB.Where(models.Admin{AdminID: "admin"}).
		Attrs(models.Admin{Password: hashed, MustChangePassword: true}).
		Assign(models.Admin{Name: "Platform Admin", Email: "admin@isparc.dev", Role: "admin", AssignedBatch: "IT2K24"}).
		FirstOrCreate(&admin).Error; err != nil {
		log.Printf("Seed admin failed: %v", err)
	}
}

// SeedDefaultSuperAdmin creates the demo super admin. A super admin has no
// assigned batch: it is not scoped to one, and sees the whole platform.
func SeedDefaultSuperAdmin() {
	hashed, err := utils.HashPassword("superadmin123")
	if err != nil {
		log.Printf("Seed super admin failed: %v", err)
		return
	}

	var superAdmin models.Admin
	if err := DB.Where(models.Admin{AdminID: "superadmin"}).
		Attrs(models.Admin{Password: hashed}).
		Assign(models.Admin{Name: "Platform Super Admin", Email: "superadmin@isparc.dev", Role: "superadmin"}).
		FirstOrCreate(&superAdmin).Error; err != nil {
		log.Printf("Seed super admin failed: %v", err)
	}
}

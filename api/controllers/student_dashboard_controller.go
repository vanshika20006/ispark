package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iips-oss/ispark/api/config"
	"github.com/iips-oss/ispark/api/models"
)

// GetDashboardStats returns stats for the student dashboard home page
func GetDashboardStats(c *fiber.Ctx) error {
	rollNo := c.Locals("roll_no").(string)

	// 1. Total activities participated (count of enrollments)
	var activitiesCount int64
	config.DB.Model(&models.Enrollment{}).Where("student_roll_no = ? AND status IN ('Enrolled', 'Completed')", rollNo).Count(&activitiesCount)

	// 2. Total certificates uploaded and pending/approved/rejected
	var certificatesCount int64
	config.DB.Model(&models.Certificate{}).Where("student_roll_no = ?", rollNo).Count(&certificatesCount)

	var pendingCertificatesCount int64
	config.DB.Model(&models.Certificate{}).Where("student_roll_no = ? AND status = 'Pending'", rollNo).Count(&pendingCertificatesCount)

	var approvedCertificatesCount int64
	config.DB.Model(&models.Certificate{}).Where("student_roll_no = ? AND status = 'Approved'", rollNo).Count(&approvedCertificatesCount)

	var rejectedCertificatesCount int64
	config.DB.Model(&models.Certificate{}).Where("student_roll_no = ? AND status = 'Rejected'", rollNo).Count(&rejectedCertificatesCount)

	// 3. Credits earned from approved certificates
	type SumResult struct {
		Total int
	}
	var sumResult SumResult
	config.DB.Raw("SELECT COALESCE(SUM(credits), 0) as total FROM certificates WHERE student_roll_no = ? AND status = 'Approved'", rollNo).Scan(&sumResult)

	// 4. Current Rank on the Leaderboard
	type StudentRank struct {
		RollNo       string
		TotalCredits int
	}
	var ranks []StudentRank
	config.DB.Raw(`
		SELECT s.roll_no, COALESCE(SUM(c.credits), 0) as total_credits
		FROM students s
		LEFT JOIN certificates c ON c.student_roll_no = s.roll_no AND c.status = 'Approved'
		GROUP BY s.roll_no
		ORDER BY total_credits DESC, s.roll_no ASC
	`).Scan(&ranks)

	rank := len(ranks)
	totalStudents := len(ranks)
	for i, r := range ranks {
		if r.RollNo == rollNo {
			rank = i + 1
			break
		}
	}

	// 5. Recent extracurricular activities list (approved/pending certificates)
	var recentActivities []models.Certificate
	config.DB.Where("student_roll_no = ?", rollNo).Order("created_at desc").Limit(5).Find(&recentActivities)

	return c.JSON(fiber.Map{
		"activities_participated": activitiesCount,
		"certificates_uploaded":   certificatesCount,
		"pending_certificates":    pendingCertificatesCount,
		"approved_certificates":   approvedCertificatesCount,
		"rejected_certificates":   rejectedCertificatesCount,
		"credits_earned":          sumResult.Total,
		"current_rank":            rank,
		"total_students":          totalStudents,
		"recent_activities":       recentActivities,
	})
}

package database

import (
	"fmt"
	"log"
	"time"

	attendance "github.com/teguh522/payslip/cmd/internal/domain/attendance/entity"
	employeeEntity "github.com/teguh522/payslip/cmd/internal/domain/employee/entity"
	overtime "github.com/teguh522/payslip/cmd/internal/domain/overtime/entity"
	payroll "github.com/teguh522/payslip/cmd/internal/domain/payroll/entity"
	payslip "github.com/teguh522/payslip/cmd/internal/domain/payslip/entity"
	reimbursement "github.com/teguh522/payslip/cmd/internal/domain/reimbursement/entity"
	userEntity "github.com/teguh522/payslip/cmd/internal/domain/user/entity"
	"github.com/teguh522/payslip/cmd/internal/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgreSQLGORM(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Jakarta",
		cfg.DataBase.DBHost, cfg.DataBase.DBUser, cfg.DataBase.DBPassword, cfg.DataBase.DBName, cfg.DataBase.DBPort, cfg.DataBase.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Successfully connected to PostgreSQL with GORM!")

	err = db.AutoMigrate(
		&userEntity.User{},
		&employeeEntity.Employee{},
		&attendance.AttendancePeriod{},
		&attendance.Attendance{},
		&overtime.Overtime{},
		&reimbursement.Reimbursement{},
		&payroll.Payroll{},
		&payslip.Payslip{},
	)
	if err != nil {
		log.Printf("Failed to auto migrate database: %v", err)
	}

	return db, nil
}

func ClosePostgreSQLGORM(db *gorm.DB) {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting underlying sql.DB for closing: %v", err)
			return
		}
		err = sqlDB.Close()
		if err != nil {
			log.Printf("Error closing PostgreSQL connection: %v", err)
		} else {
			log.Println("PostgreSQL connection closed.")
		}
	}
}

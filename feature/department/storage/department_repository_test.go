package storage

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{
		DisableAutomaticPing: true,
	})
	if err != nil {
		return nil, nil, err
	}
	return gormDB, mock, nil
}

func TestDepartmentRepository_Create(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	// Create repository instance
	repo := NewDepartmentRepository(gormDB)

	// Define test data
	department := &domain.Department{
		Name:    "HR",
		Address: "123 HR Street",
		Status:  1,
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `departments` (`name`,`address`,`status`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)")).
		WithArgs(department.Name, department.Address, department.Status, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate successful insert
	mock.ExpectCommit()

	// Call the method under test
	err = repo.Create(department)

	// Assertions
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetDepartmentByID(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewDepartmentRepository(gormDB)

	departmentID := uint(1)
	department := &domain.Department{
		ID:      departmentID,
		Name:    "Finance",
		Address: "456 Finance Street",
		Status:  456,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "address", "status"}).
		AddRow(department.ID, department.Name, department.Address, department.Status)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `departments` WHERE `departments`.`id` = ? ORDER BY `departments`.`id` LIMIT ?")).
		WithArgs(department.ID, 1).
		WillReturnRows(rows)

	result, err := repo.GetByID(department.ID)
	assert.NoError(t, err)
	assert.Equal(t, department, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateDepartment(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewDepartmentRepository(gormDB)

	department := &domain.Department{
		ID:        1,
		Name:      "IT",
		Address:   "789 IT Street",
		Status:    789,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `departments` SET `name`=?,`address`=?,`status`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(department.Name, department.Address, department.Status, sqlmock.AnyArg(), sqlmock.AnyArg(), department.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Update(department)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteDepartment(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()
	repo := NewDepartmentRepository(gormDB)

	departmentID := uint(1)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `departments` WHERE `departments`.`id` = ?").
		WithArgs(departmentID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Delete(departmentID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

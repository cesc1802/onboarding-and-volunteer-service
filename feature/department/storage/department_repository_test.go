package storage

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// setupMockDB initializes a mock database connection and returns a GORM DB instance and the mock object.
func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err) // Assert that there are no errors creating the mock DB.

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0", // Data Source Name (DSN) for the mock database.
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err) // Assert that GORM opens without error.

	return gormDB, mock
}

// TestCreateDepartment tests the Create method of DepartmentRepository.
func TestCreateDepartment(t *testing.T) {
	gormDB, mock := setupMockDB(t) // Set up mock DB and GORM instance.

	repo := NewDepartmentRepository(gormDB) // Create a new DepartmentRepository.

	// Define a department object for testing.
	department := &domain.Department{
		Name:    "HR",
		Address: "123 HR Street",
		Status:  123,
	}

	// Set expectations for SQL execution: inserting a new department record.
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `departments`").
		WithArgs(department.Name, department.Address, department.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call the Create method and assert no errors.
	err := repo.Create(department)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// TestGetDepartmentByID tests the GetByID method of DepartmentRepository.
func TestGetDepartmentByID(t *testing.T) {
	gormDB, mock := setupMockDB(t) // Set up mock DB and GORM instance.

	repo := NewDepartmentRepository(gormDB) // Create a new DepartmentRepository.

	// Define a department object with expected data.
	department := &domain.Department{
		Name:    "Finance",
		Address: "456 Finance Street",
		Status:  456,
	}
	// Set expectations for SQL query: selecting a department by ID.
	rows := sqlmock.NewRows([]string{"id", "name", "address", "status"}).
		AddRow(department.Id, department.Name, department.Address, department.Status)

	mock.ExpectQuery("SELECT * FROM `departments` WHERE `departments`.`id` = ?").
		WithArgs(department.Id).
		WillReturnRows(rows)

	// Call the GetByID method and assert the returned data matches expectations.
	result, err := repo.GetByID(department.Id)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, department.Name, result.Name)
	assert.Equal(t, department.Address, result.Address)
	assert.Equal(t, department.Status, result.Status)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// TestUpdateDepartment tests the Update method of DepartmentRepository.
func TestUpdateDepartment(t *testing.T) {
	gormDB, mock := setupMockDB(t) // Set up mock DB and GORM instance.

	repo := NewDepartmentRepository(gormDB) // Create a new DepartmentRepository.

	// Define a department object with new data to update.
	department := &domain.Department{
		Name:    "IT",
		Address: "789 IT Street",
		Status:  789,
	}
	// Set expectations for SQL execution: updating a department record.
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `departments` SET `name`=?,`address`=?,`status`=? WHERE `id` = ?").
		WithArgs(department.Name, department.Address, department.Status, department.Id).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Expecting one row to be affected.
	mock.ExpectCommit()

	// Call the Update method and assert no errors.
	err := repo.Update(department)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet()) // Assert that all mock expectations were met.
}

// TestDeleteDepartment tests the Delete method of DepartmentRepository.
func TestDeleteDepartment(t *testing.T) {
	gormDB, mock := setupMockDB(t) // Set up mock DB and GORM instance.

	repo := NewDepartmentRepository(gormDB) // Create a new DepartmentRepository.

	departmentID := uint(1) //Define the ID of the department to delete.

	// Set expectations for SQL execution: deleting a department record.
	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `departments` WHERE `departments`.`id` = ?").
		WithArgs(departmentID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call the Delete method and assert no errors.
	err := repo.Delete(departmentID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet()) // Assert that all mock expectations were met.

}

package storage

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/domain"
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
func TestCreateVolunteer(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewVolunteerRepository(gormDB)
	volunteer := &domain.Volunteer{
		UserID:       101,
		DepartmentID: 10,
		Status:       01,
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `volunteers` (`user_id`,`department_id`,`status`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)")).
		WithArgs(volunteer.UserID, volunteer.DepartmentID, volunteer.Status, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.CreateVolunteer(volunteer)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateVolunteer(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer func() { _ = mock.ExpectationsWereMet() }()

	repo := NewVolunteerRepository(db)
	volunteer := &domain.Volunteer{
		ID:           1,
		UserID:       001,
		DepartmentID: 010,
		Status:       123,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `volunteers` SET `user_id`=?,`department_id`=?,`status`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(volunteer.UserID, volunteer.DepartmentID, volunteer.Status, sqlmock.AnyArg(), sqlmock.AnyArg(), volunteer.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.UpdateVolunteer(volunteer)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteVolunteer(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer func() { _ = mock.ExpectationsWereMet() }()

	repo := NewVolunteerRepository(db)
	volunteerID := 1

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `volunteers` WHERE `volunteers`.`id` = ?").
		WithArgs(volunteerID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.DeleteVolunteer(volunteerID)
	assert.NoError(t, err)
}

func TestFindVolunteerByID(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
		_ = mock.ExpectationsWereMet()
	}()

	repo := NewVolunteerRepository(db)
	volunteerID := 1
	volunteer := &domain.Volunteer{
		ID:           1,
		UserID:       101,
		DepartmentID: 10,
		Status:       123,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "department_id", "status", "created_at", "updated_at"}).
		AddRow(volunteer.ID, volunteer.UserID, volunteer.DepartmentID, volunteer.Status, volunteer.CreatedAt, volunteer.UpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `volunteers` WHERE `volunteers`.`id` = ? ORDER BY `volunteers`.`id` LIMIT ?")).
		WithArgs(volunteerID, 1).
		WillReturnRows(rows)

	result, err := repo.FindVolunteerByID(volunteerID)
	assert.NoError(t, err)
	assert.Equal(t, volunteer.ID, result.ID)
	assert.Equal(t, volunteer.UserID, result.UserID)
	assert.Equal(t, volunteer.DepartmentID, result.DepartmentID)
	assert.Equal(t, volunteer.Status, result.Status)
	assert.WithinDuration(t, volunteer.CreatedAt, result.CreatedAt, time.Second)
	assert.WithinDuration(t, volunteer.UpdatedAt, result.UpdatedAt, time.Second)
}

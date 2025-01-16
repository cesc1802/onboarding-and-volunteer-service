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
	// Ensure to close the db connection after test completes
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewVolunteerRepository(gormDB)

	// Volunteer instance with foreign key relationships
	volunteer := &domain.Volunteer{
		UserID:             101,
		DepartmentID:       new(int), // Assuming this is an existing department with ID 1
		Status:             1,
		Dob:                time.Now(),
		Mobile:             "1234567890",
		CountryID:          new(int), // Assuming this is an existing country with ID 1
		ResidentCountryID:  nil,      // Optional, assuming it can be nil
		Avatar:             "",
		VerificationStatus: 0,
	}

	// Setting foreign key values
	*volunteer.DepartmentID = 1 // Simulating the Department ID as 1
	*volunteer.CountryID = 1    // Simulating the Country ID as 1

	// Mocking the foreign key relationships
	mock.ExpectBegin()

	// Here we simulate an insert into the `volunteer_details` table, including the foreign keys
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `volunteer_details` (`user_id`,`department_id`,`dob`,`mobile`,`country_id`,`resident_country_id`,`avatar`,`verification_status`,`status`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(volunteer.UserID, volunteer.DepartmentID, volunteer.Dob, volunteer.Mobile, volunteer.CountryID, volunteer.ResidentCountryID, volunteer.Avatar, volunteer.VerificationStatus, volunteer.Status, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate successful insert with row ID = 1
	mock.ExpectCommit()

	// Call the method to test
	err = repo.CreateVolunteer(volunteer)

	// Assert no errors
	assert.NoError(t, err)

	// Check that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateVolunteer(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer func() { _ = mock.ExpectationsWereMet() }()

	repo := NewVolunteerRepository(db)
	volunteer := &domain.Volunteer{
		ID:                 1,
		UserID:             101,
		DepartmentID:       nil, // assuming it can be nil
		Status:             1,
		Dob:                time.Now(),
		Mobile:             "1234567890",
		CountryID:          nil, // assuming it can be nil
		ResidentCountryID:  nil, // assuming it can be nil
		Avatar:             "",
		VerificationStatus: 0,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `volunteer_details` SET `user_id`=?,`department_id`=?,`dob`=?,`mobile`=?,`country_id`=?,`resident_country_id`=?,`avatar`=?,`verification_status`=?,`status`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(volunteer.UserID, volunteer.DepartmentID, volunteer.Dob, volunteer.Mobile, volunteer.CountryID, volunteer.ResidentCountryID, volunteer.Avatar, volunteer.VerificationStatus, volunteer.Status, sqlmock.AnyArg(), sqlmock.AnyArg(), volunteer.ID).
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
	mock.ExpectExec("DELETE FROM `volunteer_details` WHERE `volunteer_details`.`id` = ?").
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
		ID:                 1,
		UserID:             101,
		DepartmentID:       nil, // assuming it can be nil
		Status:             1,
		Dob:                time.Now(),
		Mobile:             "1234567890",
		CountryID:          nil, // assuming it can be nil
		ResidentCountryID:  nil, // assuming it can be nil
		Avatar:             "",
		VerificationStatus: 0,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "department_id", "dob", "mobile", "country_id", "resident_country_id", "avatar", "verification_status", "status", "created_at", "updated_at"}).
		AddRow(volunteer.ID, volunteer.UserID, volunteer.DepartmentID, volunteer.Dob, volunteer.Mobile, volunteer.CountryID, volunteer.ResidentCountryID, volunteer.Avatar, volunteer.VerificationStatus, volunteer.Status, volunteer.CreatedAt, volunteer.UpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `volunteer_details` WHERE `volunteer_details`.`id` = ? ORDER BY `volunteer_details`.`id` LIMIT ?")).
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

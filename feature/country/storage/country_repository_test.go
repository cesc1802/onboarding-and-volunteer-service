package storage

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/domain"
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

func TestCountryRepository_Create(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewCountryRepository(gormDB)

	country := &domain.Country{
		Name:   "Test Country",
		Status: 0,
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `countries` (`name`,`status`,`created_at`,`updated_at`) VALUES (?,?,?,?)")).
		WithArgs(country.Name, country.Status, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Create(country)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCountryRepository_GetByID(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewCountryRepository(gormDB)

	countryID := uint(1)
	country := &domain.Country{
		ID:     countryID,
		Name:   "Test Country",
		Status: 1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "status"}).
		AddRow(country.ID, country.Name, country.Status)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `countries` WHERE `countries`.`id` = ? ORDER BY `countries`.`id` LIMIT ?")).
		WithArgs(countryID, 1).
		WillReturnRows(rows)

	result, err := repo.GetByID(countryID)
	assert.NoError(t, err)
	assert.Equal(t, country, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCountryRepository_Update(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewCountryRepository(gormDB)

	country := &domain.Country{
		ID:        1,
		Name:      "Updated Country",
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `countries` SET `name`=?,`status`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(country.Name, country.Status, sqlmock.AnyArg(), sqlmock.AnyArg(), country.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Update(country)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCountryRepository_Delete(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewCountryRepository(gormDB)

	countryID := uint(1)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `countries` WHERE `countries`.`id` = ?")).
		WithArgs(countryID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Delete(countryID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

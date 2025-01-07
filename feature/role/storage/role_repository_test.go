package storage

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/domain"
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

func TestRoleRepository_Create(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewRoleRepository(gormDB)

	role := &domain.Role{
		Name:   "Admin",
		Status: "pending",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `roles` (`name`,`status`,`created_at`,`updated_at`) VALUES (?,?,?,?)")).
		WithArgs(role.Name, role.Status, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Create(role)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRoleRepository_GetByID(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewRoleRepository(gormDB)

	roleID := uint(1)
	role := &domain.Role{
		ID:     roleID,
		Name:   "Admin",
		Status: "approve",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "status"}).
		AddRow(role.ID, role.Name, role.Status)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `roles` WHERE `roles`.`id` = ? ORDER BY `roles`.`id` LIMIT ?")).
		WithArgs(roleID, 1).
		WillReturnRows(rows)

	result, err := repo.GetByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, role.Name, result.Name)
	assert.Equal(t, role.Status, result.Status)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRoleRepository_Update(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewRoleRepository(gormDB)

	role := &domain.Role{
		ID:        1,
		Name:      "Admin",
		Status:    "rejected",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `roles` SET `name`=?,`status`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(role.Name, role.Status, sqlmock.AnyArg(), sqlmock.AnyArg(), role.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Update(role)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRoleRepository_Delete(t *testing.T) {
	gormDB, mock, err := setupMockDB()
	if err != nil {
		t.Fatalf("failed to setup mock db: %v", err)
	}
	defer func() {
		sqlDB, _ := gormDB.DB()
		sqlDB.Close()
	}()

	repo := NewRoleRepository(gormDB)

	roleID := uint(1)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `roles` WHERE `roles`.`id` = ?")).
		WithArgs(roleID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Delete(roleID)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

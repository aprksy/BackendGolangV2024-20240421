package repository_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/google/uuid"
)

func TestCreateEstate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO estate").ExpectExec().WithArgs(sqlmock.AnyArg(), 4, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := repository.Repository{db}
	if _, err := repo.CreateEstate(context.TODO(), repository.CreateEstateInput{Length: 4, Width: 3}); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSetTreeHeight(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO plot").ExpectExec().WithArgs(sqlmock.AnyArg(), 1, 1, 10).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := repository.Repository{db}
	if _, err := repo.SetTreeHeight(context.TODO(), repository.SetTreeHeightInput{X: 1, Y: 1, Height: 10}); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetEstateStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	columns := []string{"estate_id", "x", "y", "height"}
	mock.ExpectPrepare("SELECT (.+) FROM estate_stats").ExpectQuery().WithArgs(sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows(columns))
	repo := repository.Repository{db}
	if _, err := repo.GetEstateStats(context.TODO(), repository.GetEstateStatsInput{Id: uuid.New()}); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

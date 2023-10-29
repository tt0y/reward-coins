package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"reward-coins-api/config"
	"reward-coins-api/models"
)

func TestListCoinTypes(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	config.Database = db

	rows := sqlmock.NewRows([]string{"id", "name", "description"}).
		AddRow(1, "Coin Type 1", "Description 1").
		AddRow(2, "Coin Type 2", "Description 2")

	mock.ExpectQuery("SELECT id, name, description FROM coin_types").WillReturnRows(rows)

	app := fiber.New()
	app.Get("/coin_types", ListCoinTypes)

	req := httptest.NewRequest("GET", "/coin_types", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)

	var coinTypes []models.CoinType
	err = json.NewDecoder(resp.Body).Decode(&coinTypes)
	if err != nil {
		return
	}

	assert.Len(t, coinTypes, 2)
	assert.Equal(t, "Coin Type 1", coinTypes[0].Name)
	assert.Equal(t, "Description 1", coinTypes[0].Description)
	assert.Equal(t, "Coin Type 2", coinTypes[1].Name)
	assert.Equal(t, "Description 2", coinTypes[1].Description)

	assert.NoError(t, mock.ExpectationsWereMet())
}

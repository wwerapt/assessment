package expense

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func (h *handler) CreateExpensesHandler(c echo.Context) error {
	exp := Expense{}
	err := c.Bind(&exp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := h.DB.QueryRow("INSERT INTO expenses(title, amount, note, tags) values($1, $2, $3, $4) RETURNING id", exp.Title, exp.Amount, exp.Note, pq.Array(exp.Tags))
	err = row.Scan(&exp.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, exp)
}

func (h *handler) GetIdExpensesHandler(c echo.Context) error {
	id := c.Param("id")
	stmt, err := h.DB.Prepare("SELECT id, title, amount, note, tags  FROM expenses WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query user statment:" + err.Error()})
	}

	row := stmt.QueryRow(id)
	exp := Expense{}
	err = row.Scan(&exp.ID, &exp.Title, &exp.Amount, &exp.Note, pq.Array(&exp.Tags))
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "user not found"})
	case nil:
		return c.JSON(http.StatusOK, exp)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user:" + err.Error()})
	}
}

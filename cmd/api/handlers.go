package main

import (
	"errors"
	"net/http"

	"github.com/eyasu-teshome/addis-eats-api/internal/menu"
)

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (app *application) listDishes(w http.ResponseWriter, r *http.Request) {
	dishes, err := app.dishes.All()
	if err != nil {
		app.logger.Println(err)
		writeError(w, http.StatusInternalServerError, "internal error")
		return
	}
	writeJSON(w, http.StatusOK, dishes)
}

func (app *application) showDish(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // Go 1.22 wildcard path variable
	dish, err := app.dishes.ByID(id)
	switch {
	case errors.Is(err, menu.ErrNotFound):
		writeError(w, http.StatusNotFound, "dish not found")
	case err != nil:
		app.logger.Println(err)                            // Log actual internal error
		writeError(w, http.StatusInternalServerError, "internal error") // Return safe summary
	default:
		writeJSON(w, http.StatusOK, dish)
	}
}
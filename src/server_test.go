package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// TestReadBookRoute testar att routen /books/{title} fungerar korrekt
func TestReadBookRoute(t *testing.T) {
	// 1. Skapa en ny router (samma som i main)
	r := mux.NewRouter()

	// 2. Registrera routen exakt som i server.go
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")

	// 3. Skapa en testförfrågan med en specifik titel
	req, err := http.NewRequest("GET", "/books/1984", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 4. Skapa en ResponseRecorder för att fånga svaret
	rr := httptest.NewRecorder()

	// 5. Skicka förfrågan genom routern
	r.ServeHTTP(rr, req)

	// 6. Kontrollera statuskod (ska vara 200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returnerade fel statuskod: got %v want %v", status, http.StatusOK)
	}

	// 7. Kontrollera svarsinnehållet
	expected := "Du läser boken: 1984\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returnerade felaktigt innehåll: got %v want %v", rr.Body.String(), expected)
	}
	t.Logf("Förväntat innehåll: %q", expected)
}

// TestReadBookRouteNotFound testar att en felaktig metod ger 404/405
func TestReadBookRouteMethodNotAllowed(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")

	// Försök med POST istället för GET
	req, err := http.NewRequest("POST", "/books/1984", nil)
	t.Logf("Försök med POST istället för GET")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Ska inte vara 200 OK eftersom metoden är fel
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler borde ha misslyckats för POST: got %v", status)
	}
	t.Logf("Status: %v instället för 200", rr.Code)
}   
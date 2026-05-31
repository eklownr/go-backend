package passwd

import (
	"testing"
)

func TestPassword(t *testing.T) {
    // Anropa dina funktioner här, t.ex:
    // result := HashPassword("minlösenord")
    // if result == "" { t.Error("Funktionen returnerade inget") }
	expected := print_password()
    if !expected {
        t.Error("Funktionen returnerade fel")
    }
    t.Logf("Match expect to be true: %v", expected)

}   
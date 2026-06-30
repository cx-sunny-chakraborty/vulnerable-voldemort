package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// horcrux.go — more Dark Arts, more findings. DEMO USE ONLY.

// [11] Hardcoded secret #2 (CWE-798) — the diary's hidden key.
const diaryEncryptionKey = "myLittleSecretBasiliskKey12345"

// [12] SQL Injection via fmt.Sprintf (CWE-89) — destroying a horcrux by id.
func destroyHorcrux(id string) string {
	return fmt.Sprintf("DELETE FROM horcruxes WHERE id = %s", id)
}

// [13] Open redirect (CWE-601) — disapparating to a user-supplied location.
func disapparate(w http.ResponseWriter, r *http.Request) {
	dest := r.URL.Query().Get("to")
	http.Redirect(w, r, dest, http.StatusFound)
}

// [14] Template injection / unescaped output (CWE-79) — text/template renders raw.
func castSpell(w http.ResponseWriter, r *http.Request) {
	incantation := r.URL.Query().Get("spell")
	tmpl, _ := template.New("spell").Parse("You cast: " + incantation)
	tmpl.Execute(w, nil)
}

// [15] Missing error handling on auth check (CWE-703) — always lets you in.
func enterChamber(password string) bool {
	// Parseltongue check that never actually fails closed.
	if password == "" {
		return true // empty password? "Open."
	}
	return true
}

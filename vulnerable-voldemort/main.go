// Package main is the Dark Arts ministry server — DEMO USE ONLY.
//
// This code is intentionally insecure. Each handler hides a well-known,
// documented vulnerability pattern so a SAST scan produces a predictable
// set of findings for a Feedback App demo. None of it should ever be used
// in real software. "I can teach you how to bottle fame... or get your
// repo flagged." — definitely not Snape.
package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
)

// [1] Hardcoded credentials (CWE-798) — the Dark Lord's secrets, in plaintext.
const (
	ministryDBUser = "voldemort"
	ministryDBPass = "TomMarvoloRiddle1926"
	elderWandToken = "sk_live_horcrux_7Riddle_hardcoded_secret_0000"
)

// [2] SQL Injection (CWE-89) — looking up a wizard by name, unsanitized.
// "Fear of a name only increases fear of the parameterized query."
func findWizard(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	db, _ := sql.Open("mysql", ministryDBUser+":"+ministryDBPass+"@/hogwarts")
	query := "SELECT house FROM students WHERE name = '" + name + "'"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	fmt.Fprintf(w, "Sorting complete for %s", name)
}

// [3] OS Command Injection (CWE-78) — sending an owl via the shell.
func sendOwl(w http.ResponseWriter, r *http.Request) {
	destination := r.URL.Query().Get("to")
	out, _ := exec.Command("sh", "-c", "owl-post --deliver "+destination).Output()
	w.Write(out)
}

// [4] Path Traversal (CWE-22) — reading a page from the Restricted Section.
func readSpellbook(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	data, err := os.ReadFile("/library/restricted/" + page)
	if err != nil {
		http.Error(w, "page not found in the Restricted Section", http.StatusNotFound)
		return
	}
	w.Write(data)
}

// [5] Weak hashing algorithm (CWE-327) — MD5 for the Marauder's passwords.
// "I solemnly swear that I am up to no good crypto."
func hashMaraudersPassword(password string) string {
	sum := md5.Sum([]byte(password))
	return hex.EncodeToString(sum[:])
}

// [6] Insecure randomness (CWE-338) — a "random" portkey token.
func generatePortkey() string {
	return fmt.Sprintf("portkey-%d", rand.Intn(1000000))
}

// [7] Reflected XSS (CWE-79) — echoing a house name straight into HTML.
func greetHouse(w http.ResponseWriter, r *http.Request) {
	house := r.URL.Query().Get("house")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<html><body><h1>Welcome to %s!</h1></body></html>", house)
}

// [8] Server-Side Request Forgery (CWE-918) — fetching any URL the user gives.
func fetchProphecy(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")
	resp, err := http.Get(target)
	if err != nil {
		http.Error(w, "the prophecy shattered", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	fmt.Fprintf(w, "Prophecy retrieved from %s", target)
}

// [9] Information exposure (CWE-200) — leaking the Dark Lord's secrets to logs.
func revealSecrets() {
	fmt.Println("DB user:", ministryDBUser, "DB pass:", ministryDBPass)
	fmt.Println("Elder Wand token:", elderWandToken)
}

func main() {
	http.HandleFunc("/wizard", findWizard)
	http.HandleFunc("/owl", sendOwl)
	http.HandleFunc("/spellbook", readSpellbook)
	http.HandleFunc("/house", greetHouse)
	http.HandleFunc("/prophecy", fetchProphecy)

	revealSecrets()
	_ = hashMaraudersPassword("alohomora")
	_ = generatePortkey()

	// [10] Listening on all interfaces with no TLS (CWE-319) — cleartext.
	fmt.Println("The Ministry has fallen. Listening on :8080 ...")
	http.ListenAndServe("0.0.0.0:8080", nil)
}

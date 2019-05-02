package sensei

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

func TestSensei(t *testing.T) {
	sensei, rec, req := prepare()
	key, value := "snacks", "apple"
	if err := sensei.Set(rec, req, key, value); err != nil {
		fatalln(t, "unexpected error", err, "nil")
	}

	stored, err := sensei.Get(req, key)
	if err != nil {
		fatalln(t, "unepxected error", err, "nil")
	}
	stored, ok := stored.(string)
	if !ok {
		fatalln(t, "type of stored type", fmt.Sprintf("%T", stored), "string")
	}
	if stored != value {
		errorln(t, "stored value", stored, value)
	}
}

func TestSenseiFlash(t *testing.T) {
	sensei, rec, req := prepare()
	key, value := "snacks", "banana"
	if err := sensei.SetFlash(rec, req, key, value); err != nil {
		fatalln(t, "unexpected error", err, "nil")
	}

	flashes, err := sensei.GetFlashes(rec, req, key)
	if err != nil {
		fatalln(t, "unexpected error", err, "nil")
	}

	if len(flashes) != 1 {
		fatalln(t, "length of flashes", len(flashes), 1)
	}
	stored, ok := flashes[0].(string)
	if !ok {
		fatalln(t, "type of stored value", fmt.Sprintf("%T", flashes[0]), "string")
	}
	if stored != value {
		errorln(t, "stored value", stored, value)
	}

	flashes, err = sensei.GetFlashes(rec, req, key)
	if err != nil {
		fatalln(t, "unexpected error", err, "nil")
	}
	if len(flashes) != 0 {
		fatalln(t, "length of flashes", len(flashes), 1)
	}
}

func prepare() (*Sensei, *httptest.ResponseRecorder, *http.Request) {
	return testSensei(), httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "http://sensei.com", nil)
}

func testSensei() *Sensei {
	return New(testStore, testSessionKey)
}

var testStore = sessions.NewCookieStore(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

const testSessionKey = "BananaIsIncludedIntoSnacks"

func fatalln(t *testing.T, name string, actual, expected interface{}) {
	t.Fatalf("%s: got %v, expected %v\n", name, actual, expected)
}

func errorln(t *testing.T, name string, actual, expected interface{}) {
	t.Errorf("%s: got %v, expected %v\n", name, actual, expected)
}

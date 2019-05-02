package sensei

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type Sensei struct {
	Store sessions.Store
	Key   string
}

func New(store sessions.Store, key string) *Sensei {
	return &Sensei{
		Store: store,
		Key:   key,
	}
}

func (s *Sensei) SetFlash(w http.ResponseWriter, r *http.Request, key string, value interface{}) error {
	sess, err := s.Session(r)
	if err != nil {
		return err
	}

	sess.AddFlash(value, key)

	return sess.Save(r, w)
}

func (s *Sensei) GetFlashes(w http.ResponseWriter, r *http.Request, key string) ([]interface{}, error) {
	sess, err := s.Session(r)
	if err != nil {
		return nil, err
	}

	flashes := sess.Flashes(key)
	sess.Save(r, w)

	return flashes, nil
}

func (s *Sensei) Session(r *http.Request) (*sessions.Session, error) {
	return s.Store.Get(r, s.Key)
}

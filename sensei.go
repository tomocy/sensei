package sensei

import "github.com/gorilla/sessions"

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

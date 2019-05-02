# sensei

[![CircleCI](https://circleci.com/gh/tomocy/sensei.svg?style=svg)](https://circleci.com/gh/tomocy/sensei)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

utilities for [gorilla/sessions](https://github.com/gorilla/sessions)

## Installtion
```
go get github.com/tomocy/sensei
```

## Useage
Create a `Sensei` instance with `sessions.Store` and session key
```go
var manager = sensei.New(store, sessionKey)

var store = sessions.NewCookieStore(
    securecookie.GenerateRandomKey(64),
    securecookie.GenerateRandomKey(32),
)

const sessionKey = "BananaIsIncludedIntoSnacks"
```

```go
func KeepAuthenticUserID(w http.ResponseWriter, r *http.Request, id string) error {
    return sess.Set(w, r, authenticUserID, id)
}

func FindAuthenticUserID(r *http.Request) (string, error) {
    if id, ok := manager.Get(r, authenticUserID); ok && id != "" {
        return id, nil
    }

    return "", errors.New("no authentic user id")
}

const (
    authentiUserID = "authentic_user_id"
)
```

## Author
[tomocy](https://github.com/tomocy)
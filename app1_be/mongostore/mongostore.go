// myapp/mongostore/mongostore.go
package mongostore

import (
	"github.com/gin-contrib/sessions"
	"github.com/globalsign/mgo"
	"github.com/kidstuff/mongostore"
)

func NewMongoStore(url string, dbName, collectionName string, maxAge int, ensureTTL bool, keyPairs ...[]byte) (sessions.Store, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	store := mongostore.NewMongoStore(session.DB(dbName).C(collectionName), maxAge, ensureTTL, keyPairs...)
	return store, nil
}

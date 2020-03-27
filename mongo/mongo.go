package mongo

import (
	"errors"
	"github.com/globalsign/mgo"
)

// Mongo represents a simplistic MongoDB configuration.
type Mongo struct {
	Collection string
	Database   string
	Session    *mgo.Session
	URI        string
}

// Init creates a new mgo.Session with a strong consistency and a write mode of "majority".
func (m *Mongo) Init() (session *mgo.Session, err error) {
	if session != nil {
		return nil, errors.New("session already exists")
	}

	if session, err = mgo.Dial(m.URI); err != nil {
		return nil, err
	}

	session.EnsureSafe(&mgo.Safe{WMode: "majority"})
	session.SetMode(mgo.Strong, true)
	return session, nil
}

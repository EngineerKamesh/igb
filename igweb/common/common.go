package common

import (
	"github.com/EngineerKamesh/igb/igweb/common/datastore"
	"github.com/gorilla/sessions"
	"github.com/isomorphicgo/isokit"
)

type Env struct {
	DB          datastore.Datastore
	TemplateSet *isokit.TemplateSet
	Store       *sessions.FilesystemStore
}

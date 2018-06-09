package common

import (
	"github.com/EngineerKamesh/igb/igweb/common/datastore"
	"github.com/gorilla/sessions"
	"go.isomorphicgo.org/go/isokit"
)

type Env struct {
	DB          datastore.Datastore
	TemplateSet *isokit.TemplateSet
	Store       *sessions.FilesystemStore
}

package generator

import (
	"github.com/hyperxpizza/contact-manager/server/graph/model"
	"github.com/hyperxpizza/rpiCli/server/filestorage"
)
package generator

var filestorage filestorage.Filestorage

func init() {
	filestorage = filestorage.NewFileStorage()
}

func GenerateJSON(contacts []*model.Contact) (string, error) {

}

func GenerateCSV() {

}

func GenerateVCARD() {

}

package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("5c1oykad2csm1dk")
		if err != nil {
			return err
		}

		collection.Name = "git"

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("5c1oykad2csm1dk")
		if err != nil {
			return err
		}

		collection.Name = "gitSource_duplicate"

		return dao.SaveCollection(collection)
	})
}

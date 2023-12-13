package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ojmqge2b4cj4ywj")
		if err != nil {
			return err
		}

		// add
		new_buildTime := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "gfkidnpv",
			"name": "buildTime",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), new_buildTime)
		collection.Schema.AddField(new_buildTime)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ojmqge2b4cj4ywj")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("gfkidnpv")

		return dao.SaveCollection(collection)
	})
}

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
		new_logs := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "zmw8978x",
			"name": "logs",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_logs)
		collection.Schema.AddField(new_logs)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ojmqge2b4cj4ywj")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("zmw8978x")

		return dao.SaveCollection(collection)
	})
}

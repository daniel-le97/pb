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

		collection, err := dao.FindCollectionByNameOrId("6po51li0eihk83d")
		if err != nil {
			return err
		}

		// add
		new_content := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "n2yrsue3",
			"name": "content",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_content)
		collection.Schema.AddField(new_content)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("6po51li0eihk83d")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("n2yrsue3")

		return dao.SaveCollection(collection)
	})
}

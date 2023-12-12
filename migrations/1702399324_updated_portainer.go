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

		collection, err := dao.FindCollectionByNameOrId("67nwlz2znzqld2l")
		if err != nil {
			return err
		}

		collection.Name = "templates"

		// remove
		collection.Schema.RemoveField("i73hs3kj")

		// add
		new_type := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "hep0zk5c",
			"name": "type",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"portainer",
					"caprover",
					"custom"
				]
			}
		}`), new_type)
		collection.Schema.AddField(new_type)

		// add
		new_content := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "gkblufxu",
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

		collection, err := dao.FindCollectionByNameOrId("67nwlz2znzqld2l")
		if err != nil {
			return err
		}

		collection.Name = "portainer"

		// add
		del_template := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "i73hs3kj",
			"name": "template",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 2000000
			}
		}`), del_template)
		collection.Schema.AddField(del_template)

		// remove
		collection.Schema.RemoveField("hep0zk5c")

		// remove
		collection.Schema.RemoveField("gkblufxu")

		return dao.SaveCollection(collection)
	})
}

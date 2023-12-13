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

		collection, err := dao.FindCollectionByNameOrId("tqlcamhhas2xzr7")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("yygtahnw")

		// add
		new_installCommand := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "wkcbdx8g",
			"name": "installCommand",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_installCommand)
		collection.Schema.AddField(new_installCommand)

		// add
		new_buildCommand := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "8eahdfow",
			"name": "buildCommand",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_buildCommand)
		collection.Schema.AddField(new_buildCommand)

		// add
		new_startCommand := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "7s3jdfc8",
			"name": "startCommand",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_startCommand)
		collection.Schema.AddField(new_startCommand)

		// add
		new_ports := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rb6k79dz",
			"name": "ports",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_ports)
		collection.Schema.AddField(new_ports)

		// add
		new_exposedPorts := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "btii8ft7",
			"name": "exposedPorts",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_exposedPorts)
		collection.Schema.AddField(new_exposedPorts)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("tqlcamhhas2xzr7")
		if err != nil {
			return err
		}

		// add
		del_application := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "yygtahnw",
			"name": "application",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 2000000
			}
		}`), del_application)
		collection.Schema.AddField(del_application)

		// remove
		collection.Schema.RemoveField("wkcbdx8g")

		// remove
		collection.Schema.RemoveField("8eahdfow")

		// remove
		collection.Schema.RemoveField("7s3jdfc8")

		// remove
		collection.Schema.RemoveField("rb6k79dz")

		// remove
		collection.Schema.RemoveField("btii8ft7")

		return dao.SaveCollection(collection)
	})
}

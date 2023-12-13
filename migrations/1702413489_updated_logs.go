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

		// update
		edit_project := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "hwh92s3r",
			"name": "project",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "tqlcamhhas2xzr7",
				"cascadeDelete": false,
				"minSelect": 1,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), edit_project)
		collection.Schema.AddField(edit_project)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("6po51li0eihk83d")
		if err != nil {
			return err
		}

		// update
		edit_project := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "hwh92s3r",
			"name": "project",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "tqlcamhhas2xzr7",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": null
			}
		}`), edit_project)
		collection.Schema.AddField(edit_project)

		return dao.SaveCollection(collection)
	})
}

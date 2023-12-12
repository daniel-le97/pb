package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "67nwlz2znzqld2l",
			"created": "2023-12-12 16:40:26.123Z",
			"updated": "2023-12-12 16:40:26.123Z",
			"name": "portainer",
			"type": "base",
			"system": false,
			"schema": [
				{
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
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("67nwlz2znzqld2l")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}

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
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("y3zwk9rp46aq6vy")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	}, func(db dbx.Builder) error {
		jsonData := `{
			"id": "y3zwk9rp46aq6vy",
			"created": "2023-12-12 20:40:23.821Z",
			"updated": "2023-12-12 20:40:23.821Z",
			"name": "team",
			"type": "auth",
			"system": false,
			"schema": [],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"allowEmailAuth": true,
				"allowOAuth2Auth": true,
				"allowUsernameAuth": true,
				"exceptEmailDomains": [],
				"manageRule": null,
				"minPasswordLength": 8,
				"onlyEmailDomains": [],
				"onlyVerified": false,
				"requireEmail": false
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	})
}

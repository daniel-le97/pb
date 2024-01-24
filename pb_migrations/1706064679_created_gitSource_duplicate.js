/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "5c1oykad2csm1dk",
    "created": "2024-01-24 02:51:19.316Z",
    "updated": "2024-01-24 02:51:19.316Z",
    "name": "gitSource_duplicate",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "7te191mo",
        "name": "app_id",
        "type": "text",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "gkd9vvw8",
        "name": "webhook_secret",
        "type": "text",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "yfrikmiq",
        "name": "private_key",
        "type": "text",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
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
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("5c1oykad2csm1dk");

  return dao.deleteCollection(collection);
})

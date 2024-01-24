/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("hlr0zuxvvfokdk8");

  return dao.deleteCollection(collection);
}, (db) => {
  const collection = new Collection({
    "id": "hlr0zuxvvfokdk8",
    "created": "2024-01-24 02:36:44.686Z",
    "updated": "2024-01-24 02:42:32.179Z",
    "name": "Queue",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "8evghbvx",
        "name": "status",
        "type": "select",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "maxSelect": 1,
          "values": [
            "enqueued",
            "failed",
            "completed"
          ]
        }
      },
      {
        "system": false,
        "id": "u4xyqoiq",
        "name": "project",
        "type": "relation",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "collectionId": "k65kv3llr2oynh7",
          "cascadeDelete": false,
          "minSelect": null,
          "maxSelect": 1,
          "displayFields": null
        }
      },
      {
        "system": false,
        "id": "e23yzums",
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
      },
      {
        "system": false,
        "id": "4yf2haqz",
        "name": "buildTime",
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
})

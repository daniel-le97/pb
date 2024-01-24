/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "ojmqge2b4cj4ywj",
    "created": "2024-01-24 02:51:19.316Z",
    "updated": "2024-01-24 02:51:19.316Z",
    "name": "queue",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "jblknvez",
        "name": "project",
        "type": "relation",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "collectionId": "tqlcamhhas2xzr7",
          "cascadeDelete": false,
          "minSelect": null,
          "maxSelect": 1,
          "displayFields": null
        }
      },
      {
        "system": false,
        "id": "lxzcku0m",
        "name": "active",
        "type": "bool",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {}
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
  const collection = dao.findCollectionByNameOrId("ojmqge2b4cj4ywj");

  return dao.deleteCollection(collection);
})

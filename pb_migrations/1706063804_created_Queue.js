/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "hlr0zuxvvfokdk8",
    "created": "2024-01-24 02:36:44.686Z",
    "updated": "2024-01-24 02:36:44.686Z",
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
  const collection = dao.findCollectionByNameOrId("hlr0zuxvvfokdk8");

  return dao.deleteCollection(collection);
})

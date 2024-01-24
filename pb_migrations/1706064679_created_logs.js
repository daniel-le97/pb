/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "6po51li0eihk83d",
    "created": "2024-01-24 02:51:19.315Z",
    "updated": "2024-01-24 02:51:19.315Z",
    "name": "logs",
    "type": "base",
    "system": false,
    "schema": [
      {
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
  const collection = dao.findCollectionByNameOrId("6po51li0eihk83d");

  return dao.deleteCollection(collection);
})

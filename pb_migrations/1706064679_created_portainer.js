/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "67nwlz2znzqld2l",
    "created": "2024-01-24 02:51:19.316Z",
    "updated": "2024-01-24 02:51:19.316Z",
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
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("67nwlz2znzqld2l");

  return dao.deleteCollection(collection);
})

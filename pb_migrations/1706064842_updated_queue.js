/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ojmqge2b4cj4ywj")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "vzxfbjdv",
    "name": "status",
    "type": "select",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "values": [
        "finished",
        "failed",
        "in-queue"
      ]
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ojmqge2b4cj4ywj")

  // remove
  collection.schema.removeField("vzxfbjdv")

  return dao.saveCollection(collection)
})

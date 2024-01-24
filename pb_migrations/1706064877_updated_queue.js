/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ojmqge2b4cj4ywj")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "6uhw6urm",
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
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("ojmqge2b4cj4ywj")

  // remove
  collection.schema.removeField("6uhw6urm")

  return dao.saveCollection(collection)
})

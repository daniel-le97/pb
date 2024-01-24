/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("hlr0zuxvvfokdk8")

  // add
  collection.schema.addField(new SchemaField({
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
  }))

  // add
  collection.schema.addField(new SchemaField({
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
  }))

  // add
  collection.schema.addField(new SchemaField({
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
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("hlr0zuxvvfokdk8")

  // remove
  collection.schema.removeField("u4xyqoiq")

  // remove
  collection.schema.removeField("e23yzums")

  // remove
  collection.schema.removeField("4yf2haqz")

  return dao.saveCollection(collection)
})

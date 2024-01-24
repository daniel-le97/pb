/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("tqlcamhhas2xzr7")

  // remove
  collection.schema.removeField("yygtahnw")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "wnzh6cbk",
    "name": "startCommand",
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
    "id": "jzhqgjr9",
    "name": "buildCommand",
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
    "id": "bti7wkff",
    "name": "installCommand",
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
  const collection = dao.findCollectionByNameOrId("tqlcamhhas2xzr7")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "yygtahnw",
    "name": "application",
    "type": "json",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "maxSize": 2000000
    }
  }))

  // remove
  collection.schema.removeField("wnzh6cbk")

  // remove
  collection.schema.removeField("jzhqgjr9")

  // remove
  collection.schema.removeField("bti7wkff")

  return dao.saveCollection(collection)
})

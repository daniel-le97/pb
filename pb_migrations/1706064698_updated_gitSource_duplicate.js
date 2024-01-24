/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("5c1oykad2csm1dk")

  collection.name = "gitSource"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("5c1oykad2csm1dk")

  collection.name = "gitSource_duplicate"

  return dao.saveCollection(collection)
})

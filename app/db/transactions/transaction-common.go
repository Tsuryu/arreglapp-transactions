package transactions

import "github.com/Tsuryu/arreglapp-transactions/app/db"

var database = db.Connection.Database("arreglapp")

// Collection : collection database mongo name
var Collection = database.Collection("transaction")

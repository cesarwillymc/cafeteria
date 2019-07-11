#! /bin/bash
mongorestore -d users -c users /backup/users/users/users.bson
mongorestore -d ventas -c ventas /backup/ventas/ventas/ventas.bson
mongorestore -d productos -c productos /backup/productos/productos/productos.bson
mongorestore -d reportes -c reportes /backup/reportes/reportes/reportes.bson

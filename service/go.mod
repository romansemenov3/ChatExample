module service

go 1.14

require (
	common v0.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/gorilla/websocket v1.4.2
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	go.mongodb.org/mongo-driver v1.3.2
	model v0.0.0
)

replace common v0.0.0 => ../common

replace model v0.0.0 => ../model

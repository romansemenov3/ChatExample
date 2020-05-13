module Chat

go 1.14

require (
	api v0.0.0
	common v0.0.0
	github.com/gofrs/uuid v3.2.0+incompatible
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71 // indirect
	go.mongodb.org/mongo-driver v1.3.2 // indirect
	model v0.0.0
	service v0.0.0
)

replace api v0.0.0 => ./api

replace common v0.0.0 => ./common

replace service v0.0.0 => ./service

replace model v0.0.0 => ./model

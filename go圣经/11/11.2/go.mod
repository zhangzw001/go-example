module word

go 1.13

replace word => ./word

replace storageMail => ./storageMail

require (
	github.com/streadway/amqp v1.0.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

# simple-messaging
simple example of exchanging messages/events through kafka

# requirements
- apache kafka
- confluent-kafka-client

# how to run 
- compile by running `go build main.go` in yout terminal.
- and then run `./main`

# possible errors
- `%5|1606994571.933|CONFWARN|rdkafka#producer-1| [thrd:app]: No `bootstrap.servers` configured: client will not be able to connect to Kafka cluster`
- `|1606994129.988|FAIL|rdkafka#producer-1| [thrd::9092/0]: :9092/0: Failed to resolve ':9092': nodename nor servname provided, or not known (after 0ms in state CONNECT)`

# solution
- https://github.com/confluentinc/confluent-kafka-go/issues/574

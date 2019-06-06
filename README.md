# rabbitmq-worker-pool

# Description
Example of using worker pool to consume rabbitmq message. The pattern to build the worker pool is based on these two writings :
- http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
- https://medium.com/@j.d.livni/write-a-go-worker-pool-in-15-minutes-c9b42f640923

# Usage
1. Run file `main.go` file to start consuming
2. Run file `send.go` file to send some messages

By default, the consumer will consume from the channel specified within `config.go` file. Also, the number of workers spawned is 2. 
You could change it inside the `config.go` file or add it if you want to customize it.

# Contributions
Any PR and issue are appreciated. Thanks

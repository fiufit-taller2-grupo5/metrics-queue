# Metrics Queue

This project receives the metrics triggered from other services. When something that want to be monitored happens, it sends a message to the queue. This service will enqueue it using a Redis FIFO queue, then the `metrics-consumer` will dequeue all those metrics once a minute, and saved them aggregated into a MongoDB database.

The metrics queue has a really simple API: It works with HTTP, and the only endpoint is a POST to `/api/metrics/system`, with a payload with the following form:

```json
{
  "metric_name": "user_created"
}
```

This will enqueue a metric with the name `user_created`. There are no limitations on the metric name, although the convention is to use snake case. To create new metrics you just have to enqueue it, and the `metrics-consumer` will create it automatically in the database.

## Running the project

To run it, first update the dependencies with the command `go get -u`, and then run it with `go run main.go`.

Good Luck!

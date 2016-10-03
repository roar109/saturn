# saturn

### GET
	curl http://localhost:8088
	Ping pong, hello there!

## Operations

### POST

	http://localhost:8088/job

### POST

	http://localhost:8088/sjob

**Payload for both**

	{"name":"some_job_name","payload":"some payload","pattern":"pending the format here"}

## Roadmap ##


- Other execution providers (now only g pubsub)
- Join the scheduled job and simple job in single entity
- Simple API for scheduled jobs
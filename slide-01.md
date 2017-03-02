# Distributed Tracing Workshop

- Distributed systems:
	* Collection of *independent* components/services
	* Illusion of a global single system

- Compared to a monolithic system it's harder to:
	* Debug,
	* Profile,
	* Understand.

- Logs:
	* Record of sequential events
	* Never ending append only file
	* Hard to understand when written concurrently/parallel

- Traces:
	* Tells a story (Begins, ends)
	* Requests, Transacations
	* DT, one story told by many.
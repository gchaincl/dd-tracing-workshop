

                  +---------------+         +---------------+       +---------------+         +---------------+
 POST /users/{id} |    SRV-1      |         |    SRV-2      |       |      DB       |         |   Async Job   |
 ----------------->               ----------|               --------|               ----------|               |
                  +---|-----------+         +---------------+       +---------------+         +---------------+
                      |                         |                       |                         |
                      |                         |                       |                         |
                      |                         |                       |                         |
                      v                         |                       |                         |
                      +-------------------------.-----------------------.-------------------------.---------------+
                      |                         .                       .                         .               |
                      |  StartSpan("srv-1")     .                       .                         .               |  <-----------------+
                      |                         .                       .                         .               |                    |
                      +-------------------------.-----------------------.-------------------------.---------------+                    |
                                                |                       |                     ^   |                                    |
                                                |                       |                     |   |                                    |
                                                v                       |                     |   |                                    |
                                                +-----------------------.---------------------|   |                                    |
                                                |  StartSpan("srv-2", parent("srv-1")         |   |                                    |
                                                |                       .                     |   |                                    |
                                                |                       .                     |   |                                    |
                                                +-----------------------.---------------------+   |                                    |
                                                                        |                     ^   |                                    |
                                                                        |                     |   |                                    |
                                                                        v                     |   |                                    |
                                                                        +---------------------+   +------------------------------------+
                                                                        |                     |   |                                    |
                                                                        |  StartSpan("db",    |   |                                    |
                                                                        |    parent("srv-2"), |   | StartSpan("async", parent("srv-2"))|
                                                                        |  )                  |   |                                    |
                                                                        |                     |   |                                    |
                                                                        +---------------------+   +------------------------------------+

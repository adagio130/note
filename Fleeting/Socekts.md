**Types of Sockets**:

- **Stream Sockets (TCP Sockets)**: These provide reliable, two-way, connection-oriented byte streams. They guarantee that data will be delivered in the order it was sent and without duplication.
- **Datagram Sockets (UDP Sockets)**: These are used for connectionless communication. They're suitable for scenarios where speed is crucial and occasional data loss is acceptable.
- **Raw Sockets**: These offer more control over network communication, allowing direct sending and receiving of packets at the IP layer. They're typically used for custom low-level protocol development.
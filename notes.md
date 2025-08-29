### Hashing

Hashing is a one-way cryptographic process that takes an input and produces a **fixed-length** output
called a hash or **digest** (e.g., a 32-byte value for SHA-256).

#### Purpose

Ensures data **integrity** by creating a **unique fingerprint** of the input. If the input changes even slightly, the hash changes completely.

- Hashing not involve keys meaning it doesn’t use encryption keys; it’s a public operation. 
- Hashing not prove authenticity on its own.

The process is **deterministic** (same input = same hash) and collision-resistant (hard to find two inputs with the same hash).

### Signing
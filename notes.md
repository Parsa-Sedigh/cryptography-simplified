### Encryption
- Ensures **confidentially** because only the private-key holder(receiver - recipient) can decrypt
- Doesn't prove **authenticity** because **anyone** can encrypt with the public key (it’s public!).

### Hashing

Hashing is a one-way cryptographic process that takes an input and produces a **fixed-length** output
called a hash or **digest** (e.g., a 32-byte value for SHA-256).

#### Purpose

Ensures data **integrity** by creating a **unique fingerprint** of the input. If the input changes even slightly, the hash changes completely.

- Hashing not involve keys meaning it doesn’t use encryption keys; it’s a public operation. 
- Hashing not prove authenticity on its own.

The process is **deterministic** (same input = same hash) and collision-resistant (hard to find two inputs with the same hash).

### Signing

NOTE: 
- encryption uses the public key to encrypt and private key to decrypt. But signing uses the private key to sign and public key
to verify. Now since encryption uses the public key to encrypt and everyone can have the public key, the encrypted msg doesn't
prove authenticity. But signing proves authenticity because only the trusted source can have the private key.

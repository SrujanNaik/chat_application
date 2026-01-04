
# Database Structure – Encrypted Chat System

## Overview

This document describes the database design for an **end-to-end encrypted (E2EE) chat application**.  
The system is designed so that **no plaintext user messages are ever stored on the server**.  
Even in the event of a database breach, user data remains secure.

The database is used **only for metadata, ordering, indexing, and consistency**, not for message content inspection.

---

## Design Principles

- 🔐 **End-to-End Encryption (E2EE)**
  - Messages are encrypted on the client.
  - Server stores only encrypted blobs.
  - Encryption keys are never stored in plaintext on the server.

- 📜 **Append-only message storage**
  - Messages are immutable after insertion.
  - Ensures ordering and prevents accidental overwrites.

- 📦 **Lazy loading / pagination**
  - Only recent messages are loaded initially.
  - Older messages are fetched on demand.

- 🔄 **Offline-first sync**
  - Clients can send messages while offline.
  - Messages are synced when connectivity is restored.

---

## Database Choice

- **Primary Database**: PostgreSQL (Relational Database)
- **Optional Cache / Realtime Layer**: Redis

PostgreSQL is used for its strong consistency, indexing support, transactional guarantees, and reliable pagination.

---

## Entity Relationship Overview



---

## Tables

### 1. users

Stores user identity and cryptographic public keys.


**Notes**
- `public_key` is used to encrypt chat keys.
- No passwords are stored here if using external auth or token-based auth.

---

### 2. chats

Represents a conversation (one-to-one or group).


**Notes**
- No chat name or description is required for core functionality.
- Metadata can be extended later.

---

### 3. chat_members

Links users to chats and stores encrypted chat keys.


**Notes**
- `encrypted_key` is the symmetric chat key encrypted using the user’s public key.
- Server cannot decrypt chat messages.

---

### 4. messages

Stores encrypted messages and metadata required for ordering and pagination.


**Notes**
- `ciphertext` contains encrypted message content.
- `nonce` is required for authenticated encryption (e.g., AES-GCM).
- No plaintext message data is stored.
- Messages are immutable.

---

## Pagination Strategy (Lazy Loading)

Messages are fetched using **cursor-based pagination**.

### Example:
- Initial load: last 50 messages
- Older messages:

This avoids loading entire chat histories and scales efficiently.

---

## Offline Sync Strategy

- Client assigns temporary IDs to messages while offline.
- On reconnection:
1. Client sends encrypted pending messages.
2. Server appends them to the messages table.
3. Server returns authoritative message IDs.
4. Client updates local state.

---

## Security Considerations

- Server cannot read message contents.
- Database breach exposes only encrypted data.
- Message integrity can be enhanced using:
- message hashes
- chained hashes (hash(prev_message))

---

## Future Improvements

- Add Redis for:
- online presence
- last message caching
- WebSocket fan-out
- Add message deletion flags (soft delete)
- Add delivery receipts (encrypted)
- Migrate storage backend without breaking E2EE

---

## Summary

This database design prioritizes:
- Privacy
- Security
- Correctness
- Scalability

It intentionally limits server-side intelligence to protect user data while remaining performant and maintainable.

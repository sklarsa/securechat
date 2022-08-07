# SecureChat

## Introduction

SecureChat is an e2e-encrypted, p2p chat that is built on top of IPFS for a distributed, highly available data store.

Users are identified by public keys, which are also used for encrypting and decrypting messages.  To simplify user creation, users can generate pub/priv keypairs using the application itself.

Currently, only text messaging will be supported, but the aim is to support multimedia content like images, gifs, and videos.

## Details

## Questions

1. Use built-in ipfs/ipns keystore for identity or a separate private key managed by this executable?
2. Which serialization format to use for messages?  Include public key with message (or at least a pointer to it)?
    a. Json
    b. Protobuf (<https://github.com/golang/protobuf>)
    c. gob (<https://pkg.go.dev/encoding/gob>)
3. If using pub/priv key encryption, how to support channels with multiple users?

/*
Package nutsdb implements a simple, fast, embeddable and persistent key/value store
written in pure Go. It supports fully serializable transactions.

The design of NutsDB is based on the bitcask storage engine model, and to do some optimization,
using B+ tree instead of hash index, so it supports range scanning and prefix scanning,
using mmap technology to improve write performance.

NutsDB currently works on Mac OS X and Linux.

Usage

NutsDB has the following main types: DB, BPTree, Entry, DataFile And Tx. and NutsDB supports bucket, A bucket is
a collection of unique keys that are associated with values.

All operations happen inside a Tx. Tx represents a transaction, which can
be read-only or read-write. Read-only transactions can read values for a
given key , or iterate over a set of key-value pairs (prefix scanning or range scanning).
read-write transactions can also update and delete keys from the DB.

See the examples for more usage details.
*/
package nutsdb

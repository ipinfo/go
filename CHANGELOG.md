# 2.1.1

- Fixed go module path to have "v2" at the end as necessary.

# 2.1.0

- A new field `CountryName` was added to both `ASNDetails` and `Core`, which
  is the full name of the country abbreviated in the existing `Country` field.
  For example, if `Country == "PK"`, now `CountryName == "Pakistan"` exists.

# 2.0.0

- The API for creating a client and making certain requests has changed and has
  been made generally simpler. Please see the documentation for exact details.
- go.mod now included.
- All new API data types are now available for the Core & ASN APIs.
- Cache interface now requires implementors to be concurrency-safe.

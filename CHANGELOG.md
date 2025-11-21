# 2.12.0

- Add support for IPinfo Core API
- Add support for IPinfo Plus API

# 2.11.0

- Add support for IPinfo Lite API

# 2.10.0

- Added support for looking up your v6 IP via `GetIPInfoV6`.

# 2.9.4

- Removing null values while printing in `yaml` format

# 2.9.3

- Added a `CountryFlagURL` field to `Core`.

# 2.9.2

- Custom error message on 429(Too many requests).
- Check bogon locally.

# 2.9.1

- Fixed default useragent.
- Fixed `GetIPInfoBatch` panic on empty token or invalid IP.

# 2.9.0

- Added an `CountryFlag` field to `Core`.
- Added an `CountryCurrency` field to `Core`.
- Added an `Continent` field to `Core`.

# 2.8.0

- Added an `IsEU` field to `Core`, which checks whether the IP geolocates to a
  country within the European Union (EU).

# 2.7.0

- Made batch operations limit their concurrency to 8 batches by default, but
  configurable.

# 2.6.0

- Added `Relay` and `Service` fields to `CorePrivacy`.
- Added `Relay` field to `IPSummary.Privacy` and `PrivacyServices` to
  `IPSummary`.

# 2.5.4

- Fixed issue where disabling per-batch timeouts was impossible with negative
  numbers, contrary to what the documentation says.

# 2.5.3

- Dummy release to make up for a bug in 2.5.2.

# 2.5.2 (skip this release)

- Removed the IP list length constraints on `GetIPSummary`.
  This is because the underlying API has changed.

# 2.5.1

- Added the `IPSummary.Domains` field.

# 2.5.0

- Added versioned cache keys.
  This allows more reliable changes to cached data in the future without
  causing confusing incompatibilities. This should be transparent to the user.

# 2.4.0

- Added support for IP Map API.

# 2.3.2

- Added the `Core.Bogon` boolean field.

# 2.3.1

- Added more summary fields (carrier & mobile data).

# 2.3.0

- Added support for IP summary API.

# 2.2.3

- Added CSV tags for `Core` data for easier CSV marshaling.
- Omit empty `Core` objects when encoding JSON.
- Encode `Core.CountryName` and `Core.Abuse.CountryName` properly in JSON.

# 2.2.2

- Added a function `GetCountryName` to transform country code into full name.

# 2.2.1

- Added the `Core.Anycast` boolean field.

# 2.2.0

- The following functions are now private:
  - `Client.Do`
  - `Client.NewRequest`
  - `CheckResponse`
- The following **new** functions now exist, which operate on the IPinfo
  `/batch` endpoint:
  - `Client.GetBatch`
  - `Client.GetIPInfoBatch`
  - `Client.GetIPStrInfoBatch`
  - `Client.GetASNDetailsBatch`
  - `ipinfo.GetBatch`
  - `ipinfo.GetIPInfoBatch`
  - `ipinfo.GetIPStrInfoBatch`
  - `ipinfo.GetASNDetailsBatch`

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

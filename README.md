# Geohash

Implementation (under development) of the [geohash](https://en.wikipedia.org/wiki/Geohash) concept.

Currently, it implements the `Encode()` function that let's you encode a lat/lon pair into a geohash of length 12.

```go
func Encode(lat, lon float64) string
```

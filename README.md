# gmapsign

[![Go Report Card](https://goreportcard.com/badge/github.com/sgarcez/gmapsign)](https://goreportcard.com/report/github.com/sgarcez/gmapsign)


`gmapsign` is a Go pkg and cli tool to sign Google Maps API request URLs.

## CLI Usage

```console
◊ echo "https://maps.googleapis.com/maps/api/staticmap?center=40.714%2c%20-73.998&zoom=12&size=400x400&client=myclient" | gmapsign -key bXlrZXk=

https://maps.googleapis.com/maps/api/staticmap?center=40.714%2C+-73.998&client=myclient&size=400x400&zoom=12&signature=C1UG9w-v41q7i1lISpZsw1xzOV8=
```

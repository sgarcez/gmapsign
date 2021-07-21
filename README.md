# gmapsign

[![Go Report Card](https://goreportcard.com/badge/github.com/sgarcez/gmapsign)](https://goreportcard.com/report/github.com/sgarcez/gmapsign)


_gmapsign_ is a Go pkg and cli tool to [sign Google Maps API request URLs](https://developers.google.com/maps/premium/overview#digital_signatures).

This is required when using:

- A client ID with the web service APIs, Maps Static API, or Street View Static API

- An API key with the Maps Static API or Street View Static API

## Installation

### Go Get

```console
◊ go install github.com/sgarcez/gmapsign/cmd/gmapsign@latest
```

## Usage

```console
◊ echo "https://maps.googleapis.com/maps/api/staticmap?center=40.714%2C+-73.998&client=myclient&size=400x400&zoom=12" \
  | gmapsign -key bXlrZXk=

https://maps.googleapis.com/maps/api/staticmap?center=40.714%2C+-73.998&client=myclient&size=400x400&zoom=12&signature=C1UG9w-v41q7i1lISpZsw1xzOV8=
```


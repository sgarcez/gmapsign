package gmapsign

import (
	"bytes"
	"strings"
	"testing"
)

func TestSign(t *testing.T) {
	url := "https://maps.googleapis.com/maps/api/staticmap?center=40.714%2c%20-73.998&zoom=12&size=400x400&client=myclient"

	key, err := DecodeSigningKey("bXlrZXk=")
	if err != nil {
		t.Errorf("decoding signing key: %w", err)
	}

	got, err := Sign(url, key)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	exp := "https://maps.googleapis.com/maps/api/staticmap?center=40.714%2C+-73.998&client=myclient&size=400x400&zoom=12&signature=C1UG9w-v41q7i1lISpZsw1xzOV8="

	if got != exp {
		t.Errorf("expected: %q, got: %q", exp, got)
	}
}

func TestPipeline(t *testing.T) {
	key, err := DecodeSigningKey("bXlrZXk=")
	if err != nil {
		t.Errorf("decoding signing key: %w", err)
	}

	in := strings.NewReader(`
	https://maps.googleapis.com/maps/api/staticmap?center=40.714%2c%20-73.998&zoom=12&size=400x400&client=myclient

	https://maps.googleapis.com/maps/api/distancematrix/json?departure_time=now&destinations=38.62288639085796%2C-9.223203416168769&mode=driving&origins=38.62288639085796%2C-9.223203416168769&units=metric`)

	out := &bytes.Buffer{}

	err = Pipeline(in, out, key)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	exp := `https://maps.googleapis.com/maps/api/staticmap?center=40.714%2C+-73.998&client=myclient&size=400x400&zoom=12&signature=C1UG9w-v41q7i1lISpZsw1xzOV8=
https://maps.googleapis.com/maps/api/distancematrix/json?departure_time=now&destinations=38.62288639085796%2C-9.223203416168769&mode=driving&origins=38.62288639085796%2C-9.223203416168769&units=metric&signature=kGApffVXsWUUbRv_lw3UnrejoOI=
`

	if out.String() != exp {
		t.Errorf("expected: %q, got: %q", exp, out.String())
	}
}

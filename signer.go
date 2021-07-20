// Package gmapsign provides helpers for signing requests to Google Maps APIs.
// The signature is included as the last query string var, as per Google's documentation.
package gmapsign

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha1" // nolint: gosec
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
)

// Pipeline signs incoming urls from the io.Reader and writes them in the io.Writer.
func Pipeline(in io.Reader, out io.Writer, signKey []byte) error {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		signed, err := Sign(line, signKey)
		if err != nil {
			return err
		}
		_, err = io.WriteString(out, signed+"\n")
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}

// Sign signs a string based URL.
func Sign(raw string, key []byte) (string, error) {
	u, err := parseURL(raw)
	if err != nil {
		return "", fmt.Errorf("parsing: %w", err)
	}

	if err := SignURL(u, key); err != nil {
		return "", fmt.Errorf("signing: %w", err)
	}

	return u.String(), nil
}

func parseURL(in string) (*url.URL, error) {
	in = strings.TrimSpace(in)

	// replace literal unicode e.g. \u0026 -> &
	in, err := strconv.Unquote("\"" + in + "\"")
	if err != nil {
		fmt.Println(in)
		return nil, err
	}

	u, err := url.Parse(in)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// SignURL signs a *url.URL.
func SignURL(u *url.URL, key []byte) error {
	q := u.Query()
	q.Del("signature")
	encodedQuery := q.Encode()
	s, err := sig(fmt.Sprintf("%s?%s", u.Path, encodedQuery), key)
	if err != nil {
		return err
	}
	u.RawQuery = fmt.Sprintf("%s&signature=%s", encodedQuery, s)
	return nil
}

func sig(msg string, key []byte) (string, error) {
	mac := hmac.New(sha1.New, key)
	_, err := mac.Write([]byte(msg))
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(mac.Sum(nil)), nil
}

// DecodeSigningKey base64 decodes a signing key.
func DecodeSigningKey(raw string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(raw)
}

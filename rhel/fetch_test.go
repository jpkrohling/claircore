package rhel

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/quay/claircore/libvuln/driver"
)

func TestFetch(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/Red_Hat_Enterprise_Linux_3.xml")
	}))
	defer srv.Close()

	t.Run("FetchContext", func(t *testing.T) {
		u, err := NewUpdater(3, WithClient(srv.Client()), WithURL(srv.URL, ""))
		if err != nil {
			t.Fatal(err)
		}
		rd, hint, err := u.FetchContext(ctx, driver.Fingerprint(""))
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("got fingerprint: %+v", hint)
		defer rd.Close()
		n, err := io.Copy(ioutil.Discard, rd)
		switch {
		case err != nil:
			t.Fatalf("unable to read returned data: %v", err)
		case n == 0:
			t.Fatalf("expected more data than %d bytes", n)
		}

		rd, got, err := u.FetchContext(ctx, hint)
		t.Logf("got fingerprint: %+v", got)
		t.Logf("returned expected error: %v", err)
		if err != driver.Unchanged {
			rd.Close()
			t.Log("resource changed unexpectedly")
			t.Fatalf("%x != %x", got, hint)
		}
	})

	t.Run("Fetch", func(t *testing.T) {
		u, err := NewUpdater(3, WithClient(srv.Client()), WithURL(srv.URL, ""))
		if err != nil {
			t.Fatal(err)
		}
		rd, hint, err := u.Fetch()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("got fingerprint: %+v", hint)
		defer rd.Close()
		n, err := io.Copy(ioutil.Discard, rd)
		switch {
		case err != nil:
			t.Fatalf("unable to read returned data: %v", err)
		case n == 0:
			t.Fatalf("expected more data than %d bytes", n)
		}

		rd, got, err := u.Fetch()
		t.Logf("got fingerprint: %+v", got)
		if err != nil {
			t.Fatal(err)
		}
		defer rd.Close()
		if hint != got {
			t.Log("resource changed unexpectedly")
			t.Fatalf("%x != %x", got, hint)
		}
	})
}

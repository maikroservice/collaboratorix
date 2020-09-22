package dnsrcv_test

import app "agner.io/boast"

type mockStorage struct{}

var tID = "mpqhomfbxab55m5de32mywvfoy"
var tCanary = "k2b27meg7dfifvxuxmnfnm24oa"

func (s *mockStorage) SetTest(secret []byte) (id string, canary string, err error) {
	return "", "", nil
}

func (s *mockStorage) LoadEvents(id string) (evts []app.Event, loaded bool) {
	return evts, false
}

func (s *mockStorage) SearchTest(f func(k, v string) bool) (id string, canary string) {
	return tID, tCanary
}

func (s *mockStorage) StoreEvent(evt app.Event) error {
	return nil
}

func (s *mockStorage) TotalTests() int {
	return 0
}

func (s *mockStorage) TotalEvents() int {
	return 0
}

func (s *mockStorage) StartExpire(err chan error) {}

package main

import "testing"

func TestAddNode(t *testing.T) {
	r := NewRing()

	tests := []struct {
		desc        string
		Host        string
		ExpectedLen int
	}{
		{
			desc:        "When a node is not present",
			Host:        "xyz",
			ExpectedLen: 1,
		},
		{
			desc:        "When a node is present",
			Host:        "xyz",
			ExpectedLen: 1,
		},
	}

	for _, tt := range tests {

		t.Run(tt.desc, func(t *testing.T) {

			r.addNode(tt.Host)

			if tt.ExpectedLen != r.Nodes.Len() {
				t.Errorf("Expected %d Got %d", tt.ExpectedLen, r.Nodes.Len())
			}

		})
	}

}

func TestRemoveNode(t *testing.T) {
	r := NewRing()
	r.Nodes = append(r.Nodes, NewNode("abc"))

	tests := []struct {
		desc    string
		Host    string
		WantErr bool
	}{
		{
			desc:    "When an existing node is being removed",
			Host:    "abc",
			WantErr: false,
		},
		{
			desc:    "When a non existing node is being removed",
			Host:    "xyz",
			WantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.desc, func(t *testing.T) {

			err := r.removeNode(tt.Host)

			if tt.WantErr != (err != nil) {
				t.Errorf("Unexpected err %v", err)
			}

		})
	}

}

func TestGetKey(t *testing.T) {
	r := NewRing()
	r.Nodes = append(r.Nodes, NewNode("abc"), NewNode("xyz"), NewNode("pqr"))

	tests := []struct {
		desc         string
		Key          string
		expectedHost string
	}{
		{
			desc:         "When for a given key Get is called",
			Key:          "abc",
			expectedHost: "abc",
		},
	}

	for _, tt := range tests {

		t.Run(tt.desc, func(t *testing.T) {

			host := r.Get(tt.Key)

			if host != tt.expectedHost {
				t.Errorf("Wanted %s got %s", tt.expectedHost, host)
			}

		})
	}

}

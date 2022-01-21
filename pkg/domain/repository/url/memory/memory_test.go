// Package memory is an in-memory implementation of the URL repository.
package memory

import (
	"testing"
	"url-shortner/pkg/domain/aggregate"
	"url-shortner/pkg/domain/repository/url"
)

func TestMemoryRepository_Add(t *testing.T) {
	googleURL, err := aggregate.NewURL("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		urls map[string]aggregate.URL
	}
	type args struct {
		in0 aggregate.URL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "New URL",
			fields: fields{
				urls: make(map[string]aggregate.URL),
			},
			args: args{
				in0: googleURL,
			},
			wantErr: false,
		},
		{
			name: "URL exists",
			fields: fields{
				urls: map[string]aggregate.URL{googleURL.GetID(): googleURL},
			},
			args: args{
				in0: googleURL,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := NewMemoryRepository()
			mr.urls = tt.fields.urls

			if err := mr.Add(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("MemoryRepository.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryRepository_Get(t *testing.T) {
	googleURL, err := aggregate.NewURL("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		urls map[string]aggregate.URL
	}
	type args struct {
		shortCode string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        aggregate.URL
		wantErr     bool
		expectedErr error
	}{
		{
			name: "URL exists",
			fields: fields{
				urls: map[string]aggregate.URL{googleURL.GetID(): googleURL},
			},
			args: args{
				shortCode: googleURL.GetID(),
			},
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name: "URL does not exist",
			fields: fields{
				urls: make(map[string]aggregate.URL),
			},
			args: args{
				shortCode: googleURL.GetID(),
			},
			wantErr:     true,
			expectedErr: url.ErrURLNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := NewMemoryRepository()
			mr.urls = tt.fields.urls

			_, err := mr.Get(tt.args.shortCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemoryRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr == true {
				if err != tt.expectedErr {
					t.Errorf("MemoryRepository.Get() error = %v, expectedErr %v", err, tt.expectedErr)
					return
				}
			}
		})
	}
}

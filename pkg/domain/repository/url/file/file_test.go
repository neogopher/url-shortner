// Package memory is a file-based implementation of the URL repository.
package file

import (
	"os"
	"testing"
	"url-shortner/pkg/domain/aggregate"
	"url-shortner/pkg/domain/repository/url"
)

func TestFileRepository_Add(t *testing.T) {
	googleURL, err := aggregate.NewURL("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		urls     map[string]aggregate.URL
		filename string
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
				urls:     make(map[string]aggregate.URL),
				filename: "test.txt",
			},
			args: args{
				in0: googleURL,
			},
			wantErr: false,
		},
		{
			name: "URL exists",
			fields: fields{
				urls:     map[string]aggregate.URL{googleURL.GetID(): googleURL},
				filename: "test.txt",
			},
			args: args{
				in0: googleURL,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// delete test file
			if _, err := os.Stat(tt.fields.filename); err == nil {
				err = os.Remove(tt.fields.filename)
				if err != nil {
					t.Fatal(err)
				}
			}

			fr := NewFileRepository(tt.fields.filename)
			fr.urls = tt.fields.urls
			fr.filename = tt.fields.filename

			if err := fr.Add(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("FileRepository.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileRepository_Get(t *testing.T) {
	googleURL, err := aggregate.NewURL("http://www.google.com")
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		urls     map[string]aggregate.URL
		filename string
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
				urls:     map[string]aggregate.URL{googleURL.GetID(): googleURL},
				filename: "test.txt",
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
				urls:     make(map[string]aggregate.URL),
				filename: "test.txt",
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
			// delete test file
			if _, err := os.Stat(tt.fields.filename); err == nil {
				err = os.Remove(tt.fields.filename)
				if err != nil {
					t.Fatal(err)
				}
			}

			fr := NewFileRepository(tt.fields.filename)
			fr.urls = tt.fields.urls
			fr.filename = tt.fields.filename

			_, err := fr.Get(tt.args.shortCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr == true {
				if err != tt.expectedErr {
					t.Errorf("FileRepository.Get() error = %v, expectedErr %v", err, tt.expectedErr)
					return
				}
			}
		})
	}
}

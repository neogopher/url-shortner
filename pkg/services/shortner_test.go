// Package services contains all services that work with Handlers and repositories to perform an action.
package services

import (
	"testing"
	"url-shortner/pkg/domain/aggregate"
	"url-shortner/pkg/domain/repository/url"
	"url-shortner/pkg/domain/repository/url/memory"
)

func TestShortnerService_Shorten(t *testing.T) {
	path := "http://www.google.com"
	googleURL, err := aggregate.NewURL(path)
	if err != nil {
		t.Fatal(err)
	}
	shortCode := googleURL.GetID()

	repo := memory.NewMemoryRepository()
	err = repo.Add(googleURL)
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		urls url.Repository
	}
	type args struct {
		fullPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Empty Path",
			fields: fields{
				urls: repo,
			},
			args: args{
				fullPath: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Invalid Path",
			fields: fields{
				urls: repo,
			},
			args: args{
				fullPath: "This is an invalid path",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Valid Path",
			fields: fields{
				urls: repo,
			},
			args: args{
				fullPath: path,
			},
			want:    shortCode,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewShortnerService(tt.fields.urls)
			got, err := s.Shorten(tt.args.fullPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortnerService.Shorten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ShortnerService.Shorten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortnerService_GetFullPath(t *testing.T) {
	path := "http://www.google.com"
	googleURL, err := aggregate.NewURL(path)
	if err != nil {
		t.Fatal(err)
	}
	shortCode := googleURL.GetID()

	repo := memory.NewMemoryRepository()
	err = repo.Add(googleURL)
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		urls url.Repository
	}
	type args struct {
		shortCode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "URL Exists",
			fields: fields{
				urls: repo,
			},
			args: args{
				shortCode: shortCode,
			},
			want:    path,
			wantErr: false,
		},
		{
			name: "URL Does not Exist",
			fields: fields{
				urls: repo,
			},
			args: args{
				shortCode: "abcd",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewShortnerService(tt.fields.urls)
			got, err := s.GetFullPath(tt.args.shortCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortnerService.GetFullPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ShortnerService.GetFullPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

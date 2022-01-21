// Package aggregate holds aggregates that combine multiple entities into full objects.
package aggregate

import (
	"testing"
	"url-shortner/internal/hash"
	"url-shortner/pkg/domain/entity"
)

func TestNewURL(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name        string
		args        args
		want        URL
		wantErr     bool
		expectedErr error
	}{
		{
			name: "Empty Path",
			args: args{
				path: "",
			},
			want:        URL{},
			wantErr:     true,
			expectedErr: ErrEmptyPath,
		},
		{
			name: "Invalid Path",
			args: args{
				path: "This is an invalid path",
			},
			want:        URL{},
			wantErr:     true,
			expectedErr: ErrInvalidPath,
		},
		{
			name: "Valid Path",
			args: args{
				path: "http://www.google.com",
			},
			want:        URL{},
			wantErr:     false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewURL(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr == true {
				if err != tt.expectedErr {
					t.Errorf("NewURL() error = %v, expectedErr %v", err, tt.expectedErr)
					return
				}
			}
		})
	}
}

func TestURL_GetID(t *testing.T) {
	path := "http://www.google.com"
	shortCode := hash.GenerateShortCode(path)

	googleURL := &URL{
		link: &entity.Link{
			ID:       shortCode,
			FullPath: path,
		},
	}

	tests := []struct {
		name string
		want string
	}{
		{
			name: "Valid ID",
			want: shortCode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := googleURL
			if got := u.GetID(); got != tt.want {
				t.Errorf("URL.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURL_GetFullPath(t *testing.T) {
	path := "http://www.google.com"
	shortCode := hash.GenerateShortCode(path)

	googleURL := &URL{
		link: &entity.Link{
			ID:       shortCode,
			FullPath: path,
		},
	}

	tests := []struct {
		name string
		want string
	}{
		{
			name: "Valid URL",
			want: path,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := googleURL
			if got := u.GetFullPath(); got != tt.want {
				t.Errorf("URL.GetFullPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

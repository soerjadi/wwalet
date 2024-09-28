package user

import (
	"reflect"
	"testing"

	"github.com/soerjadi/wwalet/internal/config"
	"github.com/soerjadi/wwalet/internal/repository/user"
)

func TestGetUsecase(t *testing.T) {
	type args struct {
		repository user.Repository
		cfg        *config.Config
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "soerja",
			want: &userUsecase{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUsecase(tt.args.repository, tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

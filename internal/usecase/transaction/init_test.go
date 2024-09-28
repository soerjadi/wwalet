package transaction

import (
	"reflect"
	"testing"

	"github.com/soerjadi/wwalet/internal/repository/transaction"
)

func TestGetUsecase(t *testing.T) {
	type args struct {
		repository transaction.Repository
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "soerja",
			want: &transactionUsecase{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUsecase(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

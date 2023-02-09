package mongo

import (
	"context"
	"reflect"
	"testing"
	"top-ranking-worker/infra"
	"top-ranking-worker/lineup/domain"
)

func TestCalculator_Calculate(t *testing.T) {
	type fields struct {
		Database *infra.MongoDatabase
	}
	type args struct {
		ctx          context.Context
		contents     []domain.Content
		interactions []domain.Interaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *map[int]float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &Calculator{
				Database: tt.fields.Database,
			}
			got, err := lc.Calculate(tt.args.ctx, tt.args.contents, tt.args.interactions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

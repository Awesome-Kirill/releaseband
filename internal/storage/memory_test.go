package storage

import (
	"reflect"
	"releaseband/internal/domain"
	"sync"
	"testing"
)

func TestInMemory_Get(t *testing.T) {
	type fields struct {
		data map[string]*domain.GameDate
	}
	type args struct {
		id string
	}
	var tests []struct {
		want   *domain.GameDate
		name   string
		fields fields
		args   args
		want1  bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &InMemory{
				mu:   sync.RWMutex{},
				data: tt.fields.data,
			}
			got, got1 := m.Get(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInMemory_Set(t *testing.T) {
	type fields struct {
		data map[string]*domain.GameDate
	}
	type args struct {
		input *domain.GameDate
		id    string
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &InMemory{
				mu:   sync.RWMutex{},
				data: tt.fields.data,
			}
			m.Set(tt.args.id, tt.args.input)
		})
	}
}

func TestNew(t *testing.T) {
	var tests []struct {
		want *InMemory
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

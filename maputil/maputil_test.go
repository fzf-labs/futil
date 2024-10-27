package maputil

import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want []K
	}
	var tests = []testCase[string, string]{
		{
			name: "case 1",
			args: args[string, string]{
				m: map[string]string{
					"a": "1",
					"b": "2",
				},
			},
			want: []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Keys(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want []V
	}
	tests := []testCase[string, string]{
		{
			name: "case 1",
			args: args[string, string]{
				m: map[string]string{
					"a": "1",
					"b": "2",
				},
			},
			want: []string{"1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Values(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args[K comparable, V any] struct {
		maps []map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want map[K]V
	}
	var tests = []testCase[string, string]{
		{
			name: "case 1",
			args: args[string, string]{
				maps: []map[string]string{
					{
						"a": "1",
					},
					{
						"b": "2",
					},
				},
			},
			want: map[string]string{
				"a": "1",
				"b": "2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	type args[K comparable, V any] struct {
		m        map[K]V
		iteratee func(key K, value V)
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
	}
	tests := []testCase[string, string]{
		{
			name: "case 1",
			args: args[string, string]{
				m: map[string]string{
					"a": "1",
					"b": "2",
				},
				iteratee: func(key string, value string) {
					// do nothing
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForEach(tt.args.m, tt.args.iteratee)
		})
	}
}

func TestFilter(t *testing.T) {
	type args[K comparable, V any] struct {
		m         map[K]V
		predicate func(key K, value V) bool
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want map[K]V
	}
	tests := []testCase[string, string]{
		{
			name: "case 1",
			args: args[string, string]{
				m:         map[string]string{"a": "1", "b": "2"},
				predicate: func(key string, value string) bool { return key == "a" },
			},
			want: map[string]string{"a": "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.m, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	type args[K comparable, V any] struct {
		maps []map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want map[K]V
	}
	var tests = []testCase[string, string]{
		{
			name: "case 1",
			args: args[string, string]{
				maps: []map[string]string{
					{
						"a": "1",
					},
					{
						"b": "2",
					},
				},
			},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

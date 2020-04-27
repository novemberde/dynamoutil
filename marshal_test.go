package main

import (
	"reflect"
	"testing"
)

func TestMarshalDynamo(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"string",
			args{
				map[string]interface{}{
					"attr1": map[string]interface{}{
						"S": "Test",
					},
				},
			},
			map[string]interface{}{
				"attr1": "Test",
			},
			false,
		},
		{
			"number",
			args{
				map[string]interface{}{
					"attr1": map[string]interface{}{
						"N": 123,
					},
				},
			},
			map[string]interface{}{
				"attr1": 123,
			},
			false,
		},
		{
			"list",
			args{
				map[string]interface{}{
					"attr1": map[string]interface{}{
						"L": []map[string]interface{}{
							{
								"S": "123",
							},
							{
								"N": 456,
							},
						},
					},
				},
			},
			map[string]interface{}{
				"attr1": []interface{}{
					"123", 456,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalDynamo(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalDynamo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalDynamo() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_isInteger(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{name: "1", arg: "1", want: true},
		{name: "2", arg: "9", want: true},
		{name: "3", arg: "0", want: true},
		{name: "4", arg: "5", want: true},
		{name: "5", arg: "515", want: true},
		{name: "6", arg: "515032", want: true},
		{name: "7", arg: "test", want: false},
		{name: "8", arg: "t2", want: false},
		{name: "9", arg: "32t2", want: false},
		{name: "10", arg: "-132t2", want: false},
		{name: "11", arg: "-132.2", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInteger(tt.arg); got != tt.want {
				t.Errorf("isInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcState_parse(t *testing.T) {
	type fields struct {
		Root    *Expr
		Current *Expr
	}

	tests := []struct {
		name   string
		fields fields
		tokens []string
	}{
		{
			name: "long",
			fields: fields{
				Root: &Expr{
					Operator: "GR",
					N:        1,
					Sets:     [][]int64{{1, 2, 3, 4, 5}},
					Child: &Expr{
						Operator: "EQ",
						N:        3,
						Sets:     [][]int64{{1, 2, 3}, {1, 2, 3}, {2, 3, 4}},
						Child:    nil,
					},
				},
				Current: nil,
			},
			tokens: []string{"[", "GR", "1", "c.txt", "[", "EQ", "3", "a.txt", "a.txt", "b.txt", "]", "]"},
		},
		{
			name: "simple",
			fields: fields{
				Root: &Expr{
					Operator: "EQ",
					N:        3,
					Sets:     [][]int64{{1, 2, 3}, {1, 2, 3}, {2, 3, 4}},
					Child:    nil,
				},
			},
			tokens: []string{"[", "EQ", "3", "a.txt", "a.txt", "b.txt", "]"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := &LexerCtx{}

			state.Parse(tt.tokens)

			if !reflect.DeepEqual(state.root, tt.fields.Root) {
				res, _ := json.MarshalIndent(tt.fields.Root, "", "  ")
				t.Log("EXPECTED:", string(res))
				res, _ = json.MarshalIndent(state.root, "", "  ")
				t.Log("RESULT:", string(res))
				t.Fail()
			}
		})
	}
}

func Test_extractSet(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     []int64
	}{
		{filename: "a.txt", want: []int64{1, 2, 3}},
		{filename: "b.txt", want: []int64{2, 3, 4}},
		{filename: "c.txt", want: []int64{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractSetFromFile(tt.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractSetFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

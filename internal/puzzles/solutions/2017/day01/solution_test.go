package day01

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution_Part1(t *testing.T) {
	type fields struct {
		name string
		year string
	}

	type args struct {
		input io.Reader
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: `1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2) matches the fourth digit`,
			fields: fields{
				name: "",
				year: "",
			},
			args: args{
				input: strings.NewReader("1122"),
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: `1111 produces 4 because each digit (all 1) matches the next.`,
			fields: fields{
				name: "",
				year: "",
			},
			args: args{
				input: strings.NewReader("1111"),
			},
			want:    "4",
			wantErr: false,
		},
		{
			name: `1234 produces 0 because no digit matches the next`,
			fields: fields{
				name: "",
				year: "",
			},
			args: args{
				input: strings.NewReader("1234"),
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: `91212129 produces 9 because the only digit that matches the next one is the last digit, 9`,
			fields: fields{
				name: "",
				year: "",
			},
			args: args{
				input: strings.NewReader("91212129"),
			},
			want:    "9",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			s := solution{
				name: tt.fields.name,
			}

			got, err := s.Part1(tt.args.input)
			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_solution_Part2(t *testing.T) {
	t.Skip()

	type fields struct {
		name string
	}

	type args struct {
		input io.Reader
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				name: "",
			},
			args: args{
				input: strings.NewReader(")"),
			},
			want:    "1",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			s := solution{
				name: tt.fields.name,
			}

			got, err := s.Part2(tt.args.input)
			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

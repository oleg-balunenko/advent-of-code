package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_module_fuel(t *testing.T) {
	type fields struct {
		mass int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "mass 12",
			fields: fields{
				mass: 12,
			},
			want: 2,
		},
		{
			name: "mass 14",
			fields: fields{
				mass: 14,
			},
			want: 2,
		},
		{
			name: "mass 1969",
			fields: fields{
				mass: 1969,
			},
			want: 654,
		},
		{
			name: "mass 100756",
			fields: fields{
				mass: 100756,
			},
			want: 33583,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			m := module{
				mass: tt.fields.mass,
			}

			got := m.fuel()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_calcPart1(t *testing.T) {
	in := make(chan module)
	res := make(chan int)
	done := make(chan struct{})

	go calcPart1(in, res, done)

	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "mass 12",
			input: 12,
			want:  2,
		},
		{
			name:  "mass 14",
			input: 14,
			want:  2,
		},
		{
			name:  "mass 1969",
			input: 1969,
			want:  654,
		},
		{
			name:  "mass 100756",
			input: 100756,
			want:  33583,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			in <- module{mass: tt.input}

			var got int

		LOOP:
			for {
				select {
				case r := <-res:
					got += r
				case <-done:
					break LOOP
				}
			}

			assert.Equal(t, tt.want, got)
		})
	}

	close(in)
}

func Test_calcPart2(t *testing.T) {
	in := make(chan module)
	res := make(chan int)
	done := make(chan struct{})

	go calcPart2(in, res, done)

	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "mass 12",
			input: 12,
			want:  2,
		},
		{
			name:  "mass 14",
			input: 14,
			want:  2,
		},
		{
			name:  "mass 1969",
			input: 1969,
			want:  966,
		},
		{
			name:  "mass 100756",
			input: 100756,
			want:  50346,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			in <- module{mass: tt.input}

			var got int

		LOOP:
			for {
				select {
				case r := <-res:
					got += r
				case <-done:
					break LOOP
				}
			}

			assert.Equal(t, tt.want, got)
		})
	}

	close(in)
}

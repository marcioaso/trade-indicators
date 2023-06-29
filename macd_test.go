package trade_indicators

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateMACD(t *testing.T) {
	type args struct {
		prices []float64
	}
	tests := []struct {
		name string
		args args
		want func() ([]float64, []float64, error)
	}{
		{
			name: "correct",
			args: args{
				prices: []float64{
					283.46, 280.69, 285.48, 294.08, 293.90, 299.92, 301.15, 284.45, 294.09, 302.77, 301.97, 306.85, 305.02,
					301.06, 291.97, 284.18, 286.48, 284.54, 276.82, 284.49, 275.01, 279.07, 277.85, 278.85, 283.76, 291.72,
					284.73, 291.82, 296.74, 291.13, 285.48, 294.08, 293.90, 299.92, 286.48, 284.54, 276.82, 284.49, 301.06, 291.97, 284.18,
					283.46, 280.69, 285.48, 294.08, 293.90, 299.92, 301.15, 299.92, 298, 297, 296, 295, 294, 293, 295, 294, 293, 292,
				},
			},
			want: func() ([]float64, []float64, error) {
				return []float64{-4.89, -4.53, -3.63, -2.49, -2.02, -2.07, -1.4, -0.88, 0.02, -0.34, -0.78, -1.73, -1.85, -0.59, -0.33, -0.74, -1.11, -1.61, -1.6, -0.89, -0.34, 0.58, 1.39, 1.91, 2.14, 2.22, 2.18, 2.04, 1.82, 1.56, 1.5, 1.35, 1.13, 0.87},
					[]float64{-2.43, -2.01, -1.76, -1.75, -1.77, -1.53, -1.29, -1.18, -1.17, -1.26, -1.33, -1.24, -1.06, -0.73, -0.31, 0.13, 0.53, 0.87, 1.13, 1.31, 1.41, 1.44, 1.45, 1.43, 1.37, 1.27}, nil
			},
		},
		{
			name: "err",
			args: args{
				prices: []float64{
					30,
				},
			},
			want: func() ([]float64, []float64, error) {
				return nil, nil, errors.New("prices len must be more or equal 2 periods")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			macd, signal, err := CalculateMACD(tt.args.prices)

			expectedMacd, expectedSignal, expectedErr := tt.want()
			assert.Equal(t, expectedMacd, macd)
			assert.Equal(t, expectedSignal, signal)
			assert.Equal(t, expectedErr, err)
		})
	}
}

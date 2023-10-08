package stats_test

import (
    "fmt"
    "testing"
    "slices"

    "github.com/montanaflynn/stats"
)

func TestAnscombe1LinearRegression(t *testing.T) {
    anscombe_1 := []stats.Coordinate{
        {10, 8.04},
        {8, 6.95},
        {13, 7.58},
        {9, 8.81},
        {11, 8.33},
        {14, 9.96},
        {6, 7.24},
        {4, 4.26},
        {12, 10.84},
        {7, 4.82},
        {5, 5.68},
    }

    expected_coefficients :=stats.Series{
        {10, 8.001000000000001},
        {8, 7.000818181818185},
        {13, 9.501272727272724},
        {9, 7.500909090909093},
        {11, 8.501090909090909},
        {14, 10.001363636363633},
        {6, 6.000636363636369},
        {4, 5.000454545454553},
        {12, 9.001181818181816},
        {7, 6.500727272727277},
        {5, 5.5005454545454615},
    }

    calculated_coefficients, _ := stats.LinearRegression(anscombe_1)

    if !slices.Equal(calculated_coefficients, expected_coefficients) {
        t.Errorf("Linear regression coefficients for Anscombe 2 do not match expectations\n%s\n%s",
                 fmt.Sprint(calculated_coefficients),
                 fmt.Sprint(expected_coefficients))
    }
}

func TestAnscombe2LinearRegression(t *testing.T) {
    anscombe_2 := []stats.Coordinate{
        {10, 9.14},
        {8, 8.14},
        {13, 8.74},
        {9, 8.77},
        {11, 9.26},
        {14, 8.1},
        {6, 6.13},
        {4, 3.1},
        {12, 9.13},
        {7, 7.26},
        {5, 4.74},
    }

    expected_coefficients := stats.Series{
        {10, 8.00090909090909},
        {8, 7.000909090909092},
        {13, 9.500909090909088},
        {9, 7.500909090909091},
        {11, 8.50090909090909},
        {14, 10.000909090909087},
        {6, 6.000909090909094},
        {4, 5.000909090909095},
        {12, 9.00090909090909},
        {7, 6.500909090909093},
        {5, 5.500909090909094},
    }

    calculated_coefficients, _ := stats.LinearRegression(anscombe_2)

    if !slices.Equal(calculated_coefficients, expected_coefficients) {
        t.Errorf("Linear regression coefficients for Anscombe 2 do not match expectations\n%s\n%s",
                 fmt.Sprint(calculated_coefficients),
                 fmt.Sprint(expected_coefficients))
    }
}

func TestAnscombe3LinearRegression(t *testing.T) {
    anscombe_3 := []stats.Coordinate{
        {10, 7.46},
        {8, 6.77},
        {13, 12.74},
        {9, 7.11},
        {11, 7.81},
        {14, 8.84},
        {6, 6.08},
        {4, 5.39},
        {12, 8.15},
        {7, 6.42},
        {5, 5.73},
    }

    expected_coefficients :=stats.Series{
        {10, 7.999727272727271},
        {8, 7.000272727272732},
        {13, 9.498909090909081},
        {9, 7.500000000000002},
        {11, 8.499454545454540},
        {14, 9.998636363636351},
        {6, 6.000818181818192},
        {4, 5.001363636363653},
        {12, 8.999181818181810},
        {7, 6.5005454545454615},
        {5, 5.501090909090922},
    }

    calculated_coefficients, _ := stats.LinearRegression(anscombe_3)

    if !slices.Equal(calculated_coefficients, expected_coefficients) {
        t.Errorf("Linear regression coefficients for Anscombe 3 do not match expectations\n%s\n%s",
                 fmt.Sprint(calculated_coefficients),
                 fmt.Sprint(expected_coefficients))
    }
}

func TestAnscombe4LinearRegression(t *testing.T) {
    anscombe_4 := []stats.Coordinate{
        {8, 6.58},
        {8, 5.76},
        {8, 7.71},
        {8, 8.84},
        {8, 8.47},
        {8, 7.04},
        {8, 5.25},
        {19, 12.5},
        {8, 5.56},
        {8, 7.91},
        {8, 6.89},
    }

    expected_coefficients :=stats.Series{
        {8, 7.000999999999998},
        {8, 7.000999999999998},
        {8, 7.000999999999998},
        {8, 7.000999999999998},
        {8, 7.000999999999998},
        {8, 7.000999999999998},
        {8, 7.000999999999998},
        {19, 12.500000000000018},
        {8, 7.000999999999998},
        {8, 7.000999999999998},
        {8, 7.000999999999998},
    }

    calculated_coefficients, _ := stats.LinearRegression(anscombe_4)

    if !slices.Equal(calculated_coefficients, expected_coefficients) {
        t.Errorf("Linear regression coefficients for Anscombe 4 do not match expectations\n%s\n%s",
                 fmt.Sprint(calculated_coefficients),
                 fmt.Sprint(expected_coefficients))
    }
}

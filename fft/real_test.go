// Copyright 2018 The ZikiChombo Authors. All rights reserved.  Use of this source
// code is governed by a license that can be found in the License file.

package fft

import (
	"math"
	"math/rand"
	"testing"
)

func TestRealEquivCmplx(t *testing.T) {
	testRealEquivCmplx(true, t)
	testRealEquivCmplx(false, t)
}

func testRealEquivCmplx(scale bool, t *testing.T) {
	N := 64
	df := make([]float64, N)
	rft := NewReal(N)
	ft := New(N)
	rft.Scale(scale)
	ft.Scale(scale)
	dc := ft.Win(nil)

	for i := range df {
		v := rand.Float64()
		df[i] = v
		dc[i] = complex(v, 0.)
	}

	hc := rft.Do(df)
	ft.Do(dc)
	for i := 0; i < hc.Len(); i++ {
		cmplxCmpErr(hc.Cmplx(i), dc[i], 1e-10, t)
	}
}

func TestRealInv(t *testing.T) {
	iters := 1
	N := 32
	d := make([]float64, N)
	tmp := make([]float64, N)
	for i := 0; i < iters; i++ {
		rft := NewReal(N)
		for j := 0; j < N; j++ {
			d[j] = rand.Float64()
			tmp[j] = d[j]
		}
		sp := rft.Do(d)
		dd := rft.Inv(sp)
		for j, v := range dd {
			if math.Abs(v-tmp[j]) > 1e-10 {
				t.Errorf("iter %d idx %d got %f not %f\n", i, j, v, tmp[j])
			}
		}
	}
}

func BenchmarkReal(b *testing.B) {
	b.StopTimer()
	N := 1024
	w := make([]float64, N)
	tr := NewReal(N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Do(w)
	}
}

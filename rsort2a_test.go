package main

import (
	"bytes"
	"fmt"
	"runtime"
	"slices"
	"sort"
	"testing"
)

func BenchmarkStrings(b *testing.B) {
	for _, n := range []int{1 << 4, 1 << 8, 1 << 12, 1 << 16, 1 << 20, 1 << 24} {
		b.Run(fmt.Sprint("n=", n), func(b *testing.B) {
			for _, s := range []int{8, 16, 32, 64, 128, 256, 512} {
				var rand, lines []string
				b.Run(fmt.Sprint("len=", s), func(b *testing.B) {
					if rand == nil {
						rand = randomstrings(n, s, true)
						lines = make([]string, n)
					}
					b.Run("alg=radix", func(b *testing.B) {
						b.ReportAllocs()
						for i := 0; i < b.N; i++ {
							copy(lines, rand)
							out := rsort2b(lines, 0)
							runtime.KeepAlive(&out[0])
						}
					})
					b.Run("alg=sort.Strings", func(b *testing.B) {
						b.ReportAllocs()
						for i := 0; i < b.N; i++ {
							copy(lines, rand)
							sort.Strings(lines)
						}
					})
					b.Run("alg=slices.Sort", func(b *testing.B) {
						b.ReportAllocs()
						for i := 0; i < b.N; i++ {
							copy(lines, rand)
							slices.Sort(lines)
						}
					})
				})
			}
		})
	}
}

func BenchmarkBytes(b *testing.B) {
	for _, n := range []int{1 << 4, 1 << 8, 1 << 12, 1 << 16, 1 << 20, 1 << 24} {
		b.Run(fmt.Sprint("n=", n), func(b *testing.B) {
			for _, s := range []int{8, 16, 32, 64, 128, 256, 512} {
				var rand, lines [][]byte
				var randstr []string
				b.Run(fmt.Sprint("len=", s), func(b *testing.B) {
					if rand == nil {
						randstr = randomstrings(n, s, true)
						rand = make([][]byte, n)
						for i, s := range randstr {
							rand[i] = []byte(s)
						}
						lines = make([][]byte, n)
						b.ResetTimer()
					}
					b.Run("alg=radix", func(b *testing.B) {
						b.ReportAllocs()
						for i := 0; i < b.N; i++ {
							copy(lines, rand)
							out := rsort2a(rand, 0)
							runtime.KeepAlive(&out[0])
						}
					})
					b.Run("alg=slices.Sort", func(b *testing.B) {
						b.ReportAllocs()
						for i := 0; i < b.N; i++ {
							copy(lines, rand)
							slices.SortFunc(lines, bytes.Compare)
						}
					})
				})
			}
		})
	}
}

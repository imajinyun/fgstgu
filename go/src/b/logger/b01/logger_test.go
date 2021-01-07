package logger

import (
	"io/ioutil"
	"testing"
)

func init() {
	Writer = ioutil.Discard
}

func BenchmarkError(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Error("testing error")
	}
}

func BenchmarkWarn(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Warn("testing warning")
	}
}

func BenchmarkInfo(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Info("testing info")
	}
}

func BenchmarkDebug(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Debug("testing debug")
	}
}

func BenchmarkLogger(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Logger(Writer, LogLevel, 0, "testing")
	}
}

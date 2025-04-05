package main

import "testing"

func BenchmarkWithoutReservation(b *testing.B) {
	sourceData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		targetData := make([]int, 0)
		for _, v := range sourceData {
			targetData = append(targetData, v)
		}
	}

}

//// На данном примере в 4+ раза быстрее
func BenchmarkWithReservation(b *testing.B) {
	sourceData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		targetData := make([]int, 0, len(sourceData))
		for _, v := range sourceData {
			targetData = append(targetData, v)
		}
	}
}

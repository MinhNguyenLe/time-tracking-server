package utils

import (
	"time"
)

func TimeParsed(dateString string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateString)
}

type AverageScore struct {
	Satisfaction float64
	Productivity float64
	Interested   float64
	Insight      float64
}

func CalculateAverageScore(scores []AverageScore) AverageScore {
	if len(scores) == 0 {
		return AverageScore{0, 0, 0, 0}
	}

	var sumS, sumP, sumInt, sumIns float64

	for _, item := range scores {
		sumS += item.Satisfaction
		sumP += item.Productivity
		sumInt += item.Interested
		sumIns += item.Insight
	}

	avgS := sumS / float64(len(scores))
	avgP := sumP / float64(len(scores))
	avgInt := sumInt / float64(len(scores))
	avgIns := sumIns / float64(len(scores))

	return AverageScore{Satisfaction: avgS, Productivity: avgP, Interested: avgInt, Insight: avgIns}
}

package pkg_test

import (
	"geektrust/pkg"
	"testing"
)

type TestCase struct {
	drivers        []*pkg.Driver
	riders         []*pkg.Rider
	matchTestCases []MatchTestCase
}

type MatchTestCase struct {
	riderId                     string
	radiusInKm                  float64
	maxDrivers                  int
	expectedIdsOfMatchedDrivers []string
}

type TestCases []TestCase

func TestMatchRiderWithDriver(t *testing.T) {
	t.Run("when all drivers are available for matching with riders", func(t *testing.T) {
		testCases := TestCases{
			{
				drivers: []*pkg.Driver{
					driver("D4", 2, 2),
					driver("D1", 1, 1),
					driver("D2", 4, 5),
					driver("D3", 2, 2),
				},
				riders: []*pkg.Rider{
					rider("R1", 0, 0),
				},
				matchTestCases: []MatchTestCase{
					{
						riderId:                     "R1",
						radiusInKm:                  5,
						maxDrivers:                  5,
						expectedIdsOfMatchedDrivers: []string{"D1", "D3", "D4"},
					},
				},
			},
			{
				drivers: []*pkg.Driver{
					driver("D1", 0, 1),
					driver("D4", 2, 3),
					driver("D2", 2, 3),
					driver("D3", 4, 2),
				},
				riders: []*pkg.Rider{
					rider("R1", 3, 5),
					rider("R2", 1, 1),
				},
				matchTestCases: []MatchTestCase{
					{
						riderId:                     "R1",
						radiusInKm:                  5,
						maxDrivers:                  5,
						expectedIdsOfMatchedDrivers: []string{"D2", "D4", "D3", "D1"},
					},
					{
						riderId:                     "R2",
						radiusInKm:                  5,
						maxDrivers:                  5,
						expectedIdsOfMatchedDrivers: []string{"D1", "D2", "D4", "D3"},
					},
				},
			},
		}

		for _, testCase := range testCases {
			rideSharingApp := pkg.NewRideSharingApp()
			addDrivers(t, rideSharingApp, testCase.drivers)
			addRiders(t, rideSharingApp, testCase.riders)

			for _, matchTestCase := range testCase.matchTestCases {
				input := &pkg.MatchRiderWithDriverInput{
					RiderId:    matchTestCase.riderId,
					RadiusInKm: matchTestCase.radiusInKm,
					MaxDrivers: matchTestCase.maxDrivers,
				}

				idsOfMatchedDrivers, err := rideSharingApp.MatchRiderWithDriver(input)
				if err != nil {
					t.Errorf("expected no error occur while matching rider with driver but got error: %v", err)
				}

				assertStringArrayEqual(t, idsOfMatchedDrivers, matchTestCase.expectedIdsOfMatchedDrivers)
			}
		}
	})
}

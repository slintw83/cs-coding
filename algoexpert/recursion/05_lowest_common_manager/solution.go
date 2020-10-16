package main

type OrgChart struct {
	Name          string
	DirectReports []*OrgChart
}

func GetLowestCommonManager(org, reportOne, reportTwo *OrgChart) *OrgChart {
	_, manager := recurse(org, reportOne, reportTwo)
	return manager
}

func recurse(manager, reportOne, reportTwo *OrgChart) (int, *OrgChart) {
	numImportantReport := 0
	for _, report := range manager.DirectReports {
		num, r := recurse(report, reportOne, reportTwo)
		if r != nil {
			return num, r
		}
		numImportantReport += num
	}
	if manager == reportOne || manager == reportTwo {
		numImportantReport += 1
	}

	if numImportantReport == 2 {
		return numImportantReport, manager
	} else {
		return numImportantReport, nil
	}
}

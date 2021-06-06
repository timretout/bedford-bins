package bins

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGetURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `
{
	"BinCollections": [
		[
		{
			"JobId": 35309514,
			"Uprn": 100080022975,
			"PhysicalUprn": 100080022975,
			"Usrn": 27301473,
			"BinType": "Black bin",
			"JobName": "Empty Bin 240L Black",
			"JobScheduledStart": "2020-11-18T00:00:00",
			"JobStatus": "Collected",
			"CustomerAllowedToReportMissedBin": false,
			"MissedBinUrl": "https://bedford-self.achieveservice.com/service/Missed_Bin_Collection",
			"CustomerNotAllowedToReportMissedBinMessage": "Sorry you cannot report a missed bin more than two working days after it was scheduled to be picked up.",
			"LastDateMissedBinCanBeReported": "2020-11-20T00:00:00",
			"PremisesEvents": []
		}
		],
		[
		{
			"JobId": 37579274,
			"Uprn": 100080022975,
			"PhysicalUprn": 100080022975,
			"Usrn": 27301473,
			"BinType": "Green bin",
			"JobName": "Empty Bin 240L Green",
			"JobScheduledStart": "2020-11-25T00:00:00",
			"JobStatus": "Collected",
			"CustomerAllowedToReportMissedBin": false,
			"MissedBinUrl": "https://bedford-self.achieveservice.com/service/Missed_Bin_Collection",
			"CustomerNotAllowedToReportMissedBinMessage": "Sorry you cannot report a missed bin more than two working days after it was scheduled to be picked up.",
			"LastDateMissedBinCanBeReported": "2020-11-27T00:00:00",
			"PremisesEvents": []
		},
		{
			"JobId": 35428043,
			"Uprn": 100080022975,
			"PhysicalUprn": 100080022975,
			"Usrn": 27301473,
			"BinType": "Orange bin",
			"JobName": "Empty Bin 240L Orange",
			"JobScheduledStart": "2020-11-25T00:00:00",
			"JobStatus": "Collected",
			"CustomerAllowedToReportMissedBin": false,
			"MissedBinUrl": "https://bedford-self.achieveservice.com/service/Missed_Bin_Collection",
			"CustomerNotAllowedToReportMissedBinMessage": "Sorry you cannot report a missed bin more than two working days after it was scheduled to be picked up.",
			"LastDateMissedBinCanBeReported": "2020-11-27T00:00:00",
			"PremisesEvents": []
		}
		],
		[
		{
			"JobId": 36109105,
			"Uprn": 100080022975,
			"PhysicalUprn": 100080022975,
			"Usrn": 27301473,
			"BinType": "Black bin",
			"JobName": "Empty Bin 240L Black",
			"JobScheduledStart": "2020-12-02T00:00:00",
			"JobStatus": "Collected",
			"CustomerAllowedToReportMissedBin": false,
			"MissedBinUrl": "https://bedford-self.achieveservice.com/service/Missed_Bin_Collection",
			"CustomerNotAllowedToReportMissedBinMessage": "Sorry you cannot report a missed bin more than two working days after it was scheduled to be picked up.",
			"LastDateMissedBinCanBeReported": "2020-12-04T00:00:00",
			"PremisesEvents": []
		}
		],
		[
		{
			"JobId": 36261781,
			"Uprn": 100080022975,
			"PhysicalUprn": 100080022975,
			"Usrn": 27301473,
			"BinType": "Orange bin",
			"JobName": "Empty Bin 240L Orange",
			"JobScheduledStart": "2020-12-09T00:00:00",
			"JobStatus": "Collected",
			"CustomerAllowedToReportMissedBin": false,
			"MissedBinUrl": "https://bedford-self.achieveservice.com/service/Missed_Bin_Collection",
			"CustomerNotAllowedToReportMissedBinMessage": "Sorry you cannot report a missed bin more than two working days after it was scheduled to be picked up.",
			"LastDateMissedBinCanBeReported": "2020-12-11T00:00:00",
			"PremisesEvents": []
		}
		],
		[
		{
			"JobId": 37833750,
			"Uprn": 100080022975,
			"PhysicalUprn": 100080022975,
			"Usrn": 27301473,
			"BinType": "Black bin",
			"JobName": "Empty Bin 240L Black",
			"JobScheduledStart": "2020-12-16T00:00:00",
			"JobStatus": "Collected",
			"CustomerAllowedToReportMissedBin": false,
			"MissedBinUrl": "https://bedford-self.achieveservice.com/service/Missed_Bin_Collection",
			"CustomerNotAllowedToReportMissedBinMessage": "Sorry you cannot report a missed bin more than two working days after it was scheduled to be picked up.",
			"LastDateMissedBinCanBeReported": "2020-12-18T00:00:00",
			"PremisesEvents": []
		}
		],
		[
		{
			"JobId": 37897672,
			"Uprn": 100080022975,
			"PhysicalUprn": 100080022975,
			"Usrn": 27301473,
			"BinType": "Orange bin",
			"JobName": "Empty Bin 240L Orange",
			"JobScheduledStart": "2020-12-23T00:00:00",
			"JobStatus": "Collected",
			"CustomerAllowedToReportMissedBin": true,
			"MissedBinUrl": "https://bedford-self.achieveservice.com/service/Missed_Bin_Collection?uprn=100080022975&bintype=Orange+bin&collectiondate=2020-12-23&jobstatus=Collected&nocollectionfound=false",
			"CustomerNotAllowedToReportMissedBinMessage": null,
			"LastDateMissedBinCanBeReported": "2020-12-29T00:00:00",
			"PremisesEvents": []
		}
		],
		[
		{
			"JobId": 37975967,
			"Uprn": 100080022975,
			"PhysicalUprn": 100080022975,
			"Usrn": 27301473,
			"BinType": "Black bin",
			"JobName": "Empty Bin 240L Black",
			"JobScheduledStart": "2021-01-02T00:00:00",
			"JobStatus": "Scheduled",
			"CustomerAllowedToReportMissedBin": false,
			"MissedBinUrl": "https://bedford-self.achieveservice.com/service/Missed_Bin_Collection",
			"CustomerNotAllowedToReportMissedBinMessage": null,
			"LastDateMissedBinCanBeReported": "2021-01-05T00:00:00",
			"PremisesEvents": []
		}
		],
		[
		{
			"JobId": 38120573,
			"Uprn": 100080022975,
			"PhysicalUprn": 100080022975,
			"Usrn": 27301473,
			"BinType": "Orange bin",
			"JobName": "Empty Bin 240L Orange",
			"JobScheduledStart": "2021-01-08T00:00:00",
			"JobStatus": "Scheduled",
			"CustomerAllowedToReportMissedBin": false,
			"MissedBinUrl": "https://bedford-self.achieveservice.com/service/Missed_Bin_Collection",
			"CustomerNotAllowedToReportMissedBinMessage": null,
			"LastDateMissedBinCanBeReported": "2021-01-12T00:00:00",
			"PremisesEvents": []
		}
		],
		[
		{
			"JobId": 38030444,
			"Uprn": 100080022975,
			"PhysicalUprn": 100080022975,
			"Usrn": 27301473,
			"BinType": "Black bin",
			"JobName": "Empty Bin 240L Black",
			"JobScheduledStart": "2021-01-14T00:00:00",
			"JobStatus": "Scheduled",
			"CustomerAllowedToReportMissedBin": false,
			"MissedBinUrl": "https://bedford-self.achieveservice.com/service/Missed_Bin_Collection",
			"CustomerNotAllowedToReportMissedBinMessage": null,
			"LastDateMissedBinCanBeReported": "2021-01-18T00:00:00",
			"PremisesEvents": []
		}
		],
		[
		{
			"JobId": 37903992,
			"Uprn": 100080022975,
			"PhysicalUprn": 100080022975,
			"Usrn": 27301473,
			"BinType": "Orange bin",
			"JobName": "Empty Bin 240L Orange",
			"JobScheduledStart": "2021-01-20T00:00:00",
			"JobStatus": "Scheduled",
			"CustomerAllowedToReportMissedBin": false,
			"MissedBinUrl": "https://bedford-self.achieveservice.com/service/Missed_Bin_Collection",
			"CustomerNotAllowedToReportMissedBinMessage": null,
			"LastDateMissedBinCanBeReported": "2021-01-22T00:00:00",
			"PremisesEvents": []
		}
		],
		[
		{
			"JobId": 36939375,
			"Uprn": 100080022975,
			"PhysicalUprn": 100080022975,
			"Usrn": 27301473,
			"BinType": "Black bin",
			"JobName": "Empty Bin 240L Black",
			"JobScheduledStart": "2021-01-27T00:00:00",
			"JobStatus": "Scheduled",
			"CustomerAllowedToReportMissedBin": false,
			"MissedBinUrl": "https://bedford-self.achieveservice.com/service/Missed_Bin_Collection",
			"CustomerNotAllowedToReportMissedBinMessage": null,
			"LastDateMissedBinCanBeReported": "2021-01-29T00:00:00",
			"PremisesEvents": []
		}
		]
	],
	"BinCollectionDaysByBinType": [
		{
		"ScheduledDay": "Wed",
		"BinType": "Black bin",
		"DayOrder": 3
		},
		{
		"ScheduledDay": "Wed",
		"BinType": "Green bin",
		"DayOrder": 3
		},
		{
		"ScheduledDay": "Wed",
		"BinType": "Orange bin",
		"DayOrder": 3
		}
	],
	"BinCollectionDays": [
		"Wed"
	],
	"DefaultMissedBinUrl": "https://bedford-self.achieveservice.com/service/Missed_Bin_Collection?uprn=100080022975&nocollectionfound=true",
	"HasCalendar": true,
	"CalendarUrl": "https://bbcdevwebfiles.blob.core.windows.net/webfiles/Bin%20Collection%20Calendars/collection-calendar-A.pdf"
}`)
	}))
	defer ts.Close()

	actual, err := getURL(context.Background(), ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	expected := []Collection{
		{[]BinType{Black}, time.Date(2020, time.November, 18, 0, 0, 0, 0, time.UTC)},
		{[]BinType{Green, Orange}, time.Date(2020, time.November, 25, 0, 0, 0, 0, time.UTC)},
		{[]BinType{Black}, time.Date(2020, time.December, 2, 0, 0, 0, 0, time.UTC)},
		{[]BinType{Orange}, time.Date(2020, time.December, 9, 0, 0, 0, 0, time.UTC)},
		{[]BinType{Black}, time.Date(2020, time.December, 16, 0, 0, 0, 0, time.UTC)},
		{[]BinType{Orange}, time.Date(2020, time.December, 23, 0, 0, 0, 0, time.UTC)},
		{[]BinType{Black}, time.Date(2021, time.January, 2, 0, 0, 0, 0, time.UTC)},
		{[]BinType{Orange}, time.Date(2021, time.January, 8, 0, 0, 0, 0, time.UTC)},
		{[]BinType{Black}, time.Date(2021, time.January, 14, 0, 0, 0, 0, time.UTC)},
		{[]BinType{Orange}, time.Date(2021, time.January, 20, 0, 0, 0, 0, time.UTC)},
		{[]BinType{Black}, time.Date(2021, time.January, 27, 0, 0, 0, 0, time.UTC)},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Got %v, expected %v", actual, expected)
	}
}

func TestBinTypes(t *testing.T) {
	type test struct {
		col      binCollection
		expected []BinType
	}

	tests := []test{
		{binCollection{{"Black bin", customTime{time.Now()}}}, []BinType{Black}},
		{binCollection{{"Orange bin", customTime{time.Now()}}}, []BinType{Orange}},
		{binCollection{{"Green bin", customTime{time.Now()}}}, []BinType{Green}},
		{binCollection{{"Orange bin", customTime{time.Now()}}, {"Green bin", customTime{time.Now()}}}, []BinType{Orange, Green}},
	}

	for _, tc := range tests {
		actual := tc.col.BinTypes()
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Got %s, expected %s", actual, tc.expected)
		}
	}
}

package controllers_test

import (
	"encoding/json"
	"fmt"
	"math"
	"medicare-api/types"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetMedicalData(t *testing.T) {
	Convey("When GET /medicare/data is called", t, withCleanup(func() {

		zip1 := insertZipCode(types.ZipCodeLatLong{
			ZipCode:   90001,
			Latitude:  10,
			Longitude: 11,
		})

		zip2 := insertZipCode(types.ZipCodeLatLong{
			ZipCode:   90002,
			Latitude:  12,
			Longitude: 13,
		})

		zip3 := insertZipCode(types.ZipCodeLatLong{
			ZipCode:   90003,
			Latitude:  14,
			Longitude: 15,
		})

		provider1 := insertProvider(types.Provider{
			ID:             1001,
			Name:           "provider A",
			Street:         "provider A street",
			City:           "provider A city",
			State:          "P1",
			ZipCode:        zip1.ZipCode,
			HRRDescription: "P1 - provider A",
		})

		provider2 := insertProvider(types.Provider{
			ID:             1002,
			Name:           "provider B",
			Street:         "provider B street",
			City:           "provider B city",
			State:          "P2",
			ZipCode:        zip2.ZipCode,
			HRRDescription: "P2 - provider B",
		})

		provider3 := insertProvider(types.Provider{
			ID:             1003,
			Name:           "provider C",
			Street:         "provider C street",
			City:           "provider C city",
			State:          "P3",
			ZipCode:        zip3.ZipCode,
			HRRDescription: "P3 - provider C",
		})

		procedure1 := insertProcedure(types.Procedure{
			ID:                      "001",
			AverageCoveredCharges:   1111,
			AverageTotalPayments:    1111,
			AverageMedicarePayments: 1111,
			DRGDefinition:           "001 - PROCEDURE ONE DEFINITION",
		})

		procedure2 := insertProcedure(types.Procedure{
			ID:                      "002",
			AverageCoveredCharges:   2222,
			AverageTotalPayments:    2222,
			AverageMedicarePayments: 2222,
			DRGDefinition:           "002 - PROCEDURE TWO DEFINITION",
		})

		procedure3 := insertProcedure(types.Procedure{
			ID:                      "003",
			AverageCoveredCharges:   3333,
			AverageTotalPayments:    3333,
			AverageMedicarePayments: 3333,
			DRGDefinition:           "003 - PROCEDURE THREE DEFINITION",
		})

		// Provider 1
		assignProcedureToProvider(provider1.ID, procedure1.ID)
		assignProcedureToProvider(provider1.ID, procedure2.ID)
		assignProcedureToProvider(provider1.ID, procedure3.ID)

		// Produce medical documents for the provider
		provider1procedure1 := produceMedicalData(zip1, provider1, procedure1)
		provider1procedure2 := produceMedicalData(zip1, provider1, procedure2)
		provider1procedure3 := produceMedicalData(zip1, provider1, procedure3)

		// Provider 2
		assignProcedureToProvider(provider2.ID, procedure2.ID)
		assignProcedureToProvider(provider2.ID, procedure3.ID)

		// Produce medical documents for the provider
		provider2procedure2 := produceMedicalData(zip2, provider2, procedure2)
		provider2procedure3 := produceMedicalData(zip2, provider2, procedure3)

		// Provider 3
		assignProcedureToProvider(provider3.ID, procedure1.ID)
		assignProcedureToProvider(provider3.ID, procedure3.ID)

		// Produce medical documents for the provider
		provider3procedure1 := produceMedicalData(zip3, provider3, procedure1)
		provider3procedure3 := produceMedicalData(zip3, provider3, procedure3)

		var errorResponse types.ErrorResponse
		var validResponse types.MedicareDataResponse
		rr := httptest.NewRecorder()

		Convey("When trying to query by non-existing query param", func() {

			req := prepareGetRequest("/medicare/api/data?foo=bar")
			Convey("Bad request should be returned", func() {
				r.ServeHTTP(rr, req)

				unmarshalInto(rr, &errorResponse)
				So(rr.Code, ShouldEqual, http.StatusBadRequest)
				So(errorResponse, ShouldResemble, types.ErrorResponse{
					Message: "Validation error",
					Details: map[string]string{
						"foo": "Is not allowed as an additional property",
					},
				})
			})
		})

		Convey("When trying to query by existing params", func() {
			Convey("When querying by min_price", func() {
				Convey("When min_price is not an integer", func() {
					req := prepareGetRequest("/medicare/api/data?min_price=foo")
					Convey("Validation error should be returned", func() {
						r.ServeHTTP(rr, req)

						unmarshalInto(rr, &errorResponse)
						So(rr.Code, ShouldEqual, http.StatusBadRequest)
						So(errorResponse, ShouldResemble, types.ErrorResponse{
							Message: "Validation error",
							Details: map[string]string{
								"min_price": "Invalid type. Expected: integer, given: string",
							},
						})
					})
				})

				Convey("When min_price is an integer", func() {
					url := fmt.Sprintf("/medicare/api/data?min_price=%d", int(procedure2.AverageTotalPayments))
					req := prepareGetRequest(url)

					Convey("Only the results that match should be returned", func() {
						r.ServeHTTP(rr, req)

						unmarshalInto(rr, &validResponse)
						So(rr.Code, ShouldEqual, http.StatusOK)
						So(validResponse, ShouldResemble, types.MedicareDataResponse{
							Objects: []types.MedicalData{
								// Anything above or equal to the price of procedure 2 will be returned
								// Results are sorted by averaga_total_payments ASC
								provider1procedure2,
								provider2procedure2,
								provider1procedure3,
								provider2procedure3,
								provider3procedure3,
							},
							Total:      5,
							PerPage:    20,
							PageNumber: 1,
						})
					})
				})
			})

			Convey("When querying by max_price", func() {
				Convey("When max_price is not an integer", func() {
					req := prepareGetRequest("/medicare/api/data?max_price=foo")
					Convey("Validation error should be returned", func() {
						r.ServeHTTP(rr, req)

						unmarshalInto(rr, &errorResponse)
						So(rr.Code, ShouldEqual, http.StatusBadRequest)
						So(errorResponse, ShouldResemble, types.ErrorResponse{
							Message: "Validation error",
							Details: map[string]string{
								"max_price": "Invalid type. Expected: integer, given: string",
							},
						})
					})
				})

				Convey("When max_price is an integer", func() {
					url := fmt.Sprintf("/medicare/api/data?max_price=%d", int(procedure2.AverageTotalPayments))
					req := prepareGetRequest(url)

					Convey("Only the results that match should be returned", func() {
						r.ServeHTTP(rr, req)

						unmarshalInto(rr, &validResponse)
						So(rr.Code, ShouldEqual, http.StatusOK)
						So(validResponse, ShouldResemble, types.MedicareDataResponse{
							Objects: []types.MedicalData{
								// Anything above or equal to the price of procedure 2 will be returned
								// Results are sorted by averaga_total_payments ASC
								provider1procedure1,
								provider3procedure1,
								provider1procedure2,
								provider2procedure2,
							},
							Total:      4,
							PerPage:    20,
							PageNumber: 1,
						})
					})
				})
			})
		})

		Convey("When querying by proximity", func() {
			Convey("When proximity is not an integer", func() {
				req := prepareGetRequest("/medicare/api/data?proximity=foo")
				Convey("Validation error should be returned", func() {
					r.ServeHTTP(rr, req)

					unmarshalInto(rr, &errorResponse)
					So(rr.Code, ShouldEqual, http.StatusBadRequest)
					So(errorResponse, ShouldResemble, types.ErrorResponse{
						Message: "Validation error",
						Details: map[string]string{
							"proximity": "Invalid type. Expected: integer, given: string",
						},
					})
				})
			})

			Convey("When longitude and latitude are not set", func() {
				req := prepareGetRequest("/medicare/api/data?proximity=200")
				Convey("It should return all entries and ignore proximity", func() {
					r.ServeHTTP(rr, req)

					unmarshalInto(rr, &validResponse)
					So(rr.Code, ShouldEqual, http.StatusOK)
					So(validResponse, ShouldResemble, types.MedicareDataResponse{
						Objects: []types.MedicalData{
							// Anything above or equal to the price of procedure 2 will be returned
							// Results are sorted by averaga_total_payments ASC
							provider1procedure1,
							provider3procedure1,
							provider1procedure2,
							provider2procedure2,
							provider1procedure3,
							provider2procedure3,
							provider3procedure3,
						},
						Total:      7,
						PerPage:    20,
						PageNumber: 1,
					})
				})

			})

			Convey("When longitude and latitude are set", func() {
				Convey("It should return all entries that lay in the proximity", func() {
					url := fmt.Sprintf("/medicare/api/data?lat=%f&long=%f&proximity=10",
						zip1.Latitude, zip1.Longitude)
					req := prepareGetRequest(url)
					r.ServeHTTP(rr, req)

					unmarshalInto(rr, &validResponse)
					So(rr.Code, ShouldEqual, http.StatusOK)
					So(validResponse, ShouldResemble, types.MedicareDataResponse{
						Objects: []types.MedicalData{
							// Anything above or equal to the price of procedure 2 will be returned
							// Results are sorted by averaga_total_payments ASC
							provider1procedure1,
							provider1procedure2,
							provider1procedure3,
						},
						Total:      3,
						PerPage:    20,
						PageNumber: 1,
					})

					url = fmt.Sprintf("/medicare/api/data?lat=%f&long=%f&proximity=300",
						zip2.Latitude, zip2.Longitude)
					req = prepareGetRequest(url)

					// Reset request recorder
					rr = httptest.NewRecorder()
					r.ServeHTTP(rr, req)

					unmarshalInto(rr, &validResponse)
					So(rr.Code, ShouldEqual, http.StatusOK)
					So(validResponse, ShouldResemble, types.MedicareDataResponse{
						Objects: []types.MedicalData{
							// Anything above or equal to the price of procedure 2 will be returned
							// Results are sorted by averaga_total_payments ASC
							provider1procedure1,
							provider1procedure2,
						},
						Total:      3,
						PerPage:    20,
						PageNumber: 1,
					})
				})
			})
		})
	}))
}

// Helper to reduce code duplication
func unmarshalInto(rr *httptest.ResponseRecorder, dest interface{}) {
	err := json.Unmarshal(rr.Body.Bytes(), &dest)
	if err != nil {
		panic(err)
	}
}

// Helper to reduce code duplication
func prepareGetRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	return req
}

// Combines zip, provider and procedure data into single record
// Required for testing responses
func produceMedicalData(providerZip *types.ZipCodeLatLong,
	provider *types.Provider, procedure *types.Procedure) types.MedicalData {

	record := types.MedicalData{
		ProviderID:              provider.ID,
		ProviderName:            provider.Name,
		ProviderStreet:          provider.Street,
		ProviderCity:            provider.City,
		ProviderState:           provider.State,
		ProviderZipCode:         providerZip.ZipCode,
		HRRDescription:          provider.HRRDescription,
		TotalDischarges:         procedure.TotalDischarges,
		AverageCoveredCharges:   procedure.AverageCoveredCharges,
		AverageTotalPayments:    procedure.AverageTotalPayments,
		AverageMedicarePayments: procedure.AverageMedicarePayments,
		DRGDefinition:           procedure.DRGDefinition,
		Latitude:                providerZip.Latitude,
		Longitude:               providerZip.Longitude,
	}

	return record
}

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	return dist
}

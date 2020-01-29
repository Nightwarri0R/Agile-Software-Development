package controllers_test

import (
	"medicare-api/types"
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
			ZipCode:   90001,
			Latitude:  12,
			Longitude: 13,
		})

		zip3 := insertZipCode(types.ZipCodeLatLong{
			ZipCode:   90001,
			Latitude:  14,
			Longitude: 15,
		})

		provider1 := insertProvider(types.Provider{
			ID:             1001,
			Name:           "provider one",
			Street:         "provider one street",
			City:           "provider one city",
			State:          "P1",
			ZipCode:        zip1.ZipCode,
			HRRDescription: "P1 - provider one",
		})

		provider2 := insertProvider(types.Provider{
			ID:             1002,
			Name:           "provider two",
			Street:         "provider two street",
			City:           "provider two city",
			State:          "P2",
			ZipCode:        zip2.ZipCode,
			HRRDescription: "P2 - provider two",
		})

		provider3 := insertProvider(types.Provider{
			ID:             1003,
			Name:           "provider three",
			Street:         "provider three street",
			City:           "provider three city",
			State:          "P3",
			ZipCode:        zip3.ZipCode,
			HRRDescription: "P3 - provider three",
		})

		procedure1 := insertProcedure(types.Procedure{
			ID:                      "001",
			TotalDischarges:         10,
			AverageCoveredCharges:   1111.1,
			AverageTotalPayments:    1111.1,
			AverageMedicarePayments: 1111.1,
			DRGDefinition:           "001 - PROCEDURE ONE DEFINITION",
		})

		procedure2 := insertProcedure(types.Procedure{
			ID:                      "002",
			TotalDischarges:         11,
			AverageCoveredCharges:   2222.2,
			AverageTotalPayments:    2222.2,
			AverageMedicarePayments: 2222.2,
			DRGDefinition:           "002 - PROCEDURE TWO DEFINITION",
		})

		procedure3 := insertProcedure(types.Procedure{
			ID:                      "003",
			TotalDischarges:         12,
			AverageCoveredCharges:   3333.3,
			AverageTotalPayments:    3333.3,
			AverageMedicarePayments: 3333.3,
			DRGDefinition:           "003 - PROCEDURE THREE DEFINITION",
		})

		// Provider 1
		assignProcedureToProvider(provider1.ID, procedure1.ID)
		assignProcedureToProvider(provider1.ID, procedure2.ID)
		assignProcedureToProvider(provider1.ID, procedure3.ID)

		// Provider 2
		assignProcedureToProvider(provider2.ID, procedure2.ID)
		assignProcedureToProvider(provider2.ID, procedure3.ID)

		// Provider 3
		assignProcedureToProvider(provider3.ID, procedure1.ID)
		assignProcedureToProvider(provider3.ID, procedure3.ID)

		Convey("")
	}))
}

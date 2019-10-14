package rpc

import (
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"

	erpc "github.com/Varunram/essentials/rpc"
	utils "github.com/Varunram/essentials/utils"
	consts "github.com/YaleOpenLab/opensolar/consts"
	core "github.com/YaleOpenLab/opensolar/core"
)

// setupEntityRPCs sets up all entity related endpoints
func setupEntityRPCs() {
	validateEntity()
	getStage0Contracts()
	getStage1Contracts()
	getStage2Contracts()
	addCollateral()
	createOpensolarProject()
	proposeOpensolarProject()
}

var EntityRpc = map[int][]string{
	1: []string{"/entity/validate", "GET"},                                      // GET
	2: []string{"/entity/stage0", "GET"},                                        // GET
	3: []string{"/entity/stage1", "GET"},                                        // GET
	4: []string{"/entity/stage2", "GET"},                                        // GET
	5: []string{"/entity/addcollateral", "POST", "amount", "collateral"},        // POST
	6: []string{"/entity/proposeproject/opensolar", "POST", "projIndex", "fee"}, // POST
	7: []string{"/entity/newproject/opensolar", "GET"},                          // GET
}

// EntityValidateHelper is a helper that helps validate an entity
func EntityValidateHelper(w http.ResponseWriter, r *http.Request) (core.Entity, error) {
	var prepEntity core.Entity
	if r.Method == "GET" {
		if r.URL.Query() == nil || r.URL.Query()["username"] == nil ||
			len(r.URL.Query()["token"][0]) != 32 {
			return prepEntity, errors.New("Invalid params passed")
		}

		prepEntity, err := core.ValidateEntity(r.URL.Query()["username"][0], r.URL.Query()["token"][0])
		if err != nil {
			return prepEntity, errors.Wrap(err, "Error while validating entity")
		}

		return prepEntity, nil
	} else if r.Method == "POST" {
		if r.FormValue("username") == "" || r.FormValue("password") == "" {
			return prepEntity, errors.New("Invalid params passed")
		}

		prepEntity, err := core.ValidateEntity(r.FormValue("username"), r.FormValue("token"))
		if err != nil {
			return prepEntity, errors.Wrap(err, "Error while validating entity")
		}
		return prepEntity, nil
	}
	return prepEntity, errors.New("invalid method type")
}

// validateEntity is an endpoint that vlaidates is a specific entity is registered on the platform
func validateEntity() {
	http.HandleFunc(EntityRpc[1][0], func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			return
		}
		prepEntity, err := EntityValidateHelper(w, r)
		if err != nil {
			log.Println("Error while validating entity", err)
			erpc.ResponseHandler(w, erpc.StatusUnauthorized)
			return
		}
		erpc.MarshalSend(w, prepEntity)
	})
}

// getStage0Contracts gets a list of all the pre origianted contracts on the platform
func getStage0Contracts() {
	http.HandleFunc(EntityRpc[2][0], func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			return
		}
		prepEntity, err := EntityValidateHelper(w, r)
		if err != nil {
			log.Println("Error while validating entity", err)
			erpc.ResponseHandler(w, erpc.StatusUnauthorized)
			return
		}

		x, err := core.RetrieveOriginatorProjects(core.Stage0.Number, prepEntity.U.Index)
		if err != nil {
			log.Println("Error while retrieving originator project", err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}
		erpc.MarshalSend(w, x)
	})
}

// getStage1Contracts gets a list of all the originated contracts on the platform
func getStage1Contracts() {
	http.HandleFunc(EntityRpc[3][0], func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			return
		}
		prepEntity, err := EntityValidateHelper(w, r)
		if err != nil {
			log.Println("Error while validating entity", err)
			erpc.ResponseHandler(w, erpc.StatusUnauthorized)
			return
		}

		x, err := core.RetrieveOriginatorProjects(core.Stage1.Number, prepEntity.U.Index)
		if err != nil {
			log.Println("Error while retrieving originator projects", err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}
		erpc.MarshalSend(w, x)
	})
}

// getStage2Contracts gets a list of all the proposed contracts on the platform
func getStage2Contracts() {
	http.HandleFunc(EntityRpc[4][0], func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			return
		}
		prepEntity, err := EntityValidateHelper(w, r)
		if err != nil {
			log.Println("Error while validating entity", err)
			erpc.ResponseHandler(w, erpc.StatusUnauthorized)
			return
		}

		x, err := core.RetrieveContractorProjects(core.Stage2.Number, prepEntity.U.Index)
		if err != nil {
			log.Println("Error while retrieving contractor projects", err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}
		erpc.MarshalSend(w, x)
	})
}

// addCollateral is a route that a contractor can use to add collateral
func addCollateral() {
	http.HandleFunc(EntityRpc[5][0], func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckPost(w, r)
		if err != nil {
			log.Println(err)
			return
		}
		prepEntity, err := EntityValidateHelper(w, r)
		if err != nil {
			erpc.ResponseHandler(w, erpc.StatusUnauthorized)
			return
		}

		err = checkReqdParams(w, r, EntityRpc[5][2:], EntityRpc[5][1])
		if err != nil {
			log.Println(err)
			return
		}

		amountx := r.FormValue("amount")
		collateral := r.FormValue("collateral")

		amount, err := utils.ToFloat(amountx)
		if err != nil {
			log.Println("Error while converting string to float", err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		err = prepEntity.AddCollateral(amount, collateral)
		if err != nil {
			log.Println("Error while adding collateral", err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.ResponseHandler(w, erpc.StatusOK)
	})
}

// proposeOpensolarProject creates a contract which the contractor proposes towards a particular project
func proposeOpensolarProject() {
	http.HandleFunc(EntityRpc[6][0], func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckPost(w, r)
		if err != nil {
			log.Println(err)
			return
		}

		prepEntity, err := EntityValidateHelper(w, r)
		if err != nil {
			log.Println("Error while validating entity", err)
			erpc.ResponseHandler(w, erpc.StatusUnauthorized)
			return
		}

		err = checkReqdParams(w, r, EntityRpc[6][2:], EntityRpc[6][1])
		if err != nil {
			log.Println(err)
			return
		}

		projIndexx := r.FormValue("projIndex")
		feex := r.FormValue("fee")

		projIndex, err := utils.ToInt(projIndexx)
		if err != nil {
			log.Println("project idnex not int, quitting!")
			return
		}

		x, err := core.RetrieveProject(projIndex)
		if err != nil {
			log.Println("couldn't retrieve project with index")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
		}

		fee, err := utils.ToFloat(feex)
		if err != nil {
			log.Println("fee passed not integer, quitting!")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
		}

		x.TotalValue += fee
		x.OriginatorFee = fee
		x.OriginatorIndex = prepEntity.U.Index
		x.Stage = 2

		err = x.Save()
		if err != nil {
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, x)
	})
}

// createOpensolarProject creates a contract which the originator can take to the recipient in order to be validated
// as a level 1 project.
func createOpensolarProject() {
	http.HandleFunc(EntityRpc[7][0], func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			return
		}

		prepEntity, err := EntityValidateHelper(w, r)
		if err != nil {
			log.Println("Error while validating entity", err)
			erpc.ResponseHandler(w, erpc.StatusUnauthorized)
			return
		}

		if r.URL.Query()["TotalValue"] == nil || r.URL.Query()["Years"] == nil || r.URL.Query()["InterestRate"] == nil ||
			r.URL.Query()["Location"] == nil || r.URL.Query()["PanelSize"] == nil || r.URL.Query()["Inverter"] == nil ||
			r.URL.Query()["ChargeRegulator"] == nil || r.URL.Query()["ControlPanel"] == nil || r.URL.Query()["CommBox"] == nil ||
			r.URL.Query()["ACTransfer"] == nil || r.URL.Query()["SolarCombiner"] == nil || r.URL.Query()["Batteries"] == nil ||
			r.URL.Query()["IoTHub"] == nil || r.URL.Query()["Metadata"] == nil || r.URL.Query()["OriginatorFee"] == nil ||
			r.URL.Query()["recpIndex"] == nil || r.URL.Query()["AuctionType"] == nil || r.URL.Query()["PaybackPeriod"] == nil {
			log.Println("Bad request, required params missing!")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		allProjects, err := core.RetrieveAllProjects()
		if err != nil {
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		var x core.Project
		x.TotalValue, err = utils.ToFloat(r.URL.Query()["TotalValue"][0])
		if err != nil {
			log.Println("param passed not float, quitting!")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}
		x.EstimatedAcquisition, err = utils.ToInt(r.URL.Query()["Years"][0])
		if err != nil {
			log.Println("param passed not int, quitting!")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}
		x.InterestRate, err = utils.ToFloat(r.URL.Query()["InterestRate"][0])
		if err != nil {
			log.Println("param passed not int, quitting!")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}
		x.OriginatorFee, err = utils.ToFloat(r.URL.Query()["OriginatorFee"][0])
		if err != nil {
			log.Println("ORiginator fee not float, quitting!")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		x.RecipientIndex, err = utils.ToInt(r.URL.Query()["recpIndex"][0])
		if err != nil {
			log.Println("passed recipient index not int, quitting!")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		_, err = core.RetrieveRecipient(x.RecipientIndex)
		if err != nil {
			log.Println("could not retrieve recipient, quitting!")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		x.AuctionType = r.URL.Query()["AuctionType"][0]
		paybackPeriodInt, err := utils.ToInt(r.URL.Query()["PaybackPeriod"][0])
		if err != nil {
			log.Println("payback period not integer, quitting!")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		x.PaybackPeriod = time.Duration(paybackPeriodInt) * consts.OneWeekInSecond
		x.OriginatorIndex = prepEntity.U.Index
		x.State = r.URL.Query()["Location"][0]
		x.PanelSize = r.URL.Query()["PanelSize"][0]
		x.Inverter = r.URL.Query()["Inverter"][0]
		x.ChargeRegulator = r.URL.Query()["ChargeRegulator"][0]
		x.ControlPanel = r.URL.Query()["ControlPanel"][0]
		x.CommBox = r.URL.Query()["CommBox"][0]
		x.ACTransfer = r.URL.Query()["ACTransfer"][0]
		x.SolarCombiner = r.URL.Query()["SolarCombiner"][0]
		x.Batteries = r.URL.Query()["Batteries"][0]
		x.IoTHub = r.URL.Query()["IoTHub"][0]
		x.Metadata = r.URL.Query()["Metadata"][0]

		x.Index = len(allProjects) + 1
		x.Stage = 0
		x.MoneyRaised = 0
		x.BalLeft = x.TotalValue
		x.Votes = 0
		x.Reputation = x.TotalValue
		x.InvestmentType = "munibond" // hardcode for now, expand if we have other investment models later down the road
		x.DateInitiated = utils.Timestamp()

		err = x.Save()
		if err != nil {
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, x)
	})
}

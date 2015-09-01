package spi

import (
	"encoding/base64"
	"fmt"
	"log"
)

//Common Variables -------------------------------------------------------------

const XPS_HTTPS = API_HTTPS + "/Experiments"
const REX_HTTPS = API_HTTPS + "/Realizations"

//API calls --------------------------------------------------------------------

// Create Experiment -----------------------------------------------------------

func CreateExperiment(expId, owner, topdl string) (
	*CreateExperimentResponse, error) {

	e := CreateExperimentEnvelope{}
	e.Body.CreateExperiment.EID = expId
	e.Body.CreateExperiment.Owner = owner
	e.Body.CreateExperiment.Aspects = append(e.Body.CreateExperiment.Aspects,
		ExperimentAspect{
			Data: base64.StdEncoding.EncodeToString([]byte(topdl)),
			Type: "layout",
		})
	e.Body.CreateExperiment.Profile = append(e.Body.CreateExperiment.Profile,
		DescriptionAttr{
			"description", "This is not an experiment"})

	var responseEnvelope CreateExperimentResponseEnvelope

	rsp, _, err := spiCall(XPS_HTTPS+"/createExperiment", e, &responseEnvelope)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("Server did not accept the createExperiment call - %d",
			rsp.StatusCode)
	}

	response := &responseEnvelope.Body.CreateExperimentResponse

	return response, nil
}

// Realize Experiment ----------------------------------------------------------

func RealizeExperiment(expId, circle, owner string) (
	*RealizeExperimentResponse, error) {

	e := RealizeExperimentEnvelope{}
	e.Body.RealizeExperiment.EID = expId
	e.Body.RealizeExperiment.UID = owner
	e.Body.RealizeExperiment.CID = circle

	var responseEnvelope RealizeExperimentResponseEnvelope

	rsp, _, err := spiCall(XPS_HTTPS+"/realizeExperiment", e, &responseEnvelope)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("Server did not accept the realizeExperiment call - %d",
			rsp.StatusCode)
	}

	response := &responseEnvelope.Body.RealizeExperimentResponse

	return response, nil

}

// Remove Realization ----------------------------------------------------------

func RemoveRealization(expId string) (*RemoveRealizationResponse, error) {

	e := RemoveRealizationEnvelope{}
	e.Body.RemoveRealization.Name = expId

	var responseEnvelope RemoveRealizationResponseEnvelope

	rsp, _, err := spiCall(REX_HTTPS+"/removeRealization", e, &responseEnvelope)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("Server did not accept the removeRealization call - %d",
			rsp.StatusCode)
	}

	response := &responseEnvelope.Body.RemoveRealizationResponse

	return response, nil
}

// Release Realization ----------------------------------------------------------

func ReleaseRealization(expId string) (*ReleaseRealizationResponse, error) {

	e := ReleaseRealizationEnvelope{}
	e.Body.ReleaseRealization.Name = expId

	var responseEnvelope ReleaseRealizationResponseEnvelope

	rsp, _, err := spiCall(REX_HTTPS+"/releaseRealization", e, &responseEnvelope)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("Server did not accept the releaseRealization call - %d",
			rsp.StatusCode)
	}

	response := &responseEnvelope.Body.ReleaseRealizationResponse

	return response, nil
}

// Realize Experiment ----------------------------------------------------------

func RemoveExperiment(expId string) (*RemoveExperimentResponse, error) {

	e := RemoveExperimentEnvelope{}
	e.Body.RemoveExperiment.EID = expId

	var responseEnvelope RemoveExperimentResponseEnvelope

	rsp, _, err := spiCall(XPS_HTTPS+"/removeExperiment", e, &responseEnvelope)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("Server did not accept the removeExperiment call - %d",
			rsp.StatusCode)
	}

	response := &responseEnvelope.Body.RemoveExperimentResponse

	return response, nil
}

// View Experiments ------------------------------------------------------------

func ViewExperiments(user, regex string) (*ViewExperimentsResponse, error) {

	e := ViewExperimentsEnvelope{}
	e.Body.ViewExperiments.UID = user
	e.Body.ViewExperiments.Regex = regex

	var responseEnvelope ViewExperimentsResponseEnvelope

	rsp, _, err := spiCall(XPS_HTTPS+"/viewExperiments", e, &responseEnvelope)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("Server did not accept the viewExperiments call - %d",
			rsp.StatusCode)
	}

	response := &responseEnvelope.Body.ViewExperimentsResponse

	return response, nil
}

// View Realizations -----------------------------------------------------------

func ViewRealizations(user, regex string) (*ViewRealizationsResponse, error) {

	e := ViewRealizationsEnvelope{}
	e.Body.ViewRealizations.UID = user
	e.Body.ViewRealizations.Regex = regex

	var responseEnvelope ViewRealizationsResponseEnvelope

	rsp, _, err := spiCall(REX_HTTPS+"/viewRealizations", e, &responseEnvelope)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("Server did not accept the viewRealizations call - %d",
			rsp.StatusCode)
	}

	response := &responseEnvelope.Body.ViewRealizationsResponse

	return response, nil

}

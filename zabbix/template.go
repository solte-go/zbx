package zabbix

import (
	"github.com/AlekSi/reflector"
)

type Template struct {
	TempHostID string `json:"templateid"`
	TempHost   string `json:"host"`
	TemplateID string `json:"name"`
}

type Trigger struct {
	Triggerid      string `json:"triggerid"`
	TrgExp         string `json:"expression"`
	TrgDescription string `json:"description"`
}

type Triggers []Trigger

type Templates []Template

//Get Templates
func (api *API) TemplatesGet(params Params) (res Templates, err error) {
	//func (api *API) TemplatesGet(params Params) (res Templates, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	response, err := api.CallWithError("template.get", params)
	if err != nil {
		return
	}

	reflector.MapsToStructs2(response.Result.([]interface{}), &res, reflector.Strconv, "json")
	return
}

// Gets hosts by host group Ids.
func (api *API) TemplatesGetID(ids string) (res Templates, err error) {
	return api.TemplatesGet(Params{"hostids": ids})
}

//Get Triggers
func (api *API) TriggersGet(params Params) (res Triggers, err error) {
	//func (api *API) TemplatesGet(params Params) (res Templates, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	response, err := api.CallWithError("trigger.get", params)
	if err != nil {
		return
	}

	reflector.MapsToStructs2(response.Result.([]interface{}), &res, reflector.Strconv, "json")
	return
}

// Gets hosts by host group Ids.
func (api *API) TriggersGetID(ids string) (res Triggers, err error) {
	return api.TriggersGet(Params{"hostids": ids})
}

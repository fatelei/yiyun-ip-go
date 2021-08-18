package yiyun_ip_go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type YiYun struct {
	appCode string
}

type showResBody struct {
	Region      string `json:"region,omitempty"`
	ISP         string `json:"isp,omitempty"`
	ENName      string `json:"en_name,omitempty"`
	Country     string `json:"country,omitempty"`
	IP          string `json:"ip,omitempty"`
	RetCode     int    `json:"ret_code"`
	County      string `json:"county,omitempty"`
	Continents  string `json:"continents,omitempty"`
	CityCode    string `json:"city_code,omitempty"`
	City        string `json:"city,omitempty"`
	Lnt         string `json:"lnt,omitempty"`
	ENNameShort string `json:"en_name_short,omitempty"`
	Lat         string `json:"lat,omitempty"`
	Remark      string `json:"remark,omitempty"`
}

type YiYunResponse struct {
	ShowAPIResCode  int         `json:"showapi_res_code"`
	ShowAPIResID    string      `json:"showapi_res_id,omitempty"`
	ShowAPIResError string      `json:"showapi_res_error"`
	ShowResBody     showResBody `json:"showapi_res_body"`
}

func NewYiYunClient(appCode string) YiYun {
	return YiYun{appCode: appCode}
}

func (p *YiYun) GetLocationByIP(ip string) (*YiYunResponse, error) {
	client := http.Client{
		Timeout: time.Second * 4,
	}
	url := fmt.Sprintf("https://ali-ip.showapi.com/ip?ip=%s", ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("APPCODE %s", p.appCode))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http call has error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data YiYunResponse
	if err = json.Unmarshal(body, &data); err == nil {
		return &data, nil
	} else {
		return nil, err
	}
}

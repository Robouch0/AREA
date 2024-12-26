//
// EPITECH PROJECT, 2024
// AREA
// File description:
// hfWebhookRequest
//

package hfType

type HFRepo struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type HFWebHookRequest struct {
	Watched []HFRepo `json:"watched"`
	Url     string   `json:"url"`
	Domains []string `json:"domains"`
	Secret  string   `json:"secret"`
}

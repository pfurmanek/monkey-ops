package main

import (

)

type LoginOutput struct {
	Token string `json:"token"`
	Projects []string `json:"projects"`
}

type LoginInput struct {
	User      string    `json:"user"`
	Password      string    `json:"password"`
	Url      string    `json:"url"`
}

//The third field sets a default value
type ChaosInput struct {
	Url      string    `json:"url"`
	Project string `json:"project"`
	Token string `json:"token"`
	Interval float64 `json:"interval"`
	TotalTime float64 `json:"totalTime"`
	Names      string    `json:"names"`
}

type ChaosOutput struct {
	Pods []string `json:"pods"`
}

type Scale struct {
	Kind string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata Metadata `json:"metadata"`
	Spec Spec `json:"spec"`
}

type Metadata struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
}

type Spec struct {
	Replicas float64 `json:"replicas"`
}

type DcObject struct {
	Name string `json:"name"`
	Replicas float64 `json:"replicas"`
}


type LoginInputs []LoginInput


package models

const KEY = "sensor" // ключ по которому будет храниться в key-value хранилище

type SensorsData struct {
	Value      float64 `json:"value"`
	DataIsFull bool    `json:"data_is_full"`
}

type InfoResponse struct {
	Info string `json:"info"`
}

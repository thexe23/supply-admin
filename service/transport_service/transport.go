package transport_service

import (
	"supply-admin/models"
	"time"
)

type Transport struct {
	ID       int64     `json:"id"`
	SourceID int64     `json:"sourceId"`
	TargetID int64     `json:"targetId"`
	Item     string    `json:"item"`
	Quantity int       `json:"quantity"`
	CreateAt time.Time `json:"createAt"`
	Status   int       `json:"status"`
}

func (t *Transport) Add() (int64, error) {
	return models.AddTransport(t.SourceID, t.TargetID, t.Item, t.Quantity)
}

func (t *Transport) UpdateStatus() error {
	return models.UpdateTransportStatus(t.ID, t.Status)
}

func (t *Transport) GetForSource() (error, []models.Transport) {
	return models.GetTransportForSource(t.SourceID)
}

func (t *Transport) GetForTarget() (error, []models.Transport) {
	return models.GetTransportForTarget(t.TargetID)
}

func (t *Transport)GetAll() (error, []models.Transport) {
	return models.GetTransport()
}
package store

import (
	"log"
	"sync"
	"time"

	"permit-proxy/internal/client"
	"permit-proxy/internal/enricher"
	"permit-proxy/internal/models"
)

type Store struct {
	cachedPermits []models.Permit
	mu            sync.RWMutex
}

func (s *Store) Get() ([]models.Permit, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.cachedPermits, nil
}

func (s *Store) LoadPermits() error {
	p, err := client.Paginator(client.CKANBaseUrl)
	if err != nil {
		return err
	}

	e, err := enricher.Enrich(p)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.cachedPermits = e

	return nil
}

func (s *Store) StartRefresh() {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		if err := s.LoadPermits(); err != nil {
			log.Printf("refresh failed: %v", err)
		}
	}
}

func New() *Store {
	s := Store{}

	if err := s.LoadPermits(); err != nil {
		log.Fatal(err)
	}

	go s.StartRefresh()

	return &s
}

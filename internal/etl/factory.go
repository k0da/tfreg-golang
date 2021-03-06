package etl

import (
	"github.com/k0da/tfreg-golang/internal/config"
	"github.com/k0da/tfreg-golang/internal/location"
	"github.com/k0da/tfreg-golang/internal/repo"
	"github.com/k0da/tfreg-golang/internal/storage"
	"github.com/k0da/tfreg-golang/internal/terraform"
)

// IFactory always provides valid configuration or error
type IFactory interface {
	Get() (b *Batch, err error)
}

type EtlFactory struct {
	config config.Config
}

func NewEtlFactory(config config.Config) *EtlFactory {
	return &EtlFactory{
		config: config,
	}
}

// Batch ties components of ETL. The reason is not to return all components by factory or accepts all components by ETL.
// Instead of that we use more transparent batch.
type Batch struct {
	location  location.ILocation
	repo      repo.IRepo
	storage   storage.IStorage
	terraform terraform.ITerraform
}

func (e *EtlFactory) Get() (b *Batch, err error) {
	b = new(Batch)
	b.location, err = location.NewLocation(e.config)
	if err != nil {
		return
	}
	b.storage, err = storage.NewStorage(b.location)
	if err != nil {
		return
	}
	b.terraform, err = terraform.NewTerraformProvider(b.location)
	if err != nil {
		return
	}
	b.repo, err = repo.NewGithub(b.location)
	if err != nil {
		return
	}
	return
}

package depedency_injection

import (
	"test_sawit_pro/config"
	"test_sawit_pro/controller"
	"test_sawit_pro/repository"
	"test_sawit_pro/service"
)

type Container struct {
	EstateController *controller.EstateController
}

func Init(cfg *config.Config) (*Container, error) {
	db, err := config.InitDB(cfg)
	if err != nil {
		return nil, err
	}
	// migration table
	config.Migration(db)

	// === Repository ===
	estateRepository := repository.NewEstateRepository(db)
	droneRepository := repository.NewDroneRepository(db)
	treeRepository := repository.NewTreeRepository(db)

	// === Service ===
	estateService := service.NewEstateService(estateRepository)
	droneService := service.NewDroneService(droneRepository)
	treeService := service.NewTreeService(treeRepository, estateRepository)

	// === Controller ===
	estateController := controller.NewEstateController(estateService, treeService, droneService)

	return &Container{
		EstateController: &estateController,
	}, nil
}

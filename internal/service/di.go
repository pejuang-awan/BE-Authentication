package service

import "go.uber.org/dig"

type Holder struct {
	dig.In
	Auth            Auth
	TourneyMaker    TourneyMaker
	TourneyRegistry TourneyRegistry
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewAuth); err != nil {
		return err
	}

	if err := container.Provide(NewTourneyMaker); err != nil {
		return err
	}

	if err := container.Provide(NewTourneyRegistry); err != nil {
		return err
	}

	return nil
}

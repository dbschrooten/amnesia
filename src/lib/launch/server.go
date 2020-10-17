package launch

import "amnesia/src/lib/config"

func Server() error {
	config.Load()

	if err := config.Setup(); err != nil {
		return err
	}

	return nil
}

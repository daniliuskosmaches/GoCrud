package Ingridient

type Service struct {
	repository *Repository
}

func (s *Service) IngredientService(ingridient Ingredient) error {

	s.repository.IngredientRepository(ingridient)

	if ingridient.Name == "" {
		return nil
	}

	return nil
}

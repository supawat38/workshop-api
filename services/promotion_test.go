package services_test

import (
	repo "app/app/controllers/module_promotion"
	"app/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	type testCase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expecred        int
	}

	cases := []testCase{
		{name: "applied 100", purchaseMin: 100, discountPercent: 20, amount: 100, expecred: 80},
		{name: "applied 200", purchaseMin: 100, discountPercent: 20, amount: 200, expecred: 160},
		{name: "applied 300", purchaseMin: 100, discountPercent: 20, amount: 300, expecred: 240},
		{name: "applied 50", purchaseMin: 100, discountPercent: 20, amount: 50, expecred: 50},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			promoRepo := repo.NewPomotionRepositiortMock()
			promoRepo.On("GetPromotion").Return(repo.Promotion{
				ID:              1,
				PurchaseMin:     c.purchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)

			promoService := services.NewPromotionService(promoRepo)

			//Act
			discount, _ := promoService.CalculateDiscount(c.amount)
			expected := c.expecred

			//Assert
			assert.Equal(t, expected, discount)
		})
	}

	t.Run("zero", func(t *testing.T) {
		promoRepo := repo.NewPomotionRepositiortMock()
		promoRepo.On("GetPromotion").Return(repo.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoService := services.NewPromotionService(promoRepo)

		//Act
		_, err := promoService.CalculateDiscount(0)
		assert.ErrorIs(t, err, services.ErrZeroAmount)
		promoRepo.AssertNotCalled(t, "GetPromotion")
	})

	t.Run("repo error", func(t *testing.T) {
		promoRepo := repo.NewPomotionRepositiortMock()
		promoRepo.On("GetPromotion").Return(repo.Promotion{}, errors.New("fail"))

		promoService := services.NewPromotionService(promoRepo)

		//Act
		_, err := promoService.CalculateDiscount(100)

		//Assert
		assert.ErrorIs(t, err, services.ErrRepository)
	})

}

package services

import repositories "app/app/controllers/module_promotion"

type PromotionService interface {
	CalculateDiscount(amount int) (int, error)
}

type promotionService struct {
	promoRepo repositories.PromotionRepository
}

func NewPromotionService(PromoRepo repositories.PromotionRepository) PromotionService {
	return promotionService{promoRepo: PromoRepo}
}

func (s promotionService) CalculateDiscount(amount int) (int, error) {

	if amount <= 0 {
		return 0, ErrZeroAmount
	}

	promotion, err := s.promoRepo.GetPromotion()
	if err != nil {
		return 0, ErrRepository
	}

	//ซื้อเกินราคา
	if amount >= promotion.PurchaseMin {
		return amount - (promotion.DiscountPercent * amount / 100), nil
	}

	//ซื้อไม่ถึง
	return amount, nil
}

package repository

import (
	"database/sql"
	"final-project-sanbercode/structs"
)

// SubscriptionPlanRepo is the repository for subscription plan
type SubscriptionPlanRepo struct {
	DB *sql.DB
}

// NewSubscriptionPlanRepo is the constructor for subscription plan repository
func NewSubscriptionPlanRepo(db *sql.DB) *SubscriptionPlanRepo {
	return &SubscriptionPlanRepo{DB: db}
}

// GetAll is the function to get all subscription plans
func (s *SubscriptionPlanRepo) GetAll() ([]structs.SubscriptionPlan, error) {
	var subscriptionPlans []structs.SubscriptionPlan

	rows, err := s.DB.Query("SELECT * FROM subscription_plans")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var subscriptionPlan structs.SubscriptionPlan

		err = rows.Scan(&subscriptionPlan.ID, &subscriptionPlan.Name, &subscriptionPlan.Price, &subscriptionPlan.Active_period_in_months, &subscriptionPlan.Featured, &subscriptionPlan.Created_at, &subscriptionPlan.Updated_at)
		if err != nil {
			return nil, err
		}
		subscriptionPlans = append(subscriptionPlans, subscriptionPlan)
	}

	return subscriptionPlans, nil
}

// GetByID is the function to get subscription plan by id
func (s *SubscriptionPlanRepo) GetByID(id int) (structs.SubscriptionPlan, error) {
	var subscriptionPlan structs.SubscriptionPlan

	err := s.DB.QueryRow("SELECT * FROM subscription_plans WHERE id = $1", id).Scan(&subscriptionPlan.ID, &subscriptionPlan.Name, &subscriptionPlan.Price, &subscriptionPlan.Active_period_in_months, &subscriptionPlan.Featured, &subscriptionPlan.Created_at, &subscriptionPlan.Updated_at)
	if err != nil {
		return subscriptionPlan, err
	}

	return subscriptionPlan, nil
}

// Insert is the function to insert subscription plan
func (s *SubscriptionPlanRepo) Insert(subscriptionPlan structs.SubscriptionPlan) (structs.SubscriptionPlan, error) {
	err := s.DB.QueryRow("INSERT INTO subscription_plans (name, price, active_period_in_months, features, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", subscriptionPlan.Name, subscriptionPlan.Price, subscriptionPlan.Active_period_in_months, subscriptionPlan.Featured, subscriptionPlan.Created_at, subscriptionPlan.Updated_at).Scan(&subscriptionPlan.ID)
	if err != nil {
		return subscriptionPlan, err
	}

	return subscriptionPlan, nil
}

// Update is the function to update subscription plan with id
func (s *SubscriptionPlanRepo) Update(id int, subscriptionPlan structs.SubscriptionPlan) (structs.SubscriptionPlan, error) {
	_, err := s.DB.Exec("UPDATE subscription_plans SET name = $1, price = $2, active_period_in_months = $3, features = $4, updated_at = $5 WHERE id = $6", subscriptionPlan.Name, subscriptionPlan.Price, subscriptionPlan.Active_period_in_months, subscriptionPlan.Featured, subscriptionPlan.Updated_at, id)

	if err != nil {
		return subscriptionPlan, err
	}

	return subscriptionPlan, nil
}

// Delete is the function to delete subscription plan
func (s *SubscriptionPlanRepo) Delete(id int) error {
	_, err := s.DB.Exec("DELETE FROM subscription_plans WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}

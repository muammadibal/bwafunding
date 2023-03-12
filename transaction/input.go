package transaction

import "bwafunding/user"

type TransactionDetailInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

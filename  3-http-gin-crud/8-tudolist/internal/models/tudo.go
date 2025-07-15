package models


type Todo struct {
		ID      int    `json:"id"`
	Title   string `json:"title" binding:"required"`
	Details string `json:"details"`
	UserID  int    `json:"user_id"` // Todo kimga tegishl
}

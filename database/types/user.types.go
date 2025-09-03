package types

type (
	UserAll struct {
		Id        int    `db:"id" json:"id"`
		Username  string `db:"username" json:"username"`
		Email     string `db:"email" json:"email"`
		Password  string `db:"password" json:"password"`
		CreatedAt string `db:"created_at" json:"created_at"`
	}

	User struct {
		Id        int    `db:"id" json:"id"`
		Username  string `db:"username" json:"username"`
		Email     string `db:"email" json:"email"`
		CreatedAt string `db:"created_at" json:"created_at"`
	}

	UserLogin struct {
		Email    string `db:"email" json:"email" validate:"required"`
		Password string `db:"password" json:"password" validate:"required"`
	}

	UserRegister struct {
		Username string `db:"username" json:"username" validate:"required"`
		Password string `db:"password" json:"password" validate:"required"`
		Email    string `db:"email" json:"email" validate:"required"`
	}
)

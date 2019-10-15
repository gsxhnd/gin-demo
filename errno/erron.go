package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrTypeQuery        = &Errno{Code: 10003, Message: "Error type of params."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}

	// token errors
	ErrToken           = &Errno{Code: 10101, Message: "Error occurred while signing the JSON web token."}
	ErrTokenExpire     = &Errno{Code: 10102, Message: "The token was expire."}
	ErrTokenWillExpire = &Errno{Code: 10103, Message: "The token will be expire."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrKeyInvalid        = &Errno{Code: 20105, Message: "The key was invalid."}

	//ErrPay               = &Errno{Code: 20105, Message: "Pay fail."}
)

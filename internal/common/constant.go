package common

const (
	SUCCESS_CODE               = 0
	BAD_REQUEST_ERROR_CODE     = 100
	NOT_AUTHORIZED_ERROR_CODE  = 101
	FORBIDDEN_ERROR_CODE       = 103
	INTERNAL_SERVER_ERROR_CODE = 199

	SUCCESS_DELETE_MESSAGE        = "Register with the following ID: '%d' was deleted successfully!"
	INTERNAL_SERVER_ERROR_MESSAGE = "Something is broken on our side :(. Sorry for the inconvenience!"
	FORBIDDEN_ERROR_MESSAGE       = "Hummmm... It seems you are not allowed to do such a thing. Ask for your manager help!"
	NOT_AUTHORIZED_ERROR_MESSAGE  = "Stop right there! You are unauthorized!"
	DATE_BAD_REQUEST              = "key: 'TaskCreate.Summary' Error:Field validation for 'Performed' failed on the 'format' regex"

	DATE_REGEX              = `^([0-2][0-9]|(3)[0-1])(\/)(((0)[0-9])|((1)[0-2]))(\/)\d{4}$`
	DATE_FORMAT             = "02/01/2006"
	DB_ERROR_NOT_AUTHORIZED = "record not found"
)

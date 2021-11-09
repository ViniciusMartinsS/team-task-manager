package common

const (
	SUCCESS_CODE                = 0
	BAD_REQUEST_ERROR_CODE      = 100
	NOT_AUTHORIZED_ERROR_CODE   = 101
	RECORD_NOT_FOUND_ERROR_CODE = 104
	FORBIDDEN_ERROR_CODE        = 103
	INTERNAL_SERVER_ERROR_CODE  = 199

	SUCCESS_DELETE_MESSAGE         = "Record with the following ID: '%d' was deleted successfully!"
	INTERNAL_SERVER_ERROR_MESSAGE  = "Something is broken on our side :(. Sorry for the inconvenience!"
	RECORD_NOT_FOUND_ERROR_MESSAGE = "Hmmmm... We could not find the requested record. Are you sure it exists? Are you sure it belongs to you?"
	FORBIDDEN_ERROR_MESSAGE        = "Hmmmm... It seems you are not allowed to do such a thing. Ask for your manager help!"
	NOT_AUTHORIZED_ERROR_MESSAGE   = "Stop right there! You are unauthorized!"
	DATE_BAD_REQUEST               = "key: 'TaskCreate.Summary' Error:Field validation for 'Performed' failed on the 'format' regex"
	RECORD_NOT_FOUND_LIST_MESSAGE  = "You do not have any tasks. Create a new one & let's get to work! ;)"

	DATE_REGEX          = `^([0-2][0-9]|(3)[0-1])(\/)(((0)[0-9])|((1)[0-2]))(\/)\d{4}$`
	DATE_FORMAT         = "02/01/2006"
	DB_RECORD_NOT_FOUND = "record not found"
)

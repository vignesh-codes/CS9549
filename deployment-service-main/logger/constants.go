package logger

import "net/http"

const (
	KEY_PAYLOAD        string = "payload"
	KEY_ERROR          string = "error"
	KEY_MESSAGE        string = "message"
	KEY_FUNCTION       string = "function"
	KEY_PATH           string = "path"
	KEY_KEY            string = "key"
	KEY_TIME_TAKEN     string = "time_taken"
	KEY_S3_FILE        string = "s3_file"
	KEY_ITEMS_COUNT    string = "items_count"
	KEY_SQS_QUEUE_NAME string = "sqs_queue"
	KEY_EVENT          string = "event"
)

type ErrResp struct {
	Error_Msg string
	Code      string
}

//codes
const (
	INVALID_REQ = http.StatusBadRequest
	SERVER_ERR  = http.StatusInternalServerError
	SUCCESS     = http.StatusOK
	AUTH_ERR    = http.StatusUnauthorized
)

var (
	SUCCESS_RESP = ErrResp{
		Error_Msg: "SUCCESS",
		Code:      "N10",
	}

	UNAUTHORIZED_RESP = ErrResp{
		Error_Msg: "UNAUTHORIZED",
		Code:      "N20",
	}

	INVALID_REQUEST_PAYLOAD_RESP = ErrResp{
		Error_Msg: "INVALID_REQUEST",
		Code:      "N30",
	}

	SERVER_ERROR_RESP = ErrResp{
		Error_Msg: "SERVER_ERROR",
		Code:      "N40",
	}

	UNSUPPORTED_MESSAGE_RESP = ErrResp{
		Error_Msg: "UNSUPPORTED_MESSAGE",
		Code:      "N000",
	}

	UID_MISMATCH_RESP = ErrResp{
		Error_Msg: "UID_MISMATCH",
		Code:      "N001",
	}

	USER_LIMIT_REACHED_RESP = ErrResp{
		Error_Msg: "You have reached max limit of making request, please try again after some time!",
		Code:      "N003",
	}
)

const (
	AwsSession                  = "AWS_SESSION"
	S3ConnectionError           = "S3_CONNECTION_ERROR"
	DynamoConnectionError       = "DYNAMODB_CONNECTION_ERROR"
	RedisError                  = "REDIS_ERROR"
	DynamoPUTError              = "DYNAMODB:PUT"
	DynamoGetError              = "DYNAMODB:GET"
	DynamoUpdateError           = "DYNAMODB:UPDATE"
	DynamoUpdateExpressionError = "DYNAMODB:UPDATE_EXPRESSION"
	DynamoUnmarshallError       = "DYNAMODB:UNMARSHALL"
	DynamoQueryError            = "DYNAMODB:QUERY"
	DynamoQueryExpressionError  = "DYNAMODB:QUERY_EXPRESSION"
	DynamoTransactionError      = "DYNAMODB:TRANSACTION_ERROR"
	SQSError                    = "SQS:ERROR"
	SQSHandler                  = "SQS_HANDLER"
	SQSUnknownEvent             = "SQS:UNKNOWN_EVENT"
	S3FileSaveError             = "S3:FAILED_TO_SAVE_FILE"
	S3FileDownloadError         = "S3:FAILED_TO_DOWNLOAD_FILE"
	CSVFileParse                = "CSV:FILE_PARSE_FAILED"
	StructToMap                 = "GOLANG:STRUCT_TO_MAP"
	BindJSONtoStruct            = "GOLANG:BINDING_JSON_TO_STRUCT_FAILED"
	JSONParseError              = "JSON_PARSE_ERROR"
	JSONUnmarshallError         = "JSON_UNMARSHALL_ERROR"
	NILValueError               = "NIL_VALUE_ERROR"
	TransactionHash             = "VICTOR_TRANSACTION_HASH"
	ErrExternalServiceDown      = "EXTERNAL_SERVICE_DOWN"
)

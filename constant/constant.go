package constant

const (
	HTTP_STATUS_SUCCESS = "success"
	HTTP_STATUS_FAIL    = "fail"
	HTTP_STATUS_ERROR   = "error"

	HTTP_CODE_200_OK                    = "200" //http.StatusOk
	HTTP_CODE_400_BAD_REQUEST           = "400" //http.StatusBadRequest
	HTTP_CODE_401_UNAUTHORIZED          = "401" //http.StatusUnauthorized
	HTTP_CODE_500_INTERNAL_SERVER_ERROR = "500" //http.StatusInternalServerError
	HTTP_REQUEST_HEADER_AUTHRORIZATION  = "Authorization"

	CONTRACT_ADDRESS                        = "0xE3518Afd0a45439c737823c3EDcb85611FcEbB3b"      // 合约地址
	AUCTION_ADDRESS                         = "0x144255298efF5AFd8000B9fba74e4a4F2aFD6b20"      // 竞拍池地址
	BSC_RPC_DEV                             = "https://data-seed-prebsc-1-s1.binance.org:8545/" // bsc测试网rpc
	BSC_RPC_POD                             = "https://bsc-dataseed.binance.org"                // bsc主网rpc
	BSC_HOLDER_LIST_BY_CONTRACT_ADDRESS_URL = "https://api.bscscan.com/api?module=token&action=tokenholderlist"
	BSC_API_KEY                             = "EASEE1AGT2SCG99X5PJAJAIAHYW91TGAUP" // BSC API Key

	API_GET_LAST_RANK          = "/lastRank"
	API_GET_CURRENT_RANK       = "/currentRank"
	API_GET_REFRESHTIME        = "/refreshtime"
	API_GET_STATISTICS_RECOMM  = "/recomm/:holder"
	API_POST_STATISTICS_RECOMM = "/recomm/:holder/:recomm"

	//http request
	HTTP_REQUEST_PARAMS_JSON_FORMAT_ERROR_CODE        = "500002001"
	HTTP_REQUEST_PARAMS_JSON_FORMAT_ERROR_MSG         = "Request JSON format was error"
	HTTP_REQUEST_PARAMS_NULL_ERROR_CODE               = "500002002"
	HTTP_REQUEST_PARAMS_NULL_ERROR_MSG                = "Request params value can not be null"
	PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_CODE             = "500002003"
	PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_MSG              = "Page number or page size format was wrong"
	HTTP_REQUEST_SEND_REQUEST_RETUREN_ERROR_CODE      = "500002004"
	HTTP_REQUEST_SEND_REQUEST_RETUREN_ERROR_MSG       = "Return error when sending http request"
	HTTP_REQUEST_PARSER_RESPONSE_TO_STRUCT_ERROR_CODE = "500002005"
	HTTP_REQUEST_PARSER_RESPONSE_TO_STRUCT_ERROR_MSG  = "Parse http request to structure occurred error"
	HTTP_REQUEST_PARSER_STRUCT_TO_REQUEST_ERROR_CODE  = "500002006"
	HTTP_REQUEST_PARSER_STRUCT_TO_REQUEST_ERROR_MSG   = "Parse structure to http request request occurred error"
	HTTP_REQUEST_GET_RESPONSE_ERROR_CODE              = "500002007"
	HTTP_REQUEST_GET_RESPONSE_ERROR_MSG               = "Get http response occurred error"
	HTTP_REQUEST_PARAM_TYPE_ERROR_CODE                = "500002008"
	HTTP_REQUEST_PARAM_TYPE_ERROR_MSG                 = "Http request param type is error"
	HTTP_REQUEST_PARAM_VALUE_ERROR_CODE               = "500002009"
	HTTP_REQUEST_PARAM_VALUE_ERROR_MSG                = "Request parameter invalid"

	//database err 003
	GET_RECORD_COUNT_ERROR_CODE  = "500003001"
	GET_RECORD_COUNT_ERROR_MSG   = "Get data total count from db occurred error"
	GET_RECORD_lIST_ERROR_CODE   = "500003002"
	GET_RECORD_lIST_ERROR_MSG    = "Get data records from db occurred error"
	SAVE_DATA_TO_DB_ERROR_CODE   = "500003003"
	SAVE_DATA_TO_DB_ERROR_MSG    = "Saving data to database occurred error"
	UPDATE_DATA_TO_DB_ERROR_CODE = "500003004"
	UPDATE_DATA_TO_DB_ERROR_MSG  = "Updating data to database occurred error"
)

var ADDRESS_LIST = []string{
	"0xb3a7C64F9065c0a6A9EB57597943A3d187733238",
	"0x182AD90BFBFC9b972fE5298A0825314d5dDA3642",
	"0x8C2B33a09dA1Be414591204424f36b1F7dA14241",
	"0x1ec9dcf7DCd28AFb87E96511BfF3494423d2B50A",
	"0x06aD629119493cCc0bc5423aeAd7e37cf31CEBAE",
	"0xf71BD23CF2322FF8ff2EB42Cda7b0157956b7449",
	"0x3ce4b119488C10a62fE379a26cF99ba3bEd2E834",
	"0xc49E8851983c7aD445a1697bfF0Aacc85182C4CF",
	"0xD986Cfb4c7C370A6A81e24032d61836744D63647",
	"0x144255298efF5AFd8000B9fba74e4a4F2aFD6b20",
	"0x6C116406FcD525BAC21dd52954C1eDa4aa9368B2",
}

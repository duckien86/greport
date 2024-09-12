package verifier

import (
	"fmt"
	"greport/common"
	"greport/component/hasher"
	"greport/component/myredis"
	"greport/component/sender"

	"context"
	"errors"
	"strconv"
	"time"
)

const (
	StatusInit = iota // status
	StatusOk
	StatusFail
)
const (
	Email          = "email" // verify type
	Sms            = "sms"   // verify type
	CodeLength     = 4       // verify code length
	MaxRetry       = 3
	MaxRequest     = 3
	ExpireInMinute = 5
)

type Verifier struct {
	Via        string      // they way to send verify. Ex: email or sms or ...
	Info       string      // email or phonenumber or ...
	MaxRequest int         // max request to verify per verifyInfo
	MaxRetry   int         // max retry to verify fail
	Expire     int         // expire time to verify. Ex: 5 minutes
	RelateData interface{} // relate data for orther biz
}

type verifyDataStruct struct {
	Id           string
	Code         string
	Info         string
	RetryCount   int
	RequestCount int
	MaxRequest   int // max request to verify per verifyInfo
	MaxRetry     int // max retry to verify fail
	RelateData   interface{}
}

type VerifyRequest struct {
	Id   string `json:"verify_id" validate:"required" example:"abc31233"`
	Code string `json:"verify_code" validate:"required" example:"1234"`
	Info string `json:"verify_info" validate:"required" example:"email or phone number"`
}

// NewVerifier is a constructor to create a new Verifier instance
//
// via: the way to send verify. Ex: email or sms or ...
//
// info: email or phonenumber or ...
//
// maxRetry: max retry to verify fail
//
// maxRequest: max request to verify per verifyInfo
//
// expire: expire time to verify. Ex: 5 minutes
//
// relateData: relate data for orther biz
//
// return a new Verifier instance
func NewVerifier(via, info string, maxRetry, maxRequest, expire int, relateData interface{}) *Verifier {
	return &Verifier{
		Via:        via,
		Info:       info,
		MaxRequest: maxRequest,
		MaxRetry:   maxRetry,
		Expire:     expire,
		RelateData: relateData,
	}
}

// Send verify code (OTP) request via email or sms
func (verify *Verifier) SendVerifyCode(ctx context.Context, serviceName, bizName string) (string, error) {
	myRedisCli := myredis.NewClient(myredis.DB_USER)                   // create new redis client
	defer myRedisCli.Close()                                           // close redis client
	redisKey, err := myredis.GenKey(serviceName, bizName, verify.Info) // generate redis key
	if err != nil {
		return "", common.ErrInternal(err)
	}
	sentCount := myRedisCli.HGet(context.Background(), redisKey, "RequestCount").Val()
	sentCountInt := 0
	if len(sentCount) > 0 { // check if phone is already sent
		sentCountInt, _ = strconv.Atoi(sentCount)
	}
	if sentCountInt >= verify.MaxRequest { // check if phone is already sent >= MaxRequest times at this time
		return "", ErrExceededLimitRequest
	}
	// generate code
	code := common.GenerateOTP(CodeLength)
	hasher := hasher.New(hasher.TypeSha256)
	verifyId := hasher.Hash(verify.Info + code)

	redisData := verifyDataStruct{
		Id:           verifyId,
		Code:         code,
		Info:         verify.Info,
		RetryCount:   0,
		RequestCount: sentCountInt + 1,
		MaxRequest:   verify.MaxRequest,
		MaxRetry:     verify.MaxRetry,
		RelateData:   verify.RelateData,
	}
	redisDataMap := common.StructToMap(redisData)
	// save code to redis
	if err := myRedisCli.HSet(context.Background(), redisKey, redisDataMap).Err(); err != nil {
		return "", common.ErrInternal(err)
	}
	if err := myRedisCli.Expire(context.Background(), redisKey, time.Duration(verify.Expire)*time.Minute).Err(); err != nil {
		return "", common.ErrInternal(err)
	}

	switch verify.Via {
	case Sms:
		smsData := sender.SmsData{
			To:      verify.Info,
			Content: fmt.Sprintf("Your OTP is: %s", code),
		}
		smsService := sender.NewSms()
		if err := smsService.Send(smsData, true); err != nil {
			return "", err
		}
	case Email:
		emailData := sender.EmailData{
			To:      verify.Info,
			Content: fmt.Sprintf("Your OTP is: %s", code),
		}
		emailService := sender.NewEmail()
		if err := emailService.Send(emailData, true); err != nil {
			return "", err
		}

	default:
		return "", ErrInvalidVerify
	}

	return verifyId, nil
}

// Verify code(OTP) receive from email or sms.
// serviceName: service name
// bizName: biz name
// vReqData: verify data
//
// Return relate data if verify success
func DoVerify(serviceName, bizName string, vReqData *VerifyRequest) (string, error) {
	myRedisCli := myredis.NewClient(myredis.DB_USER) // create new redis client
	defer myRedisCli.Close()                         // close redis client
	strRtn := ""
	redisKey, err := myredis.GenKey(serviceName, bizName, vReqData.Info) // generate redis key
	if err != nil {
		return strRtn, common.ErrInternal(err)
	}
	redisData := myRedisCli.HGetAll(context.Background(), redisKey).Val()
	if len(redisData) == 0 { // if key not found
		return strRtn, ErrCannotVerify
	}
	myRedisCli.HIncrBy(context.Background(), redisKey, "RetryCount", 1) // increase retry count
	retryCount, _ := strconv.Atoi(redisData["RetryCount"])              // get retry count
	maxRetry, _ := strconv.Atoi(redisData["MaxRetry"])                  // get max retry
	if retryCount >= maxRetry {                                         // if exceed max retry
		myRedisCli.Del(context.Background(), redisKey) // delete verify data from redis
		return strRtn, ErrExceededLimitRetry
	}
	if vReqData.Code != redisData["Code"] || vReqData.Id != redisData["Id"] { // if id or code not match
		return strRtn, ErrCannotVerify
	}
	myRedisCli.Del(context.Background(), redisKey) // clear unused data from redis
	strRtn = redisData["RelateData"]
	return strRtn, nil
}

var (
	ErrCannotVerify = common.NewCustomError(
		errors.New("cannot verify account"),
		"cannot verify account",
		"ErrCannotVerify",
	)
	ErrExceededLimitRequest = common.NewCustomError(
		errors.New("exceeded limit of sending verify code"),
		"exceeded limit of sending verify code",
		"ErrExceededLimitRequest",
	)
	ErrExceededLimitRetry = common.NewCustomError(
		errors.New("exceeded limit of verify code"),
		"exceeded limit of verify code",
		"ErrExceededLimitRetry",
	)
	ErrInvalidVerify = common.NewCustomError(
		errors.New("invalid verify request"),
		"invalid verify request",
		"ErrInvalidVerify",
	)

	// ErrCannotGetConfig = common.NewCustomError(
	// 	errors.New("cannot get config"),
	// 	"cannot get config",
	// 	"ErrCannotGetConfig",
	// )
)

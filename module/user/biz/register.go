package userbiz

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/verifier"
	usermodel "2ndbrand-api/module/user/model"
	"context"
	"encoding/json"
)

const (
	Email            = "email" // verify type
	Sms              = "sms"   // verify type
	ServiceName      = "verify"
	BizName          = "RegisterWithVerify" // verify action
	MaxRetryVerify   = 3
	MaxRequestVerify = 3
	ExpVerifyMinute  = 10
)

type RegisterStore interface {
	FindUser(context context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.Users, error)
	Create(context context.Context, data *usermodel.UserCreate) error
}
type Hasher interface {
	Hash(data string) string
}

// type Verifier interface {
// 	SendVerifyCode(context context.Context, serviceName, bizName, verifyBy, verifyData string) (string, error)
// }

type RegisterBiz struct {
	store  RegisterStore
	hasher Hasher
	// verifier Verifier
}

func NewRegisterBiz(store RegisterStore, hasher Hasher) *RegisterBiz {
	return &RegisterBiz{
		store:  store,
		hasher: hasher,
	}
}

// RegisterUser is method to register user with verify request
// return verifyId and error
func (biz *RegisterBiz) RegisterUser(
	ctx context.Context,
	userData *usermodel.UserCreate,
) (string, error) {
	verifyId := "" // init verifyId
	switch userData.Verify {
	case Email:
		userData.Email = userData.Username
	case Sms:
		userData.Phone = userData.Username
	}
	// validate request data
	if details, err := common.ValidateStruct(userData); err != nil {
		return verifyId, common.ErrValidationData(err, details)
	}
	// check phone is existed
	user, _ := biz.store.FindUser(ctx, map[string]interface{}{"username": userData.Username})
	if user != nil {
		return verifyId, usermodel.ErrPhoneIsExisted
	}
	jsonRelateData, _ := json.Marshal(userData)
	// init verifier
	verifier := verifier.NewVerifier(
		userData.Verify, userData.Username,
		MaxRequestVerify, MaxRetryVerify,
		ExpVerifyMinute, string(jsonRelateData),
	)
	verifyId, err := verifier.SendVerifyCode(ctx, ServiceName, BizName)
	if err != nil {
		return verifyId, err
	}

	return verifyId, nil
}

// VerifyAndCreateUser is method to verify and create user
func (biz *RegisterBiz) VerifyAndCreateUser(
	ctx context.Context,
	verifyData *verifier.VerifyRequest,
	userData *usermodel.UserCreate,
) error {
	// validate request data
	if details, err := common.ValidateStruct(verifyData); err != nil {
		return common.ErrValidationData(err, details)
	}
	// verify user info
	jsonUserData, err := verifier.DoVerify(ServiceName, BizName, verifyData)
	if err != nil {
		// return err
		return common.ErrInvalidRequest(err)
	}
	if err := json.Unmarshal([]byte(jsonUserData), &userData); err != nil { //
		return common.ErrInvalidRequest(err)
	}
	userData.Status = usermodel.StatusActive
	userData.Role = usermodel.RoleName
	userData.Salt = common.GenSalt(20)
	userData.Password = biz.hasher.Hash(userData.Salt + userData.Password)

	// create user
	if err := biz.store.Create(ctx, userData); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}

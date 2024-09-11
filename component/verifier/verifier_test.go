package verifier

import (
	"2ndbrand-api/component/myredis"
	"2ndbrand-api/component/mytest"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendVerifyCode_DoVerify_Success(t *testing.T) {
	serviceName := "verify"
	bizName := "register"
	testcaseLimit := map[string]int{ // testcase limit
		"success_sms":   20,
		"success_email": 20,
		// "fail_sms": 1,
		// "fail_email":    1,
	}
	testcaseMap := map[string][]Verifier{}
	for key, limit := range testcaseLimit {
		for i := 0; i < limit; i++ {
			testcaseMap[key] = append(testcaseMap[key], RandomVerifierStruct(key))
		}
	}
	for key, verifiers := range testcaseMap {
		for _, v := range verifiers {
			t.Run(key, func(t *testing.T) {
				ctx := context.Background()
				id, err := v.SendVerifyCode(ctx, serviceName, bizName)
				keyParts := strings.Split(key, "_")
				// Use keyParts as needed
				fmt.Println("SendVerifyCode::input::", v)
				fmt.Printf("SendVerifyCode::return::key-%s::id-%s", key, id)
				if keyParts[0] == "success" {
					assert.Nil(t, err)
					assert.NotEmpty(t, id)
					fmt.Println()
					vReqData := &VerifyRequest{
						Id:   id,
						Code: getOtpCodeFromRedis(serviceName, bizName, v),
						Info: v.Info,
					}
					userId, err := DoVerify(serviceName, bizName, vReqData)
					fmt.Printf("DoVerify::UserID::%s ", userId)
					assert.NotEmpty(t, userId)
					assert.NoError(t, err)
				} else {
					assert.NotNil(t, err)
					assert.Empty(t, id)
				}
			})
		}
	}
}

func Test_SendVerifyCode_DoVerify_Fail(t *testing.T) {
	serviceName := "verify"
	bizName := "register"
	testcaseLimit := map[string]int{ // testcase limit
		// "success_sms":   20,
		// "success_email": 20,
		"fail_sms": 1,
		// "fail_email":    1,
	}
	testcaseMap := map[string][]Verifier{}
	for key, limit := range testcaseLimit {
		for i := 0; i < limit; i++ {
			testcaseMap[key] = append(testcaseMap[key], RandomVerifierStruct(key))
		}
	}
	for key, verifiers := range testcaseMap {
		for _, v := range verifiers {
			t.Run(key, func(t *testing.T) {
				ctx := context.Background()
				id, err := v.SendVerifyCode(ctx, serviceName, bizName)
				// Use keyParts as needed
				fmt.Println("SendVerifyCode::input::", v)
				assert.NotNil(t, err)
				fmt.Println("SendVerifyCode::error::", err)
				assert.Empty(t, id)
				// fmt.Printf("SendVerifyCode::return::key-%s::id-%s", key, id)
				// fmt.Println()
				// vReqData := &VerifyRequest{
				// 	Id:   id,
				// 	Code: getOtpCodeFromRedis(serviceName, bizName, v),
				// 	Info: v.Info,
				// }
				// userId, err := DoVerify(serviceName, bizName, vReqData)
				// fmt.Printf("DoVerify::UserID::%s ", userId)
				// assert.NotEmpty(t, userId)
				// assert.NoError(t, err)

			})
		}
	}
}

func RandomVerifierStruct(testCaseType string) Verifier {
	user := map[string]string{
		"username":   mytest.RandomPhoneNumber(),
		"password":   mytest.RandomString(10),
		"first_name": mytest.RandomFirstName(),
		"last_name":  mytest.RandomLastName(),
		"phone":      mytest.RandomPhoneNumber(),
		"email":      mytest.RandomEmail(),
		"verify":     mytest.RandomVia(),
	}
	switch testCaseType {
	case "success_sms":
		user["username"] = mytest.RandomPhoneNumber()
		user["phone"] = mytest.RandomPhoneNumber()
		user["verify"] = Sms
	case "fail_sms":
		user["username"] = ""
		user["phone"] = mytest.RandomPhoneNumber()
		user["verify"] = Sms
	case "success_email":
		user["username"] = mytest.RandomEmail()
		user["email"] = mytest.RandomPhoneNumber()
		user["verify"] = Email
	case "fail_email":
		user["username"] = ""
		user["email"] = ""
		user["verify"] = Email
	}
	createUser := mytest.RandomCreateUserStruct(user["verify"], user["username"])
	jsonCreateUser, _ := json.Marshal(createUser)
	return Verifier{user["verify"], user["username"], 3, 3, 5, string(jsonCreateUser)}
}

func getOtpCodeFromRedis(serviceName, bizName string, vReqData Verifier) string {
	myRedisCli := myredis.NewClient(myredis.DB_USER)                   // create new redis client
	defer myRedisCli.Close()                                           // close redis client
	redisKey, _ := myredis.GenKey(serviceName, bizName, vReqData.Info) // generate redis key
	return myRedisCli.HGet(context.Background(), redisKey, "Code").Val()
}

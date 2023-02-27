package utils

/**
账号密码进行md5加密
*/
func EncryptIdAndSecret(AccessKeyId string, AccessKeySecret string) string {
	//1、将账号密码拼接
	str := AccessKeyId + "#" + AccessKeySecret
	if len(str) != 0 {
		return MD5V([]byte(str))
	}
	return ""
}

package common

func GetSecret(key string) string {
	secrets := map[string]string{
		"AWS_SECRET_KEY": "",
	}
	//
	return secrets[key]
}

package common

import "net/http"

func Cookie(Key string, Value string) *http.Cookie {
	return &http.Cookie{
		Name:  Key,
		Value: Value,
	}
}
func SetCookie(w http.ResponseWriter, Key string, Value string) *http.Cookie {
	return &http.Cookie{
		Name:  Key,
		Value: Value,
	}
}

func GetCookie(r *http.Request, Key string, DefaultValue string) string {
	c, e := r.Cookie(Key)
	if e != nil {
		return DefaultValue
	}
	return c.Value
}

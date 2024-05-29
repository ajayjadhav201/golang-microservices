package common

// Custom error type for email already registered
type EmailAlreadyRegistered struct {
	Email string
}

func (e *EmailAlreadyRegistered) Error() string {
	return Sprintf("The email %s is already registered", e.Email)
}

// Custom error type for mobile number already registered
type MobileNumberAlreadyRegistered struct {
	MobileNumber string
}

func (m *MobileNumberAlreadyRegistered) Error() string {
	return Sprintf("The mobile number %s is already registered", m.MobileNumber)
}

//if both Email and Mobile Number exists
type EmailMobileAlreadyRegistered struct {
	Email        string
	MobileNumber string
}

func (m *EmailMobileAlreadyRegistered) Error() string {
	return Sprintf("The email %s and mobile number %s are already registered", m.Email, m.MobileNumber)
}

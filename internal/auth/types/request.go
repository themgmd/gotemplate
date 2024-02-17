package types

type InitRegistrationRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (irr InitRegistrationRequest) Validate() error {
	return nil
}

type InitRegistrationResponse struct {
	Identifier string `json:"identifier"`
	QrCode     []byte `json:"qrCode"`
	Secret     string `json:"secret"`
}

type FinishRegistrationRequest struct {
	Identifier string `json:"identifier"`
	Email      string `json:"email"`
	Code       string `json:"code"`
}

func (frr FinishRegistrationRequest) Validate() error {
	return nil
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

func (lr LoginRequest) Validate() error {
	return nil
}

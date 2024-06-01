package helperuser

type HelperuserInterface interface {
	ConvertToNullableString(value string) *string
	DereferenceString(s *string) string
}

type helperuser struct {
}

func NewHelperService() HelperuserInterface {
	return &helperuser{}
}

func (h *helperuser) ConvertToNullableString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func (h *helperuser) DereferenceString(s *string) string {
	if s == nil {
		return "" // atau nilai default lain yang sesuai
	}
	return *s
}

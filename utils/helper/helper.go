package helper

type HelperInterface interface {
	ConvertToNullableString(value string) *string
}

type helper struct{}

func NewHelperService() HelperInterface {
	return &helper{}
}

func (h *helper) ConvertToNullableString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

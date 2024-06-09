package midtrans

import (
	"PetPalApp/app/configs"

	"github.com/veritrans/go-midtrans"
)

// GetMidtransClient initializes Midtrans client
func GetMidtransClient(cfg *configs.AppConfig) midtrans.Client {
    c := midtrans.NewClient()
    c.ServerKey = cfg.MIDTRANS_SERVER_KEY
    c.ClientKey = cfg.MIDTRANS_CLIENT_KEY
    c.APIEnvType = midtrans.Sandbox // Use midtrans.Production for production environment
    return c
}

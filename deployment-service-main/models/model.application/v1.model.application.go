package model_application

type Application struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Service      string `json:"service"`
}

func (app *Application) PrimaryKey() map[string]interface{} {
	return map[string]interface{}{
		"client_id":     app.ClientID,
		"client_secret": app.ClientSecret,
	}
}

func (app *Application) TableName() string {
	return "project-secrets"
}

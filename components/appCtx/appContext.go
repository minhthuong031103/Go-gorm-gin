package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetDb() *gorm.DB
	GetSecretKey() string
	GetServerHost() string
}

type appContext struct {
	Db         *gorm.DB
	SecretKey  string
	serverHost string
}

func NewAppContext(db *gorm.DB, secretKey, serverHost string) *appContext {
	return &appContext{
		Db:         db,
		SecretKey:  secretKey,
		serverHost: serverHost,
	}
}

func (a *appContext) GetDb() *gorm.DB {
	return a.Db
}

func (a *appContext) GetSecretKey() string {
	return a.SecretKey
}

func (a *appContext) GetServerHost() string {
	return a.serverHost
}

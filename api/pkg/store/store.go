package store

import (
	"context"

	"github.com/shellhub-io/shellhub/pkg/models"
)

type Store interface {
	ListDevices(ctx context.Context, perPage int, page int) ([]models.Device, error)
	GetDevice(ctx context.Context, uid models.UID) (*models.Device, error)
	DeleteDevice(ctx context.Context, uid models.UID) error
	AddDevice(ctx context.Context, d models.Device) error
	RenameDevice(ctx context.Context, uid models.UID, name string) error
	LookupDevice(ctx context.Context, namespace, name string) (*models.Device, error)
	UpdateDeviceStatus(ctx context.Context, uid models.UID, online bool) error
	ListSessions(ctx context.Context, perPage int, page int) ([]models.Session, error)
	GetSession(ctx context.Context, uid models.UID) (*models.Session, error)
	CreateSession(ctx context.Context, session models.Session) (*models.Session, error)
	SetSessionAuthenticated(ctx context.Context, uid models.UID, authenticated bool) error
	KeepAliveSession(ctx context.Context, uid models.UID) error
	DeactivateSession(ctx context.Context, uid models.UID) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserByTenant(ctx context.Context, tenant string) (*models.User, error)
	GetDeviceByMac(ctx context.Context, mac string, tenant string) (*models.Device, error)
	GetDeviceByName(ctx context.Context, name string, tenant string) (*models.Device, error)
	CountDevices(ctx context.Context) (int64, error)
	CountSessions(ctx context.Context) (int64, error)
	CreateToken(ctx context.Context, token models.Token) (*models.Token, error)
	GetToken(ctx context.Context, id string) (*models.Token, error)
	ListTokens(ctx context.Context, perPage int, page int) ([]models.Token, error)
}

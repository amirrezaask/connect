package testutils

import (
	"context"
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/amirrezaask/connect/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func GetDB() (*sql.DB, error) {
	return sql.Open("postgres", "user=connect password=connect dbname=connect sslmode=disable")
}

func AddUser(db *sql.DB, userID string) error {
	u := &models.User{
		ID: userID,
	}
	return u.Insert(context.TODO(), db, boil.Infer())
}

func AddHub(db *sql.DB, hubID string) error {
	h := &models.Hub{
		ID: hubID,
	}
	return h.Insert(context.TODO(), db, boil.Infer())
}

func RemoveHub(db *sql.DB, hubID string) error {
	_, err := models.Hubs(models.HubWhere.ID.EQ(hubID)).DeleteAll(context.TODO(), db)
	return err
}

func RemoveUser(db *sql.DB, userID string) error {
	_, err := models.Users(models.UserWhere.ID.EQ(userID)).DeleteAll(context.TODO(), db)
	return err
}

func RemoveChannel(db *sql.DB, channelID string) error {
	_, err := models.Channels(models.ChannelWhere.ID.EQ(channelID)).DeleteAll(context.Background(), db)
	return err
}

func AddChannel(db *sql.DB, hubID, channelID string) error {
	return (&models.Channel{ID: channelID, HubID: hubID}).Insert(context.TODO(), db, boil.Infer())
}

func MakeRequest(method string, body io.Reader) (*http.Request, *httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(method, "/", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	return req, rec, echo.New().NewContext(req, rec)
}

func AddUserToChannel(db *sql.DB, userID string, channelID string) error {
	_, err := db.Exec(`INSERT INTO channel_users (user_id, channel_id) VALUES ($1, $2)`, userID, channelID)
	return err
}
func AddUserToHub(db *sql.DB, userID string, hubID string) error {
	_, err := db.Exec(`INSERT INTO hub_users (user_id, hub_id) VALUES ($1, $2)`, userID, hubID)
	return err
}

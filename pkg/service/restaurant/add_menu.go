package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	"Go_Food_Delivery/pkg/handler/restaurant/unsplash"
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
)

var ImageUpdateLock *sync.Mutex = &sync.Mutex{}

func (restSrv *RestaurantService) AddMenu(ctx context.Context, menu *restaurant.MenuItem) (bool, error) {
	_, err := restSrv.db.NewInsert().Model(menu).Exec(ctx)
	if err != nil {
		return false, err
	}

	menuImageURL := unsplash.GetUnSplashImageURL(menu.Name)
	imageFileName := fmt.Sprintf("menu_item_%d.jpg", menu.MenuID)
	imageFileLocalPath := fmt.Sprintf("uploads/%s", imageFileName)
	imageFilePath := filepath.Join(os.Getenv("LOCAL_STORAGE_PATH"), imageFileName)
	err = unsplash.DownloadImageToDisk(menuImageURL, imageFilePath)
	if err != nil {
		slog.Info("UnSplash Failed to Download Image", "error", err)
	}

	go postSaveUpdateImage(restSrv.db, ctx, menu.MenuID, imageFileLocalPath)
	return true, nil
}

func postSaveUpdateImage(db *bun.DB, ctx context.Context, menuID int64, imageURL string) {
	ImageUpdateLock.Lock()
	defer ImageUpdateLock.Unlock()
	_, err := db.NewUpdate().Table("menu_item").
		Set("photo = ?", imageURL).
		Where("menu_id = ?", menuID).
		Exec(ctx)
	if err != nil {
		slog.Info("UnSplash DB Image Update", "error", err)
	}

}

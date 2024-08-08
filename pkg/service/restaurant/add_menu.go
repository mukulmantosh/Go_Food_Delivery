//go:build !test

package restaurant

import (
	"Go_Food_Delivery/pkg/database/models/restaurant"
	"Go_Food_Delivery/pkg/service/restaurant/unsplash"
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
)

var ImageUpdateLock *sync.Mutex = &sync.Mutex{}

func (restSrv *RestaurantService) AddMenu(ctx context.Context, menu *restaurant.MenuItem) (bool, error, int64, string) {
	_, err := restSrv.db.NewInsert().Model(menu).Exec(ctx)
	if err != nil {
		return false, err, 0, ""
	}

	menuImageURL := unsplash.GetUnSplashImageURL(menu.Name)
	imageFileName := fmt.Sprintf("menu_item_%d.jpg", menu.MenuID)
	imageFileLocalPath := fmt.Sprintf("uploads/%s", imageFileName)
	imageFilePath := filepath.Join(os.Getenv("LOCAL_STORAGE_PATH"), imageFileName)
	err = unsplash.DownloadImageToDisk(menuImageURL, imageFilePath)
	if err != nil {
		slog.Info("UnSplash Failed to Download Image", "error", err)
		return false, err, 0, ""
	}

	return true, nil, menu.MenuID, imageFileLocalPath
}

func (restSrv *RestaurantService) UpdateMenuPhoto(ctx context.Context, menuID int64, imageURL string) {
	go func() {
		ImageUpdateLock.Lock()
		defer ImageUpdateLock.Unlock()

		_, err := restSrv.db.NewUpdate().Table("menu_item").
			Set("photo = ?", imageURL).
			Where("menu_id = ?", menuID).
			Exec(ctx)
		if err != nil {
			slog.Info("UnSplash DB Image Update", "error", err)
		}
	}()
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"io"
	"net/http"

	"github.com/joho/godotenv"
)

type AlbumListStruct []struct {
	AlbumName             string    `json:"albumName"`
	Description           string    `json:"description"`
	AlbumThumbnailAssetID string    `json:"albumThumbnailAssetId"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	ID                    string    `json:"id"`
	OwnerID               string    `json:"ownerId"`
	Owner                 struct {
		ID               string    `json:"id"`
		Email            string    `json:"email"`
		Name             string    `json:"name"`
		ProfileImagePath string    `json:"profileImagePath"`
		AvatarColor      string    `json:"avatarColor"`
		ProfileChangedAt time.Time `json:"profileChangedAt"`
	} `json:"owner"`
	AlbumUsers                 []any     `json:"albumUsers"`
	Shared                     bool      `json:"shared"`
	HasSharedLink              bool      `json:"hasSharedLink"`
	StartDate                  time.Time `json:"startDate"`
	EndDate                    time.Time `json:"endDate"`
	Assets                     []any     `json:"assets"`
	AssetCount                 int       `json:"assetCount"`
	IsActivityEnabled          bool      `json:"isActivityEnabled"`
	Order                      string    `json:"order"`
	LastModifiedAssetTimestamp time.Time `json:"lastModifiedAssetTimestamp"`
}

type AlbumInfo struct {
	AlbumName             string    `json:"albumName"`
	Description           string    `json:"description"`
	AlbumThumbnailAssetID string    `json:"albumThumbnailAssetId"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	ID                    string    `json:"id"`
	OwnerID               string    `json:"ownerId"`
	Owner                 struct {
		ID               string    `json:"id"`
		Email            string    `json:"email"`
		Name             string    `json:"name"`
		ProfileImagePath string    `json:"profileImagePath"`
		AvatarColor      string    `json:"avatarColor"`
		ProfileChangedAt time.Time `json:"profileChangedAt"`
	} `json:"owner"`
	AlbumUsers    []any     `json:"albumUsers"`
	Shared        bool      `json:"shared"`
	HasSharedLink bool      `json:"hasSharedLink"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	Assets        []struct {
		ID               string    `json:"id"`
		DeviceAssetID    string    `json:"deviceAssetId"`
		OwnerID          string    `json:"ownerId"`
		DeviceID         string    `json:"deviceId"`
		LibraryID        any       `json:"libraryId"`
		Type             string    `json:"type"`
		OriginalPath     string    `json:"originalPath"`
		OriginalFileName string    `json:"originalFileName"`
		OriginalMimeType string    `json:"originalMimeType"`
		Thumbhash        string    `json:"thumbhash"`
		FileCreatedAt    time.Time `json:"fileCreatedAt"`
		FileModifiedAt   time.Time `json:"fileModifiedAt"`
		LocalDateTime    time.Time `json:"localDateTime"`
		UpdatedAt        time.Time `json:"updatedAt"`
		IsFavorite       bool      `json:"isFavorite"`
		IsArchived       bool      `json:"isArchived"`
		IsTrashed        bool      `json:"isTrashed"`
		Duration         string    `json:"duration"`
		ExifInfo         struct {
			Make             string    `json:"make"`
			Model            string    `json:"model"`
			ExifImageWidth   int       `json:"exifImageWidth"`
			ExifImageHeight  int       `json:"exifImageHeight"`
			FileSizeInByte   int       `json:"fileSizeInByte"`
			Orientation      string    `json:"orientation"`
			DateTimeOriginal time.Time `json:"dateTimeOriginal"`
			ModifyDate       time.Time `json:"modifyDate"`
			TimeZone         any       `json:"timeZone"`
			LensModel        string    `json:"lensModel"`
			FNumber          float32   `json:"fNumber"`
			FocalLength      int       `json:"focalLength"`
			Iso              int       `json:"iso"`
			ExposureTime     string    `json:"exposureTime"`
			Latitude         any       `json:"latitude"`
			Longitude        any       `json:"longitude"`
			City             any       `json:"city"`
			State            any       `json:"state"`
			Country          any       `json:"country"`
			Description      string    `json:"description"`
			ProjectionType   any       `json:"projectionType"`
			Rating           any       `json:"rating"`
		} `json:"exifInfo"`
		LivePhotoVideoID any    `json:"livePhotoVideoId"`
		People           []any  `json:"people"`
		Checksum         string `json:"checksum"`
		IsOffline        bool   `json:"isOffline"`
		HasMetadata      bool   `json:"hasMetadata"`
		DuplicateID      any    `json:"duplicateId"`
		Resized          bool   `json:"resized"`
	} `json:"assets"`
	AssetCount                 int       `json:"assetCount"`
	IsActivityEnabled          bool      `json:"isActivityEnabled"`
	Order                      string    `json:"order"`
	LastModifiedAssetTimestamp time.Time `json:"lastModifiedAssetTimestamp"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	albums := getAlbumList()

	var selected int
	fmt.Scanln(&selected)

	albumInfo := getAlbumInfo(albums[selected].ID)
	filterByFavourite(albumInfo)

}

func getAlbumList() AlbumListStruct {
	url := os.Getenv("ImmichURL") + "/api/albums"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", os.Getenv("ImmichKey"))

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(body))

	var result AlbumListStruct
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		log.Fatalf("Can not unmarshal JSON")
	}
	//fmt.Println(PrettyPrint(result))

	for i, a := range result {
		fmt.Printf("[%d]  %s\n", i, a.AlbumName)
	}
	return result
}

func PrettyPrint(i interface{}) string {

	s, _ := json.MarshalIndent(i, "", "\t")

	return string(s)

}

func getAlbumInfo(id string) AlbumInfo {
	url := os.Getenv("ImmichURL") + "/api/albums/" + id
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", os.Getenv("ImmichKey"))

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result AlbumInfo
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		log.Fatalf("Can not unmarshal JSON. Error message: %s", err)
	}

	if result.ID != id {
		log.Fatal("Recived ID \"" + result.ID + "\" doesn't match excpected id \"" + id + "\"")
	}
	//fmt.Println(PrettyPrint(result))
	return result
}

func filterByFavourite(album AlbumInfo) []string {
	assets := album.Assets
	var favouriteIDs []string
	for _, a := range assets {
		if a.IsFavorite {
			favouriteIDs = append(favouriteIDs, a.ID)
		}
	}
	//fmt.Println(favouriteIDs)
	return favouriteIDs
}

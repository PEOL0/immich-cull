package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

// Global mutex for duplicate file prompts
var duplicatePromptMutex sync.Mutex

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

type AssetInfo struct {
	ID            string `json:"id"`
	DeviceAssetID string `json:"deviceAssetId"`
	OwnerID       string `json:"ownerId"`
	Owner         struct {
		ID               string    `json:"id"`
		Email            string    `json:"email"`
		Name             string    `json:"name"`
		ProfileImagePath string    `json:"profileImagePath"`
		AvatarColor      string    `json:"avatarColor"`
		ProfileChangedAt time.Time `json:"profileChangedAt"`
	} `json:"owner"`
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
	Visibility       string    `json:"visibility"`
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
		TimeZone         string    `json:"timeZone"`
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
		Rating           int       `json:"rating"`
	} `json:"exifInfo"`
	LivePhotoVideoID any   `json:"livePhotoVideoId"`
	Tags             []any `json:"tags"`
	People           []struct {
		ID            string    `json:"id"`
		Name          string    `json:"name"`
		BirthDate     any       `json:"birthDate"`
		ThumbnailPath string    `json:"thumbnailPath"`
		IsHidden      bool      `json:"isHidden"`
		IsFavorite    bool      `json:"isFavorite"`
		UpdatedAt     time.Time `json:"updatedAt"`
		Faces         []struct {
			ID            string `json:"id"`
			ImageHeight   int    `json:"imageHeight"`
			ImageWidth    int    `json:"imageWidth"`
			BoundingBoxX1 int    `json:"boundingBoxX1"`
			BoundingBoxX2 int    `json:"boundingBoxX2"`
			BoundingBoxY1 int    `json:"boundingBoxY1"`
			BoundingBoxY2 int    `json:"boundingBoxY2"`
			SourceType    string `json:"sourceType"`
		} `json:"faces"`
	} `json:"people"`
	UnassignedFaces []any  `json:"unassignedFaces"`
	Checksum        string `json:"checksum"`
	Stack           any    `json:"stack"`
	IsOffline       bool   `json:"isOffline"`
	HasMetadata     bool   `json:"hasMetadata"`
	DuplicateID     any    `json:"duplicateId"`
	Resized         bool   `json:"resized"`
}

func main() {
	var directory string
	var err error

	flag.StringVar(&directory, "D", "", "Path to directory")
	flag.Parse()

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var wd string
	wd, err = os.Getwd()
	if err != nil {
		log.Fatal("Could not get working directory")
	}

	fmt.Printf("Directory to use:\n")
	fmt.Print(directory)

	fmt.Scanln(&directory)

	if directory != wd {
		createFolder(directory)
	}

	albums := getAlbumList()

	var selected int
	fmt.Scanln(&selected)

	albumInfo := getAlbumInfo(albums[selected].ID)
	favourites := getFavouritesInAlbum(albumInfo)

	fmt.Printf("Found %d favorites in album '%s'\n", len(favourites), albumInfo.AlbumName)
	downloadFavouriteAssets(directory, favourites)
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

func getFavouritesInAlbum(album AlbumInfo) []string {
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

func getAssetInfo(id string) (AssetInfo, error) {
	url := os.Getenv("ImmichURL") + "/api/assets/" + id
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return AssetInfo{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", os.Getenv("ImmichKey"))

	res, err := client.Do(req)
	if err != nil {
		return AssetInfo{}, fmt.Errorf("failed to execute request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return AssetInfo{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var result AssetInfo
	if err := json.Unmarshal(body, &result); err != nil {
		return AssetInfo{}, fmt.Errorf("cannot unmarshal JSON: %w", err)
	}

	if result.ID != id {
		return AssetInfo{}, fmt.Errorf("received ID %q doesn't match expected ID %q", result.ID, id)
	}

	return result, nil
}

func createFolder(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal("failed to create folder: %w", err)
	}
	fmt.Printf("Created folder: %s\n", path)
}

func downloadFavouriteAssets(path string, favouriteIDs []string) error {
	// Reduce concurrency to avoid rate limiting
	const maxConcurrent = 3
	sem := make(chan struct{}, maxConcurrent)

	// Use waitgroup to track goroutines
	var wg sync.WaitGroup
	wg.Add(len(favouriteIDs))

	// Track any errors and successful downloads
	errCh := make(chan error, len(favouriteIDs))
	successCount := 0
	var mu sync.Mutex // For thread-safe counter access

	fmt.Printf("Starting download of %d assets...\n", len(favouriteIDs))

	for i, id := range favouriteIDs {
		go func(idx int, assetID string) {
			defer wg.Done()

			// Add a small delay between starting goroutines to avoid overwhelming the server
			time.Sleep(time.Duration(idx*100) * time.Millisecond)

			// Acquire semaphore slot
			sem <- struct{}{}
			defer func() { <-sem }()

			// Make multiple attempts for each download
			var downloadErr error
			for attempt := 1; attempt <= 3; attempt++ {
				downloadErr = downloadAsset(assetID, path)
				if downloadErr == nil {
					mu.Lock()
					successCount++
					fmt.Printf("Progress: %d/%d assets downloaded\n", successCount, len(favouriteIDs))
					mu.Unlock()
					return
				}

				// Log retry attempt
				fmt.Printf("Retry %d for asset %s: %v\n", attempt, assetID, downloadErr)
				time.Sleep(2 * time.Second) // Wait before retry
			}

			errCh <- fmt.Errorf("failed to download asset %s after 3 attempts: %w", assetID, downloadErr)
		}(i, id)
	}

	wg.Wait()
	close(errCh)

	// Collect errors
	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	fmt.Printf("Download complete: %d successful, %d failed\n", successCount, len(errs))

	if len(errs) > 0 {
		fmt.Printf("Failed downloads: %v\n", errs)
		return fmt.Errorf("failed to download %d assets", len(errs))
	}

	return nil
}

func handleDuplicateFile(filename, path string) (string, bool) {
	// Ensure only one prompt is shown at a time
	duplicatePromptMutex.Lock()
	defer duplicatePromptMutex.Unlock()

	fmt.Printf("\nDuplicate file found: %s\n", filename)
	fmt.Println("Choose an option:")
	fmt.Println("1. Keep both files (add date prefix to new file)")
	fmt.Println("2. Overwrite the existing file")
	fmt.Println("3. (Default) Keep the existing file (skip download)")
	fmt.Print("Enter your choice (1-3): ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		// Keep both - add timestamp prefix
		timestamp := time.Now().Format("20060102-150405")
		newFilePath := fmt.Sprintf("%s/%s-%s", path, timestamp, filename)
		return newFilePath, true
	case 2:
		// Overwrite existing
		return fmt.Sprintf("%s/%s", path, filename), true
	case 3:
		// Skip download
		return "", false
	default:
		fmt.Println("Invalid choice. Defaulting to skip download.")
		return "", false
	}
}

func downloadAsset(id string, path string) error {
	// Get asset info to retrieve the original filename
	assetInfo, err := getAssetInfo(id)
	if err != nil {
		return fmt.Errorf("failed to get asset info for %s: %w", id, err)
	}

	filename := assetInfo.OriginalFileName

	// Handle potential duplicate filenames
	filePath := fmt.Sprintf("%s/%s", path, filename)
	if _, err := os.Stat(filePath); err == nil {
		// File exists - prompt user for choice
		newFilePath, shouldDownload := handleDuplicateFile(filename, path)
		if !shouldDownload {
			fmt.Printf("Skipped download of %s\n", filename)
			return nil
		}
		filePath = newFilePath
	}

	url := os.Getenv("ImmichURL") + "/api/assets/" + id + "/original"
	method := "GET"

	// Create a client with timeout
	client := &http.Client{
		Timeout: 60 * time.Second, // Add a 60-second timeout
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Accept", "application/octet-stream")
	req.Header.Add("x-api-key", os.Getenv("ImmichKey"))

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}

	// Check for non-200 status codes
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return fmt.Errorf("server returned status %d for asset %s", res.StatusCode, id)
	}

	defer res.Body.Close()

	// Create the output file
	outputFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", filePath, err)
	}
	defer outputFile.Close()

	// Write the response body directly to the file
	bytesWritten, err := io.Copy(outputFile, res.Body)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %w", filePath, err)
	}

	fmt.Printf("Downloaded %s (%d bytes)\n", filename, bytesWritten)
	return nil
}

package apis

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mkvone/mkv-backend/cmd/config"
)

// func updateSnapshotInfo(chain *[]config.ChainConfig, errChan chan<- error) {
// 	for i := range *chain {
// 		go func(chain *config.ChainConfig, validator *config.Validator) {
// 			if !chain.Snapshot.Enable {
// 				return
// 			}

// 			res, err := http.Get(chain.Snapshot.SnapshotURL)
// 			if err != nil {
// 				  l(err)
// 				return
// 			}
// 			defer res.Body.Close()

// 			if res.StatusCode != 200 {
// 				errChan <- fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
// 				return
// 			}

// 			doc, err := goquery.NewDocumentFromReader(res.Body)
// 			if err != nil {
// 				  l(err)
// 				return
// 			}

// 			// Clear existing files
// 			chain.Snapshot.Files = nil

// 			// 특정 파일과 정보 파싱
// 			doc.Find("tr.file").Each(func(i int, s *goquery.Selection) {
// 				name := strings.TrimSpace(s.Find("td:nth-child(2) .name").Text())
// 				link := chain.Snapshot.SnapshotURL + "/" + name
// 				size := strings.TrimSpace(s.Find("td.size .sizebar-text").Text())
// 				modifiedTime := strings.TrimSpace(s.Find("td.timestamp time").Text())
// 				modifiedTime, _ = timeAgo(modifiedTime)

// 				parts := strings.Split(name, "_")
// 				var height string
// 				if len(parts) > 1 {
// 					// Assume height is the segment before `.tar`
// 					heightPart := parts[len(parts)-1]
// 					height = strings.TrimSuffix(heightPart, ".tar.lz4")
// 					// If the height contains the date, remove it
// 					if strings.Contains(height, "-") {
// 						height = parts[len(parts)-2]
// 					}
// 				}

// 				chain.Snapshot.Files = append(chain.Snapshot.Files, struct {
// 					Name   string
// 					URL    string
// 					Size   string
// 					Date   string
// 					Height string
// 				}{Name: name, URL: link, Size: size, Date: modifiedTime, Height: height})
// 			})
// 		}(&(*chain)[i], &(*chain)[i].Validator)
// 	}

// }
func updateSnapshotInfo(chain *[]config.ChainConfig) {
	for i := range *chain {
		go func(chain *config.ChainConfig, validator *config.Validator) {
			if !chain.Snapshot.Enable {
				return
			}

			res, err := http.Get(chain.Snapshot.SnapshotURL)
			if err != nil {
				l(err)
				return
			}
			defer res.Body.Close()

			if res.StatusCode != 200 {
				l("status code error: %d %s", res.StatusCode, res.Status)
				return
			}

			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				l(err)
				return
			}

			// Clear existing files
			chain.Snapshot.Files = nil

			// 특정 파일과 정보 파싱
			doc.Find("tr.file").Each(func(i int, s *goquery.Selection) {
				name := strings.TrimSpace(s.Find("td:nth-child(2) .name").Text())
				link := chain.Snapshot.SnapshotURL + "/" + name
				size := strings.TrimSpace(s.Find("td.size .sizebar-text").Text())
				modifiedTime := strings.TrimSpace(s.Find("td.timestamp time").Text())
				modifiedTime, _ = timeAgo(modifiedTime)

				parts := strings.Split(name, "_")
				var height string
				if len(parts) > 1 {
					// Assume height is the segment before `.tar`
					heightPart := parts[len(parts)-1]
					height = strings.TrimSuffix(heightPart, ".tar.lz4")
					// If the height contains the date, remove it
					if strings.Contains(height, "-") {
						height = parts[len(parts)-2]
					}
				}

				chain.Snapshot.Files = append(chain.Snapshot.Files, config.File{
					Name:   name,
					URL:    link,
					Size:   size,
					Date:   modifiedTime,
					Height: height,
				})

			})
		}(&(*chain)[i], &(*chain)[i].Validator)
	}

}

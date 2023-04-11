package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"
)

const path = "songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	Album     string `json:"album"`
	PlayCount int64  `json:"play_count"`
}

// makePlaylist makes the merged sorted list of songs
func makePlaylist(albums [][]Song) (result []Song) {
	switch len(albums) {
	case 0:
		result = []Song{}
	case 1:
		onlyAlbum := albums[0]
		result = make([]Song, len(onlyAlbum))
		copy(result, onlyAlbum)
	default:
		songs := map[Song]int64{}
		for _, album := range albums {
			for _, song := range album {
				playCount := song.PlayCount
				song.PlayCount = 0 // because we use the song as key
				songs[song] += playCount
			}
		}
		result = make([]Song, 0, len(songs))
		for song, count := range songs {
			song.PlayCount = count
			result = append(result, song)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].PlayCount > result[j].PlayCount
	})
	return result
}

func main() {
	albums := importData()
	printTable(makePlaylist(albums))
}

// printTable prints merged playlist as a table
func printTable(songs []Song) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "####\tSong\tAlbum\tPlay count")
	for i, s := range songs {
		fmt.Fprintf(w, "[%d]:\t%s\t%s\t%d\n", i+1, s.Name, s.Album, s.PlayCount)
	}
	w.Flush()

}

// importData reads the input data from file and creates the friends map
func importData() [][]Song {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data [][]Song
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

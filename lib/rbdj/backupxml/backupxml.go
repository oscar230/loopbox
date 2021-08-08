package backupxml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"loopbox/utils/logger"
	"os"
)

type Backupxml struct {
	XMLName xml.Name `xml:"DJ_PLAYLISTS"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"Version,attr"`
	PRODUCT struct {
		Text    string `xml:",chardata"`
		Name    string `xml:"Name,attr"`
		Version string `xml:"Version,attr"`
		Company string `xml:"Company,attr"`
	} `xml:"PRODUCT"`
	COLLECTION struct {
		Text    string `xml:",chardata"`
		Entries string `xml:"Entries,attr"`
		TRACK   []struct {
			Text        string `xml:",chardata"`
			TrackID     string `xml:"TrackID,attr"`
			Name        string `xml:"Name,attr"`
			Artist      string `xml:"Artist,attr"`
			Composer    string `xml:"Composer,attr"`
			Album       string `xml:"Album,attr"`
			Grouping    string `xml:"Grouping,attr"`
			Genre       string `xml:"Genre,attr"`
			Kind        string `xml:"Kind,attr"`
			Size        string `xml:"Size,attr"`
			TotalTime   string `xml:"TotalTime,attr"`
			DiscNumber  string `xml:"DiscNumber,attr"`
			TrackNumber string `xml:"TrackNumber,attr"`
			Year        string `xml:"Year,attr"`
			AverageBpm  string `xml:"AverageBpm,attr"`
			DateAdded   string `xml:"DateAdded,attr"`
			BitRate     string `xml:"BitRate,attr"`
			SampleRate  string `xml:"SampleRate,attr"`
			Comments    string `xml:"Comments,attr"`
			PlayCount   string `xml:"PlayCount,attr"`
			Rating      string `xml:"Rating,attr"`
			Location    string `xml:"Location,attr"`
			Remixer     string `xml:"Remixer,attr"`
			Tonality    string `xml:"Tonality,attr"`
			Label       string `xml:"Label,attr"`
			Mix         string `xml:"Mix,attr"`
			TEMPO       []struct {
				Text    string `xml:",chardata"`
				Inizio  string `xml:"Inizio,attr"`
				Bpm     string `xml:"Bpm,attr"`
				Metro   string `xml:"Metro,attr"`
				Battito string `xml:"Battito,attr"`
			} `xml:"TEMPO"`
			POSITIONMARK []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"Name,attr"`
				Type  string `xml:"Type,attr"`
				Start string `xml:"Start,attr"`
				Num   string `xml:"Num,attr"`
				Red   string `xml:"Red,attr"`
				Green string `xml:"Green,attr"`
				Blue  string `xml:"Blue,attr"`
				End   string `xml:"End,attr"`
			} `xml:"POSITION_MARK"`
		} `xml:"TRACK"`
	} `xml:"COLLECTION"`
	PLAYLISTS struct {
		Text string `xml:",chardata"`
		NODE struct {
			Text  string `xml:",chardata"`
			Type  string `xml:"Type,attr"`
			Name  string `xml:"Name,attr"`
			Count string `xml:"Count,attr"`
			NODE  []struct {
				Text    string `xml:",chardata"`
				Name    string `xml:"Name,attr"`
				Type    string `xml:"Type,attr"`
				Count   string `xml:"Count,attr"`
				KeyType string `xml:"KeyType,attr"`
				Entries string `xml:"Entries,attr"`
				NODE    []struct {
					Text    string `xml:",chardata"`
					Name    string `xml:"Name,attr"`
					Type    string `xml:"Type,attr"`
					Count   string `xml:"Count,attr"`
					KeyType string `xml:"KeyType,attr"`
					Entries string `xml:"Entries,attr"`
					NODE    []struct {
						Text    string `xml:",chardata"`
						Name    string `xml:"Name,attr"`
						Type    string `xml:"Type,attr"`
						KeyType string `xml:"KeyType,attr"`
						Entries string `xml:"Entries,attr"`
						Count   string `xml:"Count,attr"`
						TRACK   []struct {
							Text string `xml:",chardata"`
							Key  string `xml:"Key,attr"`
						} `xml:"TRACK"`
						NODE []struct {
							Text    string `xml:",chardata"`
							Name    string `xml:"Name,attr"`
							Type    string `xml:"Type,attr"`
							KeyType string `xml:"KeyType,attr"`
							Entries string `xml:"Entries,attr"`
							TRACK   []struct {
								Text string `xml:",chardata"`
								Key  string `xml:"Key,attr"`
							} `xml:"TRACK"`
						} `xml:"NODE"`
					} `xml:"NODE"`
					TRACK []struct {
						Text string `xml:",chardata"`
						Key  string `xml:"Key,attr"`
					} `xml:"TRACK"`
				} `xml:"NODE"`
				TRACK []struct {
					Text string `xml:",chardata"`
					Key  string `xml:"Key,attr"`
				} `xml:"TRACK"`
			} `xml:"NODE"`
		} `xml:"NODE"`
	} `xml:"PLAYLISTS"`
}

func (b *Backupxml) Parse(filepath string) error {
	if f, fErr := os.Open(filepath); fErr != nil {
		logger.Error("[backupxml] Open file. " + fErr.Error())
		return fErr
	} else {
		defer f.Close()
		logger.Info("[backupxml] Opened file " + f.Name())
		if bytes, iErr := ioutil.ReadAll(f); iErr != nil {
			logger.Error("[backupxml] Read all as bytes. " + iErr.Error())
			return iErr
		} else {
			if xErr := xml.Unmarshal(bytes, b); xErr != nil {
				logger.Error("[backupxml] XML Unmarshal. " + xErr.Error())
				return xErr
			} else {
				logger.Info("[backupxml] Successfully read and unmarshalled XML file " + f.Name() + " with " + fmt.Sprint(b.TrackCount()))
				return nil
			}
		}
	}
}

func (b *Backupxml) TrackCount() int {
	return len(b.COLLECTION.TRACK)
}

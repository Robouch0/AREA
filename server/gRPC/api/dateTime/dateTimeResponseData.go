//
// EPITECH PROJECT, 2024
// AREA
// File description:
// dateTimeResponseData
//

package dateTime

type AimylogicDateTime struct {
	Timezone  string `json:"timezone"`
	Formatted string `json:"formatted"`
	Timestamp int    `json:"timestamp"`
	WeekDay   int    `json:"week_day"`
	Day       int    `json:"day"`
	Month     int    `json:"month"`
	Year      int    `json:"year"`
	Hour      int    `json:"hour"`
	Minute    int    `json:"minute"`
}

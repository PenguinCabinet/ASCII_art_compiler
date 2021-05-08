package main

type setting_file_t struct {
	Font_size  int `json:"font_size"`
	Top_offset int `json:"top_offset"`
}

func new_setting_file_t() setting_file_t {
	A := setting_file_t{Font_size: 90, Top_offset: 200}

	return A
}

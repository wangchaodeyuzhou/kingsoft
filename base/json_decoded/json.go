package json_decoded

type (
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedURL"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheURL"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

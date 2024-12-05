package google_recaptcha

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

func ValidateRecaptcha(ctx context.Context, recaptchaToken string) (bool, error){
	const secretKey = "6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe"

	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", 
		url.Values{
			"secret": {secretKey},
			"response": {recaptchaToken},
		},
	)
	if err != nil {
        return false, err
    }
    defer resp.Body.Close()

	var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return false, err
    }

    success, ok := result["success"].(bool)
    return success && ok, nil
}
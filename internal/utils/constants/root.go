package constants

import "time"

const ACCESS_TOKEN_DURATION = 10 * time.Second
const REFRESH_TOKEN_DURATION = 30 * 24 * time.Hour // 30 days
const COOKIE_DURATION = 2629800 // 1 month
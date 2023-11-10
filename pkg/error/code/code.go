package code

import "net/http"

type ErrorCode string

const (
	// ErrorCodeOutdatedApp アプリのバージョン更新が必要
	ErrorCodeOutdatedApp ErrorCode = "OutdatedApp"
	// ErrorCodeOutdatedMasterData マスタデータの更新が必要
	ErrorCodeOutdatedMasterData = "OutdatedMasterData"
	// ErrorCodeDateChanged 日付が変わったので再ログインが必要
	ErrorCodeDateChanged = "DateChanged"
	// ErrorCodeInMaintenance メンテナンス中
	ErrorCodeInMaintenance = "InMaintenance"
	// ErrorCodeInvalidArgument パラメータの指定が不正
	ErrorCodeInvalidArgument = "InvalidArgument"
	// ErrorCodeInternalServerError 内部サーバエラー
	ErrorCodeInternalServerError = "InternalServerError"
	// ErrorCodeUnauthenticated 認証エラー
	ErrorCodeUnauthenticated = "Unauthenticated"
)

type ActionType string

const (
	// Retry 通信をリトライ
	Retry ActionType = "retry"
	// Title タイトル画面へ遷移
	Title ActionType = "title"
	// Store App Store / Google Play ストアへ遷移
	Store ActionType = "store"
	// Home ホーム画面へ遷移
	Home ActionType = "home"
	// Continue 処理を継続
	Continue ActionType = "continue"
)

type ErrorPattern struct {
	ErrorCode      ErrorCode
	HTTPStatusCode int
	ActionType     ActionType
}

var (
	OutdatedApp = ErrorPattern{
		ErrorCode:      ErrorCodeOutdatedApp,
		HTTPStatusCode: http.StatusBadRequest,
		ActionType:     Store,
	}
	OutdatedMasterData = ErrorPattern{
		ErrorCode:      ErrorCodeOutdatedMasterData,
		HTTPStatusCode: http.StatusBadRequest,
		ActionType:     Title,
	}
	DateChanged = ErrorPattern{
		ErrorCode:      ErrorCodeDateChanged,
		HTTPStatusCode: http.StatusBadRequest,
		ActionType:     Title,
	}
	InMaintenance = ErrorPattern{
		ErrorCode:      ErrorCodeInMaintenance,
		HTTPStatusCode: http.StatusServiceUnavailable,
		ActionType:     Title,
	}
	InvalidArgument = ErrorPattern{
		ErrorCode:      ErrorCodeInvalidArgument,
		HTTPStatusCode: http.StatusBadRequest,
		ActionType:     Title,
	}
	InternalServerError = ErrorPattern{
		ErrorCode:      ErrorCodeInternalServerError,
		HTTPStatusCode: http.StatusInternalServerError,
		ActionType:     Title,
	}
	Unauthenticated = ErrorPattern{
		ErrorCode:      ErrorCodeUnauthenticated,
		HTTPStatusCode: http.StatusUnauthorized,
		ActionType:     Title,
	}
)

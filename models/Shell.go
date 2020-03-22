package models

type Shell interface {
	EncodeHistoryItem(item HistoryItem) string
	DecodeHistoryItem(encodedHistoryItem string) (HistoryItem, error)
	GetShellHistoryFromBytes(historyBytes []byte) (ShellHistory, error)
}

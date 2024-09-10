package lsp

type TextDocumentDidSaveNotificiation struct {
	Notification
	Params DidSaveTextDocumentParams `json:"params"`
}

type DidSaveTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

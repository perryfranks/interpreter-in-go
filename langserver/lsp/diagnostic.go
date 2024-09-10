package lsp

type PublishDiagnosticsNotification struct {
	Notification
	Params PublishDiagParams `json:"Params"`
}

type PublishDiagParams struct {
	TextDocumentIdentifier
	// VersionTextDocumentIdentifier
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type Diagnostic struct {
	Range Range  `json:"Range"`
	Code  string `json:"Code"` // error code
	// CodeDescription DiagnosticCodeDescription `json:"codeDescription"`
	Message string `json:"message"`

	// and more but see how these are displayed first
}

// is just a uri. not sure how you interface with this
// type DiagnosticCodeDescription struct {
//
// }

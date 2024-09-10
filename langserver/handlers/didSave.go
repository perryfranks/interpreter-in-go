package handlers

import (
	"io"
	"log"

	"github.com/perryfranks/monkey-lang/entries"
	"github.com/perryfranks/monkey-lang/langserver/analysis"
	"github.com/perryfranks/monkey-lang/langserver/lsp"
)

func HandleDidSave(logger *log.Logger, writer io.Writer, state analysis.State, request lsp.DidChangeTextDocumentNotification) (int, error) {
	//  get the page from the statehanders
	prog := state.Documents[request.Params.TextDocument.URI]
	logger.Printf("Program:", prog)

	// call eval and check for errors
	progErrors, err := entries.EvalErrors(prog)
	if err != nil {
		return 0, err
	}

	// log the amount of errors
	logger.Printf("Eval ran. Num errors: %d. errors: %v", len(progErrors), progErrors)

	// TODO: spin off the textDocument/publishDiagnostics notification to the client

	return len(progErrors), nil

}

package scripting

import (
	"log"
	"net/http"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/wendehals/bricks/api"
	"github.com/wendehals/bricks/scripting/parser"
)

type BricksScript struct {
	scriptPath  string
	credentials *api.Credentials

	client    *http.Client
	bricksAPI *api.BricksAPI
	usersAPI  *api.UsersAPI
	verbose   bool
}

func NewBricksScript(credentials *api.Credentials, scriptPath string, verbose bool) *BricksScript {
	b := &BricksScript{}
	b.scriptPath = scriptPath
	b.credentials = credentials
	b.verbose = verbose

	return b
}

func (b *BricksScript) Execute() {
	input, err := antlr.NewFileStream(b.scriptPath)
	if err != nil {
		log.Fatalf("opening file '%s' failed: %s", b.scriptPath, err.Error())
	}

	lexer := parser.NewBricksLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	parser := parser.NewBricksParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true

	treeListener := newBricksTreeListener(b.getUsersAPI(), b.getBricksAPI())
	antlr.ParseTreeWalkerDefault.Walk(treeListener, parser.Bricks())
}

func (b *BricksScript) getClient() *http.Client {
	if b.client == nil {
		b.client = &http.Client{
			Timeout: time.Second * 5,
		}
	}
	return b.client
}

func (b *BricksScript) getBricksAPI() *api.BricksAPI {
	if b.bricksAPI == nil {
		b.bricksAPI = api.NewBricksAPI(b.getClient(), b.credentials.APIKey, b.verbose)
	}
	return b.bricksAPI
}

func (b *BricksScript) getUsersAPI() *api.UsersAPI {
	if b.usersAPI == nil {
		b.usersAPI = api.NewUsersAPI(b.getClient(), b.credentials, b.verbose)
	}
	return b.usersAPI
}

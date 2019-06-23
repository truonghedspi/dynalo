package policy

import (
	"context"
	"dynalo"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/loader"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type PolicyEval struct {
	store    storage.Store
	compiler *ast.Compiler
}

func newPolicyEval() *PolicyEval {
	pe := PolicyEval{}
	return &pe
}

func (pe *PolicyEval) withStore(s storage.Store) *PolicyEval {
	pe.store = s
	return pe
}

func (pe *PolicyEval) withCompiler(c *ast.Compiler) *PolicyEval {
	pe.compiler = c
	return pe
}

//Eval evaluate dynamic logic
func (pe *PolicyEval) Eval(ctx context.Context, logic dynalo.DynamicLogic) (interface{}, error) {
	rego := rego.New(
		rego.Store(pe.store),
		rego.Compiler(pe.compiler),
		rego.Query(logic.PolicyQuery),
		rego.Input(logic.Input),
	)

	rs, err := rego.Eval(ctx)
	if err != nil {
		return nil, err
	}
	if len(rs) == 0 {
		return nil, nil
	}
	return rs[0].Expressions[0].Value, nil
}

//NewPolicyEval create policy evaluation object from path
func NewPolicyEval(path string) (*PolicyEval, error) {
	loaded, err := loadFrom(path)
	if err != nil {
		return nil, err
	}

	store, err := loaded.Store()
	if err != nil {
		return nil, err
	}

	compiler, err := loaded.Compiler()
	if err != nil {
		return nil, err
	}

	pe := newPolicyEval().withStore(store).withCompiler(compiler)
	return pe, nil

}

//loadFrom return Result object loaded from specified path.
func loadFrom(path string) (*loader.Result, error) {
	files, err := loader.Paths(path, false)
	if err != nil {
		return nil, errors.Wrapf(err, "error when read policies in directory %s", path)
	}
	log.Infof("load policies in file : %s", files)

	return loader.All(files)
}

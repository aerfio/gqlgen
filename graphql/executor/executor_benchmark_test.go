package executor

import (
	"context"
	"testing"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/99designs/gqlgen/graphql"
)

type benchmarkSchema struct {
	schema *ast.Schema
}

func (s benchmarkSchema) Schema() *ast.Schema {
	return s.schema
}

func (benchmarkSchema) Complexity(context.Context, string, string, int, map[string]any) (int, bool) {
	panic("not used")
}

func (benchmarkSchema) Exec(context.Context) graphql.ResponseHandler {
	panic("not used")
}

func BenchmarkParseQuery(b *testing.B) {
	schema := gqlparser.MustLoadSchema(&ast.Source{Input: "type Query { name: String! }"})
	exec := New(benchmarkSchema{schema: schema})
	ctx := context.Background()
	stats := &graphql.Stats{}

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		_, errs := exec.parseQuery(ctx, stats, "{name}")
		if len(errs) != 0 {
			b.Fatal(errs)
		}
	}
}

// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package parser_test

import (
	"strings"
	"testing"

	. "github.com/pingcap/check"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/format"
)

func TestRPSGame(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&testRPSGameSuite{})

type testRPSGameSuite struct {
}

func assertRestore(c *C, node ast.Node, expected string) {
	var sb strings.Builder
	err := node.Restore(format.NewRestoreCtx(format.DefaultRestoreFlags, &sb))
	c.Assert(err, IsNil)
	c.Assert(sb.String(), Equals, expected)
}

func (s *testParserSuite) TestParseCreateRPSGame(c *C) {
	p := parser.New()
	n, err := p.ParseOneStmt("CREATE RPS GAME game1", "", "")
	c.Assert(err, IsNil)
	stmt, ok := n.(*ast.CreateRPSGameStmt)
	c.Assert(ok, IsTrue)
	c.Assert(stmt.Name.O, Equals, "game1")
	assertRestore(c, stmt, "CREATE RPS GAME `game1`")
}

func (s *testParserSuite) TestShowCreateRPSGame(c *C) {
	p := parser.New()
	n, err := p.ParseOneStmt("SHOW CREATE RPS GAME game1", "", "")
	c.Assert(err, IsNil)
	stmt, ok := n.(*ast.ShowStmt)
	c.Assert(ok, IsTrue)
	c.Assert(stmt.Tp, Equals, ast.ShowStmtType(ast.ShowCreateRPSGame))
	c.Assert(stmt.GameName.O, Equals, "game1")
	assertRestore(c, stmt, "SHOW CREATE RPS GAME `game1`")
}

func (s *testParserSuite) TestShowCreateRPSGameStatus(c *C) {
	p := parser.New()
	n, err := p.ParseOneStmt("SHOW RPS GAME STATUS", "", "")
	c.Assert(err, IsNil)
	stmt, ok := n.(*ast.ShowStmt)
	c.Assert(ok, IsTrue)
	c.Assert(stmt.Tp, Equals, ast.ShowStmtType(ast.ShowRPSGameStatus))
	assertRestore(c, stmt, "SHOW RPS GAME STATUS")

	n, err = p.ParseOneStmt("SHOW RPS GAME STATUS LIKE game1", "", "")
	c.Assert(err, IsNil)
	stmt, ok = n.(*ast.ShowStmt)
	c.Assert(ok, IsTrue)
	c.Assert(stmt.Tp, Equals, ast.ShowStmtType(ast.ShowRPSGameStatus))
	assertRestore(c, stmt, "SHOW RPS GAME STATUS LIKE `game1`")
}

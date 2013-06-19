package librarian

import (
  "librarian/nodes"
)

type SelectManager struct {
  Tree     *nodes.SelectStatementNode
  Context  *nodes.SelectCoreNode
  Relation *nodes.RelationNode
}

func (mgmt *SelectManager) Project(projections ...interface{}) *SelectManager {
  for _, projection := range projections {
    switch projection.(type) {
    case string:
      projection = nodes.Attribute(projection.(string), mgmt.Relation)
    default:
    }
    mgmt.Context.Projections = append(mgmt.Context.Projections, projection)
  }
  return mgmt
}

func (mgmt *SelectManager) Where(comparison interface{}) *SelectManager {
  mgmt.Context.Wheres = append(mgmt.Context.Wheres, comparison)
  return mgmt
}

func (mgmt *SelectManager) InnerJoin(expression interface{}) *SelectManager {
  mgmt.Context.Source.Right = append(mgmt.Context.Source.Right,
    nodes.InnerJoin(expression))
  return mgmt
}

func (mgmt *SelectManager) Limit(take interface{}) *SelectManager {
  mgmt.Tree.Limit = nodes.Limit(take)
  return mgmt
}

func (mgmt *SelectManager) Offset(skip interface{}) *SelectManager {
  mgmt.Tree.Offset = nodes.Offset(skip)
  return mgmt
}

func (mgmt *SelectManager) ToSql() string {
  return ""
}

func NewSelectManager(relation *nodes.RelationNode) *SelectManager {
  tree := nodes.SelectStatement()
  core := nodes.SelectCore(relation)
  tree.Cores = append(tree.Cores, core)
  return &SelectManager{tree, core, relation}
}

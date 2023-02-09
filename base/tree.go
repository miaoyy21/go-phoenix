package base

type RelationTree struct {
	values   map[string]map[string]string // Id :: Rows
	relation map[string]string            // Id :: ParentId

	children []string
}

func NewRelationTree(rows []map[string]string, children []string) *RelationTree {
	values := make(map[string]map[string]string)
	relation := make(map[string]string)

	for _, row := range rows {
		id, ok := row["id"]
		if !ok {
			continue
		}

		// Id && ParentId
		values[id] = row
		if parentId, ok := row["parent_id_"]; ok {
			relation[id] = parentId
		}
	}

	return &RelationTree{
		values:   values,
		relation: relation,

		children: children,
	}
}

func (o *RelationTree) Build() []map[string]string {
	ids := make([]string, 0, len(o.children))
	for _, child := range o.children {
		ids = append(ids, o.getParents(child)...)
	}

	res := make([]map[string]string, 0, len(o.values))
	for _, id := range NewStringSet(ids).Values() {
		res = append(res, o.values[id])
	}

	return res
}

func (o *RelationTree) getParents(id string) []string {
	parentId, ok := o.relation[id]
	if !ok {
		return []string{id}
	}

	supers := o.getParents(parentId)

	news := make([]string, 0, len(supers)+1)
	news = append(news, id)
	news = append(news, supers...)

	return news
}

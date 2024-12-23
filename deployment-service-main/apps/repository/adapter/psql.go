package adapter

type BuildQueryHelper interface {
	Where(key string) QueryBuilder
	And(key string) QueryBuilder
	Or(key string) QueryBuilder
	LEFT_JOIN(key string) QueryBuilder
	GROUP_BY(key string) QueryBuilder
	UNION_ALL(key string) QueryBuilder
	HAVING(key string) QueryBuilder
	LIMIT(key string) QueryBuilder
	OFFSET(key string) QueryBuilder
}

type QueryBuilder struct {
	V       string
	QHelper BuildQueryHelper
}

func (s *PSql) Select(key string) QueryBuilder {
	return QueryBuilder{V: "SELECT " + key, QHelper: &QueryBuilder{V: "SELECT " + key}}
}

func (s *QueryBuilder) Where(key string) QueryBuilder {
	return QueryBuilder{V: s.V + " WHERE " + key, QHelper: &QueryBuilder{V: s.V + " WHERE " + key}}
}

func (s *QueryBuilder) And(key string) QueryBuilder {
	return QueryBuilder{V: s.V + " AND " + key, QHelper: &QueryBuilder{V: s.V + " AND " + key}}

}

func (s *QueryBuilder) Or(key string) QueryBuilder {
	return QueryBuilder{V: s.V + " OR " + key, QHelper: &QueryBuilder{V: s.V + " OR " + key}}
}

func (s *QueryBuilder) LEFT_JOIN(key string) QueryBuilder {
	return QueryBuilder{V: s.V + " LEFT JOIN " + key, QHelper: &QueryBuilder{V: s.V + " LEFT JOIN " + key}}
}

func (s *QueryBuilder) GROUP_BY(key string) QueryBuilder {
	return QueryBuilder{V: s.V + " GROUP BY " + key, QHelper: &QueryBuilder{V: s.V + " GROUP BY " + key}}
}

func (s *QueryBuilder) UNION_ALL(key string) QueryBuilder {
	return QueryBuilder{V: s.V + " UNION ALL " + key, QHelper: &QueryBuilder{V: s.V + " UNION ALL " + key}}
}

func (s *QueryBuilder) HAVING(key string) QueryBuilder {
	return QueryBuilder{V: s.V + " HAVING " + key, QHelper: &QueryBuilder{V: s.V + " HAVING " + key}}
}

func (s *QueryBuilder) LIMIT(key string) QueryBuilder {
	return QueryBuilder{V: s.V + " LIMIT " + key, QHelper: &QueryBuilder{V: s.V + " LIMIT " + key}}
}

func (s *QueryBuilder) OFFSET(key string) QueryBuilder {
	return QueryBuilder{V: s.V + " OFFSET " + key, QHelper: &QueryBuilder{V: s.V + " OFFSET " + key}}
}

func (s *PSql) RawQuery(queryString string) []map[string]interface{} {
	queryString += ";"
	var output []map[string]interface{}
	_ = s.connection.Raw(queryString).Scan(&output)
	return output
}

func (s *PSql) Exec(queryString string) {
	queryString += ";"
	_ = s.connection.Exec(queryString)
}

package handle

var whiteRoutes map[string]map[string]route

func AddWhiteRoute(method string, path string, params map[string]string, values map[string]string) {
	if whiteRoutes == nil {
		whiteRoutes = make(map[string]map[string]route)
	}

	m1, ok := whiteRoutes[method]
	if !ok {
		m1 = make(map[string]route)
	}

	m2, ok := m1[path]
	if !ok {
		m2 = route{Params: params, Values: values}
	}

	m1[path] = m2
	whiteRoutes[method] = m1
}

type route struct {
	Params map[string]string // 包含匹配
	Values map[string]string // 包含匹配
}

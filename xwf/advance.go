package xwf

//
//func advance(tx *sql.Tx, ctx *handle.Context, diagramId string, key int, instanceId string, values string, start int, backwards []ExecuteBackward) (interface{}, error) {
//
//	keys := make(map[int]struct{})
//
//BREAK:
//	for _, backward := range backwards {
//		for _, route := range backward.Routes {
//			// 已经执行的节点忽略
//			if _, ok := keys[route]; ok {
//				continue
//			}
//
//			node, err := flow.NewNode(tx, ctx, diagramId, route)
//			if err != nil {
//				return nil, err
//			}
//
//			// Start
//			if _, ok := node.(flow.StartFlowable); ok {
//				logrus.Panic("unreachable")
//			}
//
//			// Execute
//			if execute, ok := node.(flow.ExecuteFlowable); ok {
//				// Start
//				if route == backward.Key {
//					if err := execute.ExecuteStart(backward.Executors); err != nil {
//						return nil, err
//					}
//
//					continue
//				}
//
//				// End
//				if route == start {
//					if err := execute.ExecuteFinished(); err != nil {
//						return nil, err
//					}
//				}
//
//				logrus.Panic("unreachable")
//			}
//
//			// End
//			if end, ok := node.(flow.EndFlowable); ok {
//				if err := end.End(); err != nil {
//					return nil, err
//				}
//
//				break BREAK
//			}
//
//			keys[route] = struct{}{}
//		}
//	}
//
//	return map[string]interface{}{"status": "success"}, nil
//}

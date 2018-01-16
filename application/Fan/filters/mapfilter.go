package filters

var MapFilter map[string]interface{}

func init() {
	MapFilter = make(map[string]interface{})
	MapFilter["/admin/login"] = struct{}{}
	MapFilter["/admin/new"] = struct{}{}
	MapFilter["/home"] = struct{}{}
}

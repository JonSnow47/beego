package filters

var MapFilter map[string]interface{}

func init() {
	MapFilter = make(map[string]interface{})
	MapFilter["/home"] = struct{}{}
	MapFilter["/admin/new"] = struct{}{}
	MapFilter["/admin/login"] = struct{}{}
	MapFilter["/user/register"] = struct{}{}
	MapFilter["/user/login"] = struct{}{}
}

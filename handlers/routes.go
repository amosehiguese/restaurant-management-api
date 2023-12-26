package handlers

var routes = []route{
	// menu
	newRoute("GET", "/menu", GetMenu),
	newRoute("POST", "/menu", CreateMenu),
	newRoute("GET", "/menu/([0-9a-fA-F]{24})", RetrieveMenu),
	newRoute("PUT", "/menu/([0-9a-fA-F]{24})", UpdateMenu),
	newRoute("DELETE", "/menu/([0-9a-fA-F]{24})", DeleteMenu),

	// dishes
}


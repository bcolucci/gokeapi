package endpoints

type Endpoint interface {
	FetchResourceList() *ResourceList
	FetchNextResourceList(current *ResourceList) *ResourceList
	FetchPreviousResourceList(current *ResourceList) *ResourceList
	GetItem(current *ResourceList, idx int) interface{}
	GetItems(current *ResourceList) []interface{}
	GetItemById(id int) interface{}
	GetItemByName(name string) interface{}
}

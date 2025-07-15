package dm

type DataModelOptions struct {
	DmRoot          string `default:"connectedhomeip/data_model/master" aliases:"dmRoot" help:"where to place the data model files" group:"Data Model:"`
	IgnoreHierarchy bool   `default:"true" help:"don't parse the cluster hierarchy when building data model" group:"Data Model:"`
}

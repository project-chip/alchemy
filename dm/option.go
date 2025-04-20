package dm

type DataModelOptions struct {
	DmRoot string `default:"connectedhomeip/data_model/master" aliases:"dmRoot" help:"where to place the data model files" group:"Data Model:"`
}

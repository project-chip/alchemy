package dm

type DataModelOptions struct {
	DmRoot string `default:"connectedhomeip/data_model/master" aliases:"dmRoot" help:"where to place the data model files" group:"Data Model:"`
	Force  bool   `default:"false" help:"write data model XML even if there were parsing errors" group:"ZAP:"  group:"Data Model:"`
}

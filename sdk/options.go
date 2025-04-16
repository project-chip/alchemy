package sdk

type SDKOptions struct {
	SdkRoot string `default:"connectedhomeip" aliases:"sdkRoot" help:"the root of your clone of project-chip/connectedhomeip" group:"SDK:"`
}

package idl

import (
	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter/spec"
)

type Endpoint struct {
	ID int
	EndpointType

	Servers []*ClusterInfo
	Clients []*ClusterInfo
}

func endpointServersHelper(spec *spec.Specification, filter ProvisionalFilter) func(Endpoint, *raymond.Options) raymond.SafeString {
	return func(endpoint Endpoint, options *raymond.Options) raymond.SafeString {
		servers := make([]ClusterRef, 0, len(endpoint.Clusters))
		for _, clusterRef := range endpoint.Clusters {
			if clusterRef.Side != "server" {
				continue
			}
			cluster, ok := spec.ClustersByID[uint64(clusterRef.Code)]
			if !ok {
				continue
			}
			if !entityShouldBeIncluded(spec, filter, cluster) {
				continue
			}
			clusterRef.Cluster = cluster
			servers = append(servers, clusterRef)
		}
		return enumerateHelper(servers, spec, options)
	}
}

func endpointClientsHelper(spec *spec.Specification, filter ProvisionalFilter) func(Endpoint, *raymond.Options) raymond.SafeString {
	return func(endpoint Endpoint, options *raymond.Options) raymond.SafeString {
		clients := make([]ClusterRef, 0, len(endpoint.Clusters))
		for _, clusterRef := range endpoint.Clusters {
			if clusterRef.Side != "client" {
				continue
			}
			cluster, ok := spec.ClustersByID[uint64(clusterRef.Code)]
			if !ok {
				continue
			}
			if !entityShouldBeIncluded(spec, filter, cluster) {
				continue
			}
			clusterRef.Cluster = cluster
			clients = append(clients, clusterRef)
		}
		return enumerateHelper(clients, spec, options)
	}
}

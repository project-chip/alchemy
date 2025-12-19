package regen

import (
	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/zap"
)

type Endpoint struct {
	ID int
	zap.EndpointType

	Servers []*ClusterInfo
	Clients []*ClusterInfo
}

func endpointServersHelper(spec *spec.Specification) func(Endpoint, *raymond.Options) raymond.SafeString {
	return func(endpoint Endpoint, options *raymond.Options) raymond.SafeString {
		servers := make([]zap.ClusterRef, 0, len(endpoint.Clusters))
		for _, clusterRef := range endpoint.Clusters {
			if clusterRef.Side != "server" {
				continue
			}
			cluster, ok := spec.ClustersByID[uint64(clusterRef.Code)]
			if !ok {
				continue
			}
			clusterRef.Cluster = cluster
			servers = append(servers, clusterRef)
		}
		return enumerateHelper(servers, spec, options)
	}
}

func endpointClientsHelper(spec *spec.Specification) func(Endpoint, *raymond.Options) raymond.SafeString {
	return func(endpoint Endpoint, options *raymond.Options) raymond.SafeString {
		clients := make([]zap.ClusterRef, 0, len(endpoint.Clusters))
		for _, clusterRef := range endpoint.Clusters {
			if clusterRef.Side != "client" {
				continue
			}
			cluster, ok := spec.ClustersByID[uint64(clusterRef.Code)]
			if !ok {
				continue
			}
			clusterRef.Cluster = cluster
			clients = append(clients, clusterRef)
		}
		return enumerateHelper(clients, spec, options)

	}
}

package testplan

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/matter"
)

var header = `= *%s Cluster Test Plan*
ifdef::env-github[]
:tip-caption: :bulb:
:note-caption: :information_source:
:important-caption: :heavy_exclamation_mark:
:caution-caption: :fire:
:warning-caption: :warning:
:imagesdir: https://github.com/chip-csg/chip-test-plans/tree/master/src/images
:toc:
:toclevels: 3
:sectnumlevels: 3
endif::[]
ifndef::env-github[]
:imagesdir: images
endif::[]
:sectanchors:
:doctype: book
:author: Matter CSG Test Plans Tiger Team
:sectnums:
:picsCode: %s
:clustername: %s

// Common AsciiDocAttributes
include::../common/cluster_common.adoc[]
include::../common/spec_ref_common.adoc[]

== PICS Definition
This section covers the {clustername} Cluster Test Plan related PICS items that are referenced in the following test cases.  Support for an item is considered as "true" for conditional statements within the test case steps.

=== Role

|===
| *Variable* | *Description*                                   | *Mandatory/Optional* | *Notes/Additional Constraints*
| {PICS_S}   | {devimp} the {clustername} cluster as a server? | O                    |
|===

`

func renderHeader(cluster *matter.Cluster, b *strings.Builder) (err error) {

	header := fmt.Sprintf(header, cluster.Name, cluster.PICS, cluster.Name)
	_, err = b.WriteString(header)
	return
}

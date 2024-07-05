package testplan

import (
	"fmt"
	"strings"
	"time"

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
:revnumber: 0.1
:revdate: %s
:author: Matter CSG Test Plans Tiger Team
:sectnums:
:picsCode: %s
:clustername: %s

ifdef::env-github[]
*Document History*
|===
|*Rev*|*Date*|*Author*|*Description*
|0.1|%s|<author>|Initial Test Plan for {picsCode}
|===

*Common Introduction*
include::intro.adoc[Introduction]
endif::[]

// Common AsciiDocAttributes
include::../common/cluster_common.adoc[]
include::../common/spec_ref_common.adoc[]

== PICS Definition
This section covers the {clustername} Cluster Test Plan related PICS items that are referenced in the following test cases.  Support for an item is considered as "true" for conditional statements within the test case steps.

=== Role

|===
| *Variable* | *Description*                                   | *Mandatory/Optional* | *Notes/Additional Constraints*
| {PICS_S}   | {devimp} the {clustername} cluster as a server? | O                    |
| {PICS_C}   | {devimp} the {clustername} cluster as a client? | O                    |
|===

`

func renderHeader(cluster *matter.Cluster, b *strings.Builder) (err error) {

	now := time.Now()
	longDate := now.Format("02-Jan-2006")
	shortDate := now.Format("2006-01-02")
	header := fmt.Sprintf(header, cluster.Name, longDate, cluster.PICS, cluster.Name, shortDate)
	_, err = b.WriteString(header)

	return
}

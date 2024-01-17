package ascii

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
	"github.com/hasty/alchemy/parse"
)

func (s *Section) toDeviceTypes(d *Doc) (entities []mattertypes.Entity, err error) {
	var deviceTypes []*matter.DeviceType
	var description string
	p := parse.FindFirst[*types.Paragraph](s.Elements)
	if p != nil {
		se := parse.FindFirst[*types.StringElement](p.Elements)
		if se != nil {
			description = strings.ReplaceAll(se.Content, "\n", " ")
		}
	}

	for _, s := range parse.Skim[*Section](s.Elements) {
		switch s.SecType {
		case matter.SectionClassification:
			deviceTypes, err = readDeviceTypeIDs(d, s)
		}
		if err != nil {
			return nil, err
		}
	}

	for _, c := range deviceTypes {
		c.Description = description

		elements := parse.Skim[*Section](s.Elements)
		for _, s := range elements {
			switch s.SecType {
			case matter.SectionClusterRequirements:
				var crs []*matter.ClusterRequirement
				crs, err = s.toClusterRequirements(d)
				if err == nil {
					c.ClusterRequirements = append(c.ClusterRequirements, crs...)
				}
			case matter.SectionElementRequirements:
				c.ElementRequirements, err = s.toElementRequirements(d)
			case matter.SectionConditions:
				c.Conditions, err = s.toConditions(d)
			case matter.SectionRevisionHistory:
				c.Revisions, err = readRevisionHistory(d, s)
			default:
			}
			if err != nil {
				return nil, fmt.Errorf("error reading section in %s: %w", d.Path, err)
			}
		}
	}
	for _, c := range deviceTypes {
		entities = append(entities, c)
	}
	return entities, nil
}

func readDeviceTypeIDs(doc *Doc, s *Section) ([]*matter.DeviceType, error) {
	rows, headerRowIndex, columnMap, _, err := parseFirstTable(doc, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading device type ID: %w", err)
	}
	var deviceTypes []*matter.DeviceType
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		c := &matter.DeviceType{}
		c.ID, err = readRowID(row, columnMap, matter.TableColumnID)
		if err != nil {
			return nil, err
		}
		c.Name, err = readRowValue(row, columnMap, matter.TableColumnDeviceName)
		if err != nil {
			return nil, err
		}
		c.Superset, err = readRowValue(row, columnMap, matter.TableColumnSuperset)
		if err != nil {
			return nil, err
		}
		c.Class, err = readRowValue(row, columnMap, matter.TableColumnClass)
		if err != nil {
			return nil, err
		}
		c.Scope, err = readRowValue(row, columnMap, matter.TableColumnScope)
		if err != nil {
			return nil, err
		}
		deviceTypes = append(deviceTypes, c)
	}

	return deviceTypes, nil
}

func (d *Doc) toBaseDeviceType() (baseDeviceType *matter.DeviceType, err error) {
	for _, e := range d.Elements {
		switch e := e.(type) {
		case *Section:
			var baseClusterRequirements, elementRequirements *Section
			parse.Traverse(e, e.Elements, func(sec *Section, parent parse.HasElements, index int) bool {
				switch sec.Name {
				case "Base Cluster Requirements":
					baseClusterRequirements = sec
				case "Element Requirements":
					elementRequirements = sec
				}
				return false
			})
			if baseClusterRequirements != nil || elementRequirements != nil {
				baseClusterRequirements.toClusterRequirements(d)
				baseDeviceType = &matter.DeviceType{}
				if baseClusterRequirements != nil {
					baseDeviceType.ClusterRequirements, err = baseClusterRequirements.toClusterRequirements(d)
					if err != nil {
						return
					}
				}
				if elementRequirements != nil {
					baseDeviceType.ElementRequirements, err = elementRequirements.toElementRequirements(d)
					if err != nil {
						return
					}
				}
				return
			}
		}
	}
	return
}

/*
{para}:
	{se}: ifeval::["
	{se}: {docname}
	{se}: " == "device_library"]
{delim kind=comment}:
	{se}: endif::[]\nCopyright …== "device_library"]
{blank}
{attrib}: toc	unknown render element type: <nil>

{para}:
	{se}: endif::[]
{blank}
{SEC 0 (DeviceType)}:
	{attr:
		 id=ref_BaseDevice
	}
	{title:}
		{se}: Base Device Type
	{body:}
		{blank}
		{SEC 1 (RevisionHistory)}:
			{attr:
				 id=_Revision_History
			}
			{title:}
				{se}: Revision History
			{body:}
				{para}:
					{se}: This is the revision…he description here.
				{blank}
				{tab}:
					{attr:
						 options={
							header,
						}

						 valign=middle
					}
					{head}:
						{cell}:
							{para}:
								{se}: Revision
						{cell}:
							{para}:
								{se}: Description
					{body}:
						{row}:
							{cell}:
								{para}:
									{se}: 0
							{cell}:
								{para}:
									{se}: Represents device de…ype revision numbers
						{row}:
							{cell}:
								{para}:
									{se}: 1
							{cell}:
								{para}:
									{se}: Initial release of this document
						{row}:
							{cell}:
								{para}:
									{se}: 2
							{cell}:
								{para}:
									{se}: Duplicate condition …s Multiple condition
				{blank}
		{SEC 1 (Unknown)}:
			{attr:
				 id=_Overview
			}
			{title:}
				{se}: Overview
			{body:}
				{blank}
				{para}:
					{se}: This defines common … but not limited to:
				{blank}
				{list els}
					{uole bs=1asterisk cs=nocheck}:
						{para}:
							{se}: Certification progra…igbee, Matter, etc.)
					{uole bs=1asterisk cs=nocheck}:
						{para}:
							{se}: Underlying protocol …e PRO, IPv6, TCP/IP)
					{uole bs=1asterisk cs=nocheck}:
						{para}:
							{se}: Regional regulations
					{uole bs=1asterisk cs=nocheck}:
						{para}:
							{se}: Interfaces (UI, cloud, etc.)
					{uole bs=1asterisk cs=nocheck}:
						{para}:
							{se}: Scale (e.g. residential vs commercial)
					{uole bs=1asterisk cs=nocheck}:
						{para}:
							{se}: Other common limitat…ed or sleepy nodes).
					{uole bs=1asterisk cs=nocheck}:
						{para}:
							{se}: etc.
				{blank}
		{SEC 1 (Conditions)}:
			{attr:
				 id=_Conditions
			}
			{title:}
				{se}: Conditions
			{body:}
				{para}:
					{se}: Each section below i…ading purposes only.
				{blank}
				{SEC 2 (Conditions)}:
					{attr:
						 id=_Certification_Program_Conditions
					}
					{title:}
						{se}: Certification Program Conditions
					{body:}
						{para}:
							{se}: At the time of the f…Automation standard.
						{blank}
						{tab}:
							{attr:
								 options={
									header,
								}

								 valign=middle
							}
							{head}:
								{cell}:
									{para}:
										{se}: Certification Program
								{cell}:
									{para}:
										{se}: Tag
								{cell}:
									{para}:
										{se}: Description
							{body}:
								{row}:
									{cell}:
										{para}:
											{se}: Zigbee Home Automation
									{cell}:
										{para}:
											{se}: ZHA
									{cell}:
										{para}:
											{se}: Zigbee Home Automation standard
								{row}:
									{cell}:
										{para}:
											{se}: Zigbee Smart Energy
									{cell}:
										{para}:
											{se}: ZSE
									{cell}:
										{para}:
											{se}: Zigbee Smart Energy standard
								{row}:
									{cell}:
										{para}:
											{se}: Green Power
									{cell}:
										{para}:
											{se}: GP
									{cell}:
										{para}:
											{se}: Zigbee Green Power standard
								{row}:
									{cell}:
										{para}:
											{se}: Zigbee
									{cell}:
										{para}:
											{se}: Zigbee
									{cell}:
										{para}:
											{se}: Zigbee standard
								{row}:
									{cell}:
										{para}:
											{se}: SuZi
									{cell}:
										{para}:
											{se}: SuZi
									{cell}:
										{para}:
											{se}: Zigbee PRO Sub-GHz standard
								{row}:
									{cell}:
										{para}:
											{se}: Matter
									{cell}:
										{para}:
											{se}: Matter
									{cell}:
										{para}:
											{se}: Matter standard
						{blank}
				{SEC 2 (Conditions)}:
					{attr:
						 id=_Protocol_Conditions
					}
					{title:}
						{se}: Protocol Conditions
					{body:}
						{blank}
						{tab}:
							{attr:
								 valign=middle
								 options={
									header,
								}

							}
							{head}:
								{cell}:
									{para}:
										{se}: Protocol Tag
							{body}:
								{row}:
									{cell}:
										{para}:
											{se}: Ethernet
								{row}:
									{cell}:
										{para}:
											{se}: Wi-Fi
								{row}:
									{cell}:
										{para}:
											{se}: Thread
								{row}:
									{cell}:
										{para}:
											{se}: TCP
								{row}:
									{cell}:
										{para}:
											{se}: UDP
								{row}:
									{cell}:
										{para}:
											{se}: IP
								{row}:
									{cell}:
										{para}:
											{se}: IPv4
								{row}:
									{cell}:
										{para}:
											{se}: IPv6
						{blank}
				{SEC 2 (Conditions)}:
					{attr:
						 id=_Interface_Conditions
					}
					{title:}
						{se}: Interface Conditions
					{body:}
						{blank}
						{tab}:
							{attr:
								 valign=middle
								 options={
									header,
								}

							}
							{head}:
								{cell}:
									{para}:
										{se}: Interface Tag
								{cell}:
									{para}:
										{se}: Description
							{body}:
								{row}:
									{cell}:
										{para}:
											{se}: LanguageLocale
									{cell}:
										{para}:
											{se}: The node supports lo…ing text to the user
								{row}:
									{cell}:
										{para}:
											{se}: TimeLocale
									{cell}:
										{para}:
											{se}: The node supports lo…ing time to the user
								{row}:
									{cell}:
										{para}:
											{se}: UnitLocale
									{cell}:
										{para}:
											{se}: The node supports lo… measure to the user
						{para}:
							{se}: Note that "supports …luster interactions.
						{blank}
		{SEC 1 (Unknown)}:
			{attr:
				 id=_Common_Capability_Conditions
			}
			{title:}
				{se}: Common Capability Conditions
			{body:}
				{para}:
					{se}: This category is for…abilities of a node.
				{blank}
				{tab}:
					{attr:
						 options={
							header,
						}

						 valign=middle
					}
					{head}:
						{cell}:
							{para}:
								{se}: Capability Tag
						{cell}:
							{para}:
								{se}: Description
					{body}:
						{row}:
							{cell}:
								{para}:
									{se}: SIT
							{cell}:
								{para}:
									{se}: The node is a short …tly connected device
						{row}:
							{cell}:
								{para}:
									{se}: LIT
							{cell}:
								{para}:
									{se}: The node is a long i…tly connected device
						{row}:
							{cell}:
								{para}:
									{se}: Active
							{cell}:
								{para}:
									{se}: The node is always able to communicate
						{row}:
							{cell}:
								{para}:
									{se}: Simplex
							{cell}:
								{para}:
									{se}: One way communication, client to server
				{blank}
		{SEC 1 (Unknown)}:
			{attr:
				 id=_Device_Type_Class_Conditions
			}
			{title:}
				{se}: Device Type Class Conditions
			{body:}
				{para}:
					{se}: This category is for…on other conditions.
				{blank}
				{tab}:
					{attr:
						 options={
							header,
						}

						 valign=middle
					}
					{head}:
						{cell}:
							{para}:
								{se}: Class Tag
						{cell}:
							{para}:
								{se}: Summary
					{body}:
						{row}:
							{cell}:
								{para}:
									{se}: Node
							{cell}:
								{para}:
									{se}: the device type is c…Model specification)
						{row}:
							{cell}:
								{para}:
									{se}: App
							{cell}:
								{para}:
									{se}: the device type is c…Model specification)
						{row}:
							{cell}:
								{para}:
									{se}: Simple
							{cell}:
								{para}:
									{se}: the device type is c…Model specification)
						{row}:
							{cell}:
								{para}:
									{se}: Dynamic
							{cell}:
								{para}:
									{se}: the device type is c…Model specification)
						{row}:
							{cell}:
								{para}:
									{se}: Client
							{cell}:
								{para}:
									{se}: there exists a clien…ster on the endpoint
						{row}:
							{cell}:
								{para}:
									{se}: Server
							{cell}:
								{para}:
									{se}: there exists a serve…ster on the endpoint
						{row}:
							{cell}:
								{para}:
									{se}: Composed
							{cell}:
								{para}:
									{se}: the device type is c…Model specification)
						{row}:
							{cell}:
								{para}:
									{se}: Duplicate
							{cell}:
								{para}:
									{se}: see
									{xref id:ref_Duplicate label  Duplicate Condition}
						{row}:
							{cell}:
								{para}:
									{se}: EZ-Initiator
							{cell}:
								{para}:
									{se}: the endpoint is an I…bee EZ-Mode Finding
									{sc: &}
									{se}:  Binding
						{row}:
							{cell}:
								{para}:
									{se}: EZ-Target
							{cell}:
								{para}:
									{se}: the endpoint is a Ta…bee EZ-Mode Finding
									{sc: &}
									{se}:  Binding
						{row}:
							{cell}:
								{para}:
									{se}: BridgedPowerSourceInfo
							{cell}:
								{para}:
									{se}: the endpoint represe…ilable to the Bridge
				{blank}
				{SEC 2 (Unknown)}:
					{attr:
						 id=ref_Duplicate,Duplicate Condition in Base Device
					}
					{title:}
						{se}: Duplicate Condition
					{body:}
						{para}:
							{se}: The endpoint and at …odel specification).
						{blank}
						{para}:
							{se}: ifdef::zigbee[]\n== B…or all device types.
						{blank}
				{SEC 2 (Unknown)}:
					{attr:
						 id=_Base_Cluster_Requirements
					}
					{title:}
						{se}: Base Cluster Requirements
					{body:}
						{blank}
						{para}:
							{se}: Each device type def…Conformance column).
						{blank}
						{tab}:
							{attr:
								 options={
									header,
								}

								 valign=middle
							}
							{head}:
								{cell}:
									{para}:
										{se}: Cluster
								{cell}:
									{para}:
										{se}: Client/Server
								{cell}:
									{para}:
										{se}: Quality
								{cell}:
									{para}:
										{se}: Conformance
							{body}:
								{row}:
									{cell}:
										{para}:
											{se}: Basic
									{cell}:
										{para}:
											{se}: Server
									{cell}:
										{para}:
											{se}: I
									{cell}:
										{para}:
								{row}:
									{cell}:
										{para}:
											{se}: Identify
									{cell}:
										{para}:
											{se}: Server
									{cell}:
										{para}:
									{cell}:
										{para}:
											{se}: Simple
								{row}:
									{cell}:
										{para}:
											{se}: Identify
									{cell}:
										{para}:
											{se}: Client
									{cell}:
										{para}:
									{cell}:
										{para}:
											{se}: EZ-Initiator
						{blank}
				{SEC 2 (Unknown)}:
					{attr:
						 id=_Element_Requirements
					}
					{title:}
						{se}: Element Requirements
					{body:}
						{para}:
							{se}: The table below list…try means no change.
						{blank}
						{tab}:
							{attr:
								 options={
									header,
								}

								 valign=middle
							}
							{head}:
								{cell}:
									{para}:
										{se}: ID
								{cell}:
									{para}:
										{se}: Cluster
								{cell}:
									{para}:
										{se}: Element
								{cell}:
									{para}:
										{se}: Name
								{cell}:
									{para}:
										{se}: Quality
								{cell}:
									{para}:
										{se}: Constraint
								{cell}:
									{para}:
										{se}: Access
								{cell}:
									{para}:
										{se}: Conformance
							{body}:
								{row}:
									{cell}:
										{para}:
											{se}: 0x0003
									{cell}:
										{para}:
											{se}: Identify
									{cell}:
										{para}:
											{se}: Feature
									{cell}:
										{para}:
											{se}: QRY
									{cell}:
										{para}:
									{cell}:
										{para}:
									{cell}:
										{para}:
									{cell}:
										{para}:
											{se}: M
						{para}:
							{se}: endif::[]
						{blank}
		{SEC 1 (Unknown)}:
			{attr:
				 id=_Base_Device_Type_Requirements_for_Matter
			}
			{title:}
				{se}: Base Device Type Requirements for Matter
			{body:}
				{para}:
					{se}: These are the base d…or all device types.
				{blank}
				{SEC 2 (Unknown)}:
					{attr:
						 id=_Base_Cluster_Requirements_2
					}
					{title:}
						{se}: Base Cluster Requirements
					{body:}
						{blank}
						{para}:
							{se}: Each device type def…Conformance column).
						{blank}
						{tab}:
							{attr:
								 valign=middle
								 options={
									header,
								}

							}
							{head}:
								{cell}:
									{para}:
										{se}: Cluster
								{cell}:
									{para}:
										{se}: Client/Server
								{cell}:
									{para}:
										{se}: Quality
								{cell}:
									{para}:
										{se}: Conformance
							{body}:
								{row}:
									{cell}:
										{para}:
											{se}: Descriptor
									{cell}:
										{para}:
											{se}: Server
									{cell}:
										{para}:
									{cell}:
										{para}:
											{se}: M
								{row}:
									{cell}:
										{para}:
											{se}: Binding
									{cell}:
										{para}:
											{se}: Server
									{cell}:
										{para}:
									{cell}:
										{para}:
											{se}: Simple
											{sc: &}
											{se}:  Client
								{row}:
									{cell}:
										{para}:
											{se}: Fixed Label
									{cell}:
										{para}:
											{se}: Server
									{cell}:
										{para}:
									{cell}:
										{para}:
											{se}: O
								{row}:
									{cell}:
										{para}:
											{se}: User Label
									{cell}:
										{para}:
											{se}: Server
									{cell}:
										{para}:
									{cell}:
										{para}:
											{se}: O
						{blank}
				{SEC 2 (Unknown)}:
					{attr:
						 id=_Element_Requirements_2
					}
					{title:}
						{se}: Element Requirements
					{body:}
						{para}:
							{se}: Below list qualities…try means no change.
						{blank}
						{tab}:
							{attr:
								 options={
									header,
								}

								 valign=middle
							}
							{head}:
								{cell}:
									{para}:
										{se}: ID
								{cell}:
									{para}:
										{se}: Cluster
								{cell}:
									{para}:
										{se}: Element
								{cell}:
									{para}:
										{se}: Name
								{cell}:
									{para}:
										{se}: Quality
								{cell}:
									{para}:
										{se}: Constraint
								{cell}:
									{para}:
										{se}: Access
								{cell}:
									{para}:
										{se}: Conformance
							{body}:
								{row}:
									{cell}:
										{para}:
											{se}: 0x001D
									{cell}:
										{para}:
											{se}: Descriptor
									{cell}:
										{para}:
											{se}: Feature
									{cell}:
										{para}:
											{se}: TAGLIST
									{cell}:
										{para}:
									{cell}:
										{para}:
									{cell}:
										{para}:
									{cell}:
										{para}:
											{se}: Duplicate

*/

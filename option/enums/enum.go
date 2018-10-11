package enums

type OpCode uint8

func (o OpCode) String() string {

	switch o {
	default:
		return "??"

	case TypePad:
		return "Pad"

	case TypeSubnetMask:
		return "SubnetMask"

	case TypeTimeOffset:
		return "TimeOffset"

	case TypeGateway:
		return "Gateway"

	case TypeTimeServer:
		return "TimeServer"

	case TypeNameServer:
		return "NameServer"

	case TypeDomainNameServer:
		return "DomainNameServer"

	case TypeLogServer:
		return "LogServer"

	case TypeCookieServer:
		return "CookieServer"

	case TypeLPRServer:
		return "LPRServer"

	case TypeImpressServer:
		return "ImpressServer"

	case TypeResourceLocationServer:
		return "ResourceLocationServer"

	case TypeHostName:
		return "HostName"

	case TypeBootFileSize:
		return "BootFileSize"

	case TypeMeritDumpFile:
		return "MeritDumpFile"

	case TypeDomainName:
		return "DomainName"

	case TypeSwapServer:
		return "SwapServer"

	case TypeRootPath:
		return "RootPath"

	case TypeExtensionPath:
		return "ExtensionPath"

	case TypeUseIPForwarding:
		return "UseIPForwarding"

	case TypeUseNonLocalSourceRouting:
		return "UseNonLocalSourceRouting"

	case TypePolicyFilter:
		return "PolicyFilter"

	case TypeMaximumDatagramReassembly:
		return "MaximumDatagramReassembly"

	case TypeDefaultIPTimeToLive:
		return "DefaultIPTimeToLive"

	case TypePathMTUAgingTimeout:
		return "PathMTUAgingTimeout"

	case TypePathMTUPlateauTable:
		return "PathMTUPlateauTable"

	case TypeInterfaceMTU:
		return "InterfaceMTU"

	case TypeAllSubnetsAreLocal:
		return "AllSubnetsAreLocal"

	case TypeBroadcastAddress:
		return "BroadcastAddress"

	case TypePerformMaskDiscovery:
		return "PerformMaskDiscovery"

	case TypeClientShouldRespondsToSubnetMaskICMPRequests:
		return "ClientShouldRespondsToSubnetMaskICMPRequests"

	case TypePerformRouterDiscovery:
		return "PerformRouterDiscovery"

	case TypeRouterSolicitationAddress:
		return "RouterSolicitationAddress"

	case TypeStaticRoute:
		return "StaticRoute"

	case TypeUseTrailerEncapsulation:
		return "UseTrailerEncapsulation"

	case TypeARPCacheTimeout:
		return "ARPCacheTimeout"

	case TypeUseEthernetEncapsulation:
		return "UseEthernetEncapsulation"

	case TypeTCPDefaultTTL:
		return "TCPDefaultTTL"

	case TypeTCPKeepaliveInterval:
		return "TCPKeepaliveInterval"

	case TypeSendTCPKeepaliveGarbageOctet:
		return "SendTCPKeepaliveGarbageOctet"

	case TypeNetworkInformationServiceDomain:
		return "NetworkInformationServiceDomain"

	case TypeNetworkInformationServiceServers:
		return "NetworkInformationServiceServers"

	case TypeNetworkTimeProtocolServers:
		return "NetworkTimeProtocolServers"

	case TypeVendorSpecificInformation:
		return "VendorSpecificInformation"

	case TypeNetBIOSOverTCPIPNameServer:
		return "NetBIOSOverTCPIPNameServer"

	case TypeNetBIOSOverTCPIPDatagramDistributionServer:
		return "NetBIOSOverTCPIPDatagramDistributionServer"

	case TypeNetBIOSOverTCPIPNodeType:
		return "NetBIOSOverTCPIPNodeType"

	case TypeNetBIOSOverTCPIPScope:
		return "NetBIOSOverTCPIPScope"

	case TypeXWindowSystemFontServer:
		return "XWindowSystemFontServer"

	case TypeXWindowSystemDisplayManager:
		return "XWindowSystemDisplayManager"

	case TypeRequestedIPAddress:
		return "RequestedIPAddress"

	case TypeIPAddressLeaseTime:
		return "IPAddressLeaseTime"

	case TypeOverload:
		return "Overload"

	case TypeMessageType:
		return "MessageType"

	case TypeServerIdentifier:
		return "ServerIdentifier"

	case TypeParameterRequestList:
		return "ParameterRequestList"

	case TypeMessage:
		return "Message"

	case TypeMaximumDHCPMessageSize:
		return "MaximumDHCPMessageSize"

	case TypeRenewalTime:
		return "RenewalTime"

	case TypeRebindingTime:
		return "RebindingTime"

	case TypeVendorClassIdentifier:
		return "VendorClassIdentifier"

	case TypeClientIdentifier:
		return "ClientIdentifier"

	case TypeNetWareIpDomainName:
		return "NetWareIpDomainName"

	case TypeNetWareIpInformation:
		return "NetWareIpInformation"

	case TypeNetworkInformationServicePlusDomain:
		return "NetworkInformationServicePlusDomain"

	case TypeNetworkInformationServicePlusServers:
		return "NetworkInformationServicePlusServers"

	case TypeTFTPServerName:
		return "TFTPServerName"

	case TypeBootFileName:
		return "BootFileName"

	case TypeMobileIPHomeAgent:
		return "MobileIPHomeAgent"

	case TypeSimpleMailTransportProtocolServer:
		return "SimpleMailTransportProtocolServer"

	case TypePostOfficeProtocolServer:
		return "PostOfficeProtocolServer"

	case TypeNetworkNewsTransportProtocolServer:
		return "NetworkNewsTransportProtocolServer"

	case TypeDefaultWorldWideWebServer:
		return "DefaultWorldWideWebServer"

	case TypeDefaultFingerServer:
		return "DefaultFingerServer"

	case TypeDefaultInternetRelayChat:
		return "DefaultInternetRelayChat"

	case TypeStreetTalkServer:
		return "StreetTalkServer"

	case TypeStreetTalkDirectoryAssistanceServer:
		return "StreetTalkDirectoryAssistanceServer"

	case TypeUserClassIdentity:
		return "UserClassIdentity"

	case TypeServiceLocationProtocolDirectoryAgent:
		return "ServiceLocationProtocolDirectoryAgent"

	case TypeServiceLocationProtocolServiceScope:
		return "ServiceLocationProtocolServiceScope"

	case TypeRapidCommit:
		return "RapidCommit"

	case TypeFullyQualifiedDomainName:
		return "FullyQualifiedDomainName"

	case TypeRelayAgentCircuitInformation:
		return "RelayAgentCircuitInformation"

	case TypeInternetStorageNameServiceLocation:
		return "InternetStorageNameServiceLocation"

	case TypeNovellDirectoryServicesServers:
		return "NovellDirectoryServicesServers"

	case TypeNovellDirectoryServicesTreeName:
		return "NovellDirectoryServicesTreeName"

	case TypeNovellDirectoryServicesContext:
		return "NovellDirectoryServicesContext"

	case TypeBroadcastAndMulticastServiceControllerDomainNameList:
		return "BroadcastAndMulticastServiceControllerDomainNameList"

	case TypeBroadcastAndMulticastServiceControllerIPv4Address:
		return "BroadcastAndMulticastServiceControllerIPv4Address"

	case TypeAuthentication:
		return "Authentication"

	case TypeLeaseQueryClientLastTransactionTime:
		return "LeaseQueryClientLastTransactionTime"

	case TypeLeaseQueryAssociatedIpAddresses:
		return "LeaseQueryAssociatedIpAddresses"

	case TypeClientSystemArchitectureType:
		return "ClientSystemArchitectureType"

	case TypeClientNetworkInterfaceIdentifier:
		return "ClientNetworkInterfaceIdentifier"

	case TypeLightweightDirectoryAccessProtocolAddress:
		return "LightweightDirectoryAccessProtocolAddress"

	case TypeClientMachineIdentifier:
		return "ClientMachineIdentifier"

	case TypeOpenGroupUserAuthenticationProtocol:
		return "OpenGroupUserAuthenticationProtocol"

	case TypeCivicLocation:
		return "CivicLocation"

	case TypeTimezone:
		return "Timezone"

	case TypeTimezoneDatabase:
		return "TimezoneDatabase"

	case TypeNetInfoParentServerAddress:
		return "NetInfoParentServerAddress"

	case TypeNetInfoParentServerTag:
		return "NetInfoParentServerTag"

	case TypeUrl:
		return "Url"

	case TypeUseStatelessAutoConfigure:
		return "UseStatelessAutoConfigure"

	case TypeNameServiceSearchOrder:
		return "NameServiceSearchOrder"

	case TypeIPv4SubnetSelection:
		return "IPv4SubnetSelection"

	case TypeDomainSearch:
		return "DomainSearch"

	case TypeSessionInitiationProtocolServers:
		return "SessionInitiationProtocolServers"

	case TypeClasslessStaticRoute:
		return "ClasslessStaticRoute"

	case TypeCableLabsClientConfiguration:
		return "CableLabsClientConfiguration"

	case TypeCoordinateBasedLocationConfigurationInformationConfiguration:
		return "CoordinateBasedLocationConfigurationInformationConfiguration"

	case TypeVendorIdentifyingVendorClass:
		return "VendorIdentifyingVendorClass"

	case TypeVendorIdentifyingVendorSpecificInformation:
		return "VendorIdentifyingVendorSpecificInformation"

	case TypeIntelPreBootExecutionEnvironment128:
		return "IntelPreBootExecutionEnvironment128"

	case TypeIntelPreBootExecutionEnvironment129:
		return "IntelPreBootExecutionEnvironment129"

	case TypeIntelPreBootExecutionEnvironment130:
		return "IntelPreBootExecutionEnvironment130"

	case TypeIntelPreBootExecutionEnvironment131:
		return "IntelPreBootExecutionEnvironment131"

	case TypeIntelPreBootExecutionEnvironment132:
		return "IntelPreBootExecutionEnvironment132"

	case TypeIntelPreBootExecutionEnvironment133:
		return "IntelPreBootExecutionEnvironment133"

	case TypeIntelPreBootExecutionEnvironment134:
		return "IntelPreBootExecutionEnvironment134"

	case TypeIntelPreBootExecutionEnvironment135:
		return "IntelPreBootExecutionEnvironment135"

	case TypeProtocolForCarryingAuthenticationForNetworkAccessAgent:
		return "ProtocolForCarryingAuthenticationForNetworkAccessAgent"

	case TypeLocationToServiceTranslationServerLocation:
		return "LocationToServiceTranslationServerLocation"

	case TypeControlAndProvisioningOfWirelessAccessPointsAccessControllerAddresses:
		return "ControlAndProvisioningOfWirelessAccessPointsAccessControllerAddresses"

	case TypeMobilityServicesDiscoveryServerAddresses:
		return "MobilityServicesDiscoveryServerAddresses"

	case TypeMobilityServicesDiscoveryServerDomainAddresses:
		return "MobilityServicesDiscoveryServerDomainAddresses"

	case TypeSessionInitiationProtocolUserAgentConfigurationServiceDomains:
		return "SessionInitiationProtocolUserAgentConfigurationServiceDomains"

	case TypeAccessNetworkDiscoveryAndSelectionFunctionIPv4Addresses:
		return "AccessNetworkDiscoveryAndSelectionFunctionIPv4Addresses"

	case TypeCoordinateBasedLocationConfigurationInformationLocation:
		return "CoordinateBasedLocationConfigurationInformationLocation"

	case TypeForceRenewNonceProtocolCapabilities:
		return "ForceRenewNonceProtocolCapabilities"

	case TypeRecursiveDnsServersSelection:
		return "RecursiveDnsServersSelection"

	case TypeTftpServerAddresses:
		return "TftpServerAddresses"

	case TypeBulkLeaseQueryStatusCode:
		return "BulkLeaseQueryStatusCode"

	case TypeBulkLeaseQueryBaseTime:
		return "BulkLeaseQueryBaseTime"

	case TypeBulkLeaseQueryStateStartSeconds:
		return "BulkLeaseQueryStateStartSeconds"

	case TypeBulkLeaseQueryStartTime:
		return "BulkLeaseQueryStartTime"

	case TypeBulkLeaseQueryEndTime:
		return "BulkLeaseQueryEndTime"

	case TypeBulkLeaseQueryDhcpState:
		return "BulkLeaseQueryDhcpState"

	case TypeBulkLeaseQueryDataSource:
		return "BulkLeaseQueryDataSource"

	case TypePortControlProtocolServers:
		return "PortControlProtocolServers"

	case TypeDynamicAllocationOfSharedIpv4AddressesPortParameters:
		return "DynamicAllocationOfSharedIpv4AddressesPortParameters"

	case TypeCaptivePortalIdentificationUri:
		return "CaptivePortalIdentificationUri"

	case TypeManufacturerUsageDescriptionUrl:
		return "ManufacturerUsageDescriptionUrl"

	case TypeLinuxPreBootExecutionEnvironmentConfigurationFile:
		return "LinuxPreBootExecutionEnvironmentConfigurationFile"

	case TypeLinuxPreBootExecutionEnvironmentPathPrefix:
		return "LinuxPreBootExecutionEnvironmentPathPrefix"

	case TypeLinuxPreBootExecutionEnvironmentRebootTimeSeconds:
		return "LinuxPreBootExecutionEnvironmentRebootTimeSeconds"

	case TypeIpv6RapidDeploymentAddresses:
		return "Ipv6RapidDeploymentAddresses"

	case TypeLocalLocationInformationServerAccessNetworkDomainName:
		return "LocalLocationInformationServerAccessNetworkDomainName"

	case TypeCiscoSystemsSubnetAllocation:
		return "CiscoSystemsSubnetAllocation"

	case TypeVirtualSubnetSelectionInformation:
		return "VirtualSubnetSelectionInformation"

	case TypeMicrosoftClasslessStaticRoute:
		return "MicrosoftClasslessStaticRoute"

	case TypeEnd:
		return "End"
	}
}

// * = minimum implementation
const (
	TypePad                                                                   OpCode = 0   // 0*
	TypeSubnetMask                                                            OpCode = 1   // 1*
	TypeTimeOffset                                                            OpCode = 2   // 2 deprecated by 100 and 101
	TypeGateway                                                               OpCode = 3   // 3* GW
	TypeTimeServer                                                            OpCode = 4   // 4
	TypeNameServer                                                            OpCode = 5   // 5 do not confuse with 6
	TypeDomainNameServer                                                      OpCode = 6   // 6* DNS do not confuse with 5
	TypeLogServer                                                             OpCode = 7   // 7
	TypeCookieServer                                                          OpCode = 8   // 8
	TypeLPRServer                                                             OpCode = 9   // 9
	TypeImpressServer                                                         OpCode = 10  // 10
	TypeResourceLocationServer                                                OpCode = 11  // 11
	TypeHostName                                                              OpCode = 12  // 12
	TypeBootFileSize                                                          OpCode = 13  // 13
	TypeMeritDumpFile                                                         OpCode = 14  // 14
	TypeDomainName                                                            OpCode = 15  // 15
	TypeSwapServer                                                            OpCode = 16  // 16
	TypeRootPath                                                              OpCode = 17  // 17
	TypeExtensionPath                                                         OpCode = 18  // 18
	TypeUseIPForwarding                                                       OpCode = 19  // 19
	TypeUseNonLocalSourceRouting                                              OpCode = 20  // 20
	TypePolicyFilter                                                          OpCode = 21  // 21
	TypeMaximumDatagramReassembly                                             OpCode = 22  // 22
	TypeDefaultIPTimeToLive                                                   OpCode = 23  // 23
	TypePathMTUAgingTimeout                                                   OpCode = 24  // 24
	TypePathMTUPlateauTable                                                   OpCode = 25  // 25
	TypeInterfaceMTU                                                          OpCode = 26  // 26
	TypeAllSubnetsAreLocal                                                    OpCode = 27  // 27
	TypeBroadcastAddress                                                      OpCode = 28  // 28
	TypePerformMaskDiscovery                                                  OpCode = 29  // 29
	TypeClientShouldRespondsToSubnetMaskICMPRequests                          OpCode = 30  // 30
	TypePerformRouterDiscovery                                                OpCode = 31  // 31
	TypeRouterSolicitationAddress                                             OpCode = 32  // 32
	TypeStaticRoute                                                           OpCode = 33  // 33
	TypeUseTrailerEncapsulation                                               OpCode = 34  // 34
	TypeARPCacheTimeout                                                       OpCode = 35  // 35
	TypeUseEthernetEncapsulation                                              OpCode = 36  // 36
	TypeTCPDefaultTTL                                                         OpCode = 37  // 37
	TypeTCPKeepaliveInterval                                                  OpCode = 38  // 38
	TypeSendTCPKeepaliveGarbageOctet                                          OpCode = 39  // 39
	TypeNetworkInformationServiceDomain                                       OpCode = 40  // 40
	TypeNetworkInformationServiceServers                                      OpCode = 41  // 41
	TypeNetworkTimeProtocolServers                                            OpCode = 42  // 42 NTP
	TypeVendorSpecificInformation                                             OpCode = 43  // 43 see also 60
	TypeNetBIOSOverTCPIPNameServer                                            OpCode = 44  // 44
	TypeNetBIOSOverTCPIPDatagramDistributionServer                            OpCode = 45  // 45
	TypeNetBIOSOverTCPIPNodeType                                              OpCode = 46  // 46
	TypeNetBIOSOverTCPIPScope                                                 OpCode = 47  // 47
	TypeXWindowSystemFontServer                                               OpCode = 48  // 48
	TypeXWindowSystemDisplayManager                                           OpCode = 49  // 49
	TypeRequestedIPAddress                                                    OpCode = 50  // 50*
	TypeIPAddressLeaseTime                                                    OpCode = 51  // 51*
	TypeOverload                                                              OpCode = 52  // 52
	TypeMessageType                                                           OpCode = 53  // 53*
	TypeServerIdentifier                                                      OpCode = 54  // 54*
	TypeParameterRequestList                                                  OpCode = 55  // 55*
	TypeMessage                                                               OpCode = 56  // 56
	TypeMaximumDHCPMessageSize                                                OpCode = 57  // 57
	TypeRenewalTime                                                           OpCode = 58  // 58*
	TypeRebindingTime                                                         OpCode = 59  // 59*
	TypeVendorClassIdentifier                                                 OpCode = 60  // 60 see also 43
	TypeClientIdentifier                                                      OpCode = 61  // 61*
	TypeNetWareIpDomainName                                                   OpCode = 62  // 62
	TypeNetWareIpInformation                                                  OpCode = 63  // 63
	TypeNetworkInformationServicePlusDomain                                   OpCode = 64  // 64
	TypeNetworkInformationServicePlusServers                                  OpCode = 65  // 65
	TypeTFTPServerName                                                        OpCode = 66  // 66
	TypeBootFileName                                                          OpCode = 67  // 67
	TypeMobileIPHomeAgent                                                     OpCode = 68  // 68
	TypeSimpleMailTransportProtocolServer                                     OpCode = 69  // 69 SMTP
	TypePostOfficeProtocolServer                                              OpCode = 70  // 70 POP3
	TypeNetworkNewsTransportProtocolServer                                    OpCode = 71  // 71 NNTP
	TypeDefaultWorldWideWebServer                                             OpCode = 72  // 72
	TypeDefaultFingerServer                                                   OpCode = 73  // 73
	TypeDefaultInternetRelayChat                                              OpCode = 74  // 74 IRC
	TypeStreetTalkServer                                                      OpCode = 75  // 75
	TypeStreetTalkDirectoryAssistanceServer                                   OpCode = 76  // 76
	TypeUserClassIdentity                                                     OpCode = 77  // 77
	TypeServiceLocationProtocolDirectoryAgent                                 OpCode = 78  // 78
	TypeServiceLocationProtocolServiceScope                                   OpCode = 79  // 79
	TypeRapidCommit                                                           OpCode = 80  // 80
	TypeFullyQualifiedDomainName                                              OpCode = 81  // 81
	TypeRelayAgentCircuitInformation                                          OpCode = 82  // 82
	TypeInternetStorageNameServiceLocation                                    OpCode = 83  // 83
	TypeNovellDirectoryServicesServers                                        OpCode = 85  // 85
	TypeNovellDirectoryServicesTreeName                                       OpCode = 86  // 86
	TypeNovellDirectoryServicesContext                                        OpCode = 87  // 87
	TypeBroadcastAndMulticastServiceControllerDomainNameList                  OpCode = 88  // 88
	TypeBroadcastAndMulticastServiceControllerIPv4Address                     OpCode = 89  // 89
	TypeAuthentication                                                        OpCode = 90  // 90
	TypeLeaseQueryClientLastTransactionTime                                   OpCode = 91  // 91
	TypeLeaseQueryAssociatedIpAddresses                                       OpCode = 92  // 92
	TypeClientSystemArchitectureType                                          OpCode = 93  // 93
	TypeClientNetworkInterfaceIdentifier                                      OpCode = 94  // 94
	TypeLightweightDirectoryAccessProtocolAddress                             OpCode = 95  // 95
	TypeClientMachineIdentifier                                               OpCode = 97  // 97
	TypeOpenGroupUserAuthenticationProtocol                                   OpCode = 98  // 98
	TypeCivicLocation                                                         OpCode = 99  // 99
	TypeTimezone                                                              OpCode = 100 // 100
	TypeTimezoneDatabase                                                      OpCode = 101 // 101
	TypeNetInfoParentServerAddress                                            OpCode = 112 // 112
	TypeNetInfoParentServerTag                                                OpCode = 113 // 113
	TypeUrl                                                                   OpCode = 114 // 114
	TypeUseStatelessAutoConfigure                                             OpCode = 116 // 116
	TypeNameServiceSearchOrder                                                OpCode = 117 // 117
	TypeIPv4SubnetSelection                                                   OpCode = 118 // 118
	TypeDomainSearch                                                          OpCode = 119 // 119
	TypeSessionInitiationProtocolServers                                      OpCode = 120 // 120
	TypeClasslessStaticRoute                                                  OpCode = 121 // 121
	TypeCableLabsClientConfiguration                                          OpCode = 122 // 122
	TypeCoordinateBasedLocationConfigurationInformationConfiguration          OpCode = 123 // 123
	TypeVendorIdentifyingVendorClass                                          OpCode = 124 // 124
	TypeVendorIdentifyingVendorSpecificInformation                            OpCode = 125 // 125
	TypeIntelPreBootExecutionEnvironment128                                   OpCode = 128 // 128 PXE
	TypeIntelPreBootExecutionEnvironment129                                   OpCode = 129 // 129 PXE
	TypeIntelPreBootExecutionEnvironment130                                   OpCode = 130 // 130 PXE
	TypeIntelPreBootExecutionEnvironment131                                   OpCode = 131 // 131 PXE
	TypeIntelPreBootExecutionEnvironment132                                   OpCode = 132 // 132 PXE
	TypeIntelPreBootExecutionEnvironment133                                   OpCode = 133 // 133 PXE
	TypeIntelPreBootExecutionEnvironment134                                   OpCode = 134 // 134 PXE
	TypeIntelPreBootExecutionEnvironment135                                   OpCode = 135 // 135 PXE
	TypeProtocolForCarryingAuthenticationForNetworkAccessAgent                OpCode = 136 // 136
	TypeLocationToServiceTranslationServerLocation                            OpCode = 137 // 137
	TypeControlAndProvisioningOfWirelessAccessPointsAccessControllerAddresses OpCode = 138 // 138
	TypeMobilityServicesDiscoveryServerAddresses                              OpCode = 139 // 139
	TypeMobilityServicesDiscoveryServerDomainAddresses                        OpCode = 140 // 140
	TypeSessionInitiationProtocolUserAgentConfigurationServiceDomains         OpCode = 141 // 141
	TypeAccessNetworkDiscoveryAndSelectionFunctionIPv4Addresses               OpCode = 142 // 142
	TypeCoordinateBasedLocationConfigurationInformationLocation               OpCode = 144 // 144
	TypeForceRenewNonceProtocolCapabilities                                   OpCode = 145 // 145
	TypeRecursiveDnsServersSelection                                          OpCode = 146 // 146
	TypeTftpServerAddresses                                                   OpCode = 150 // 150
	TypeBulkLeaseQueryStatusCode                                              OpCode = 151 // 151
	TypeBulkLeaseQueryBaseTime                                                OpCode = 152 // 152
	TypeBulkLeaseQueryStateStartSeconds                                       OpCode = 153 // 153
	TypeBulkLeaseQueryStartTime                                               OpCode = 154 // 154
	TypeBulkLeaseQueryEndTime                                                 OpCode = 155 // 155
	TypeBulkLeaseQueryDhcpState                                               OpCode = 156 // 156
	TypeBulkLeaseQueryDataSource                                              OpCode = 157 // 157
	TypePortControlProtocolServers                                            OpCode = 158 // 158
	TypeDynamicAllocationOfSharedIpv4AddressesPortParameters                  OpCode = 159 // 159
	TypeCaptivePortalIdentificationUri                                        OpCode = 160 // 160
	TypeManufacturerUsageDescriptionUrl                                       OpCode = 161 // 161
	TypeLinuxPreBootExecutionEnvironmentConfigurationFile                     OpCode = 209 // 209
	TypeLinuxPreBootExecutionEnvironmentPathPrefix                            OpCode = 210 // 210
	TypeLinuxPreBootExecutionEnvironmentRebootTimeSeconds                     OpCode = 211 // 211
	TypeIpv6RapidDeploymentAddresses                                          OpCode = 212 // 212
	TypeLocalLocationInformationServerAccessNetworkDomainName                 OpCode = 213 // 213
	TypeCiscoSystemsSubnetAllocation                                          OpCode = 220 // 220
	TypeVirtualSubnetSelectionInformation                                     OpCode = 221 // 221
	TypeMicrosoftClasslessStaticRoute                                         OpCode = 249 // 249 see 121
	TypeEnd                                                                   OpCode = 255 // 255*
)

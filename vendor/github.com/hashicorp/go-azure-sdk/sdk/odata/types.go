package odata

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShortType = string

const (
	ShortTypeAccessPackage                               ShortType = "accessPackage"
	ShortTypeAccessPackageAssignmentPolicy               ShortType = "accessPackageAssignmentPolicy"
	ShortTypeAccessPackageCatalog                        ShortType = "accessPackageCatalog"
	ShortTypeAccessPackageResourceRequest                ShortType = "accessPackageResourceRequest"
	ShortTypeAccessPackageQuestion                       ShortType = "accessPackageQuestion"
	ShortTypeAccessPackageTextInputQuestion              ShortType = "accessPackageTextInputQuestion"
	ShortTypeAccessPackageMultipleChoiceQuestion         ShortType = "accessPackageMultipleChoiceQuestion"
	ShortTypeAdministrativeUnit                          ShortType = "administrativeUnit"
	ShortTypeApplication                                 ShortType = "application"
	ShortTypeConditionalAccessPolicy                     ShortType = "conditionalAccessPolicy"
	ShortTypeConnectedOrganizationMembers                ShortType = "connectedOrganizationMembers"
	ShortTypeConnectionInfo                              ShortType = "connectionInfo"
	ShortTypeCountryNamedLocation                        ShortType = "countryNamedLocation"
	ShortTypeDevice                                      ShortType = "device"
	ShortTypeDirectoryRole                               ShortType = "directoryRole"
	ShortTypeDirectoryRoleTemplate                       ShortType = "directoryRoleTemplate"
	ShortTypeDomain                                      ShortType = "domain"
	ShortTypeEmailAuthenticationMethod                   ShortType = "emailAuthenticationMethod"
	ShortTypeExternalSponsors                            ShortType = "externalSponsors"
	ShortTypeFido2AuthenticationMethod                   ShortType = "fido2AuthenticationMethod"
	ShortTypeGroup                                       ShortType = "group"
	ShortTypeGroupMembers                                ShortType = "groupMembers"
	ShortTypeIpNamedLocation                             ShortType = "ipNamedLocation"
	ShortTypeInternalSponsors                            ShortType = "internalSponsors"
	ShortTypeNamedLocation                               ShortType = "namedLocation"
	ShortTypeMicrosoftAuthenticatorAuthenticationMethod  ShortType = "microsoftAuthenticatorAuthenticationMethod"
	ShortTypeOrganization                                ShortType = "organization"
	ShortTypePasswordAuthenticationMethod                ShortType = "passwordAuthenticationMethod"
	ShortTypePhoneAuthenticationMethod                   ShortType = "phoneAuthenticationMethod"
	ShortTypeRequestorManager                            ShortType = "requestorManager"
	ShortTypeServicePrincipal                            ShortType = "servicePrincipal"
	ShortTypeSingleUser                                  ShortType = "singleUser"
	ShortTypeSocialIdentityProvider                      ShortType = "socialIdentityProvider"
	ShortTypeTemporaryAccessPassAuthenticationMethod     ShortType = "temporaryAccessPassAuthenticationMethod"
	ShortTypeUser                                        ShortType = "user"
	ShortTypeWindowsHelloForBusinessAuthenticationMethod ShortType = "windowsHelloForBusinessAuthenticationMethod"
)

type Type = string

const (
	TypeAccessPackage                                    Type = "#microsoft.graph.accessPackage"
	TypeAccessPackageAssignmentPolicy                    Type = "#microsoft.graph.accessPackageAssignmentPolicy"
	TypeAccessPackageCatalog                             Type = "#microsoft.graph.accessPackageCatalog"
	TypeAccessPackageMultipleChoiceQuestion              Type = "#microsoft.graph.accessPackageMultipleChoiceQuestion"
	TypeAccessPackageQuestion                            Type = "#microsoft.graph.accessPackageQuestion"
	TypeAccessPackageResourceRequest                     Type = "#microsoft.graph.accessPackageResourceRequest"
	TypeAccessPackageTextInputQuestion                   Type = "#microsoft.graph.accessPackageTextInputQuestion"
	TypeActiveDirectoryWindowsAutopilotDeploymentProfile Type = "#microsoft.graph.activeDirectoryWindowsAutopilotDeploymentProfile"
	TypeAdministrativeUnit                               Type = "#microsoft.graph.administrativeUnit"
	TypeAndroidForWorkApp                                Type = "#microsoft.graph.androidForWorkApp"
	TypeAndroidLobApp                                    Type = "#microsoft.graph.androidLobApp"
	TypeAndroidManagedStoreApp                           Type = "#microsoft.graph.androidManagedStoreApp"
	TypeAndroidManagedStoreWebApp                        Type = "#microsoft.graph.androidManagedStoreWebApp"
	TypeAndroidStoreApp                                  Type = "#microsoft.graph.androidStoreApp"
	TypeApplication                                      Type = "#microsoft.graph.application"
	TypeAzureActiveDirectoryTenant                       Type = "#microsoft.graph.azureActiveDirectoryTenant"
	TypeAzureADWindowsAutopilotDeploymentProfile         Type = "#microsoft.graph.azureADWindowsAutopilotDeploymentProfile"
	TypeConditionalAccessPolicy                          Type = "#microsoft.graph.conditionalAccessPolicy"
	TypeConnectedOrganizationMembers                     Type = "#microsoft.graph.connectedOrganizationMembers"
	TypeConnectionInfo                                   Type = "#microsoft.graph.connectionInfo"
	TypeCountryNamedLocation                             Type = "#microsoft.graph.countryNamedLocation"
	TypeDevice                                           Type = "#microsoft.graph.device"
	TypeDirectoryRole                                    Type = "#microsoft.graph.directoryRole"
	TypeDirectoryRoleTemplate                            Type = "#microsoft.graph.directoryRoleTemplate"
	TypeDomain                                           Type = "#microsoft.graph.domain"
	TypeDomainIdentitySource                             Type = "#microsoft.graph.domainIdentitySource"
	TypeEmailAuthenticationMethod                        Type = "#microsoft.graph.emailAuthenticationMethod"
	TypeExternalDomainFederation                         Type = "#microsoft.graph.externalDomainFederation"
	TypeExternalSponsors                                 Type = "#microsoft.graph.externalSponsors"
	TypeFido2AuthenticationMethod                        Type = "#microsoft.graph.fido2AuthenticationMethod"
	TypeGroup                                            Type = "#microsoft.graph.group"
	TypeGroupMembers                                     Type = "#microsoft.graph.groupMembers"
	TypeInternalSponsors                                 Type = "#microsoft.graph.internalSponsors"
	TypeIOSLobApp                                        Type = "#microsoft.graph.iosLobApp"
	TypeIOSiPadOSWebClip                                 Type = "#microsoft.graph.iosiPadOSWebClip"
	TypeIOSStoreApp                                      Type = "#microsoft.graph.iosStoreApp"
	TypeIOSVppApp                                        Type = "#microsoft.graph.iosVppApp"
	TypeIpNamedLocation                                  Type = "#microsoft.graph.ipNamedLocation"
	TypeMacOSDmgApp                                      Type = "#microsoft.graph.macOSDmgApp"
	TypeMacOSLobApp                                      Type = "#microsoft.graph.macOSLobApp"
	TypeMacOSMdatpApp                                    Type = "#microsoft.graph.macOSMdatpApp"
	TypeMacOSMicrosoftDefenderApp                        Type = "#microsoft.graph.macOSMicrosoftDefenderApp"
	TypeMacOSMicrosoftEdgeApp                            Type = "#microsoft.graph.macOSMicrosoftEdgeApp"
	TypeMacOSOfficeSuiteApp                              Type = "#microsoft.graph.macOSOfficeSuiteApp"
	TypeMacOsVppApp                                      Type = "#microsoft.graph.macOsVppApp"
	TypeManagedAndroidLobApp                             Type = "#microsoft.graph.managedAndroidLobApp"
	TypeManagedAndroidStoreApp                           Type = "#microsoft.graph.managedAndroidStoreApp"
	TypeManagedIOSLobApp                                 Type = "#microsoft.graph.managedIOSLobApp"
	TypeManagedIOSStoreApp                               Type = "#microsoft.graph.managedIOSStoreApp"
	TypeMicrosoftAuthenticatorAuthenticationMethod       Type = "#microsoft.graph.microsoftAuthenticatorAuthenticationMethod"
	TypeMicrosoftStoreForBusinessApp                     Type = "#microsoft.graph.microsoftStoreForBusinessApp"
	TypeNamedLocation                                    Type = "#microsoft.graph.namedLocation"
	TypeOfficeSuiteApp                                   Type = "#microsoft.graph.officeSuiteApp"
	TypeOrganization                                     Type = "#microsoft.graph.organization"
	TypePasswordAuthenticationMethod                     Type = "#microsoft.graph.passwordAuthenticationMethod"
	TypePhoneAuthenticationMethod                        Type = "#microsoft.graph.phoneAuthenticationMethod"
	TypeRequestorManager                                 Type = "#microsoft.graph.requestorManager"
	TypeServicePrincipal                                 Type = "#microsoft.graph.servicePrincipal"
	TypeSingleUser                                       Type = "#microsoft.graph.singleUser"
	TypeSocialIdentityProvider                           Type = "#microsoft.graph.socialIdentityProvider"
	TypeTemporaryAccessPassAuthenticationMethod          Type = "#microsoft.graph.temporaryAccessPassAuthenticationMethod"
	TypeUser                                             Type = "#microsoft.graph.user"
	TypeWebApp                                           Type = "#microsoft.graph.webApp"
	TypeWinGetApp                                        Type = "#microsoft.graph.winGetApp"
	TypeWin32LobApp                                      Type = "#microsoft.graph.win32LobApp"
	TypeWindowsAppX                                      Type = "#microsoft.graph.windowsAppX"
	TypeWindowsHelloForBusinessAuthenticationMethod      Type = "#microsoft.graph.windowsHelloForBusinessAuthenticationMethod"
	TypeWindowsMicrosoftEdgeApp                          Type = "#microsoft.graph.windowsMicrosoftEdgeApp"
	TypeWindowsMobileMSI                                 Type = "#microsoft.graph.windowsMobileMSI"
	TypeWindowsPhone81AppX                               Type = "#microsoft.graph.windowsPhone81AppX"
	TypeWindowsPhone81AppXBundle                         Type = "#microsoft.graph.windowsPhone81AppXBundle"
	TypeWindowsPhone81StoreApp                           Type = "#microsoft.graph.windowsPhone81StoreApp"
	TypeWindowsPhoneXAP                                  Type = "#microsoft.graph.windowsPhoneXAP"
	TypeWindowsStoreApp                                  Type = "#microsoft.graph.windowsStoreApp"
	TypeWindowsUniversalAppX                             Type = "#microsoft.graph.windowsUniversalAppX"
	TypeWindowsWebApp                                    Type = "#microsoft.graph.windowsWebApp"
)

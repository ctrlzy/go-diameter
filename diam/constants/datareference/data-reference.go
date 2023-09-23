package datareference

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

// The Data-Reference AVP is of type Enumerated, and indicates the type of the requested user data
// in the operation UDR and SNR. Its exact values and meaning is defined in 3GPP TS 29.328 [1].
// The following values are defined (more details are given in 3GPP TS 29.328 [1]):
/*
RepositoryData (0)
IMSPublicIdentity (10)
IMSUserState (11)
S-CSCFName (12)
InitialFilterCriteria (13)		This value is used to request initial filter criteria relevant to the requesting AS
LocationInformation (14)
UserState (15)
ChargingInformation (16)
MSISDN (17)
PSIActivation (18)
DSAI (19)
ServiceLevelTraceInfo (21)
IPAddressSecureBindingInformation (22)
ServicePriorityLevel (23)
SMSRegistrationInfo (24)
UEReachabilityForIP (25)
TADSinformation (26)
STN-SR (27)
UE-SRVCC-Capability (28)
ExtendedPriority (29)
CSRN (30)
ReferenceLocationInformation (31)
IMSI (32)
IMSPrivateUserIdentity (33)
IMEISV (34)
UE-5G-SRVCC-Capability (35)
NOTE: Value 20 is reserved.
*/
const (
	REPOSITORY_DATA                       = datatype.Enumerated(0)
	IMS_PUBLIC_IDENTITY                   = datatype.Enumerated(10)
	IMS_USER_STATE                        = datatype.Enumerated(11)
	SCSCF_NAME                            = datatype.Enumerated(12)
	INITIAL_FILTER_CRITERIA               = datatype.Enumerated(13)
	LOCATION_INFORMATION                  = datatype.Enumerated(14)
	USER_STATE                            = datatype.Enumerated(15)
	CHARGING_INFORMATION                  = datatype.Enumerated(16)
	MSISDN                                = datatype.Enumerated(17)
	PSI_ACTIVATION                        = datatype.Enumerated(18)
	DSAI                                  = datatype.Enumerated(19)
	SERVICE_LEVEL_TRACE_INFO              = datatype.Enumerated(21)
	IP_ADDRESS_SECURE_BINDING_INFORMATION = datatype.Enumerated(22)
	SERVICE_PRIORITY_LEVEL                = datatype.Enumerated(23)
	SMS_REGISTRATION_INFO                 = datatype.Enumerated(24)
	UE_REACHABILITY_FOR_IP                = datatype.Enumerated(25)
	TADS_INFORMATION                      = datatype.Enumerated(26)
	STN_SR                                = datatype.Enumerated(27)
	UE_SRVCC_CAPABILITY                   = datatype.Enumerated(28)
	EXTENDED_PRIORITY                     = datatype.Enumerated(29)
	CSRN                                  = datatype.Enumerated(30)
	REFERENCE_LOCATION_INFORMATION        = datatype.Enumerated(31)
	IMSI                                  = datatype.Enumerated(32)
	IMS_PRIVATE_USER_IDENTITY             = datatype.Enumerated(33)
	IMEISV                                = datatype.Enumerated(34)
	UE_5G_SRVCC_CAPABILITY                = datatype.Enumerated(35)
)

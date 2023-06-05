package resultcode

// RFC 6733 The Result-Code AVP (AVP Code 268) is of type Unsigned32 and
// indicates whether a particular request was completed successfully or
// an error occurred.
// 3GPP.29.338 result-code define

type ResultCode int32

const (
	// Infomational 1xxx
	// This informational error is returned by a Diameter server to
	// inform the access device that the authentication mechanism being
	// used requires multiple round trips, and a subsequent request needs
	// to be issued in order for access to be granted.
	ResultCode_DIAMETER_MULTI_ROUND_AUTH ResultCode = 1001

	// Success 2xxx
	// The request was successfully completed.
	ResultCode_DIAMETER_SUCCESS ResultCode = 2001

	// When returned, the request was successfully completed, but
	// additional processing is required by the application in order to
	// provide service to the user.
	ResultCode_DIAMETER_LIMITED_SUCCESS ResultCode = 2002

	// Protocol Error 3xxx
	// This error code is used when a Diameter entity receives a message
	// with a Command Code that it does not support.
	ResultCode_DIAMETER_COMMAND_UNSUPPORTED ResultCode = 3001

	// This error is given when Diameter cannot deliver the message to
	// the destination, either because no host within the realm
	// supporting the required application was available to process the
	// request or because the Destination-Host AVP was given without the
	// associated Destination-Realm AVP.
	ResultCode_DIAMETER_UNABLE_TO_DELIVER ResultCode = 3002

	// The intended realm of the request is not recognized.
	ResultCode_DIAMETER_REALM_NOT_SERVED ResultCode = 3003

	// When returned, a Diameter node SHOULD attempt to send the message
	// to an alternate peer.  This error MUST only be used when a
	// specific server is requested, and it cannot provide the requested
	// service.
	ResultCode_DIAMETER_TOO_BUSY ResultCode = 3004

	// An agent detected a loop while trying to get the message to the
	// intended recipient.  The message MAY be sent to an alternate peer,
	// if one is available, but the peer reporting the error has
	// identified a configuration problem.
	ResultCode_DIAMETER_LOOP_DETECTED ResultCode = 3005

	// A redirect agent has determined that the request could not be
	// satisfied locally, and the initiator of the request SHOULD direct
	// the request directly to the server, whose contact information has
	// been added to the response.  When set, the Redirect-Host AVP MUST
	// be present.
	ResultCode_DIAMETER_REDIRECT_INDICATION ResultCode = 3006

	// A request was sent for an application that is not supported
	ResultCode_DIAMETER_APPLICATION_UNSUPPORTED ResultCode = 3007

	// A request was received whose bits in the Diameter header were set
	// either to an invalid combination or to a value that is
	// inconsistent with the Command Code's definition.
	ResultCode_DIAMETER_INVALID_HDR_BITS ResultCode = 3008

	// A request was received that included an AVP whose flag bits are
	// set to an unrecognized value or that is inconsistent with the
	// AVP's definition.
	ResultCode_DIAMETER_INVALID_AVP_BITS ResultCode = 3009

	// A CER was received from an unknown peer.
	ResultCode_DIAMETER_UNKNOWN_PEER ResultCode = 3010

	// Transient Failures 4xxx
	// Errors that fall within the transient failures category are used to
	// inform a peer that the request could not be satisfied at the time it
	// was received but MAY be able to satisfy the request in the future.
	// Note that these errors MUST be used in answer messages whose 'E' bit
	// is not set.
	ResultCode_DIAMETER_AUTHENTICATION_REJECTED ResultCode = 4001

	// A Diameter node received the accounting request but was unable to
	// commit it to stable storage due to a temporary lack of space.
	ResultCode_DIAMETER_OUT_OF_SPACE ResultCode = 4002

	// The peer has determined that it has lost the election process and
	// has therefore disconnected the transport connection.
	ResultCode_ELECTION_LOST ResultCode = 4003

	// Permanent Failures 5xxx
	// The peer received a message that contained an AVP that is not
	// recognized or supported and was marked with the 'M' (Mandatory)
	// bit.  A Diameter message with this error MUST contain one or more
	// Failed-AVP AVPs containing the AVPs that caused the failure.
	ResultCode_DIAMETER_AVP_UNSUPPORTED ResultCode = 5001

	// The request contained an unknown Session-Id.
	ResultCode_DIAMETER_UNKNOWN_SESSION_ID ResultCode = 5002

	// A request was received for which the user could not be authorized.
	// This error could occur if the service requested is not permitted
	// to the user.
	ResultCode_DIAMETER_AUTHORIZATION_REJECTED ResultCode = 5003

	// The request contained an AVP with an invalid value in its data
	// portion.  A Diameter message indicating this error MUST include
	// the offending AVPs within a Failed-AVP AVP.
	ResultCode_DIAMETER_INVALID_AVP_VALUE ResultCode = 5004

	// The request did not contain an AVP that is required by the Command
	// Code definition.  If this value is sent in the Result-Code AVP, a
	// Failed-AVP AVP SHOULD be included in the message.  The Failed-AVP
	// AVP MUST contain an example of the missing AVP complete with the
	// Vendor-Id if applicable.  The value field of the missing AVP
	// should be of correct minimum length and contain zeroes.
	ResultCode_DIAMETER_MISSING_AVP ResultCode = 5005

	// A request was received that cannot be authorized because the user
	// has already expended allowed resources.  An example of this error
	// condition is when a user that is restricted to one dial-up PPP
	// port attempts to establish a second PPP connection.
	ResultCode_DIAMETER_RESOURCES_EXCEEDED ResultCode = 5006

	// The Home Diameter server has detected AVPs in the request that
	// contradicted each other, and it is not willing to provide service
	// to the user.  The Failed-AVP AVP MUST be present, which contain
	// the AVPs that contradicted each other.
	ResultCode_DIAMETER_CONTRADICTING_AVPS ResultCode = 5007

	// A message was received with an AVP that MUST NOT be present.  The
	// Failed-AVP AVP MUST be included and contain a copy of the
	// offending AVP.
	ResultCode_DIAMETER_AVP_NOT_ALLOWED ResultCode = 5008

	// A message was received that included an AVP that appeared more
	// often than permitted in the message definition.  The Failed-AVP
	// AVP MUST be included and contain a copy of the first instance of
	// the offending AVP that exceeded the maximum number of occurrences.
	ResultCode_DIAMETER_AVP_OCCURS_TOO_MANY_TIMES ResultCode = 5009

	// This error is returned by a Diameter node that receives a CER
	// whereby no applications are common between the CER sending peer
	// and the CER receiving peer.
	ResultCode_DIAMETER_NO_COMMON_APPLICATION ResultCode = 5010

	// This error is returned when a request was received, whose version
	// number is unsupported.
	ResultCode_DIAMETER_UNSUPPORTED_VERSION ResultCode = 5011

	// This error is returned when a request is rejected for unspecified
	// reasons.
	ResultCode_DIAMETER_UNABLE_TO_COMPLY ResultCode = 5012

	// This error is returned when a reserved bit in the Diameter header
	// is set to one (1) or the bits in the Diameter header are set
	// incorrectly.
	ResultCode_DIAMETER_INVALID_BIT_IN_HEADER ResultCode = 5013

	// The request contained an AVP with an invalid length.  A Diameter
	// message indicating this error MUST include the offending AVPs
	// within a Failed-AVP AVP.  In cases where the erroneous AVP length
	// value exceeds the message length or is less than the minimum AVP
	// header length, it is sufficient to include the offending AVP
	// header and a zero filled payload of the minimum required length
	// for the payloads data type.  If the AVP is a Grouped AVP, the
	// Grouped AVP header with an empty payload would be sufficient to
	// indicate the offending AVP.  In the case where the offending AVP
	// header cannot be fully decoded when the AVP length is less than
	// the minimum AVP header length, it is sufficient to include an
	// offending AVP header that is formulated by padding the incomplete
	// AVP header with zero up to the minimum AVP header length.
	ResultCode_DIAMETER_INVALID_AVP_LENGTH ResultCode = 5014

	// This error is returned when a request is received with an invalid
	// message length.
	ResultCode_DIAMETER_INVALID_MESSAGE_LENGTH ResultCode = 5015

	// The request contained an AVP with which is not allowed to have the
	// given value in the AVP Flags field.  A Diameter message indicating
	// this error MUST include the offending AVPs within a Failed-AVP
	// AVP.
	ResultCode_DIAMETER_INVALID_AVP_BIT_COMBO ResultCode = 5016

	// This error is returned when a CER message is received, and there
	// are no common security mechanisms supported between the peers.  A
	// Capabilities-Exchange-Answer (CEA) message MUST be returned with
	// the Result-Code AVP set to DIAMETER_NO_COMMON_SECURITY.
	ResultCode_DIAMETER_NO_COMMON_SECURITY ResultCode = 5017

	// This result code shall be sent by the MME over the SGd interface or
	// by the SGSN over the Gdd interface to indicate that the user identified
	// by the IMSI is unknown.
	// This result code shall be sent by the SMS-IWMSC over the SGd interface to
	// indicate that the user identified by the MSISDN is unknown.
	// This result code shall be sent by the HSS or the SMS Router over the S6c
	// interface to indicate that the user identified by the MSISDN is unknown.
	ResultCode_DIAMETER_ERROR_USER_UNKNOWN ResultCode = 5001

	// This result code shall be sent by the MME over the SGd interface or by the
	// SGSN over the Gdd interface to indicate that the UE is not reachable.
	// This result code shall be sent by the HSS or the SMS Router over the S6c
	// interface to indicate that the UE is not reachable.
	ResultCode_DIAMETER_ERROR_ABSENT_USER ResultCode = 5550

	// This result code shall be sent by the MME or the SGSN when the user is busy
	// for MT SMS.
	ResultCode_DIAMETER_ERROR_USER_BUSY_FOR_MT_SMS ResultCode = 5551

	// This result code shall be sent to indicate a requested facility is not supported.
	// NOTE: This code corresponds to the Facility Not Supported MAP error and may be
	// used by an IWF.
	ResultCode_DIAMETER_ERROR_FACILITY_NOT_SUPPORTED ResultCode = 5552

	// This result code shall be sent by the MME or the SGSN to indicate that the delivery
	// of the mobile terminated short message failed because the mobile station failed
	// authentication.
	ResultCode_DIAMETER_ERROR_ILLEGAL_USER ResultCode = 5553

	// This result code shall be sent by the MME or the SGSN to indicate that the
	// delivery of the mobile terminated short message failed because an IMEI
	// check failed, i.e. the IMEI was prohibited-listed or not permitted-listed.
	ResultCode_DIAMETER_ERROR_ILLEGAL_EQUIPMENT ResultCode = 5554

	// This result code shall be sent by the MME or the SGSN or the SMS-IWMSC to
	// indicate that the delivery of the mobile terminated short message failed.
	ResultCode_DIAMETER_ERROR_SM_DELIVERY_FAILURE ResultCode = 5555

	// This result code shall be sent by the HSS or the SMS Router over the S6c
	// interface to indicate that the MT SMS Teleservice is not part of the subscription.
	ResultCode_DIAMETER_ERROR_SERVICE_NOT_SUBSCRIBED ResultCode = 5556

	// This result code shall be sent by the HSS or the SMS Router over the S6c
	// interface to indicate that the MT SMS Teleservice is barred.
	// This result code shall be sent by the MME to indicate that the delivery of
	// the mobile terminated short message failed because of the barring of the
	// SMS service.
	ResultCode_DIAMETER_ERROR_SERVICE_BARRED ResultCode = 5557

	// This result code shall be sent by the HSS over the S6c interface to indicate
	// that the Message Waiting List is full.
	ResultCode_DIAMETER_ERROR_MWD_LIST_FULL ResultCode = 5558
)

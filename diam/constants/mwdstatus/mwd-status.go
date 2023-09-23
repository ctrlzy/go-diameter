package mwdstatus

// 3GPP TS 29.338 section 5.3.3.8
// The MWD-Status AVP is of type Unsigned32 and it shall contain a bit mask. The meaning of the bits shall be as defined in table 5.3.3.8/1:
// Table 5.3.3.8/1: MWD Status
// bit name Description
// 0 SC-Address Not included This bit when set shall indicate that the SC Address has not been added to the Message Waiting Data in the HSS.
// 1 MNRF-Set This bit, when set, shall indicate that the MNRF flag is set in the HSS
// 2 MCEF-Set This bit, when set, shall indicate that the MCEF flag is set in the HSS.
// 3 MNRG-Set This bit, when set, shall indicate that the MNRG flag is set in the HSS
// 4 MNR5G-Set This bit, when set, shall indicate that the HSS/UDM is waiting for a reachability notification / registration from 5G serving nodes.
// NOTE: Bits not defined in this table shall be cleared by the sending HSS and discarded by the receiving MME

const (
	SCAddressNotincluded = 1
	MNRFSet              = 1 << 1
	MCEFSet              = 1 << 2
	MNRGSet              = 1 << 3
	MNR5GSet             = 1 << 4
)

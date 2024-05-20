// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package sm_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
	"github.com/ctrlzy/go-diameter/v4/diam/diamtest"
	"github.com/ctrlzy/go-diameter/v4/diam/dict"
	"github.com/ctrlzy/go-diameter/v4/diam/sm"
	"github.com/ctrlzy/go-diameter/v4/diam/sm/smparser"
)

var (
	baseCERIdx = diam.CommandIndex{AppID: 0, Code: diam.CapabilitiesExchange, Request: true}
	baseCEAIdx = diam.CommandIndex{AppID: 0, Code: diam.CapabilitiesExchange, Request: false}
	baseDWRIdx = diam.CommandIndex{AppID: 0, Code: diam.DeviceWatchdog, Request: true}
)

func TestClient_Dial_MissingStateMachine(t *testing.T) {
	cli := &sm.Client{}
	_, err := cli.Dial("")
	if err != sm.ErrMissingStateMachine {
		t.Fatal(err)
	}
}

func TestClient_Dial_InvalidAddress(t *testing.T) {
	cli := &sm.Client{
		Handler: sm.New(clientSettings),
		AcctApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0,
				datatype.Unsigned32(0)),
		},
	}
	c, err := cli.Dial(":0")
	if err == nil {
		c.Close()
		t.Fatal("Invalid client address succeeded")
	}
}

func TestClient_DialTLS_InvalidAddress(t *testing.T) {
	cli := &sm.Client{
		Handler: sm.New(clientSettings),
		AcctApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(0)),
		},
	}
	c, err := cli.DialTLS(":0", "", "")
	if err == nil {
		c.Close()
		t.Fatal("Invalid client address succeeded")
	}
}

func TestClient_Handshake(t *testing.T) {
	srv := diamtest.NewServer(sm.New(serverSettings), dict.Default)
	defer srv.Close()
	cli := &sm.Client{
		Handler: sm.New(clientSettings),
		SupportedVendorID: []*diam.AVP{
			diam.NewAVP(avp.SupportedVendorID, avp.Mbit, 0, clientSettings.VendorID),
		},
		AcctApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(3)),
		},
		AuthApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(4)),
		},
		VendorSpecificApplicationID: []*diam.AVP{
			diam.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
				AVP: []*diam.AVP{
					diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(4)),
				},
			}),
		},
	}
	c, err := cli.Dial(srv.Addr)
	if err != nil {
		t.Fatal(err)
	}
	c.Close()
}

func TestClient_Handshake_CustomIP_TCP(t *testing.T) {
	testClient_Handshake_CustomIP(t, "tcp")
}

func testClient_Handshake_CustomIP(t *testing.T, network string) {
	srv := diamtest.NewServerNetwork(network, sm.New(serverSettings), dict.Default)
	defer srv.Close()
	cli := &sm.Client{
		RetransmitInterval: time.Second * 3,
		Handler:            sm.New(clientSettings2),
		SupportedVendorID: []*diam.AVP{
			diam.NewAVP(avp.SupportedVendorID, avp.Mbit, 0, clientSettings.VendorID),
		},
		AcctApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(3)),
		},
		AuthApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(4)),
		},
		VendorSpecificApplicationID: []*diam.AVP{
			diam.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
				AVP: []*diam.AVP{
					diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(4)),
				},
			}),
		},
	}
	c, err := cli.DialNetwork(network, srv.Addr)
	if err != nil {
		t.Fatal(err)
	}
	c.Close()
}

func TestClient_Handshake_Notify(t *testing.T) {
	srv := diamtest.NewServer(sm.New(serverSettings), dict.Default)
	defer srv.Close()
	cli := &sm.Client{
		Handler: sm.New(clientSettings),
		SupportedVendorID: []*diam.AVP{
			diam.NewAVP(avp.SupportedVendorID, avp.Mbit, 0, clientSettings.VendorID),
		},
		AcctApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(3)),
		},
		AuthApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(4)),
		},
		VendorSpecificApplicationID: []*diam.AVP{
			diam.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
				AVP: []*diam.AVP{
					diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(4)),
				},
			}),
		},
	}
	handshakeOK := make(chan struct{})
	go func() {
		<-cli.Handler.HandshakeNotify()
		close(handshakeOK)
	}()
	c, err := cli.Dial(srv.Addr)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	select {
	case <-handshakeOK:
	case <-time.After(time.Second):
		t.Fatal("Handshake timed out")
	}
}

func TestClient_Handshake_FailParseCEA(t *testing.T) {
	mux := diam.NewServeMux()
	mux.HandleFunc("CER", func(c diam.Conn, m *diam.Message) {
		a := m.Answer(diam.Success)
		// Missing Origin-Host and other mandatory AVPs.
		a.WriteTo(c)
	})
	srv := diamtest.NewServer(mux, dict.Default)
	defer srv.Close()
	cli := &sm.Client{
		Handler: sm.New(clientSettings),
		AcctApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(3)),
		},
	}
	_, err := cli.Dial(srv.Addr)
	if err != smparser.ErrMissingOriginHost {
		t.Fatal(err)
	}
}

func TestClient_Handshake_FailedResultCode(t *testing.T) {
	mux := diam.NewServeMux()
	mux.HandleFunc("CER", func(c diam.Conn, m *diam.Message) {
		cer := new(smparser.CER)
		if _, err := cer.Parse(m, smparser.Server); err != nil {
			panic(err)
		}
		a := m.Answer(diam.NoCommonApplication)
		a.NewAVP(avp.OriginHost, avp.Mbit, 0, clientSettings.OriginHost)
		a.NewAVP(avp.OriginRealm, avp.Mbit, 0, clientSettings.OriginRealm)
		if cer.OriginStateID != nil {
			a.AddAVP(cer.OriginStateID)
		}
		a.AddAVP(cer.AcctApplicationID[0]) // The one we send below.
		a.WriteTo(c)
	})
	srv := diamtest.NewServer(mux, dict.Default)
	defer srv.Close()
	cli := &sm.Client{
		Handler: sm.New(clientSettings),
		AcctApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(3)),
		},
	}
	_, err := cli.Dial(srv.Addr)
	if err == nil {
		t.Fatal("Unexpected CER worked")
	}
	e, ok := err.(*smparser.ErrFailedResultCode)
	if !ok {
		t.Fatal(err)
	}
	if !strings.Contains(e.Error(), "failed Result-Code AVP") {
		t.Fatal(e.Error())
	}
}

func TestClient_Handshake_RetransmitTimeout(t *testing.T) {
	mux := diam.NewServeMux()
	var retransmits uint32
	mux.HandleFunc("CER", func(c diam.Conn, m *diam.Message) {
		// Do nothing to force timeout.
		atomic.AddUint32(&retransmits, 1)
	})
	srv := diamtest.NewServer(mux, dict.Default)
	defer srv.Close()
	cli := &sm.Client{
		Handler:            sm.New(clientSettings),
		MaxRetransmits:     3,
		RetransmitInterval: time.Millisecond,
		AcctApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(3)),
		},
	}
	_, err := cli.Dial(srv.Addr)
	if err == nil {
		t.Fatal("Unexpected CER worked")
	}
	if err != sm.ErrHandshakeTimeout {
		t.Fatal(err)
	}
	if n := atomic.LoadUint32(&retransmits); n != 4 {
		t.Fatalf("Unexpected # of retransmits. Want 4, have %d", n)
	}
}

type testDWAHandler struct {
	Dwac chan struct{}
}

func (h testDWAHandler) ServeDIAM(c diam.Conn, m *diam.Message) {
	if m.Header.CommandCode == diam.DeviceWatchdog {
		h.Dwac <- struct{}{}
	}
	if m.Header.CommandCode == diam.CapabilitiesExchange {
		a := m.Answer(diam.Success)
		// Fix for Same H2H and E2E Identifier in success response
		a.Header.HopByHopID = m.Header.HopByHopID
		a.Header.EndToEndID = m.Header.EndToEndID
		a.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
		a.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
		a.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415))
		a.NewAVP(avp.SupportedVendorID, avp.Mbit, 0, datatype.Unsigned32(10415))
		a.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
			AVP: []*diam.AVP{
				diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
				diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(0)),
			},
		})
		_, _ = a.WriteTo(c)
	}
}

func TestClient_Watchdog(t *testing.T) {
	handler := testDWAHandler{Dwac: make(chan struct{}, 1)}
	srv := diamtest.NewServer(handler, dict.Default)
	defer srv.Close()
	cli := &sm.Client{
		EnableWatchdog:   true,
		WatchdogInterval: 100 * time.Millisecond,
		Handler:          sm.New(clientSettings),
		AcctApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(3)),
		},
	}
	c, err := cli.Dial(srv.Addr)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	select {
	case <-handler.Dwac:
	case <-time.After(200 * time.Millisecond):
		t.Fatal("Timeout waiting for DWA")
	}
}

type timeoutDWAHandler struct {
}

func (h timeoutDWAHandler) ServeDIAM(c diam.Conn, m *diam.Message) {
	if m.Header.CommandCode == diam.DeviceWatchdog {
		m.Answer(diam.UnableToComply).WriteTo(c)
	}
	if m.Header.CommandCode == diam.CapabilitiesExchange {
		a := m.Answer(diam.Success)
		// Fix for Same H2H and E2E Identifier in success response
		a.Header.HopByHopID = m.Header.HopByHopID
		a.Header.EndToEndID = m.Header.EndToEndID
		a.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
		a.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
		a.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415))
		a.NewAVP(avp.SupportedVendorID, avp.Mbit, 0, datatype.Unsigned32(10415))
		a.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
			AVP: []*diam.AVP{
				diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
				diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(0)),
			},
		})
		_, _ = a.WriteTo(c)
	}
}

func TestClient_Watchdog_Timeout(t *testing.T) {
	handler := timeoutDWAHandler{}
	srv := diamtest.NewServer(handler, dict.Default)
	defer srv.Close()
	cli := &sm.Client{
		MaxRetransmits:     3,
		RetransmitInterval: 50 * time.Millisecond,
		EnableWatchdog:     true,
		WatchdogInterval:   50 * time.Millisecond,
		Handler:            sm.New(clientSettings),
		AcctApplicationID: []*diam.AVP{
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(3)),
		},
	}
	c, err := cli.Dial(srv.Addr)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	select {
	case <-c.(diam.CloseNotifier).CloseNotify():
	case <-time.After(500 * time.Millisecond):
		t.Fatal("Timeout waiting for watchdog to disconnect client")
	}
}

// Type matching interface: net.Addr
type testLocalAddr struct {
	value string
}

func (a testLocalAddr) Network() string { return "tcp" }
func (a testLocalAddr) String() string  { return a.value }

// Type matching interface: diam.Conn
type testLocalAddrDiamConn struct {
	localAddr *testLocalAddr
}

func (d testLocalAddrDiamConn) Write(b []byte) (int, error)                    { return 0, nil }
func (d testLocalAddrDiamConn) WriteStream(b []byte, stream uint) (int, error) { return 0, nil }
func (d testLocalAddrDiamConn) Close()                                         {}
func (d testLocalAddrDiamConn) LocalAddr() net.Addr                            { return d.localAddr }
func (d testLocalAddrDiamConn) RemoteAddr() net.Addr                           { return nil }
func (d testLocalAddrDiamConn) TLS() *tls.ConnectionState                      { return nil }
func (d testLocalAddrDiamConn) Dictionary() *dict.Parser                       { return nil }
func (d testLocalAddrDiamConn) Context() context.Context                       { return context.Background() }
func (d testLocalAddrDiamConn) SetContext(c context.Context)                   {}
func (d testLocalAddrDiamConn) Connection() net.Conn                           { return nil }

func newTestLocalAddrDiamConn(localAddrValue string) diam.Conn {
	return testLocalAddrDiamConn{
		localAddr: &testLocalAddr{
			value: localAddrValue,
		},
	}
}

func TestClient_Conn_LocalAddresses_Loopback(t *testing.T) {
	c := newTestLocalAddrDiamConn("127.0.0.1:3868")

	addrList, err := getLocalAddresses(c)
	if err != nil {
		t.Fatalf("Failed to parse local addresses: %v", err)
	}
	if len(addrList) != 1 {
		t.Fatal("The only available loopback address was skipped")
	}
}

func TestClient_Conn_LocalAddresses_Complex(t *testing.T) {
	c := newTestLocalAddrDiamConn("127.0.0.1/[::1%lo]/10.0.0.3/[fe80::78ef:0efb:a57b:15b9%eth0]:3868")

	addrList, err := getLocalAddresses(c)
	if err != nil {
		t.Fatalf("Failed to parse local addresses: %v", err)
	}
	if len(addrList) != 1 {
		t.Fatal("Failed to parse valid IP address or failed to skip loopback")
	}

	actual := net.IP(addrList[0]).String()
	expected := "10.0.0.3"
	if actual != expected {
		t.Fatalf("Wrong IP address found in list of local addresses, expected: %s, actual: %s", expected, actual)
	}
}

func getLocalAddresses(c diam.Conn) ([]datatype.Address, error) {
	var (
		addr, addrStr string
		loopback      net.IP
		err           error
	)
	if c.LocalAddr() != nil {
		addrStr = c.LocalAddr().String()
	}
	if addrStr != "" {
		addr, err = getHostsWithoutPort(addrStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse local ip %s [%q]: %s", addrStr, c.LocalAddr(), err)
		}
	}
	hostIPs := strings.Split(addr, "/")
	addresses := make([]datatype.Address, 0, len(hostIPs))
	for _, ipStr := range hostIPs {
		ip := net.ParseIP(ipStr)
		if ip != nil {
			if ip.IsLoopback() {
				loopback = ip
			} else {
				addresses = append(addresses, datatype.Address(ip))
			}
		}
	}
	if len(addresses) == 0 && loopback != nil {
		addresses = append(addresses, datatype.Address(loopback))
	}
	return addresses, nil
}

func getHostsWithoutPort(hosts string) (string, error) {
	i := len(hosts) - 1
	for ; i >= 0 && hosts[i] != ':'; i-- {
		if hosts[i] < '0' || hosts[i] > '9' {
			return "", fmt.Errorf("found non numerical character in port at position %d", i+1)
		}
	}
	return hosts[:i], nil
}

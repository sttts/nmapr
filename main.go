/*
  Copyright (c) 2013 José Carlos Nieto, http://xiam.menteslibres.org/

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package nmapr

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Address struct {
	Addr     string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"`
	Level    string `xml:"level,attr"`
}

type Verbose struct {
	Level string `xml:"level,attr"`
}

type Finished struct {
	time string `xml:"time,attr"`
}

type RunStats struct {
	Finished Finished `xml:"finished,attr"`
}

type Hostname struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

type PortState struct {
	State  string `xml:"state,attr"`
	Reason string `xml:"reason,attr"`
}

type PortService struct {
	Name    string `xml:"name,attr"`
	Product string `xml:"product,attr"`
	Version string `xml:"version,attr"`
	Extra   string `xml:"extrainfo,attr"`
}

type Match struct {
	Name     string `xml:"name,attr"`
	Accuracy uint   `xml:"accuracy,attr"`
}

type Class struct {
	Type     string `xml:"type,attr"`
	Vendor   string `xml:"vendor,attr"`
	Family   string `xml:"osfamily,attr"`
	Gen      string `xml:"osgen,attr"`
	Accuracy uint   `xml:"accuracy,attr"`
}

type OS struct {
	Match Match `xml:"osmatch"`
	Class Class `xml:"osclass"`
}

type Port struct {
	Protocol string      `xml:"protocol,attr"`
	PortID   uint        `xml:"portid,attr"`
	State    PortState   `xml:"state"`
	Service  PortService `xml:"service"`
}

type Host struct {
	StartTime string     `xml:"starttime,attr"`
	Address   []Address  `xml:"address"`
	Hostnames []Hostname `xml:"hostnames>hostname"`
	Ports     []Port     `xml:"ports>port"`
	OS        []OS       `xml:"os"`
}

type Report struct {
	XMLName  xml.Name `xml:"nmaprun"`
	Scanner  string   `xml:"scanner,attr"`
	Args     string   `xml:"args,attr"`
	Verbose  Verbose  `xml:"verbose"`
	Start    uint64   `xml:"start,attr"`
	StartStr string   `xml:"startstr,attr"`
	Host     []Host   `xml:"host"`
}

func Open(file string) (*Report, error) {

	fp, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	defer fp.Close()

	info, err := fp.Stat()

	if err != nil {
		return nil, err
	}

	if info.IsDir() == true {
		return nil, fmt.Errorf("Could not open %s, is a directory.", file)
	}

	buf := make([]byte, info.Size())

	fp.Read(buf)

	report := Report{}

	err = xml.Unmarshal(buf, &report)

	if err != nil {
		return nil, err
	}

	return &report, nil
}
